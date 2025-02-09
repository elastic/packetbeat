// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package beater

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/elastic/beats/v7/filebeat/backup"
	"github.com/elastic/beats/v7/filebeat/channel"
	cfg "github.com/elastic/beats/v7/filebeat/config"
	"github.com/elastic/beats/v7/filebeat/fileset"
	_ "github.com/elastic/beats/v7/filebeat/include"
	"github.com/elastic/beats/v7/filebeat/input"
	"github.com/elastic/beats/v7/filebeat/input/filestream"
	"github.com/elastic/beats/v7/filebeat/input/filestream/takeover"
	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	"github.com/elastic/beats/v7/filebeat/input/v2/compat"
	"github.com/elastic/beats/v7/filebeat/registrar"
	"github.com/elastic/beats/v7/libbeat/autodiscover"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/beats/v7/libbeat/esleg/eslegclient"
	"github.com/elastic/beats/v7/libbeat/management"
	"github.com/elastic/beats/v7/libbeat/monitoring/inputmon"
	"github.com/elastic/beats/v7/libbeat/outputs/elasticsearch"
	"github.com/elastic/beats/v7/libbeat/publisher/pipetool"
	"github.com/elastic/beats/v7/libbeat/statestore"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/monitoring"
	"github.com/elastic/elastic-agent-libs/paths"
	"github.com/elastic/go-concert/unison"

	// Add filebeat level processors
	_ "github.com/elastic/beats/v7/filebeat/processor/add_kubernetes_metadata"
	_ "github.com/elastic/beats/v7/libbeat/processors/decode_csv_fields"

	// include all filebeat specific autodiscover features
	_ "github.com/elastic/beats/v7/filebeat/autodiscover"
)

const pipelinesWarning = "Filebeat is unable to load the ingest pipelines for the configured" +
	" modules because the Elasticsearch output is not configured/enabled. If you have" +
	" already loaded the ingest pipelines or are using Logstash pipelines, you" +
	" can ignore this warning."

var once = flag.Bool("once", false, "Run filebeat only once until all harvesters reach EOF")

// Filebeat is a beater object. Contains all objects needed to run the beat
type Filebeat struct {
	config         *cfg.Config
	moduleRegistry *fileset.ModuleRegistry
	pluginFactory  PluginFactory
	done           chan struct{}
	stopOnce       sync.Once // wraps the Stop() method
	pipeline       beat.PipelineConnector
}

type PluginFactory func(beat.Info, *logp.Logger, StateStore) []v2.Plugin

type StateStore interface {
	// Access returns the storage registry depending on the type.
	// The value of typ is expected to have been obtained from
	// cursor.InputManager.Type and represents the input type.
	Access(typ string) (*statestore.Store, error)
	CleanupInterval() time.Duration
}

// New creates a new Filebeat pointer instance.
func New(plugins PluginFactory) beat.Creator {
	return func(b *beat.Beat, rawConfig *conf.C) (beat.Beater, error) {
		return newBeater(b, plugins, rawConfig)
	}
}

func newBeater(b *beat.Beat, plugins PluginFactory, rawConfig *conf.C) (beat.Beater, error) {
	config := cfg.DefaultConfig
	if err := rawConfig.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	if err := cfgwarn.CheckRemoved6xSettings(
		rawConfig,
		"prospectors",
		"config.prospectors",
		"registry_file",
		"registry_file_permissions",
		"registry_flush",
	); err != nil {
		return nil, err
	}

	enableAllFilesets, _ := b.BeatConfig.Bool("config.modules.enable_all_filesets", -1)
	forceEnableModuleFilesets, _ := b.BeatConfig.Bool("config.modules.force_enable_module_filesets", -1)
	filesetOverrides := fileset.FilesetOverrides{
		EnableAllFilesets:         enableAllFilesets,
		ForceEnableModuleFilesets: forceEnableModuleFilesets,
	}
	moduleRegistry, err := fileset.NewModuleRegistry(config.Modules, b.Info, true, filesetOverrides)
	if err != nil {
		return nil, err
	}

	moduleInputs, err := moduleRegistry.GetInputConfigs()
	if err != nil {
		return nil, err
	}

	if err := config.FetchConfigs(); err != nil {
		return nil, err
	}

	if b.API != nil {
		if err = inputmon.AttachHandler(b.API.Router()); err != nil {
			return nil, fmt.Errorf("failed attach inputs api to monitoring endpoint server: %w", err)
		}
	}

	if b.Manager != nil {
		b.Manager.RegisterDiagnosticHook("input_metrics", "Metrics from active inputs.",
			"input_metrics.json", "application/json", func() []byte {
				data, err := inputmon.MetricSnapshotJSON()
				if err != nil {
					logp.L().Warnw("Failed to collect input metric snapshot for Agent diagnostics.", "error", err)
					return []byte(err.Error())
				}
				return data
			})

		b.Manager.RegisterDiagnosticHook(
			"registry",
			"Filebeat's registry",
			"registry.tar.gz",
			"application/octet-stream",
			gzipRegistry)
	}

	// Add inputs created by the modules
	config.Inputs = append(config.Inputs, moduleInputs...)

	enabledInputs := config.ListEnabledInputs()
	var haveEnabledInputs bool
	if len(enabledInputs) > 0 {
		haveEnabledInputs = true
	}

	if !config.ConfigInput.Enabled() && !config.ConfigModules.Enabled() && !haveEnabledInputs && config.Autodiscover == nil && !b.Manager.Enabled() {
		if !b.InSetupCmd {
			return nil, fmt.Errorf("no modules or inputs enabled and configuration reloading disabled. What files do you want me to watch?")
		}

		// in the `setup` command, log this only as a warning
		logp.Warn("Setup called, but no modules enabled.")
	}

	if *once && config.ConfigInput.Enabled() && config.ConfigModules.Enabled() {
		return nil, fmt.Errorf("input configs and --once cannot be used together")
	}

	if config.IsInputEnabled("stdin") && len(enabledInputs) > 1 {
		return nil, fmt.Errorf("stdin requires to be run in exclusive mode, configured inputs: %s", strings.Join(enabledInputs, ", "))
	}

	fb := &Filebeat{
		done:           make(chan struct{}),
		config:         &config,
		moduleRegistry: moduleRegistry,
		pluginFactory:  plugins,
	}

	err = fb.setupPipelineLoaderCallback(b)
	if err != nil {
		return nil, err
	}

	return fb, nil
}

// setupPipelineLoaderCallback sets the callback function for loading pipelines during setup.
func (fb *Filebeat) setupPipelineLoaderCallback(b *beat.Beat) error {
	if b.Config.Output.Name() != "elasticsearch" && !b.Manager.Enabled() {
		logp.Warn(pipelinesWarning)
		return nil
	}

	overwritePipelines := true
	b.OverwritePipelinesCallback = func(esConfig *conf.C) error {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		esClient, err := eslegclient.NewConnectedClient(ctx, esConfig, "Filebeat")
		if err != nil {
			return err
		}

		// When running the subcommand setup, configuration from modules.d directories
		// have to be loaded using cfg.Reloader. Otherwise those configurations are skipped.
		pipelineLoaderFactory := newPipelineLoaderFactory(ctx, b.Config.Output.Config())
		enableAllFilesets, _ := b.BeatConfig.Bool("config.modules.enable_all_filesets", -1)
		forceEnableModuleFilesets, _ := b.BeatConfig.Bool("config.modules.force_enable_module_filesets", -1)
		filesetOverrides := fileset.FilesetOverrides{
			EnableAllFilesets:         enableAllFilesets,
			ForceEnableModuleFilesets: forceEnableModuleFilesets,
		}

		modulesFactory := fileset.NewSetupFactory(b.Info, pipelineLoaderFactory, filesetOverrides)
		if fb.config.ConfigModules.Enabled() {
			if enableAllFilesets {
				// All module configs need to be loaded to enable all the filesets
				// contained in the modules.  The default glob just loads the enabled
				// ones.  Switching the glob pattern from *.yml to * achieves this.
				origPath, _ := fb.config.ConfigModules.String("path", -1)
				newPath := strings.TrimSuffix(origPath, ".yml")
				_ = fb.config.ConfigModules.SetString("path", -1, newPath)
			}
			modulesLoader := cfgfile.NewReloader(logp.L().Named("module.reloader"), fb.pipeline, fb.config.ConfigModules)
			modulesLoader.Load(modulesFactory)
		}

		return fb.moduleRegistry.LoadPipelines(esClient, overwritePipelines)
	}
	return nil
}

// loadModulesPipelines is called when modules are configured to do the initial
// setup.
func (fb *Filebeat) loadModulesPipelines(b *beat.Beat) error {
	if b.Config.Output.Name() != "elasticsearch" {
		logp.Warn(pipelinesWarning)
		return nil
	}

	overwritePipelines := fb.config.OverwritePipelines
	if b.InSetupCmd {
		overwritePipelines = true
	}

	// register pipeline loading to happen every time a new ES connection is
	// established
	callback := func(esClient *eslegclient.Connection) error {
		return fb.moduleRegistry.LoadPipelines(esClient, overwritePipelines)
	}
	_, err := elasticsearch.RegisterConnectCallback(callback)

	return err
}

// Run allows the beater to be run as a beat.
func (fb *Filebeat) Run(b *beat.Beat) error {
	var err error
	config := fb.config

	if !fb.moduleRegistry.Empty() {
		err = fb.loadModulesPipelines(b)
		if err != nil {
			return err
		}
	}

	waitFinished := newSignalWait()
	waitEvents := newSignalWait()

	// count active events for waiting on shutdown
	var reg *monitoring.Registry

	if b.Info.Monitoring.Namespace != nil {
		reg = b.Info.Monitoring.Namespace.GetRegistry().GetRegistry("stats")
		if reg == nil {
			reg = b.Info.Monitoring.Namespace.GetRegistry().NewRegistry("stats")
		}
	}
	wgEvents := &eventCounter{
		count: monitoring.NewInt(reg, "filebeat.events.active"), // Gauge
		added: monitoring.NewUint(reg, "filebeat.events.added"),
		done:  monitoring.NewUint(reg, "filebeat.events.done"),
	}
	finishedLogger := newFinishedLogger(wgEvents)

	registryMigrator := registrar.NewMigrator(config.Registry)
	if err := registryMigrator.Run(); err != nil {
		logp.Err("Failed to migrate registry file: %+v", err)
		return err
	}

	// Use context, like normal people do, hooking up to the beat.done channel
	ctx, cn := context.WithCancel(context.Background())
	go func() {
		<-fb.done
		cn()
	}()

	stateStore, err := openStateStore(ctx, b.Info, logp.NewLogger("filebeat"), config.Registry)
	if err != nil {
		logp.Err("Failed to open state store: %+v", err)
		return err
	}
	defer stateStore.Close()

	// If notifier is set, configure the listener for output configuration
	// The notifier passes the elasticsearch output configuration down to the Elasticsearch backed state storage
	// in order to allow it fully configure
	if stateStore.notifier != nil {
		b.OutputConfigReloader = reload.ReloadableFunc(func(r *reload.ConfigWithMeta) error {
			outCfg := conf.Namespace{}
			if err := r.Config.Unpack(&outCfg); err != nil || outCfg.Name() != "elasticsearch" {
				logp.Err("Failed to unpack the output config: %v", err)
				return nil
			}

			stateStore.notifier.Notify(outCfg.Config())
			return nil
		})
	}

	err = filestream.ValidateInputIDs(config.Inputs, logp.NewLogger("input.filestream"))
	if err != nil {
		logp.Err("invalid filestream configuration: %+v", err)
		return err
	}
	err = processLogInputTakeOver(stateStore, config)
	if err != nil {
		logp.Err("Failed to attempt filestream state take over: %+v", err)
		return err
	}

	// Setup registrar to persist state
	registrar, err := registrar.New(stateStore, finishedLogger, config.Registry.FlushTimeout)
	if err != nil {
		logp.Err("Could not init registrar: %v", err)
		return err
	}

	// Make sure all events that were published in
	registrarChannel := newRegistrarLogger(registrar)

	// setup event counting for startup and a global common ACKer, such that all events will be
	// routed to the reigstrar after they've been ACKed.
	// Events with Private==nil or the type of private != file.State are directly
	// forwarded to `finishedLogger`. Events from the `logs` input will first be forwarded
	// to the registrar via `registrarChannel`, which finally forwards the events to finishedLogger as well.
	// The finishedLogger decrements the counters in wgEvents after all events have been securely processed
	// by the registry.
	fb.pipeline = withPipelineEventCounter(b.Publisher, wgEvents)
	fb.pipeline = pipetool.WithACKer(fb.pipeline, eventACKer(finishedLogger, registrarChannel))

	// Filebeat by default required infinite retry. Let's configure this for all
	// inputs by default.  Inputs (and InputController) can overwrite the sending
	// guarantees explicitly when connecting with the pipeline.
	fb.pipeline = pipetool.WithDefaultGuarantees(fb.pipeline, beat.GuaranteedSend)

	outDone := make(chan struct{}) // outDone closes down all active pipeline connections
	pipelineConnector := channel.NewOutletFactory(outDone).Create

	inputsLogger := logp.NewLogger("input")
	v2Inputs := fb.pluginFactory(b.Info, inputsLogger, stateStore)
	v2InputLoader, err := v2.NewLoader(inputsLogger, v2Inputs, "type", cfg.DefaultType)
	if err != nil {
		panic(err) // loader detected invalid state.
	}

	var inputTaskGroup unison.TaskGroup
	defer func() {
		_ = inputTaskGroup.Stop()
	}()

	// Store needs to be fully configured at this point
	if err := v2InputLoader.Init(&inputTaskGroup); err != nil {
		logp.Err("Failed to initialize the input managers: %v", err)
		return err
	}

	inputLoader := channel.RunnerFactoryWithCommonInputSettings(b.Info, compat.Combine(
		compat.RunnerFactory(inputsLogger, b.Info, v2InputLoader),
		input.NewRunnerFactory(pipelineConnector, registrar, fb.done),
	))

	// Create a ES connection factory for dynamic modules pipeline loading
	var pipelineLoaderFactory fileset.PipelineLoaderFactory
	// The pipelineFactory needs a context to control the connections to ES,
	// when the pipelineFactory/ESClient are not needed any more the context
	// must be cancelled. This pipeline factory will be used by the moduleLoader
	// that is run by a crawler, whenever this crawler is stopped we also cancel
	// the context.
	pipelineFactoryCtx, cancelPipelineFactoryCtx := context.WithCancel(context.Background())
	defer cancelPipelineFactoryCtx()
	if b.Config.Output.Name() == "elasticsearch" {
		pipelineLoaderFactory = newPipelineLoaderFactory(pipelineFactoryCtx, b.Config.Output.Config())
	} else {
		logp.Warn(pipelinesWarning)
	}
	moduleLoader := fileset.NewFactory(inputLoader, b.Info, pipelineLoaderFactory, config.OverwritePipelines)
	crawler, err := newCrawler(inputLoader, moduleLoader, config.Inputs, fb.done, *once)
	if err != nil {
		logp.Err("Could not init crawler: %v", err)
		return err
	}

	// The order of starting and stopping is important. Stopping is inverted to the starting order.
	// The current order is: registrar, publisher, spooler, crawler
	// That means, crawler is stopped first.

	// Start the registrar
	err = registrar.Start()
	if err != nil {
		return fmt.Errorf("Could not start registrar: %w", err)
	}

	// Stopping registrar will write last state
	defer registrar.Stop()

	// Stopping publisher (might potentially drop items)
	defer func() {
		// Closes first the registrar logger to make sure not more events arrive at the registrar
		// registrarChannel must be closed first to potentially unblock (pretty unlikely) the publisher
		registrarChannel.Close()
		close(outDone) // finally close all active connections to publisher pipeline
	}()

	// Wait for all events to be processed or timeout
	defer waitEvents.Wait()

	if config.OverwritePipelines {
		logp.Debug("modules", "Existing Ingest pipelines will be updated")
	}

	err = crawler.Start(fb.pipeline, config.ConfigInput, config.ConfigModules)
	if err != nil {
		crawler.Stop()
		cancelPipelineFactoryCtx()
		return fmt.Errorf("Failed to start crawler: %w", err)
	}

	// If run once, add crawler completion check as alternative to done signal
	if *once {
		runOnce := func() {
			logp.Info("Running filebeat once. Waiting for completion ...")
			crawler.WaitForCompletion()
			logp.Info("All data collection completed. Shutting down.")
		}
		waitFinished.Add(runOnce)
	}

	// Register reloadable list of inputs and modules
	inputs := cfgfile.NewRunnerList(management.DebugK, inputLoader, fb.pipeline)
	b.Registry.MustRegisterInput(inputs)

	modules := cfgfile.NewRunnerList(management.DebugK, moduleLoader, fb.pipeline)

	var adiscover *autodiscover.Autodiscover
	if fb.config.Autodiscover != nil {
		adiscover, err = autodiscover.NewAutodiscover(
			"filebeat",
			fb.pipeline,
			cfgfile.MultiplexedRunnerFactory(
				cfgfile.MatchHasField("module", moduleLoader),
				cfgfile.MatchDefault(inputLoader),
			),
			autodiscover.QueryConfig(),
			config.Autodiscover,
			b.Keystore,
		)
		if err != nil {
			return err
		}
	}
	adiscover.Start()

	// We start the manager when all the subsystem are initialized and ready to received events.
	if err := b.Manager.Start(); err != nil {
		return err
	}

	// Add done channel to wait for shutdown signal
	waitFinished.AddChan(fb.done)
	waitFinished.Wait()

	// Stop reloadable lists, autodiscover -> Stop crawler -> stop inputs -> stop harvesters
	// Note: waiting for crawlers to stop here in order to install wgEvents.Wait
	//       after all events have been enqueued for publishing. Otherwise wgEvents.Wait
	//       or publisher might panic due to concurrent updates.
	inputs.Stop()
	modules.Stop()
	adiscover.Stop()
	crawler.Stop()
	cancelPipelineFactoryCtx()

	timeout := fb.config.ShutdownTimeout
	// Checks if on shutdown it should wait for all events to be published
	waitPublished := fb.config.ShutdownTimeout > 0 || *once
	if waitPublished {
		// Wait for registrar to finish writing registry
		waitEvents.Add(withLog(wgEvents.Wait,
			"Continue shutdown: All enqueued events being published."))
		// Wait for either timeout or all events having been ACKed by outputs.
		if fb.config.ShutdownTimeout > 0 {
			logp.Info("Shutdown output timer started. Waiting for max %v.", timeout)
			waitEvents.Add(withLog(waitDuration(timeout),
				"Continue shutdown: Time out waiting for events being published."))
		} else {
			waitEvents.AddChan(fb.done)
		}
	}

	// Stop the manager and stop the connection to any dependent services.
	// The Manager started to have a working implementation when
	// https://github.com/elastic/beats/pull/34416 was merged.
	// This is intended to enable TLS certificates reload on a long
	// running Beat.
	//
	// However calling b.Manager.Stop() here messes up the behavior of the
	// --once flag because it makes Filebeat exit early.
	// So if --once is passed, we don't call b.Manager.Stop().
	if !*once {
		b.Manager.Stop()
	}

	return nil
}

// Stop is called on exit to stop the crawling, spooling and registration processes.
func (fb *Filebeat) Stop() {
	logp.Info("Stopping filebeat")

	// Stop Filebeat
	fb.stopOnce.Do(func() { close(fb.done) })
}

// Create a new pipeline loader (es client) factory
func newPipelineLoaderFactory(ctx context.Context, esConfig *conf.C) fileset.PipelineLoaderFactory {
	pipelineLoaderFactory := func() (fileset.PipelineLoader, error) {
		esClient, err := eslegclient.NewConnectedClient(ctx, esConfig, "Filebeat")
		if err != nil {
			return nil, fmt.Errorf("Error creating Elasticsearch client: %w", err)
		}
		return esClient, nil
	}
	return pipelineLoaderFactory
}

// some of the filestreams might want to take over the loginput state
// if their `take_over` flag is set to `true`.
func processLogInputTakeOver(stateStore StateStore, config *cfg.Config) error {
	inputs, err := fetchInputConfiguration(config)
	if err != nil {
		return fmt.Errorf("Failed to fetch input configuration when attempting take over: %w", err)
	}
	if len(inputs) == 0 {
		return nil
	}

	store, err := stateStore.Access("")
	if err != nil {
		return fmt.Errorf("Failed to access state when attempting take over: %w", err)
	}
	defer store.Close()
	logger := logp.NewLogger("filestream-takeover")

	registryHome := paths.Resolve(paths.Data, config.Registry.Path)
	registryHome = filepath.Join(registryHome, "filebeat")

	backuper := backup.NewRegistryBackuper(logger, registryHome)

	return takeover.TakeOverLogInputStates(logger, store, backuper, inputs)
}

// fetches all the defined input configuration available at Filebeat startup including external files.
func fetchInputConfiguration(config *cfg.Config) (inputs []*conf.C, err error) {
	if len(config.Inputs) == 0 {
		inputs = []*conf.C{}
	} else {
		inputs = config.Inputs
	}

	// reading external input configuration if defined
	var dynamicInputCfg cfgfile.DynamicConfig
	if config.ConfigInput != nil {
		err = config.ConfigInput.Unpack(&dynamicInputCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack the dynamic input configuration: %w", err)
		}
	}
	if dynamicInputCfg.Path == "" {
		return inputs, nil
	}

	cfgPaths, err := filepath.Glob(dynamicInputCfg.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve external input configuration paths: %w", err)
	}

	if len(cfgPaths) == 0 {
		return inputs, nil
	}

	// making a copy so we can safely extend the slice
	inputs = make([]*conf.C, len(config.Inputs))
	copy(inputs, config.Inputs)

	for _, p := range cfgPaths {
		externalInputs, err := cfgfile.LoadList(p)
		if err != nil {
			return nil, fmt.Errorf("failed to load external input configuration: %w", err)
		}
		inputs = append(inputs, externalInputs...)
	}

	return inputs, nil
}
