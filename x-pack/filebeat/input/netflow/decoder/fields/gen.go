// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build ignore

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	outputFile = flag.String("output", "zfields.go", "Output file")
	export     = flag.String("export", "fields", "Name used to export this fields")
	nameCol    = flag.Int("column-name", 0, "Index of column with field name")
	penCol     = flag.Int("column-pen", 0, "Index of column with PEN ID")
	idCol      = flag.Int("column-id", 0, "Index of column with field ID")
	typeCol    = flag.Int("column-type", 0, "Index of column with field type")
)

const fileHeader = `// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// go run gen.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT.

package fields

`

var TypeNames = []string{
	"OctetArray",
	"Unsigned8",
	"Unsigned16",
	"Unsigned32",
	"Unsigned64",
	"Signed8",
	"Signed16",
	"Signed32",
	"Signed64",
	"Float32",
	"Float64",
	"Boolean",
	"MacAddress",
	"String",
	"DateTimeSeconds",
	"DateTimeMilliseconds",
	"DateTimeMicroseconds",
	"DateTimeNanoseconds",
	"Ipv4Address",
	"Ipv6Address",
	"BasicList",
	"SubTemplateList",
	"SubTemplateMultiList",
}

func write(w io.Writer, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Fprintf(os.Stderr, "Failed writing to %s: %v\n", *outputFile, err)
		os.Exit(4)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: gen [-output file.go] [-export name] [--column-{name|pen|id|type}=N]* <input.csv>\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func requireColumn(colFlag *int, argument string) {
	if *colFlag <= 0 {
		fmt.Fprintf(os.Stderr, "Required argument %s not provided\n", argument)
		usage()
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "No CSV file to parse provided\n")
		usage()
	}
	csvFile := flag.Args()[0]
	if len(csvFile) == 0 {
		fmt.Fprintf(os.Stderr, "Argument -input is required\n")
		os.Exit(2)
	}

	requireColumn(nameCol, "--column-name")
	requireColumn(idCol, "--column-id")
	requireColumn(typeCol, "--column-type")

	fHandle, err := os.Open(csvFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s: %v\n", csvFile, err)
		os.Exit(2)
	}
	defer fHandle.Close()

	outHandle, err := os.Create(*outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create %s: %v\n", *outputFile, err)
		os.Exit(3)
	}
	defer outHandle.Close()

	write(outHandle, fileHeader)
	write(outHandle, fmt.Sprintf("var %s = FieldDict {\n", *export))

	typeMap := make(map[string]string)
	for _, n := range TypeNames {
		typeMap[strings.ToLower(n)] = n
	}

	filtered := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(fHandle)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 || scanner.Bytes()[0] != ';' {
			filtered.Write(scanner.Bytes())
			filtered.WriteByte('\n')
		}
	}
	reader := csv.NewReader(filtered)
	for lineNum := 1; ; lineNum++ {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "read of %s failed: %v\n", csvFile, err)
			os.Exit(5)
		}
		n := len(record)
		vars := make(map[string]string)
		for _, f := range []struct {
			column int
			name   string
		}{
			{*idCol, "id"},
			{*penCol, "pen"},
			{*nameCol, "name"},
			{*typeCol, "type"},
		} {
			if f.column > 0 {
				if f.column > n {
					fmt.Fprintf(os.Stderr, "%s column is out of range in line %d\n", f.name, lineNum)
					os.Exit(6)
				}
				vars[f.name] = record[f.column-1]
			} else {
				vars[f.name] = "0"
			}
		}
		if len(vars["type"]) == 0 {
			write(outHandle, fmt.Sprintf("\t// Field %s: %s\n", vars["id"], vars["name"]))
			continue
		}
		ttype, found := typeMap[strings.ToLower(vars["type"])]
		if !found {
			fmt.Fprintf(os.Stderr, "unknown type %s in line %d\n", vars["type"], lineNum)
			os.Exit(7)
		}
		write(outHandle, fmt.Sprintf("\tKey{EnterpriseID: %s, FieldID: %s}: {Name: \"%s\", Decoder: %s},\n",
			vars["pen"], vars["id"], vars["name"], ttype))
	}
	write(outHandle, fmt.Sprintf(`}

func init() {
  if err := RegisterFields(%s); err != nil {
    panic(err)
  }
}
`, *export))
}
