// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated from Eql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr" 
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 35, 230,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 108, 10, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 114, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 124, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 136, 10, 19, 3, 20, 5,
	20, 139, 10, 20, 3, 20, 6, 20, 142, 10, 20, 13, 20, 14, 20, 143, 3, 20,
	3, 20, 6, 20, 148, 10, 20, 13, 20, 14, 20, 149, 3, 21, 5, 21, 153, 10,
	21, 3, 21, 6, 21, 156, 10, 21, 13, 21, 14, 21, 157, 3, 22, 6, 22, 161,
	10, 22, 13, 22, 14, 22, 162, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 5, 23, 173, 10, 23, 3, 24, 3, 24, 7, 24, 177, 10, 24, 12,
	24, 14, 24, 180, 11, 24, 3, 25, 6, 25, 183, 10, 25, 13, 25, 14, 25, 184,
	3, 25, 3, 25, 6, 25, 189, 10, 25, 13, 25, 14, 25, 190, 7, 25, 193, 10,
	25, 12, 25, 14, 25, 196, 11, 25, 3, 26, 3, 26, 7, 26, 200, 10, 26, 12,
	26, 14, 26, 203, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27, 209, 10, 27,
	12, 27, 14, 27, 212, 11, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 2, 2, 35, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 3, 2, 10, 3, 2, 47,
	47, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 5, 2, 67, 92, 97, 97, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 7, 2, 48, 48, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 12, 12, 15, 15, 41, 41, 5, 2, 12, 12, 15, 15, 36,
	36, 2, 246, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 3, 69, 3, 2, 2, 2,
	5, 71, 3, 2, 2, 2, 7, 73, 3, 2, 2, 2, 9, 75, 3, 2, 2, 2, 11, 78, 3, 2,
	2, 2, 13, 81, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2, 19, 88,
	3, 2, 2, 2, 21, 91, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25, 95, 3, 2, 2, 2,
	27, 97, 3, 2, 2, 2, 29, 99, 3, 2, 2, 2, 31, 107, 3, 2, 2, 2, 33, 113, 3,
	2, 2, 2, 35, 123, 3, 2, 2, 2, 37, 135, 3, 2, 2, 2, 39, 138, 3, 2, 2, 2,
	41, 152, 3, 2, 2, 2, 43, 160, 3, 2, 2, 2, 45, 172, 3, 2, 2, 2, 47, 174,
	3, 2, 2, 2, 49, 182, 3, 2, 2, 2, 51, 197, 3, 2, 2, 2, 53, 206, 3, 2, 2,
	2, 55, 215, 3, 2, 2, 2, 57, 217, 3, 2, 2, 2, 59, 219, 3, 2, 2, 2, 61, 221,
	3, 2, 2, 2, 63, 223, 3, 2, 2, 2, 65, 225, 3, 2, 2, 2, 67, 227, 3, 2, 2,
	2, 69, 70, 7, 126, 2, 2, 70, 4, 3, 2, 2, 2, 71, 72, 7, 46, 2, 2, 72, 6,
	3, 2, 2, 2, 73, 74, 7, 60, 2, 2, 74, 8, 3, 2, 2, 2, 75, 76, 7, 63, 2, 2,
	76, 77, 7, 63, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 35, 2, 2, 79, 80, 7,
	63, 2, 2, 80, 12, 3, 2, 2, 2, 81, 82, 7, 64, 2, 2, 82, 14, 3, 2, 2, 2,
	83, 84, 7, 62, 2, 2, 84, 16, 3, 2, 2, 2, 85, 86, 7, 64, 2, 2, 86, 87, 7,
	63, 2, 2, 87, 18, 3, 2, 2, 2, 88, 89, 7, 62, 2, 2, 89, 90, 7, 63, 2, 2,
	90, 20, 3, 2, 2, 2, 91, 92, 7, 45, 2, 2, 92, 22, 3, 2, 2, 2, 93, 94, 7,
	47, 2, 2, 94, 24, 3, 2, 2, 2, 95, 96, 7, 44, 2, 2, 96, 26, 3, 2, 2, 2,
	97, 98, 7, 49, 2, 2, 98, 28, 3, 2, 2, 2, 99, 100, 7, 39, 2, 2, 100, 30,
	3, 2, 2, 2, 101, 102, 7, 99, 2, 2, 102, 103, 7, 112, 2, 2, 103, 108, 7,
	102, 2, 2, 104, 105, 7, 67, 2, 2, 105, 106, 7, 80, 2, 2, 106, 108, 7, 70,
	2, 2, 107, 101, 3, 2, 2, 2, 107, 104, 3, 2, 2, 2, 108, 32, 3, 2, 2, 2,
	109, 110, 7, 113, 2, 2, 110, 114, 7, 116, 2, 2, 111, 112, 7, 81, 2, 2,
	112, 114, 7, 84, 2, 2, 113, 109, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114,
	34, 3, 2, 2, 2, 115, 116, 7, 118, 2, 2, 116, 117, 7, 116, 2, 2, 117, 118,
	7, 119, 2, 2, 118, 124, 7, 103, 2, 2, 119, 120, 7, 86, 2, 2, 120, 121,
	7, 84, 2, 2, 121, 122, 7, 87, 2, 2, 122, 124, 7, 71, 2, 2, 123, 115, 3,
	2, 2, 2, 123, 119, 3, 2, 2, 2, 124, 36, 3, 2, 2, 2, 125, 126, 7, 104, 2,
	2, 126, 127, 7, 99, 2, 2, 127, 128, 7, 110, 2, 2, 128, 129, 7, 117, 2,
	2, 129, 136, 7, 103, 2, 2, 130, 131, 7, 72, 2, 2, 131, 132, 7, 67, 2, 2,
	132, 133, 7, 78, 2, 2, 133, 134, 7, 85, 2, 2, 134, 136, 7, 71, 2, 2, 135,
	125, 3, 2, 2, 2, 135, 130, 3, 2, 2, 2, 136, 38, 3, 2, 2, 2, 137, 139, 9,
	2, 2, 2, 138, 137, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 141, 3, 2, 2,
	2, 140, 142, 9, 3, 2, 2, 141, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 147,
	7, 48, 2, 2, 146, 148, 9, 3, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 40, 3, 2, 2, 2,
	151, 153, 9, 2, 2, 2, 152, 151, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153,
	155, 3, 2, 2, 2, 154, 156, 9, 3, 2, 2, 155, 154, 3, 2, 2, 2, 156, 157,
	3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 42, 3, 2,
	2, 2, 159, 161, 9, 4, 2, 2, 160, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2,
	162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164,
	165, 8, 22, 2, 2, 165, 44, 3, 2, 2, 2, 166, 167, 7, 80, 2, 2, 167, 168,
	7, 81, 2, 2, 168, 173, 7, 86, 2, 2, 169, 170, 7, 112, 2, 2, 170, 171, 7,
	113, 2, 2, 171, 173, 7, 118, 2, 2, 172, 166, 3, 2, 2, 2, 172, 169, 3, 2,
	2, 2, 173, 46, 3, 2, 2, 2, 174, 178, 9, 5, 2, 2, 175, 177, 9, 6, 2, 2,
	176, 175, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178,
	179, 3, 2, 2, 2, 179, 48, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 183, 9,
	7, 2, 2, 182, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2, 2,
	2, 184, 185, 3, 2, 2, 2, 185, 194, 3, 2, 2, 2, 186, 188, 7, 48, 2, 2, 187,
	189, 9, 6, 2, 2, 188, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 188,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193, 3, 2, 2, 2, 192, 186, 3, 2,
	2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2,
	195, 50, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 201, 7, 41, 2, 2, 198,
	200, 10, 8, 2, 2, 199, 198, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199,
	3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2, 2, 203, 201, 3, 2,
	2, 2, 204, 205, 7, 41, 2, 2, 205, 52, 3, 2, 2, 2, 206, 210, 7, 36, 2, 2,
	207, 209, 10, 9, 2, 2, 208, 207, 3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210,
	208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2, 2, 2, 212, 210,
	3, 2, 2, 2, 213, 214, 7, 36, 2, 2, 214, 54, 3, 2, 2, 2, 215, 216, 7, 42,
	2, 2, 216, 56, 3, 2, 2, 2, 217, 218, 7, 43, 2, 2, 218, 58, 3, 2, 2, 2,
	219, 220, 7, 93, 2, 2, 220, 60, 3, 2, 2, 2, 221, 222, 7, 95, 2, 2, 222,
	62, 3, 2, 2, 2, 223, 224, 7, 125, 2, 2, 224, 64, 3, 2, 2, 2, 225, 226,
	7, 127, 2, 2, 226, 66, 3, 2, 2, 2, 227, 228, 7, 38, 2, 2, 228, 229, 7,
	125, 2, 2, 229, 68, 3, 2, 2, 2, 20, 2, 107, 113, 123, 135, 138, 143, 149,
	152, 157, 162, 172, 178, 184, 190, 194, 201, 210, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn) 

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'|'", "','", "':'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "", "", "", "", "", "", "", "", "",
	"", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'", "'${'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "EQ", "NEQ", "GT", "LT", "GTE", "LTE", "ADD", "SUB", "MUL",
	"DIV", "MOD", "AND", "OR", "TRUE", "FALSE", "FLOAT", "NUMBER", "WHITESPACE",
	"NOT", "NAME", "VNAME", "STEXT", "DTEXT", "LPAR", "RPAR", "LARR", "RARR",
	"LDICT", "RDICT", "BEGIN_VARIABLE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "EQ", "NEQ", "GT", "LT", "GTE", "LTE", "ADD", "SUB",
	"MUL", "DIV", "MOD", "AND", "OR", "TRUE", "FALSE", "FLOAT", "NUMBER", "WHITESPACE",
	"NOT", "NAME", "VNAME", "STEXT", "DTEXT", "LPAR", "RPAR", "LARR", "RARR",
	"LDICT", "RDICT", "BEGIN_VARIABLE",
}

type EqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEqlLexer(input antlr.CharStream) *EqlLexer {

	l := new(EqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Eql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EqlLexer tokens.
const (
	EqlLexerT__0           = 1
	EqlLexerT__1           = 2
	EqlLexerT__2           = 3
	EqlLexerEQ             = 4
	EqlLexerNEQ            = 5
	EqlLexerGT             = 6
	EqlLexerLT             = 7
	EqlLexerGTE            = 8
	EqlLexerLTE            = 9
	EqlLexerADD            = 10
	EqlLexerSUB            = 11
	EqlLexerMUL            = 12
	EqlLexerDIV            = 13
	EqlLexerMOD            = 14
	EqlLexerAND            = 15
	EqlLexerOR             = 16
	EqlLexerTRUE           = 17
	EqlLexerFALSE          = 18
	EqlLexerFLOAT          = 19
	EqlLexerNUMBER         = 20
	EqlLexerWHITESPACE     = 21
	EqlLexerNOT            = 22
	EqlLexerNAME           = 23
	EqlLexerVNAME          = 24
	EqlLexerSTEXT          = 25
	EqlLexerDTEXT          = 26
	EqlLexerLPAR           = 27
	EqlLexerRPAR           = 28
	EqlLexerLARR           = 29
	EqlLexerRARR           = 30
	EqlLexerLDICT          = 31
	EqlLexerRDICT          = 32
	EqlLexerBEGIN_VARIABLE = 33
)
