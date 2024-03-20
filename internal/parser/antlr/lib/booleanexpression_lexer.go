// Code generated from /Users/sid/Desktop/filter1/BooleanExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lib

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BooleanExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BooleanExpressionLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func booleanexpressionlexerLexerInit() {
	staticData := &BooleanExpressionLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'^'", "'!='", "'>'", "'>='", "'<'", "'<='", "'='", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "CONTAINS_ALL",
		"CONTAINS_ANY", "ADD", "SUBTRACT", "MULTIPLY", "DIVIDE", "MODULUS",
		"EXPONENT", "NE", "GT", "GE", "LT", "LE", "EQ", "LPAREN", "RPAREN",
		"DECIMAL", "APP_VERSION", "INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ",
		"DQ",
	}
	staticData.RuleNames = []string{
		"T__0", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "CONTAINS_ALL",
		"CONTAINS_ANY", "ADD", "SUBTRACT", "MULTIPLY", "DIVIDE", "MODULUS",
		"EXPONENT", "NE", "GT", "GE", "LT", "LE", "EQ", "LPAREN", "RPAREN",
		"DECIMAL", "APP_VERSION", "INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ",
		"DQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 294, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 72, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 88, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 96, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 104, 8, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 114, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 126, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 152, 8,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 3, 9, 178, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 4, 24, 212, 8, 24, 11, 24, 12, 24, 213, 1, 24,
		1, 24, 4, 24, 218, 8, 24, 11, 24, 12, 24, 219, 1, 25, 1, 25, 1, 25, 4,
		25, 225, 8, 25, 11, 25, 12, 25, 226, 1, 26, 4, 26, 230, 8, 26, 11, 26,
		12, 26, 231, 1, 27, 4, 27, 235, 8, 27, 11, 27, 12, 27, 236, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 4, 28, 244, 8, 28, 11, 28, 12, 28, 245, 1, 28, 1,
		28, 1, 28, 1, 28, 5, 28, 252, 8, 28, 10, 28, 12, 28, 255, 9, 28, 1, 28,
		1, 28, 3, 28, 259, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 4, 28, 265, 8, 28,
		11, 28, 12, 28, 266, 1, 28, 1, 28, 3, 28, 271, 8, 28, 3, 28, 273, 8, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 279, 8, 30, 10, 30, 12, 30, 282, 9,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 288, 8, 31, 10, 31, 12, 31, 291,
		9, 31, 1, 31, 1, 31, 2, 280, 289, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 1,
		0, 5, 1, 0, 48, 57, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 46, 46, 95, 95,
		2, 0, 45, 46, 95, 95, 3, 0, 48, 57, 65, 90, 97, 122, 325, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 1, 65, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 87, 1,
		0, 0, 0, 9, 95, 1, 0, 0, 0, 11, 103, 1, 0, 0, 0, 13, 113, 1, 0, 0, 0, 15,
		125, 1, 0, 0, 0, 17, 151, 1, 0, 0, 0, 19, 177, 1, 0, 0, 0, 21, 179, 1,
		0, 0, 0, 23, 181, 1, 0, 0, 0, 25, 183, 1, 0, 0, 0, 27, 185, 1, 0, 0, 0,
		29, 187, 1, 0, 0, 0, 31, 189, 1, 0, 0, 0, 33, 191, 1, 0, 0, 0, 35, 194,
		1, 0, 0, 0, 37, 196, 1, 0, 0, 0, 39, 199, 1, 0, 0, 0, 41, 201, 1, 0, 0,
		0, 43, 204, 1, 0, 0, 0, 45, 206, 1, 0, 0, 0, 47, 208, 1, 0, 0, 0, 49, 211,
		1, 0, 0, 0, 51, 221, 1, 0, 0, 0, 53, 229, 1, 0, 0, 0, 55, 234, 1, 0, 0,
		0, 57, 272, 1, 0, 0, 0, 59, 274, 1, 0, 0, 0, 61, 276, 1, 0, 0, 0, 63, 285,
		1, 0, 0, 0, 65, 66, 5, 44, 0, 0, 66, 2, 1, 0, 0, 0, 67, 68, 5, 73, 0, 0,
		68, 72, 5, 78, 0, 0, 69, 70, 5, 105, 0, 0, 70, 72, 5, 110, 0, 0, 71, 67,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5, 84, 0, 0,
		74, 78, 5, 79, 0, 0, 75, 76, 5, 116, 0, 0, 76, 78, 5, 111, 0, 0, 77, 73,
		1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 65, 0, 0,
		80, 81, 5, 78, 0, 0, 81, 88, 5, 68, 0, 0, 82, 83, 5, 97, 0, 0, 83, 84,
		5, 110, 0, 0, 84, 88, 5, 100, 0, 0, 85, 86, 5, 38, 0, 0, 86, 88, 5, 38,
		0, 0, 87, 79, 1, 0, 0, 0, 87, 82, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 8,
		1, 0, 0, 0, 89, 90, 5, 79, 0, 0, 90, 96, 5, 82, 0, 0, 91, 92, 5, 111, 0,
		0, 92, 96, 5, 114, 0, 0, 93, 94, 5, 124, 0, 0, 94, 96, 5, 124, 0, 0, 95,
		89, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 10, 1, 0, 0,
		0, 97, 98, 5, 78, 0, 0, 98, 99, 5, 79, 0, 0, 99, 104, 5, 84, 0, 0, 100,
		101, 5, 110, 0, 0, 101, 102, 5, 111, 0, 0, 102, 104, 5, 116, 0, 0, 103,
		97, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0, 104, 12, 1, 0, 0, 0, 105, 106, 5,
		84, 0, 0, 106, 107, 5, 82, 0, 0, 107, 108, 5, 85, 0, 0, 108, 114, 5, 69,
		0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 114, 0, 0, 111, 112, 5, 117,
		0, 0, 112, 114, 5, 101, 0, 0, 113, 105, 1, 0, 0, 0, 113, 109, 1, 0, 0,
		0, 114, 14, 1, 0, 0, 0, 115, 116, 5, 70, 0, 0, 116, 117, 5, 65, 0, 0, 117,
		118, 5, 76, 0, 0, 118, 119, 5, 83, 0, 0, 119, 126, 5, 69, 0, 0, 120, 121,
		5, 102, 0, 0, 121, 122, 5, 97, 0, 0, 122, 123, 5, 108, 0, 0, 123, 124,
		5, 115, 0, 0, 124, 126, 5, 101, 0, 0, 125, 115, 1, 0, 0, 0, 125, 120, 1,
		0, 0, 0, 126, 16, 1, 0, 0, 0, 127, 128, 5, 67, 0, 0, 128, 129, 5, 79, 0,
		0, 129, 130, 5, 78, 0, 0, 130, 131, 5, 84, 0, 0, 131, 132, 5, 65, 0, 0,
		132, 133, 5, 73, 0, 0, 133, 134, 5, 78, 0, 0, 134, 135, 5, 83, 0, 0, 135,
		136, 5, 95, 0, 0, 136, 137, 5, 65, 0, 0, 137, 138, 5, 76, 0, 0, 138, 152,
		5, 76, 0, 0, 139, 140, 5, 99, 0, 0, 140, 141, 5, 111, 0, 0, 141, 142, 5,
		110, 0, 0, 142, 143, 5, 116, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5,
		105, 0, 0, 145, 146, 5, 110, 0, 0, 146, 147, 5, 115, 0, 0, 147, 148, 5,
		95, 0, 0, 148, 149, 5, 97, 0, 0, 149, 150, 5, 108, 0, 0, 150, 152, 5, 108,
		0, 0, 151, 127, 1, 0, 0, 0, 151, 139, 1, 0, 0, 0, 152, 18, 1, 0, 0, 0,
		153, 154, 5, 67, 0, 0, 154, 155, 5, 79, 0, 0, 155, 156, 5, 78, 0, 0, 156,
		157, 5, 84, 0, 0, 157, 158, 5, 65, 0, 0, 158, 159, 5, 73, 0, 0, 159, 160,
		5, 78, 0, 0, 160, 161, 5, 83, 0, 0, 161, 162, 5, 95, 0, 0, 162, 163, 5,
		65, 0, 0, 163, 164, 5, 78, 0, 0, 164, 178, 5, 89, 0, 0, 165, 166, 5, 99,
		0, 0, 166, 167, 5, 111, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169, 5, 116,
		0, 0, 169, 170, 5, 97, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 110,
		0, 0, 172, 173, 5, 115, 0, 0, 173, 174, 5, 95, 0, 0, 174, 175, 5, 97, 0,
		0, 175, 176, 5, 110, 0, 0, 176, 178, 5, 121, 0, 0, 177, 153, 1, 0, 0, 0,
		177, 165, 1, 0, 0, 0, 178, 20, 1, 0, 0, 0, 179, 180, 5, 43, 0, 0, 180,
		22, 1, 0, 0, 0, 181, 182, 5, 45, 0, 0, 182, 24, 1, 0, 0, 0, 183, 184, 5,
		42, 0, 0, 184, 26, 1, 0, 0, 0, 185, 186, 5, 47, 0, 0, 186, 28, 1, 0, 0,
		0, 187, 188, 5, 37, 0, 0, 188, 30, 1, 0, 0, 0, 189, 190, 5, 94, 0, 0, 190,
		32, 1, 0, 0, 0, 191, 192, 5, 33, 0, 0, 192, 193, 5, 61, 0, 0, 193, 34,
		1, 0, 0, 0, 194, 195, 5, 62, 0, 0, 195, 36, 1, 0, 0, 0, 196, 197, 5, 62,
		0, 0, 197, 198, 5, 61, 0, 0, 198, 38, 1, 0, 0, 0, 199, 200, 5, 60, 0, 0,
		200, 40, 1, 0, 0, 0, 201, 202, 5, 60, 0, 0, 202, 203, 5, 61, 0, 0, 203,
		42, 1, 0, 0, 0, 204, 205, 5, 61, 0, 0, 205, 44, 1, 0, 0, 0, 206, 207, 5,
		40, 0, 0, 207, 46, 1, 0, 0, 0, 208, 209, 5, 41, 0, 0, 209, 48, 1, 0, 0,
		0, 210, 212, 7, 0, 0, 0, 211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217,
		5, 46, 0, 0, 216, 218, 7, 0, 0, 0, 217, 216, 1, 0, 0, 0, 218, 219, 1, 0,
		0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 50, 1, 0, 0, 0,
		221, 224, 7, 0, 0, 0, 222, 223, 5, 46, 0, 0, 223, 225, 3, 53, 26, 0, 224,
		222, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 52, 1, 0, 0, 0, 228, 230, 7, 0, 0, 0, 229, 228, 1, 0,
		0, 0, 230, 231, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0,
		232, 54, 1, 0, 0, 0, 233, 235, 7, 1, 0, 0, 234, 233, 1, 0, 0, 0, 235, 236,
		1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 1, 0,
		0, 0, 238, 239, 6, 27, 0, 0, 239, 56, 1, 0, 0, 0, 240, 244, 3, 59, 29,
		0, 241, 244, 3, 61, 30, 0, 242, 244, 3, 63, 31, 0, 243, 240, 1, 0, 0, 0,
		243, 241, 1, 0, 0, 0, 243, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245,
		243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 253, 1, 0, 0, 0, 247, 252,
		3, 59, 29, 0, 248, 252, 7, 2, 0, 0, 249, 252, 3, 61, 30, 0, 250, 252, 3,
		63, 31, 0, 251, 247, 1, 0, 0, 0, 251, 248, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 251, 250, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0,
		253, 254, 1, 0, 0, 0, 254, 273, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256,
		259, 3, 61, 30, 0, 257, 259, 3, 63, 31, 0, 258, 256, 1, 0, 0, 0, 258, 257,
		1, 0, 0, 0, 259, 264, 1, 0, 0, 0, 260, 265, 3, 59, 29, 0, 261, 265, 7,
		3, 0, 0, 262, 265, 3, 61, 30, 0, 263, 265, 3, 63, 31, 0, 264, 260, 1, 0,
		0, 0, 264, 261, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		270, 1, 0, 0, 0, 268, 271, 3, 61, 30, 0, 269, 271, 3, 63, 31, 0, 270, 268,
		1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 243, 1, 0,
		0, 0, 272, 258, 1, 0, 0, 0, 273, 58, 1, 0, 0, 0, 274, 275, 7, 4, 0, 0,
		275, 60, 1, 0, 0, 0, 276, 280, 5, 39, 0, 0, 277, 279, 9, 0, 0, 0, 278,
		277, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 280, 278,
		1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 284, 5, 39,
		0, 0, 284, 62, 1, 0, 0, 0, 285, 289, 5, 34, 0, 0, 286, 288, 9, 0, 0, 0,
		287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 292, 293,
		5, 34, 0, 0, 293, 64, 1, 0, 0, 0, 26, 0, 71, 77, 87, 95, 103, 113, 125,
		151, 177, 213, 219, 226, 231, 236, 243, 245, 251, 253, 258, 264, 266, 270,
		272, 280, 289, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BooleanExpressionLexerInit initializes any static state used to implement BooleanExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBooleanExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BooleanExpressionLexerInit() {
	staticData := &BooleanExpressionLexerLexerStaticData
	staticData.once.Do(booleanexpressionlexerLexerInit)
}

// NewBooleanExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBooleanExpressionLexer(input antlr.CharStream) *BooleanExpressionLexer {
	BooleanExpressionLexerInit()
	l := new(BooleanExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BooleanExpressionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "BooleanExpression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BooleanExpressionLexer tokens.
const (
	BooleanExpressionLexerT__0         = 1
	BooleanExpressionLexerIN           = 2
	BooleanExpressionLexerTO           = 3
	BooleanExpressionLexerAND          = 4
	BooleanExpressionLexerOR           = 5
	BooleanExpressionLexerNOT          = 6
	BooleanExpressionLexerTRUE         = 7
	BooleanExpressionLexerFALSE        = 8
	BooleanExpressionLexerCONTAINS_ALL = 9
	BooleanExpressionLexerCONTAINS_ANY = 10
	BooleanExpressionLexerADD          = 11
	BooleanExpressionLexerSUBTRACT     = 12
	BooleanExpressionLexerMULTIPLY     = 13
	BooleanExpressionLexerDIVIDE       = 14
	BooleanExpressionLexerMODULUS      = 15
	BooleanExpressionLexerEXPONENT     = 16
	BooleanExpressionLexerNE           = 17
	BooleanExpressionLexerGT           = 18
	BooleanExpressionLexerGE           = 19
	BooleanExpressionLexerLT           = 20
	BooleanExpressionLexerLE           = 21
	BooleanExpressionLexerEQ           = 22
	BooleanExpressionLexerLPAREN       = 23
	BooleanExpressionLexerRPAREN       = 24
	BooleanExpressionLexerDECIMAL      = 25
	BooleanExpressionLexerAPP_VERSION  = 26
	BooleanExpressionLexerINTEGER      = 27
	BooleanExpressionLexerWS           = 28
	BooleanExpressionLexerWORD         = 29
	BooleanExpressionLexerALPHANUMERIC = 30
	BooleanExpressionLexerSQ           = 31
	BooleanExpressionLexerDQ           = 32
)
