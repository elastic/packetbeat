package add_kubernetes_metadata

import (
	"time"

	"github.com/elastic/beats/libbeat/common"
)

type kubeAnnotatorConfig struct {
	InCluster bool `config:"in_cluster"`
	// If beat deployed as a k8s daemonset, Annotator will only cache localhost meta.
	AsDaemonSet bool          `config:"as_daemonset"`
	KubeConfig  string        `config:"kube_config"`
	Host        string        `config:"host"`
	SyncPeriod  time.Duration `config:"sync_period"`
	// Annotations are kept after pod is removed, until they haven't been accessed
	// for a full `cleanup_timeout`:
	CleanupTimeout     time.Duration `config:"cleanup_timeout"`
	Indexers           PluginConfig  `config:"indexers"`
	Matchers           PluginConfig  `config:"matchers"`
	DefaultMatchers    Enabled       `config:"default_matchers"`
	DefaultIndexers    Enabled       `config:"default_indexers"`
	IncludeLabels      []string      `config:"include_labels"`
	ExcludeLabels      []string      `config:"exclude_labels"`
	IncludeAnnotations []string      `config:"include_annotations"`
}

type Enabled struct {
	Enabled bool `config:"enabled"`
}

type PluginConfig []map[string]common.Config

func defaultKubernetesAnnotatorConfig() kubeAnnotatorConfig {
	return kubeAnnotatorConfig{
		InCluster:       true,
		AsDaemonSet:     true,
		SyncPeriod:      1 * time.Second,
		CleanupTimeout:  60 * time.Second,
		DefaultMatchers: Enabled{true},
		DefaultIndexers: Enabled{true},
	}
}
