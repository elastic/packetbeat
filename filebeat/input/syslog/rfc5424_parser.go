//line parser/rfc5424_parser.rl:1
// Code generated by ragel DO NOT EDIT.
package syslog

//line rfc5424_parser.go:8
const syslog_rfc5424_start int = 1
const syslog_rfc5424_first_final int = 12
const syslog_rfc5424_error int = 0

const syslog_rfc5424_en_main int = 1

//line parser/rfc5424_parser.rl:9

func Parse5424(data []byte, event *event) {
	var p, cs int
	pe := len(data)
	tok := 0
	eof := len(data)

//line rfc5424_parser.go:25
	{
		cs = syslog_rfc5424_start
	}

//line rfc5424_parser.go:30
	{
		if (p) == (pe) {
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
		case 12:
			goto st_case_12
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
		}
		goto st_out
	st_case_1:
		if data[(p)] == 60 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if (p)++; (p) == (pe) {
			goto _test_eof2
		}
	st_case_2:
		switch data[(p)] {
		case 48:
			goto tr2
		case 49:
			goto tr3
		}
		if 50 <= data[(p)] && data[(p)] <= 57 {
			goto tr4
		}
		goto st0
	tr2:
//line parser/common.rl:3

		tok = p

		goto st3
	st3:
		if (p)++; (p) == (pe) {
			goto _test_eof3
		}
	st_case_3:
//line rfc5424_parser.go:99
		if data[(p)] == 62 {
			goto tr5
		}
		goto st0
	tr5:
//line parser/common.rl:7

		event.SetPriority(data[tok:p])

		goto st4
	st4:
		if (p)++; (p) == (pe) {
			goto _test_eof4
		}
	st_case_4:
//line rfc5424_parser.go:115
		if 49 <= data[(p)] && data[(p)] <= 57 {
			goto tr6
		}
		goto st0
	tr6:
//line parser/common.rl:3

		tok = p

		goto st5
	st5:
		if (p)++; (p) == (pe) {
			goto _test_eof5
		}
	st_case_5:
//line rfc5424_parser.go:131
		if data[(p)] == 32 {
			goto tr7
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st7
		}
		goto st0
	tr7:
//line parser/common.rl:82

		event.SetVersion(data[tok:p])

		goto st6
	st6:
		if (p)++; (p) == (pe) {
			goto _test_eof6
		}
	st_case_6:
//line rfc5424_parser.go:150
		goto tr9
	tr9:
//line parser/common.rl:3

		tok = p

		goto st12
	st12:
		if (p)++; (p) == (pe) {
			goto _test_eof12
		}
	st_case_12:
//line rfc5424_parser.go:163
		goto st12
	st7:
		if (p)++; (p) == (pe) {
			goto _test_eof7
		}
	st_case_7:
		if data[(p)] == 32 {
			goto tr7
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if (p)++; (p) == (pe) {
			goto _test_eof8
		}
	st_case_8:
		if data[(p)] == 32 {
			goto tr7
		}
		goto st0
	tr3:
//line parser/common.rl:3

		tok = p

		goto st9
	st9:
		if (p)++; (p) == (pe) {
			goto _test_eof9
		}
	st_case_9:
//line rfc5424_parser.go:197
		switch data[(p)] {
		case 57:
			goto st11
		case 62:
			goto tr5
		}
		if 48 <= data[(p)] && data[(p)] <= 56 {
			goto st10
		}
		goto st0
	tr4:
//line parser/common.rl:3

		tok = p

		goto st10
	st10:
		if (p)++; (p) == (pe) {
			goto _test_eof10
		}
	st_case_10:
//line rfc5424_parser.go:219
		if data[(p)] == 62 {
			goto tr5
		}
		if 48 <= data[(p)] && data[(p)] <= 57 {
			goto st3
		}
		goto st0
	st11:
		if (p)++; (p) == (pe) {
			goto _test_eof11
		}
	st_case_11:
		if data[(p)] == 62 {
			goto tr5
		}
		if 48 <= data[(p)] && data[(p)] <= 49 {
			goto st3
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
	_test_eof12:
		cs = 12
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

	_test_eof:
		{
		}
		if (p) == eof {
			switch cs {
			case 12:
//line parser/common.rl:11

				event.SetMessage(data[tok:p])

//line rfc5424_parser.go:260
			}
		}

	_out:
		{
		}
	}

//line parser/rfc5424_parser.rl:24

}
