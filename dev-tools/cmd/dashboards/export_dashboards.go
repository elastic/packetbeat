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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common/transport/httpcommon"
	"github.com/elastic/beats/v7/libbeat/dashboards"
	"github.com/elastic/beats/v7/libbeat/kibana"
)

var (
	indexPattern = false
	quiet        = false
)

const (
	kibanaTimeout = 90 * time.Second
)

func main() {
	kibanaURL := flag.String("kibana", "http://localhost:5601", "Kibana URL")
	spaceID := flag.String("space-id", "", "Space ID")
	dashboard := flag.String("dashboard", "", "Dashboard ID")
	fileOutput := flag.String("output", "output.ndjson", "Output NDJSON file, when exporting dashboards for Beats, please use -folder instead")
	folderOutput := flag.String("folder", "", "Output folder to save all assets to more human friendly JSON format")
	ymlFile := flag.String("yml", "", "Path to the module.yml file containing the dashboards")
	flag.BoolVar(&indexPattern, "indexPattern", false, "include index-pattern in output")
	flag.BoolVar(&quiet, "quiet", false, "be quiet")

	flag.Parse()
	log.SetFlags(0)

	u, err := url.Parse(*kibanaURL)
	if err != nil {
		log.Fatalf("Error parsing Kibana URL: %v", err)
	}

	var user, pass string
	if u.User != nil {
		user = u.User.Username()
		pass, _ = u.User.Password()
	}

	transport := httpcommon.DefaultHTTPTransportSettings()
	transport.Timeout = kibanaTimeout

	client, err := kibana.NewClientWithConfig(&kibana.ClientConfig{
		Protocol:  u.Scheme,
		Host:      u.Host,
		Username:  user,
		Password:  pass,
		Path:      u.Path,
		SpaceID:   *spaceID,
		Transport: transport,
	})
	if err != nil {
		log.Fatalf("Error while connecting to Kibana: %v", err)
	}

	if len(*ymlFile) == 0 && len(*dashboard) == 0 {
		flag.Usage()
		log.Fatalf("Please specify a dashboard ID (-dashboard) or a manifest file (-yml)")
	}

	if len(*ymlFile) > 0 {
		err = exportDashboardsFromYML(client, *ymlFile)
		if err != nil {
			log.Fatalf("Failed to export dashboards from YML file: %v", err)
		}
		log.Println("Done exporting dashboards from", *ymlFile)
		return
	}

	if len(*dashboard) > 0 {
		err = exportSingleDashboard(client, *dashboard, *folderOutput, *fileOutput)
		if err != nil {
			log.Fatalf("Failed to export the dashboard: %v", err)
		}
		if !quiet {
			if *folderOutput != "" {
				log.Printf("The dashboard %s was exported to '%s'\n", *dashboard, *folderOutput)
			} else {
				log.Printf("The dashboard %s was exported under '%s'\n", *dashboard, *fileOutput)
			}
		}
		return
	}
}

func exportDashboardsFromYML(client *kibana.Client, ymlFile string) error {
	results, info, err := dashboards.ExportAllFromYml(client, ymlFile)
	if err != nil {
		return err
	}
	for i, r := range results {
		log.Printf("id=%s, name=%s\n", info.Dashboards[i].ID, info.Dashboards[i].File)
		r = dashboards.DecodeExported(r)
		err = dashboards.SaveToFile(r, info.Dashboards[i].File, filepath.Dir(ymlFile), client.GetVersion())
		if err != nil {
			return err
		}
	}
	return nil
}

func exportSingleDashboard(client *kibana.Client, dashboard, folder, output string) error {
	result, err := dashboards.Export(client, dashboard)
	if err != nil {
		return fmt.Errorf("failed to export the dashboard: %+v", err)
	}

	if folder != "" {
		return writeAssetsToFolder(client, result, folder)
	} else if output != "" {
		return writeDashboardToFile(result, output)

	}
	return nil
}

func writeDashboardToFile(dashboard []byte, output string) error {
	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return errors.Wrap(err, "failed to create directory for dashboard")
	}

	if err := ioutil.WriteFile(output, dashboard, dashboards.OutputPermission); err != nil {
		return fmt.Errorf("failed to save the dashboard: %+v", err)
	}
	return nil
}

func writeAssetsToFolder(client *kibana.Client, dashboard []byte, folder string) error {
	return dashboards.SaveToFolder(dashboard, folder, client.GetVersion())
}
