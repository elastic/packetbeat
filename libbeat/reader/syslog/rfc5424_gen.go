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

// Code generated by ragel. DO NOT EDIT.
package syslog

import (
	"io"

	"go.uber.org/multierr"
)

const rfc5424_start int = 1
const rfc5424_first_final int = 24
const rfc5424_error int = 0

const rfc5424_en_main int = 1

type machineState struct {
	sdID           string
	sdParamName    string
	sdValueEscapes []int
}

// ParseRFC5424 parses an RFC 5424-formatted syslog message.
func parseRFC5424(data string) (message, error) {
	var errs error
	var p, cs, tok int

	pe := len(data)
	eof := len(data)
	m := message{
		priority: -1,
	}

	{
		cs = rfc5424_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 27:
			goto st_case_27
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		}
		goto st_out
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto tr0
	tr0:

		errs = multierr.Append(errs, &ParseError{Err: io.ErrUnexpectedEOF, Pos: p + 1})
		p--

		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr2
		}
		goto tr0
	tr2:

		tok = p

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 62 {
			goto tr4
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st3
		}
		goto tr0
	tr4:

		if err := m.setPriority(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 62 {
			goto tr6
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr5
		}
		goto tr0
	tr5:

		tok = p

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 32:
			goto tr7
		case 62:
			goto tr9
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st5
		}
		goto tr0
	tr7:

		if err := m.setVersion(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 45 {
			goto st7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr11
		}
		goto tr0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 32 {
			goto st8
		}
		goto tr0
	tr32:

		if err := m.setTimestampRFC3339(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr13
		}
		goto tr0
	tr13:

		tok = p

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 32 {
			goto tr14
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st9
		}
		goto tr0
	tr14:

		m.setHostname(data[tok:p])

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr16
		}
		goto tr0
	tr16:

		tok = p

		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 32 {
			goto tr17
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st11
		}
		goto tr0
	tr17:

		m.setAppName(data[tok:p])

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr19
		}
		goto tr0
	tr19:

		tok = p

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 32 {
			goto tr20
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st13
		}
		goto tr0
	tr20:

		m.setProcID(data[tok:p])

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr22
		}
		goto tr0
	tr22:

		tok = p

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 32 {
			goto tr23
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st15
		}
		goto tr0
	tr23:

		m.setMsgID(data[tok:p])

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 45:
			goto st24
		case 91:
			goto tr26
		}
		goto tr0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 32 {
			goto st25
		}
		goto tr0
	tr37:

		m.setRawSDValue(data[tok:p])

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		goto tr35
	tr35:

		tok = p

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		goto st26
	tr26:

		tok = p

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 92:
			goto st19
		case 93:
			goto tr0
		}
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 92:
			goto st19
		case 93:
			goto st27
		}
		goto st18
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 93 {
			goto st18
		}
		goto tr0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 32:
			goto tr37
		case 91:
			goto st17
		}
		goto tr0
	tr11:

		tok = p

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 43:
			goto st21
		case 58:
			goto st21
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st21
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st21
				}
			case data[p] >= 65:
				goto st21
			}
		default:
			goto st22
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 32 {
			goto tr32
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 32:
			goto tr32
		case 43:
			goto st21
		case 58:
			goto st21
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st21
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st21
				}
			case data[p] >= 65:
				goto st21
			}
		default:
			goto st22
		}
		goto tr0
	tr9:

		if err := m.setPriority(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st23
	tr6:

		if err := m.setPriority(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		tok = p

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 32:
			goto tr7
		case 62:
			goto tr6
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto tr5
		}
		goto tr0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 26:

				m.setMsg(data[tok:p])

			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:

				errs = multierr.Append(errs, &ParseError{Err: io.ErrUnexpectedEOF, Pos: p + 1})
				p--

			case 27:

				m.setRawSDValue(data[tok:p])

			case 25:

				tok = p

				m.setMsg(data[tok:p])

			}
		}

	_out:
		{
		}
	}

	return m, errs
}

const check_start int = 1
const check_first_final int = 10
const check_error int = 0

const check_en_main int = 1

// isRFC5424 returns true if data is formatted as an RFC 5424 syslog message.
func isRFC5424(data string) bool {
	var isRFC5424 bool
	var p, cs int

	pe := len(data)

	{
		cs = check_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		}
		goto st_out
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 62 {
			goto st4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 32 {
			goto st6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
	tr6:

		isRFC5424 = true

		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if 48 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 48 <= data[p] && data[p] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	return isRFC5424
}

const parse_sd_start int = 1
const parse_sd_first_final int = 73
const parse_sd_error int = 0

const parse_sd_en_main int = 1

// parseStructuredData performs a best effort parsing of the raw structured data value
// extracted from the syslog message. If the raw structured data value is an empty
// string or the nil value ('-'), nil is returned. Otherwise, the value is parsed
// using the format defined by RFC 5424. If the value cannot be parsed, then nil
// is returned.
func parseStructuredData(data string) map[string]interface{} {
	var s machineState
	var p, cs, tok int

	pe := len(data)
	structuredData := map[string]interface{}{}

	if data == "" || data == "-" {
		return nil
	}

	{
		cs = parse_sd_start
	}

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 73:
			goto st_case_73
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		}
		goto st_out
	st_case_1:
		if data[p] == 91 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 33 {
			goto tr2
		}
		switch {
		case data[p] < 62:
			if 35 <= data[p] && data[p] <= 60 {
				goto tr2
			}
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
	tr2:

		tok = p

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st42
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st42
			}
		case data[p] >= 35:
			goto st42
		}
		goto st0
	tr3:

		s.sdID = data[tok:p]
		if _, ok := structuredData[s.sdID]; !ok {
			structuredData[s.sdID] = map[string]interface{}{}
		}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 33 {
			goto tr6
		}
		switch {
		case data[p] < 62:
			if 35 <= data[p] && data[p] <= 60 {
				goto tr6
			}
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto st0
	tr6:

		tok = p

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 33:
			goto st6
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st6
			}
		case data[p] >= 35:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 33:
			goto st7
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st7
			}
		case data[p] >= 35:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 33:
			goto st8
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st8
			}
		case data[p] >= 35:
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 33:
			goto st9
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st9
			}
		case data[p] >= 35:
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 33:
			goto st10
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st10
			}
		case data[p] >= 35:
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 33:
			goto st11
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st11
			}
		case data[p] >= 35:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 33:
			goto st12
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st12
			}
		case data[p] >= 35:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 33:
			goto st13
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st13
			}
		case data[p] >= 35:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 33:
			goto st14
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st14
			}
		case data[p] >= 35:
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 33:
			goto st15
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st15
			}
		case data[p] >= 35:
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 33:
			goto st16
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st16
			}
		case data[p] >= 35:
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 33:
			goto st17
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st17
			}
		case data[p] >= 35:
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 33:
			goto st18
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st18
			}
		case data[p] >= 35:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 33:
			goto st19
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st19
			}
		case data[p] >= 35:
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 33:
			goto st20
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st20
			}
		case data[p] >= 35:
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 33:
			goto st21
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st21
			}
		case data[p] >= 35:
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 33:
			goto st22
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st22
			}
		case data[p] >= 35:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 33:
			goto st23
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st23
			}
		case data[p] >= 35:
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 33:
			goto st24
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st24
			}
		case data[p] >= 35:
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 33:
			goto st25
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st25
			}
		case data[p] >= 35:
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 33:
			goto st26
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st26
			}
		case data[p] >= 35:
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 33:
			goto st27
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st27
			}
		case data[p] >= 35:
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 33:
			goto st28
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st28
			}
		case data[p] >= 35:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 33:
			goto st29
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st29
			}
		case data[p] >= 35:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 33:
			goto st30
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st30
			}
		case data[p] >= 35:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 33:
			goto st31
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st31
			}
		case data[p] >= 35:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 33:
			goto st32
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st32
			}
		case data[p] >= 35:
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 33:
			goto st33
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st33
			}
		case data[p] >= 35:
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 33:
			goto st34
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st34
			}
		case data[p] >= 35:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 33:
			goto st35
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st35
			}
		case data[p] >= 35:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 33:
			goto st36
		case 61:
			goto tr8
		}
		switch {
		case data[p] > 92:
			if 94 <= data[p] && data[p] <= 126 {
				goto st36
			}
		case data[p] >= 35:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 61 {
			goto tr8
		}
		goto st0
	tr8:

		s.sdParamName = data[tok:p]

		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 34 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 34:
			goto st0
		case 92:
			goto tr41
		case 93:
			goto st0
		}
		goto tr40
	tr40:

		tok = p

		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 34:
			goto tr43
		case 92:
			goto tr44
		case 93:
			goto st0
		}
		goto st39
	tr43:

		if subMap, ok := structuredData[s.sdID].(map[string]interface{}); ok {
			subMap[s.sdParamName] = removeBytes(data[tok:p], s.sdValueEscapes, tok)
		}

		s.sdValueEscapes = nil

		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 32:
			goto st4
		case 93:
			goto st73
		}
		goto st0
	tr5:

		s.sdID = data[tok:p]
		if _, ok := structuredData[s.sdID]; !ok {
			structuredData[s.sdID] = map[string]interface{}{}
		}

		goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 91 {
			goto st2
		}
		goto st0
	tr41:

		tok = p

		s.sdValueEscapes = append(s.sdValueEscapes, p)

		goto st41
	tr44:

		s.sdValueEscapes = append(s.sdValueEscapes, p)

		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 34 {
			goto st39
		}
		if 92 <= data[p] && data[p] <= 93 {
			goto st39
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st43
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st43
			}
		case data[p] >= 35:
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st44
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st44
			}
		case data[p] >= 35:
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st45
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st45
			}
		case data[p] >= 35:
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st46
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st46
			}
		case data[p] >= 35:
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st47
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st47
			}
		case data[p] >= 35:
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st48
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st48
			}
		case data[p] >= 35:
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st49
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st49
			}
		case data[p] >= 35:
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st50
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st50
			}
		case data[p] >= 35:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st51
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st51
			}
		case data[p] >= 35:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st52
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st52
			}
		case data[p] >= 35:
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st53
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st53
			}
		case data[p] >= 35:
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st54
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st54
			}
		case data[p] >= 35:
			goto st54
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st55
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st55
			}
		case data[p] >= 35:
			goto st55
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st56
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st56
			}
		case data[p] >= 35:
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st57
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st57
			}
		case data[p] >= 35:
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st58
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st58
			}
		case data[p] >= 35:
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st59
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st59
			}
		case data[p] >= 35:
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st60
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st60
			}
		case data[p] >= 35:
			goto st60
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st61
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st61
			}
		case data[p] >= 35:
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st62
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st62
			}
		case data[p] >= 35:
			goto st62
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st63
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st63
			}
		case data[p] >= 35:
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st64
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st64
			}
		case data[p] >= 35:
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st65
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st65
			}
		case data[p] >= 35:
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st66
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st66
			}
		case data[p] >= 35:
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st67
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st67
			}
		case data[p] >= 35:
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st68
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st68
			}
		case data[p] >= 35:
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st69
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st69
			}
		case data[p] >= 35:
			goto st69
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st70
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st70
			}
		case data[p] >= 35:
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st71
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st71
			}
		case data[p] >= 35:
			goto st71
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st72
		case 93:
			goto tr5
		}
		switch {
		case data[p] > 60:
			if 62 <= data[p] && data[p] <= 126 {
				goto st72
			}
		case data[p] >= 35:
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 32:
			goto tr3
		case 93:
			goto tr5
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	if len(structuredData) == 0 {
		return nil
	}

	return structuredData
}
