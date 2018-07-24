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

package fields

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// YmlFile holds the info on files and how to write them into the global fields.yml
type YmlFile struct {
	Path   string
	Indent int
}

func collectCommonFiles(esBeatsPath, beatPath string, fieldFiles []*YmlFile) ([]*YmlFile, error) {
	commonFields := []string{
		filepath.Join(esBeatsPath, "libbeat", "_meta", "fields.common.yml"),
		filepath.Join(beatPath, "_meta", "fields.common.yml"),
		filepath.Join(beatPath, "_meta", "fields.yml"),
	}

	var files []*YmlFile
	for _, cf := range commonFields {
		_, err := os.Stat(cf)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil, err
		}
		files = append(files, &YmlFile{
			Path:   cf,
			Indent: 0,
		})
	}

	libbeatModulesPath := filepath.Join(esBeatsPath, "libbeat", "processors")
	libbeatFieldFiles, err := CollectModuleFiles(libbeatModulesPath)
	if err != nil {
		return nil, err
	}

	files = append(files, libbeatFieldFiles...)

	return append(files, fieldFiles...), nil
}

func writeGeneratedFieldsYml(beatsPath string, fieldFiles []*YmlFile) error {
	outPath := path.Join(beatsPath, "fields.yml")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := GenerateFieldsYml(fieldFiles)
	if err != nil {
		return err
	}
	fw := bufio.NewWriter(f)
	fw.Write(data)
	return fw.Flush()
}

// GenerateFieldsYml generates a fields.yml based on the given files
func GenerateFieldsYml(fieldFiles []*YmlFile) ([]byte, error) {
	buf := bytes.NewBufferString("")
	for _, p := range fieldFiles {
		file, err := os.Open(p.Path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fs := bufio.NewScanner(file)
		for fs.Scan() {
			err = writeIndentedLine(buf, fs.Text()+"\n", p.Indent)
			if err != nil {
				return nil, err
			}
		}
		if err := fs.Err(); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func writeIndentedLine(buf *bytes.Buffer, line string, indent int) error {
	newLine := strings.Repeat(" ", indent) + line
	_, err := buf.WriteString(newLine)
	return err
}

// Generate collects fields.yml files and concatenates them into one global file.
func Generate(esBeatsPath, beatPath string, files []*YmlFile) error {
	files, err := collectCommonFiles(esBeatsPath, beatPath, files)
	if err != nil {
		return err
	}

	return writeGeneratedFieldsYml(beatPath, files)
}
