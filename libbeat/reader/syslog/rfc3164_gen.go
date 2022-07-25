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
	"time"

	"go.uber.org/multierr"
)

const rfc3164_parser_start int = 1
const rfc3164_parser_first_final int = 24
const rfc3164_parser_error int = 0

const rfc3164_parser_en_main int = 1

// parseRFC3164 parses an RFC 3164-formatted syslog message. loc is used to enrich
// timestamps that lack a time zone.
func parseRFC3164(data string, loc *time.Location) (message, error) {
	var errs error
	var p, cs, tok int

	pe := len(data)
	eof := len(data)
	m := message{
		priority: -1,
	}

	{
		cs = rfc3164_parser_start
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
		}
		goto st_out
	st_case_1:
		if data[p] == 60 {
			goto st8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr3
			}
		default:
			goto tr3
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
	tr1:

		tok = p

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 43:
			goto st3
		case 58:
			goto st3
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 65:
				goto st3
			}
		default:
			goto st7
		}
		goto tr0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 32 {
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st2
		}
		goto tr0
	tr6:

		if err := m.setTimestampRFC3339(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st4
	tr30:

		if err := m.setTimestampBSD(data[tok:p], loc); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 33 <= data[p] && data[p] <= 126 {
			goto tr8
		}
		goto tr0
	tr8:

		tok = p

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 32 {
			goto tr9
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st5
		}
		goto tr0
	tr9:

		m.setHostname(data[tok:p])

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 59:
			if 33 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 92 <= data[p] && data[p] <= 126 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto tr11
	tr11:

		tok = p

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		goto st24
	tr12:

		tok = p

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 58:
			goto tr34
		case 91:
			goto tr35
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st25
		}
		goto st24
	tr34:

		m.setTag(data[tok:p])

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 32 {
			goto st27
		}
		goto st24
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		goto tr11
	tr35:

		m.setTag(data[tok:p])

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if 32 <= data[p] && data[p] <= 126 {
			goto tr37
		}
		goto st24
	tr37:

		tok = p

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 93 {
			goto tr39
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto st29
		}
		goto st24
	tr39:

		m.setContent(data[tok:p])

		goto st30
	tr42:

		m.setContent(data[tok:p])

		tok = p

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 58:
			goto st31
		case 93:
			goto tr39
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto st29
		}
		goto st24
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 32:
			goto st32
		case 93:
			goto tr39
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st29
		}
		goto st24
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 93 {
			goto tr42
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto tr37
		}
		goto tr11
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 32:
			goto tr6
		case 43:
			goto st3
		case 58:
			goto st3
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 65:
				goto st3
			}
		default:
			goto st7
		}
		goto tr0
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
		if data[p] == 62 {
			goto tr15
		}
		if 33 <= data[p] && data[p] <= 126 {
			goto st9
		}
		goto tr0
	tr15:

		if err := m.setPriority(data[tok:p]); err != nil {
			errs = multierr.Append(errs, &ValidationError{Err: err, Pos: tok + 1})
		}

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 62 {
			goto tr15
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 48:
				if 33 <= data[p] && data[p] <= 47 {
					goto st9
				}
			case data[p] > 57:
				if 58 <= data[p] && data[p] <= 64 {
					goto st9
				}
			default:
				goto tr16
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st9
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			default:
				goto tr17
			}
		default:
			goto tr17
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
		switch data[p] {
		case 43:
			goto st12
		case 47:
			goto st9
		case 58:
			goto st12
		case 62:
			goto tr15
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 45:
				if 33 <= data[p] && data[p] <= 44 {
					goto st9
				}
			case data[p] > 46:
				if 48 <= data[p] && data[p] <= 57 {
					goto st13
				}
			default:
				goto st12
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st12
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st9
					}
				case data[p] >= 97:
					goto st12
				}
			default:
				goto st9
			}
		default:
			goto st9
		}
		goto tr0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 32:
			goto tr6
		case 62:
			goto tr15
		}
		switch {
		case data[p] < 48:
			if 33 <= data[p] && data[p] <= 47 {
				goto st9
			}
		case data[p] > 57:
			if 58 <= data[p] && data[p] <= 126 {
				goto st9
			}
		default:
			goto st11
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 32:
			goto tr6
		case 43:
			goto st12
		case 47:
			goto st9
		case 58:
			goto st12
		case 62:
			goto tr15
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] < 45:
				if 33 <= data[p] && data[p] <= 44 {
					goto st9
				}
			case data[p] > 46:
				if 48 <= data[p] && data[p] <= 57 {
					goto st13
				}
			default:
				goto st12
			}
		case data[p] > 64:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto st12
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st9
					}
				case data[p] >= 97:
					goto st12
				}
			default:
				goto st9
			}
		default:
			goto st9
		}
		goto tr0
	tr17:

		tok = p

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 32:
			goto st15
		case 62:
			goto tr15
		}
		switch {
		case data[p] < 91:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto st14
				}
			case data[p] >= 33:
				goto st9
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st9
				}
			case data[p] >= 97:
				goto st14
			}
		default:
			goto st9
		}
		goto tr0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 32 {
			goto st15
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 32 {
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto tr0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 32 {
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st18
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 58 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st18
		}
		goto tr0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto tr0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 58 {
			goto st21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st22
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 32 {
			goto tr30
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st22
		}
		goto tr0
	tr3:

		tok = p

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 32 {
			goto st15
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		case data[p] >= 65:
			goto st23
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

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 24, 25, 26, 27, 28, 29, 30, 31, 32:

				m.setMsg(data[tok:p])

			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23:

				errs = multierr.Append(errs, &ParseError{Err: io.ErrUnexpectedEOF, Pos: p + 1})
				p--

			}
		}

	_out:
		{
		}
	}

	return m, errs
}
