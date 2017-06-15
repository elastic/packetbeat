// Copyright 2017 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"text/template"
)

const includeErrno = `
#include <asm-generic/errno.h>
`

type ErrorNumber struct {
	Name  string
	Value int
}

// TemplateParams is the data used in evaluating the template.
type TemplateParams struct {
	Command   string
	NameToNum []ErrorNumber
	NumToName []ErrorNumber
}

const header = `// go run {{.Command}}.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

// Copyright 2017 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
`

var headerTmpl = template.Must(template.New("header").Parse(header))

const godefsTemplate = `
package auparse

/*
#include <asm-generic/errno.h>
*/
import "C"

// AuditErrnoToNum contains a mapping of POSIX error names to errnos (error numbers).
var AuditErrnoToNum = map[string]int{
{{- range $err := .NameToNum }}
	"{{ $err.Name }}":	C.{{ $err.Name }},
{{- end }}
}

// AuditErrnoToName contains a mapping of errnos (error numbers) to POSIX names.
var AuditErrnoToName = map[int]string{
{{- range $err := .NumToName }}
	{{ $err.Value }}:	"{{ $err.Name }}",
{{- end }}
}
`

var tmpl = template.Must(template.New("message_types").Parse(godefsTemplate))

var (
	errnoDefRegex = regexp.MustCompile(`^#define\s+(E\w+)\s+(\w+)`)
)

func readErrorNumbers() ([]ErrorNumber, error) {
	cmd := exec.Command("gcc", "-E", "-dD", "-")
	cmd.Stdin = bytes.NewBufferString(includeErrno)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	errorToNum := map[string]int{}
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		matches := errnoDefRegex.FindStringSubmatch(s.Text())
		if len(matches) != 3 {
			continue
		}
		errno, err := strconv.Atoi(matches[2])
		if err != nil {
			errorToNum[matches[1]] = -1
			continue
		}

		errorToNum[matches[1]] = errno
	}

	var errnos []ErrorNumber
	for name, value := range errorToNum {
		errnos = append(errnos, ErrorNumber{
			Name:  name,
			Value: value,
		})
	}

	sort.Slice(errnos, func(i, j int) bool {
		return errnos[i].Value < errnos[j].Value
	})

	return errnos, nil
}

func run() error {
	tmp, err := ioutil.TempDir("", "mk_audit_exit_codes")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	if err := os.Chdir(tmp); err != nil {
		return err
	}

	errnos, err := readErrorNumbers()
	if err != nil {
		return err
	}

	// Filter duplicates and sort by name.
	var numToName []ErrorNumber
	for _, errno := range errnos {
		if errno.Value >= 0 {
			numToName = append(numToName, errno)
		}
	}

	// Create output file.
	f, err := os.Create("defs.go")
	if err != nil {
		return err
	}
	defer f.Close()

	// Evaluate template.
	r := TemplateParams{
		Command:   filepath.Base(os.Args[0]),
		NameToNum: errnos,
		NumToName: numToName,
	}
	if err := tmpl.Execute(f, r); err != nil {
		return err
	}

	output, err := exec.Command("go", "tool", "cgo", "-godefs", "defs.go").Output()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = headerTmpl.Execute(buf, r); err != nil {
		return nil
	}

	s := bufio.NewScanner(bytes.NewReader(output))
	for s.Scan() {
		if !bytes.HasPrefix(s.Bytes(), []byte("//")) {
			buf.Write(s.Bytes())
			buf.WriteByte('\n')
		}
	}

	if err = ioutil.WriteFile(flagOut, buf.Bytes(), 0644); err != nil {
		return err
	}

	_, err = exec.Command("gofmt", "-w", "-s", flagOut).Output()
	if err != nil {
		return err
	}

	return nil
}

var flagOut string

func main() {
	flag.StringVar(&flagOut, "out", "zaudit_exit_codes.go", "output file")
	flag.Parse()

	var err error
	flagOut, err = filepath.Abs(flagOut)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
