
//line parser.rl:1
// Code generated by ragel DO NOT EDIT.
package syslog


//line parser.go:8
const syslog_start int = 0
const syslog_first_final int = 1
const syslog_error int = -1

const syslog_en_main int = 0


//line parser.rl:9


// syslog
//<34>Oct 11 22:14:15 wopr su: 'su root' failed for foobar
//<13>Feb  5 17:32:18 10.0.0.99 Use the quad dmg.
func Parse(data []byte, event *event) {
    var p, cs int
    pe := len(data)
    tok := 0
    eof := len(data)
    
//line parser.go:28
	{
	cs = syslog_start
	}

//line parser.go:33
	{
	if ( p) == ( pe) {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
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
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	}
	goto st_out
	st_case_0:
		switch data[( p)] {
		case 60:
			goto tr2
		case 65:
			goto tr3
		case 70:
			goto tr4
		case 74:
			goto tr5
		case 77:
			goto tr6
		case 78:
			goto tr7
		case 79:
			goto tr8
		case 83:
			goto tr9
		case 101:
			goto tr10
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr1
		}
		goto tr0
tr0:
//line parser.rl:20

        tok = p
      
	goto st1
	st1:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:289
		goto st1
tr1:
//line parser.rl:20

        tok = p
      
	goto st2
	st2:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof2
		}
	st_case_2:
//line parser.go:302
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st3
		}
		goto st1
	st3:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st4
		}
		goto st1
	st4:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st5
		}
		goto st1
	st5:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof5
		}
	st_case_5:
		if data[( p)] == 45 {
			goto tr15
		}
		goto st1
tr15:
//line parser.rl:36

        event.SetYear(data[tok:p])
      
	goto st6
	st6:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:345
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr16
		}
		goto st1
tr16:
//line parser.rl:20

        tok = p
      
	goto st7
	st7:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:361
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st8
		}
		goto st1
	st8:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof8
		}
	st_case_8:
		if data[( p)] == 45 {
			goto tr18
		}
		goto st1
tr18:
//line parser.rl:40

        event.SetMonthNumeric(data[tok:p])
      
	goto st9
	st9:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof9
		}
	st_case_9:
//line parser.go:386
		if 48 <= data[( p)] && data[( p)] <= 51 {
			goto tr19
		}
		goto st1
tr19:
//line parser.rl:20

        tok = p
      
	goto st10
	st10:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:402
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st11
		}
		goto st1
	st11:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof11
		}
	st_case_11:
		switch data[( p)] {
		case 32:
			goto tr21
		case 84:
			goto tr21
		case 116:
			goto tr21
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr21
		}
		goto st1
tr21:
//line parser.rl:44

        event.SetDay(data[tok:p])
      
	goto st12
	st12:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:435
		if data[( p)] == 50 {
			goto tr23
		}
		if 48 <= data[( p)] && data[( p)] <= 49 {
			goto tr22
		}
		goto st1
tr22:
//line parser.rl:20

        tok = p
      
	goto st13
	st13:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof13
		}
	st_case_13:
//line parser.go:454
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st14
		}
		goto st1
	st14:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof14
		}
	st_case_14:
		if data[( p)] == 58 {
			goto tr25
		}
		goto st1
tr25:
//line parser.rl:48

        event.SetHour(data[tok:p])
      
	goto st15
	st15:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:479
		if 48 <= data[( p)] && data[( p)] <= 53 {
			goto tr26
		}
		goto st1
tr26:
//line parser.rl:20

        tok = p
      
	goto st16
	st16:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:495
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st17
		}
		goto st1
	st17:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof17
		}
	st_case_17:
		if data[( p)] == 58 {
			goto tr28
		}
		goto st1
tr28:
//line parser.rl:52

        event.SetMinute(data[tok:p])
      
	goto st18
	st18:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:520
		if 48 <= data[( p)] && data[( p)] <= 53 {
			goto tr29
		}
		goto st1
tr29:
//line parser.rl:20

        tok = p
      
	goto st19
	st19:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:536
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st20
		}
		goto st1
	st20:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof20
		}
	st_case_20:
		switch data[( p)] {
		case 32:
			goto tr31
		case 43:
			goto tr32
		case 45:
			goto tr32
		case 46:
			goto tr33
		case 90:
			goto tr34
		case 122:
			goto tr34
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr31
		}
		goto st1
tr31:
//line parser.rl:56

        event.SetSecond(data[tok:p])
      
	goto st21
tr49:
//line parser.rl:76

        event.SetTimeZone(data[tok:p])
      
	goto st21
tr54:
//line parser.rl:60

        event.SetNanosecond(data[tok:p])
      
	goto st21
	st21:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:587
		switch {
		case data[( p)] > 95:
			if 97 <= data[( p)] && data[( p)] <= 122 {
				goto tr35
			}
		case data[( p)] >= 46:
			goto tr35
		}
		goto tr0
tr35:
//line parser.rl:20

        tok = p
      
	goto st22
	st22:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof22
		}
	st_case_22:
//line parser.go:608
		if data[( p)] == 32 {
			goto tr36
		}
		switch {
		case data[( p)] < 46:
			if 9 <= data[( p)] && data[( p)] <= 13 {
				goto tr36
			}
		case data[( p)] > 95:
			if 97 <= data[( p)] && data[( p)] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto st1
tr36:
//line parser.rl:64

        event.SetHostname(data[tok:p])
      
	goto st23
	st23:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:636
		switch data[( p)] {
		case 32:
			goto tr0
		case 91:
			goto tr0
		case 93:
			goto tr0
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr0
		}
		goto tr38
tr38:
//line parser.rl:20

        tok = p
      
	goto st24
	st24:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof24
		}
	st_case_24:
//line parser.go:660
		switch data[( p)] {
		case 32:
			goto st1
		case 58:
			goto tr40
		case 91:
			goto tr41
		case 93:
			goto st1
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto st1
		}
		goto st24
tr40:
//line parser.rl:68

        event.SetProgram(data[tok:p])
      
	goto st25
	st25:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:686
		switch data[( p)] {
		case 32:
			goto st26
		case 58:
			goto tr40
		case 91:
			goto tr41
		case 93:
			goto st1
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto st26
		}
		goto st24
	st26:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof26
		}
	st_case_26:
		goto tr0
tr41:
//line parser.rl:68

        event.SetProgram(data[tok:p])
      
	goto st27
	st27:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof27
		}
	st_case_27:
//line parser.go:718
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr43
		}
		goto st1
tr43:
//line parser.rl:20

        tok = p
      
	goto st28
	st28:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:734
		if data[( p)] == 93 {
			goto tr45
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st28
		}
		goto st1
tr45:
//line parser.rl:72

        event.SetPid(data[tok:p])
      
	goto st29
	st29:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:753
		if data[( p)] == 58 {
			goto st30
		}
		goto st1
	st30:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof30
		}
	st_case_30:
		if data[( p)] == 32 {
			goto st26
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto st26
		}
		goto st1
tr32:
//line parser.rl:56

        event.SetSecond(data[tok:p])
      
//line parser.rl:20

        tok = p
      
	goto st31
tr55:
//line parser.rl:60

        event.SetNanosecond(data[tok:p])
      
//line parser.rl:20

        tok = p
      
	goto st31
	st31:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:795
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st32
		}
		goto st1
	st32:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof32
		}
	st_case_32:
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st33
		}
		goto st1
	st33:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof33
		}
	st_case_33:
		switch data[( p)] {
		case 32:
			goto tr49
		case 58:
			goto st36
		}
		switch {
		case data[( p)] > 13:
			if 48 <= data[( p)] && data[( p)] <= 57 {
				goto st34
			}
		case data[( p)] >= 9:
			goto tr49
		}
		goto st1
	st34:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof34
		}
	st_case_34:
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st35
		}
		goto st1
	st35:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof35
		}
	st_case_35:
		if data[( p)] == 32 {
			goto tr49
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr49
		}
		goto st1
	st36:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof36
		}
	st_case_36:
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st34
		}
		goto st1
tr33:
//line parser.rl:56

        event.SetSecond(data[tok:p])
      
	goto st37
	st37:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof37
		}
	st_case_37:
//line parser.go:870
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr53
		}
		goto st1
tr53:
//line parser.rl:20

        tok = p
      
	goto st38
	st38:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof38
		}
	st_case_38:
//line parser.go:886
		switch data[( p)] {
		case 32:
			goto tr54
		case 43:
			goto tr55
		case 45:
			goto tr55
		case 90:
			goto tr57
		case 122:
			goto tr57
		}
		switch {
		case data[( p)] > 13:
			if 48 <= data[( p)] && data[( p)] <= 57 {
				goto st38
			}
		case data[( p)] >= 9:
			goto tr54
		}
		goto st1
tr34:
//line parser.rl:56

        event.SetSecond(data[tok:p])
      
//line parser.rl:20

        tok = p
      
	goto st39
tr57:
//line parser.rl:60

        event.SetNanosecond(data[tok:p])
      
//line parser.rl:20

        tok = p
      
	goto st39
	st39:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof39
		}
	st_case_39:
//line parser.go:933
		switch data[( p)] {
		case 32:
			goto tr49
		case 43:
			goto st31
		case 45:
			goto st31
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr49
		}
		goto st1
tr23:
//line parser.rl:20

        tok = p
      
	goto st40
	st40:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof40
		}
	st_case_40:
//line parser.go:957
		if 48 <= data[( p)] && data[( p)] <= 51 {
			goto st14
		}
		goto st1
tr2:
//line parser.rl:20

        tok = p
      
	goto st41
	st41:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof41
		}
	st_case_41:
//line parser.go:973
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr59
		}
		goto st1
tr59:
//line parser.rl:20

        tok = p
      
	goto st42
	st42:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof42
		}
	st_case_42:
//line parser.go:989
		if data[( p)] == 62 {
			goto tr61
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st43
		}
		goto st1
	st43:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof43
		}
	st_case_43:
		if data[( p)] == 62 {
			goto tr61
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st44
		}
		goto st1
	st44:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof44
		}
	st_case_44:
		if data[( p)] == 62 {
			goto tr61
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st45
		}
		goto st1
	st45:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof45
		}
	st_case_45:
		if data[( p)] == 62 {
			goto tr61
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st46
		}
		goto st1
	st46:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof46
		}
	st_case_46:
		if data[( p)] == 62 {
			goto tr61
		}
		goto st1
tr61:
//line parser.rl:24

        event.SetPriority(data[tok:p])
      
	goto st47
	st47:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof47
		}
	st_case_47:
//line parser.go:1053
		switch data[( p)] {
		case 65:
			goto tr3
		case 70:
			goto tr4
		case 74:
			goto tr5
		case 77:
			goto tr6
		case 78:
			goto tr7
		case 79:
			goto tr8
		case 83:
			goto tr9
		case 101:
			goto tr10
		}
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr1
		}
		goto tr0
tr3:
//line parser.rl:20

        tok = p
      
	goto st48
	st48:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof48
		}
	st_case_48:
//line parser.go:1087
		switch data[( p)] {
		case 112:
			goto st49
		case 117:
			goto st70
		}
		goto st1
	st49:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof49
		}
	st_case_49:
		if data[( p)] == 114 {
			goto st50
		}
		goto st1
	st50:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof50
		}
	st_case_50:
		switch data[( p)] {
		case 32:
			goto tr68
		case 105:
			goto st68
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
tr68:
//line parser.rl:32

        event.SetMonth(data[tok:p])
      
	goto st51
	st51:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof51
		}
	st_case_51:
//line parser.go:1130
		switch data[( p)] {
		case 32:
			goto st52
		case 51:
			goto tr72
		}
		switch {
		case data[( p)] < 49:
			if 9 <= data[( p)] && data[( p)] <= 13 {
				goto st52
			}
		case data[( p)] > 50:
			if 52 <= data[( p)] && data[( p)] <= 57 {
				goto tr73
			}
		default:
			goto tr71
		}
		goto st1
	st52:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof52
		}
	st_case_52:
		if 49 <= data[( p)] && data[( p)] <= 57 {
			goto tr73
		}
		goto st1
tr73:
//line parser.rl:20

        tok = p
      
	goto st53
	st53:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof53
		}
	st_case_53:
//line parser.go:1170
		if data[( p)] == 32 {
			goto tr74
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr74
		}
		goto st1
tr74:
//line parser.rl:44

        event.SetDay(data[tok:p])
      
	goto st54
	st54:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof54
		}
	st_case_54:
//line parser.go:1189
		if data[( p)] == 50 {
			goto tr76
		}
		if 48 <= data[( p)] && data[( p)] <= 49 {
			goto tr75
		}
		goto st1
tr75:
//line parser.rl:20

        tok = p
      
	goto st55
	st55:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof55
		}
	st_case_55:
//line parser.go:1208
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st56
		}
		goto st1
	st56:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof56
		}
	st_case_56:
		if data[( p)] == 58 {
			goto tr78
		}
		goto st1
tr78:
//line parser.rl:48

        event.SetHour(data[tok:p])
      
	goto st57
	st57:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof57
		}
	st_case_57:
//line parser.go:1233
		if 48 <= data[( p)] && data[( p)] <= 53 {
			goto tr79
		}
		goto st1
tr79:
//line parser.rl:20

        tok = p
      
	goto st58
	st58:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof58
		}
	st_case_58:
//line parser.go:1249
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st59
		}
		goto st1
	st59:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof59
		}
	st_case_59:
		if data[( p)] == 58 {
			goto tr81
		}
		goto st1
tr81:
//line parser.rl:52

        event.SetMinute(data[tok:p])
      
	goto st60
	st60:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof60
		}
	st_case_60:
//line parser.go:1274
		if 48 <= data[( p)] && data[( p)] <= 53 {
			goto tr82
		}
		goto st1
tr82:
//line parser.rl:20

        tok = p
      
	goto st61
	st61:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof61
		}
	st_case_61:
//line parser.go:1290
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto st62
		}
		goto st1
	st62:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof62
		}
	st_case_62:
		switch data[( p)] {
		case 32:
			goto tr31
		case 46:
			goto tr84
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr31
		}
		goto st1
tr84:
//line parser.rl:56

        event.SetSecond(data[tok:p])
      
	goto st63
	st63:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof63
		}
	st_case_63:
//line parser.go:1321
		if 48 <= data[( p)] && data[( p)] <= 57 {
			goto tr85
		}
		goto st1
tr85:
//line parser.rl:20

        tok = p
      
	goto st64
	st64:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof64
		}
	st_case_64:
//line parser.go:1337
		if data[( p)] == 32 {
			goto tr54
		}
		switch {
		case data[( p)] > 13:
			if 48 <= data[( p)] && data[( p)] <= 57 {
				goto st64
			}
		case data[( p)] >= 9:
			goto tr54
		}
		goto st1
tr76:
//line parser.rl:20

        tok = p
      
	goto st65
	st65:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof65
		}
	st_case_65:
//line parser.go:1361
		if 48 <= data[( p)] && data[( p)] <= 51 {
			goto st56
		}
		goto st1
tr71:
//line parser.rl:20

        tok = p
      
	goto st66
	st66:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof66
		}
	st_case_66:
//line parser.go:1377
		if data[( p)] == 32 {
			goto tr74
		}
		switch {
		case data[( p)] > 13:
			if 48 <= data[( p)] && data[( p)] <= 57 {
				goto st53
			}
		case data[( p)] >= 9:
			goto tr74
		}
		goto st1
tr72:
//line parser.rl:20

        tok = p
      
	goto st67
	st67:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof67
		}
	st_case_67:
//line parser.go:1401
		if data[( p)] == 32 {
			goto tr74
		}
		switch {
		case data[( p)] > 13:
			if 48 <= data[( p)] && data[( p)] <= 49 {
				goto st53
			}
		case data[( p)] >= 9:
			goto tr74
		}
		goto st1
	st68:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof68
		}
	st_case_68:
		if data[( p)] == 108 {
			goto st69
		}
		goto st1
	st69:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof69
		}
	st_case_69:
		if data[( p)] == 32 {
			goto tr68
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st70:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof70
		}
	st_case_70:
		if data[( p)] == 103 {
			goto st71
		}
		goto st1
	st71:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof71
		}
	st_case_71:
		switch data[( p)] {
		case 32:
			goto tr68
		case 117:
			goto st72
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st72:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof72
		}
	st_case_72:
		if data[( p)] == 115 {
			goto st73
		}
		goto st1
	st73:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof73
		}
	st_case_73:
		if data[( p)] == 116 {
			goto st69
		}
		goto st1
tr4:
//line parser.rl:20

        tok = p
      
	goto st74
	st74:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof74
		}
	st_case_74:
//line parser.go:1488
		if data[( p)] == 101 {
			goto st75
		}
		goto st1
	st75:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof75
		}
	st_case_75:
		if data[( p)] == 98 {
			goto st76
		}
		goto st1
	st76:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof76
		}
	st_case_76:
		switch data[( p)] {
		case 32:
			goto tr68
		case 114:
			goto st77
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st77:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof77
		}
	st_case_77:
		if data[( p)] == 117 {
			goto st78
		}
		goto st1
	st78:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof78
		}
	st_case_78:
		if data[( p)] == 97 {
			goto st79
		}
		goto st1
	st79:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof79
		}
	st_case_79:
		if data[( p)] == 114 {
			goto st80
		}
		goto st1
	st80:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof80
		}
	st_case_80:
		if data[( p)] == 121 {
			goto st69
		}
		goto st1
tr5:
//line parser.rl:20

        tok = p
      
	goto st81
	st81:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof81
		}
	st_case_81:
//line parser.go:1564
		switch data[( p)] {
		case 97:
			goto st82
		case 117:
			goto st84
		}
		goto st1
	st82:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof82
		}
	st_case_82:
		if data[( p)] == 110 {
			goto st83
		}
		goto st1
	st83:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof83
		}
	st_case_83:
		switch data[( p)] {
		case 32:
			goto tr68
		case 117:
			goto st78
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st84:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof84
		}
	st_case_84:
		switch data[( p)] {
		case 108:
			goto st85
		case 110:
			goto st86
		}
		goto st1
	st85:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof85
		}
	st_case_85:
		switch data[( p)] {
		case 32:
			goto tr68
		case 121:
			goto st69
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st86:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof86
		}
	st_case_86:
		switch data[( p)] {
		case 32:
			goto tr68
		case 101:
			goto st69
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
tr6:
//line parser.rl:20

        tok = p
      
	goto st87
	st87:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof87
		}
	st_case_87:
//line parser.go:1649
		if data[( p)] == 97 {
			goto st88
		}
		goto st1
	st88:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof88
		}
	st_case_88:
		switch data[( p)] {
		case 32:
			goto tr68
		case 114:
			goto st89
		case 121:
			goto st69
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st89:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof89
		}
	st_case_89:
		switch data[( p)] {
		case 32:
			goto tr68
		case 99:
			goto st90
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st90:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof90
		}
	st_case_90:
		if data[( p)] == 104 {
			goto st69
		}
		goto st1
tr7:
//line parser.rl:20

        tok = p
      
	goto st91
	st91:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof91
		}
	st_case_91:
//line parser.go:1706
		if data[( p)] == 111 {
			goto st92
		}
		goto st1
	st92:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof92
		}
	st_case_92:
		if data[( p)] == 118 {
			goto st93
		}
		goto st1
	st93:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof93
		}
	st_case_93:
		switch data[( p)] {
		case 32:
			goto tr68
		case 101:
			goto st94
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st94:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof94
		}
	st_case_94:
		if data[( p)] == 109 {
			goto st95
		}
		goto st1
	st95:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof95
		}
	st_case_95:
		if data[( p)] == 98 {
			goto st96
		}
		goto st1
	st96:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof96
		}
	st_case_96:
		if data[( p)] == 101 {
			goto st97
		}
		goto st1
	st97:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof97
		}
	st_case_97:
		if data[( p)] == 114 {
			goto st69
		}
		goto st1
tr8:
//line parser.rl:20

        tok = p
      
	goto st98
	st98:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof98
		}
	st_case_98:
//line parser.go:1782
		if data[( p)] == 99 {
			goto st99
		}
		goto st1
	st99:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof99
		}
	st_case_99:
		if data[( p)] == 116 {
			goto st100
		}
		goto st1
	st100:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof100
		}
	st_case_100:
		switch data[( p)] {
		case 32:
			goto tr68
		case 111:
			goto st95
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
tr9:
//line parser.rl:20

        tok = p
      
	goto st101
	st101:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof101
		}
	st_case_101:
//line parser.go:1822
		if data[( p)] == 101 {
			goto st102
		}
		goto st1
	st102:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof102
		}
	st_case_102:
		if data[( p)] == 112 {
			goto st103
		}
		goto st1
	st103:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof103
		}
	st_case_103:
		switch data[( p)] {
		case 32:
			goto tr68
		case 116:
			goto st104
		}
		if 9 <= data[( p)] && data[( p)] <= 13 {
			goto tr68
		}
		goto st1
	st104:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof104
		}
	st_case_104:
		if data[( p)] == 101 {
			goto st94
		}
		goto st1
tr10:
//line parser.rl:20

        tok = p
      
	goto st105
	st105:
		if ( p)++; ( p) == ( pe) {
			goto _test_eof105
		}
	st_case_105:
//line parser.go:1871
		if data[( p)] == 99 {
			goto st93
		}
		goto st1
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof

	_test_eof: {}
	if ( p) == eof {
		switch cs {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105:
//line parser.rl:28

        event.SetMessage(data[tok:p])
      
//line parser.go:1991
		}
	}

	}

//line parser.rl:84

}
