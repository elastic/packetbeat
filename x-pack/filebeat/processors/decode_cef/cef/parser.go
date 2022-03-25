// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//line parser.rl:1
// Code generated by ragel DO NOT EDIT.

package cef

import (
	"fmt"
	"strconv"

	"go.uber.org/multierr"
)

//line parser.go:16
var _cef_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 25,
	25, 0, 0, 0, 28, 30, 34, 34,
	34, 30, 34, 34, 34, 37, 37, 0,
}

const cef_start int = 1
const cef_first_final int = 36
const cef_error int = 0

const cef_en_gobble_extension int = 33
const cef_en_main int = 1
const cef_en_main_cef_extensions int = 29

//line parser.rl:17

// unpack unpacks a CEF message.
func (e *Event) unpack(data string) error {
	cs, p, pe, eof := 0, 0, len(data), len(data)
	mark, mark_slash := 0, 0

	// state related to CEF values.
	var state cefState

	// flag for completion of CEF header.
	complete := false

	// recoveredErrs are problems with the message that the parser was able to
	// recover from (though the parsing might not be "correct").
	var recoveredErrs []error

	e.init(data)

//line parser.go:56
	{
		cs = cef_start
	}

//line parser.go:61
	{
		if (p) == (pe) {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		switch cs {
		case 1:
			if data[(p)] == 67 {
				goto tr0
			}
			goto tr1
		case 0:
			goto _out
		case 2:
			if data[(p)] == 69 {
				goto tr2
			}
			goto tr1
		case 3:
			if data[(p)] == 70 {
				goto tr3
			}
			goto tr1
		case 4:
			if data[(p)] == 58 {
				goto tr4
			}
			goto tr1
		case 5:
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto tr5
			}
			goto tr1
		case 6:
			if data[(p)] == 124 {
				goto tr7
			}
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto tr6
			}
			goto tr1
		case 7:
			switch data[(p)] {
			case 92:
				goto tr9
			case 124:
				goto tr10
			}
			goto tr8
		case 8:
			switch data[(p)] {
			case 92:
				goto tr12
			case 124:
				goto tr13
			}
			goto tr11
		case 9:
			switch data[(p)] {
			case 92:
				goto tr14
			case 124:
				goto tr14
			}
			goto tr1
		case 10:
			switch data[(p)] {
			case 92:
				goto tr16
			case 124:
				goto tr17
			}
			goto tr15
		case 11:
			switch data[(p)] {
			case 92:
				goto tr19
			case 124:
				goto tr20
			}
			goto tr18
		case 12:
			switch data[(p)] {
			case 92:
				goto tr22
			case 124:
				goto tr23
			}
			goto tr21
		case 13:
			switch data[(p)] {
			case 92:
				goto tr24
			case 124:
				goto tr24
			}
			goto tr1
		case 14:
			switch data[(p)] {
			case 92:
				goto tr26
			case 124:
				goto tr27
			}
			goto tr25
		case 15:
			switch data[(p)] {
			case 92:
				goto tr29
			case 124:
				goto tr30
			}
			goto tr28
		case 16:
			switch data[(p)] {
			case 92:
				goto tr32
			case 124:
				goto tr33
			}
			goto tr31
		case 17:
			switch data[(p)] {
			case 92:
				goto tr34
			case 124:
				goto tr34
			}
			goto tr1
		case 18:
			switch data[(p)] {
			case 92:
				goto tr36
			case 124:
				goto tr37
			}
			goto tr35
		case 19:
			switch data[(p)] {
			case 92:
				goto tr39
			case 124:
				goto tr40
			}
			goto tr38
		case 20:
			switch data[(p)] {
			case 92:
				goto tr42
			case 124:
				goto tr43
			}
			goto tr41
		case 21:
			switch data[(p)] {
			case 92:
				goto tr44
			case 124:
				goto tr44
			}
			goto tr1
		case 22:
			switch data[(p)] {
			case 92:
				goto tr46
			case 124:
				goto tr47
			}
			goto tr45
		case 23:
			switch data[(p)] {
			case 92:
				goto tr49
			case 124:
				goto tr50
			}
			goto tr48
		case 24:
			switch data[(p)] {
			case 92:
				goto tr52
			case 124:
				goto tr53
			}
			goto tr51
		case 25:
			switch data[(p)] {
			case 92:
				goto tr54
			case 124:
				goto tr54
			}
			goto tr1
		case 26:
			switch data[(p)] {
			case 92:
				goto tr56
			case 124:
				goto tr57
			}
			goto tr55
		case 27:
			switch data[(p)] {
			case 45:
				goto tr58
			case 124:
				goto tr59
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr58
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr58
				}
			default:
				goto tr58
			}
			goto tr1
		case 28:
			switch data[(p)] {
			case 45:
				goto tr60
			case 124:
				goto tr61
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr60
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr60
				}
			default:
				goto tr60
			}
			goto tr1
		case 36:
			switch data[(p)] {
			case 32:
				goto tr73
			case 95:
				goto tr74
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr74
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr74
				}
			default:
				goto tr74
			}
			goto tr1
		case 29:
			switch data[(p)] {
			case 32:
				goto tr62
			case 95:
				goto tr63
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr63
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr63
				}
			default:
				goto tr63
			}
			goto tr1
		case 30:
			switch data[(p)] {
			case 44:
				goto tr64
			case 46:
				goto tr64
			case 61:
				goto tr65
			case 93:
				goto tr64
			case 95:
				goto tr64
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr64
				}
			case data[(p)] > 91:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr64
				}
			default:
				goto tr64
			}
			goto tr1
		case 37:
			switch data[(p)] {
			case 32:
				goto tr77
			case 61:
				goto tr66
			case 92:
				goto tr78
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr76
			}
			goto tr75
		case 38:
			switch data[(p)] {
			case 32:
				goto tr81
			case 61:
				goto tr66
			case 92:
				goto tr82
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr80
			}
			goto tr79
		case 39:
			switch data[(p)] {
			case 32:
				goto tr81
			case 61:
				goto tr66
			case 92:
				goto tr82
			case 95:
				goto tr83
			}
			switch {
			case data[(p)] < 48:
				if 9 <= data[(p)] && data[(p)] <= 13 {
					goto tr80
				}
			case data[(p)] > 57:
				switch {
				case data[(p)] > 90:
					if 97 <= data[(p)] && data[(p)] <= 122 {
						goto tr83
					}
				case data[(p)] >= 65:
					goto tr83
				}
			default:
				goto tr83
			}
			goto tr79
		case 40:
			switch data[(p)] {
			case 32:
				goto tr81
			case 44:
				goto tr84
			case 46:
				goto tr84
			case 61:
				goto tr85
			case 92:
				goto tr82
			case 95:
				goto tr84
			}
			switch {
			case data[(p)] < 48:
				if 9 <= data[(p)] && data[(p)] <= 13 {
					goto tr80
				}
			case data[(p)] > 57:
				switch {
				case data[(p)] > 93:
					if 97 <= data[(p)] && data[(p)] <= 122 {
						goto tr84
					}
				case data[(p)] >= 65:
					goto tr84
				}
			default:
				goto tr84
			}
			goto tr79
		case 41:
			switch data[(p)] {
			case 32:
				goto tr88
			case 61:
				goto tr66
			case 92:
				goto tr89
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr87
			}
			goto tr86
		case 42:
			switch data[(p)] {
			case 32:
				goto tr92
			case 61:
				goto tr66
			case 92:
				goto tr93
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr91
			}
			goto tr90
		case 43:
			switch data[(p)] {
			case 32:
				goto tr92
			case 61:
				goto tr66
			case 92:
				goto tr93
			case 95:
				goto tr94
			}
			switch {
			case data[(p)] < 48:
				if 9 <= data[(p)] && data[(p)] <= 13 {
					goto tr91
				}
			case data[(p)] > 57:
				switch {
				case data[(p)] > 90:
					if 97 <= data[(p)] && data[(p)] <= 122 {
						goto tr94
					}
				case data[(p)] >= 65:
					goto tr94
				}
			default:
				goto tr94
			}
			goto tr90
		case 44:
			switch data[(p)] {
			case 32:
				goto tr92
			case 44:
				goto tr95
			case 46:
				goto tr95
			case 61:
				goto tr85
			case 92:
				goto tr93
			case 95:
				goto tr95
			}
			switch {
			case data[(p)] < 48:
				if 9 <= data[(p)] && data[(p)] <= 13 {
					goto tr91
				}
			case data[(p)] > 57:
				switch {
				case data[(p)] > 93:
					if 97 <= data[(p)] && data[(p)] <= 122 {
						goto tr95
					}
				case data[(p)] >= 65:
					goto tr95
				}
			default:
				goto tr95
			}
			goto tr90
		case 31:
			switch data[(p)] {
			case 61:
				goto tr67
			case 92:
				goto tr67
			case 110:
				goto tr67
			case 114:
				goto tr67
			}
			goto tr66
		case 45:
			switch data[(p)] {
			case 32:
				goto tr98
			case 61:
				goto tr66
			case 92:
				goto tr99
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr97
			}
			goto tr96
		case 32:
			switch data[(p)] {
			case 61:
				goto tr68
			case 92:
				goto tr68
			case 110:
				goto tr68
			case 114:
				goto tr68
			}
			goto tr66
		case 46:
			switch data[(p)] {
			case 32:
				goto tr102
			case 61:
				goto tr66
			case 92:
				goto tr103
			}
			if 9 <= data[(p)] && data[(p)] <= 13 {
				goto tr101
			}
			goto tr100
		case 33:
			if data[(p)] == 32 {
				goto tr70
			}
			goto tr69
		case 34:
			switch data[(p)] {
			case 32:
				goto tr70
			case 95:
				goto tr71
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr71
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr71
				}
			default:
				goto tr71
			}
			goto tr69
		case 35:
			switch data[(p)] {
			case 32:
				goto tr70
			case 44:
				goto tr71
			case 46:
				goto tr71
			case 61:
				goto tr72
			case 93:
				goto tr71
			case 95:
				goto tr71
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr71
				}
			case data[(p)] > 91:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr71
				}
			default:
				goto tr71
			}
			goto tr69
		case 47:
			if data[(p)] == 32 {
				goto tr70
			}
			goto tr69
		}

	tr1:
		cs = 0
		goto _again
	tr66:
		cs = 0
		goto f24
	tr0:
		cs = 2
		goto _again
	tr2:
		cs = 3
		goto _again
	tr3:
		cs = 4
		goto _again
	tr4:
		cs = 5
		goto _again
	tr6:
		cs = 6
		goto _again
	tr5:
		cs = 6
		goto f0
	tr7:
		cs = 7
		goto f1
	tr11:
		cs = 8
		goto _again
	tr8:
		cs = 8
		goto f0
	tr15:
		cs = 8
		goto f6
	tr9:
		cs = 9
		goto f2
	tr12:
		cs = 9
		goto f4
	tr16:
		cs = 9
		goto f7
	tr14:
		cs = 10
		goto _again
	tr10:
		cs = 11
		goto f3
	tr13:
		cs = 11
		goto f5
	tr17:
		cs = 11
		goto f8
	tr21:
		cs = 12
		goto _again
	tr18:
		cs = 12
		goto f0
	tr25:
		cs = 12
		goto f6
	tr19:
		cs = 13
		goto f2
	tr22:
		cs = 13
		goto f4
	tr26:
		cs = 13
		goto f7
	tr24:
		cs = 14
		goto _again
	tr20:
		cs = 15
		goto f9
	tr23:
		cs = 15
		goto f10
	tr27:
		cs = 15
		goto f11
	tr31:
		cs = 16
		goto _again
	tr28:
		cs = 16
		goto f0
	tr35:
		cs = 16
		goto f6
	tr29:
		cs = 17
		goto f2
	tr32:
		cs = 17
		goto f4
	tr36:
		cs = 17
		goto f7
	tr34:
		cs = 18
		goto _again
	tr30:
		cs = 19
		goto f12
	tr33:
		cs = 19
		goto f13
	tr37:
		cs = 19
		goto f14
	tr41:
		cs = 20
		goto _again
	tr38:
		cs = 20
		goto f0
	tr45:
		cs = 20
		goto f6
	tr39:
		cs = 21
		goto f2
	tr42:
		cs = 21
		goto f4
	tr46:
		cs = 21
		goto f7
	tr44:
		cs = 22
		goto _again
	tr40:
		cs = 23
		goto f15
	tr43:
		cs = 23
		goto f16
	tr47:
		cs = 23
		goto f17
	tr51:
		cs = 24
		goto _again
	tr48:
		cs = 24
		goto f0
	tr55:
		cs = 24
		goto f6
	tr49:
		cs = 25
		goto f2
	tr52:
		cs = 25
		goto f4
	tr56:
		cs = 25
		goto f7
	tr54:
		cs = 26
		goto _again
	tr50:
		cs = 27
		goto f18
	tr53:
		cs = 27
		goto f19
	tr57:
		cs = 27
		goto f20
	tr60:
		cs = 28
		goto _again
	tr58:
		cs = 28
		goto f0
	tr62:
		cs = 29
		goto _again
	tr73:
		cs = 29
		goto f27
	tr64:
		cs = 30
		goto _again
	tr63:
		cs = 30
		goto f0
	tr74:
		cs = 30
		goto f28
	tr93:
		cs = 31
		goto f4
	tr99:
		cs = 31
		goto f7
	tr89:
		cs = 31
		goto f32
	tr82:
		cs = 32
		goto f4
	tr103:
		cs = 32
		goto f7
	tr78:
		cs = 32
		goto f32
	tr69:
		cs = 33
		goto _again
	tr70:
		cs = 34
		goto f0
	tr71:
		cs = 35
		goto _again
	tr59:
		cs = 36
		goto f21
	tr61:
		cs = 36
		goto f22
	tr65:
		cs = 37
		goto f23
	tr80:
		cs = 38
		goto _again
	tr101:
		cs = 38
		goto f6
	tr79:
		cs = 38
		goto f25
	tr75:
		cs = 38
		goto f30
	tr76:
		cs = 38
		goto f31
	tr100:
		cs = 38
		goto f37
	tr81:
		cs = 39
		goto _again
	tr102:
		cs = 39
		goto f6
	tr77:
		cs = 39
		goto f31
	tr84:
		cs = 40
		goto f25
	tr83:
		cs = 40
		goto f34
	tr85:
		cs = 41
		goto f23
	tr91:
		cs = 42
		goto _again
	tr97:
		cs = 42
		goto f6
	tr90:
		cs = 42
		goto f25
	tr86:
		cs = 42
		goto f30
	tr87:
		cs = 42
		goto f31
	tr96:
		cs = 42
		goto f37
	tr92:
		cs = 43
		goto _again
	tr98:
		cs = 43
		goto f6
	tr88:
		cs = 43
		goto f31
	tr95:
		cs = 44
		goto f25
	tr94:
		cs = 44
		goto f35
	tr67:
		cs = 45
		goto f25
	tr68:
		cs = 46
		goto f25
	tr72:
		cs = 47
		goto f26

	f0:
//line cef_actions.rl:9

		mark = p

		goto _again
	f4:
//line cef_actions.rl:12

		mark_slash = p

		goto _again
	f6:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

		goto _again
	f1:
//line cef_actions.rl:18

		e.Version, _ = strconv.Atoi(data[mark:p])

		goto _again
	f5:
//line cef_actions.rl:21

		e.DeviceVendor = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f10:
//line cef_actions.rl:25

		e.DeviceProduct = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f13:
//line cef_actions.rl:29

		e.DeviceVersion = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f16:
//line cef_actions.rl:33

		e.DeviceEventClassID = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f19:
//line cef_actions.rl:37

		e.Name = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f22:
//line cef_actions.rl:41

		e.Severity = data[mark:p]

		goto _again
	f27:
//line cef_actions.rl:44

		complete = true

		goto _again
	f23:
//line cef_actions.rl:51

		// A new extension key marks the end of the last extension value.
		if len(state.key) != 0 && state.valueStart < mark {
			// We should not be here, but purge the escapes and handle them.
			e.pushExtension(state.key, replaceEscapes(data[state.valueStart:mark-1], state.valueStart, state.escapes))
			state.reset()
		}
		state.key = data[mark:p]

		goto _again
	f31:
//line cef_actions.rl:60

		if len(state.escapes) != 0 {
			e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
			state.reset()
		}
		state.valueStart = p
		state.valueEnd = p

		goto _again
	f25:
//line cef_actions.rl:68

		state.valueEnd = p + 1

		goto _again
	f24:
//line cef_actions.rl:78

		recoveredErrs = append(recoveredErrs, fmt.Errorf("malformed value for %s at pos %d", state.key, p+1))
		(p)--
		cs = 33

		goto _again
	f26:
//line cef_actions.rl:82

		state.reset()
		// Resume processing at p, the start of the next extension key.
		p = mark
		cs = 29

		goto _again
	f2:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:12

		mark_slash = p

		goto _again
	f3:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:21

		e.DeviceVendor = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f9:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:25

		e.DeviceProduct = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f12:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:29

		e.DeviceVersion = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f15:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:33

		e.DeviceEventClassID = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f18:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:37

		e.Name = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f21:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:41

		e.Severity = data[mark:p]

		goto _again
	f35:
//line cef_actions.rl:9

		mark = p

//line cef_actions.rl:68

		state.valueEnd = p + 1

		goto _again
	f7:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:12

		mark_slash = p

		goto _again
	f8:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:21

		e.DeviceVendor = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f11:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:25

		e.DeviceProduct = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f14:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:29

		e.DeviceVersion = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f17:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:33

		e.DeviceEventClassID = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f20:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:37

		e.Name = replaceEscapes(data[mark:p], mark, state.escapes)
		state.reset()

		goto _again
	f37:
//line cef_actions.rl:15

		state.pushEscape(mark_slash, p)

//line cef_actions.rl:68

		state.valueEnd = p + 1

		goto _again
	f28:
//line cef_actions.rl:44

		complete = true

//line cef_actions.rl:9

		mark = p

		goto _again
	f32:
//line cef_actions.rl:60

		if len(state.escapes) != 0 {
			e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
			state.reset()
		}
		state.valueStart = p
		state.valueEnd = p

//line cef_actions.rl:12

		mark_slash = p

		goto _again
	f30:
//line cef_actions.rl:60

		if len(state.escapes) != 0 {
			e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
			state.reset()
		}
		state.valueStart = p
		state.valueEnd = p

//line cef_actions.rl:68

		state.valueEnd = p + 1

		goto _again
	f34:
//line cef_actions.rl:68

		state.valueEnd = p + 1

//line cef_actions.rl:9

		mark = p

		goto _again

	_again:
		if cs == 0 {
			goto _out
		}
		if (p)++; (p) != (pe) {
			goto _resume
		}
	_test_eof:
		{
		}
		if (p) == eof {
			switch _cef_eof_actions[cs] {
			case 28:
//line cef_actions.rl:44

				complete = true

			case 34:
//line cef_actions.rl:71

				// Reaching the EOF marks the end of the final extension value.
				if len(state.key) != 0 && state.valueStart < state.valueEnd {
					e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
					state.reset()
				}

			case 25:
//line cef_actions.rl:78

				recoveredErrs = append(recoveredErrs, fmt.Errorf("malformed value for %s at pos %d", state.key, p+1))
				(p)--
				cs = 33

			case 37:
//line cef_actions.rl:15

				state.pushEscape(mark_slash, p)

//line cef_actions.rl:71

				// Reaching the EOF marks the end of the final extension value.
				if len(state.key) != 0 && state.valueStart < state.valueEnd {
					e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
					state.reset()
				}

			case 30:
//line cef_actions.rl:60

				if len(state.escapes) != 0 {
					e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
					state.reset()
				}
				state.valueStart = p
				state.valueEnd = p

//line cef_actions.rl:71

				// Reaching the EOF marks the end of the final extension value.
				if len(state.key) != 0 && state.valueStart < state.valueEnd {
					e.pushExtension(state.key, replaceEscapes(data[state.valueStart:state.valueEnd], state.valueStart, state.escapes))
					state.reset()
				}

//line parser.go:1156
			}
		}

	_out:
		{
		}
	}

//line parser.rl:54

	// Check if state machine completed.
	if cs < cef_first_final {
		// Reached an early end.
		if p == pe {
			if complete {
				return multierr.Append(multierr.Combine(recoveredErrs...), errUnexpectedEndOfEvent)
			}
			return multierr.Append(multierr.Combine(recoveredErrs...), multierr.Combine(errUnexpectedEndOfEvent, errIncompleteHeader))
		}

		// Encountered invalid input.
		return multierr.Append(multierr.Combine(recoveredErrs...), fmt.Errorf("error in CEF event at pos %d", p+1))
	}

	return multierr.Combine(recoveredErrs...)
}
