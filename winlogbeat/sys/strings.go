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

package sys

import (
	"strings"

	"github.com/elastic/beats/v7/libbeat/common"
	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

var ansiDecoder *encoding.Decoder

func init() {
	ansiCP := windows.GetACP()
	for _, enc := range charmap.All {
		cm, ok := enc.(*charmap.Charmap)
		cmID, _ := cm.ID()
		if !ok || uint32(cmID) != ansiCP {
			continue
		}
		ansiDecoder = cm.NewDecoder()
		return
	}
	ansiDecoder = charmap.Windows1250.NewDecoder()
}

// UTF16BytesToString converts the given UTF-16 bytes to a string.
func UTF16BytesToString(b []byte) (string, error) {
	// Use space from the ByteBuffer pool as working memory for the conversion.
	bb := NewPooledByteBuffer()
	defer bb.Free()

	if err := common.UTF16ToUTF8Bytes(b, bb); err != nil {
		return "", err
	}

	// This copies the UTF-8 bytes to create a string.
	return string(bb.Bytes()), nil
}

// RemoveWindowsLineEndings replaces carriage return line feed (CRLF) with
// line feed (LF) and trims any newline character that may exist at the end
// of the string.
func RemoveWindowsLineEndings(s string) string {
	s = strings.Replace(s, "\r\n", "\n", -1)
	return strings.TrimRight(s, "\n")
}

func ANSIBytesToString(enc []byte) (string, error) {
	out, err := ansiDecoder.Bytes(enc)
	return string(out), err
}

func BinaryToString(bin []byte) string {
	if len(bin) == 0 {
		return ""
	}

	const hexTable = "0123456789ABCDEF"

	size := len(bin) * 2
	buffer := make([]byte, size)

	for i := 0; i < len(bin); i++ {
		for j := 0; j < 2; j++ {
			idx := (bin[i] >> (j * 4) & 0x0F)
			buffer[2*i+(1-j)] = hexTable[idx]
		}
	}

	return string(buffer)
}
