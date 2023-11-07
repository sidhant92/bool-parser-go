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
		"", "','", "", "", "", "", "", "", "", "", "", "'!='", "'>'", "'>='",
		"'<'", "'<='", "'='", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "CONTAINS_ALL",
		"CONTAINS_ANY", "NE", "GT", "GE", "LT", "LE", "EQ", "LPAREN", "RPAREN",
		"DECIMAL", "APP_VERSION", "INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ",
		"DQ",
	}
	staticData.RuleNames = []string{
		"T__0", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "CONTAINS_ALL",
		"CONTAINS_ANY", "NE", "GT", "GE", "LT", "LE", "EQ", "LPAREN", "RPAREN",
		"DECIMAL", "APP_VERSION", "INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ",
		"DQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 244, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 60, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 66, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		76, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 92, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 102, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 114, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 140, 8, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 166,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		4, 18, 188, 8, 18, 11, 18, 12, 18, 189, 1, 18, 1, 18, 4, 18, 194, 8, 18,
		11, 18, 12, 18, 195, 1, 19, 1, 19, 1, 19, 4, 19, 201, 8, 19, 11, 19, 12,
		19, 202, 1, 20, 4, 20, 206, 8, 20, 11, 20, 12, 20, 207, 1, 21, 4, 21, 211,
		8, 21, 11, 21, 12, 21, 212, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 4,
		22, 221, 8, 22, 11, 22, 12, 22, 222, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24,
		229, 8, 24, 10, 24, 12, 24, 232, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 5,
		25, 238, 8, 25, 10, 25, 12, 25, 241, 9, 25, 1, 25, 1, 25, 2, 230, 239,
		0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 1, 0, 4, 1, 0,
		48, 57, 3, 0, 9, 10, 12, 13, 32, 32, 2, 0, 45, 46, 95, 95, 3, 0, 48, 57,
		65, 90, 97, 122, 265, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 1, 53, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5, 65, 1, 0, 0, 0, 7, 75, 1, 0,
		0, 0, 9, 83, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15, 113,
		1, 0, 0, 0, 17, 139, 1, 0, 0, 0, 19, 165, 1, 0, 0, 0, 21, 167, 1, 0, 0,
		0, 23, 170, 1, 0, 0, 0, 25, 172, 1, 0, 0, 0, 27, 175, 1, 0, 0, 0, 29, 177,
		1, 0, 0, 0, 31, 180, 1, 0, 0, 0, 33, 182, 1, 0, 0, 0, 35, 184, 1, 0, 0,
		0, 37, 187, 1, 0, 0, 0, 39, 197, 1, 0, 0, 0, 41, 205, 1, 0, 0, 0, 43, 210,
		1, 0, 0, 0, 45, 220, 1, 0, 0, 0, 47, 224, 1, 0, 0, 0, 49, 226, 1, 0, 0,
		0, 51, 235, 1, 0, 0, 0, 53, 54, 5, 44, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56,
		5, 73, 0, 0, 56, 60, 5, 78, 0, 0, 57, 58, 5, 105, 0, 0, 58, 60, 5, 110,
		0, 0, 59, 55, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 4, 1, 0, 0, 0, 61, 62,
		5, 84, 0, 0, 62, 66, 5, 79, 0, 0, 63, 64, 5, 116, 0, 0, 64, 66, 5, 111,
		0, 0, 65, 61, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68,
		5, 65, 0, 0, 68, 69, 5, 78, 0, 0, 69, 76, 5, 68, 0, 0, 70, 71, 5, 97, 0,
		0, 71, 72, 5, 110, 0, 0, 72, 76, 5, 100, 0, 0, 73, 74, 5, 38, 0, 0, 74,
		76, 5, 38, 0, 0, 75, 67, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 75, 73, 1, 0,
		0, 0, 76, 8, 1, 0, 0, 0, 77, 78, 5, 79, 0, 0, 78, 84, 5, 82, 0, 0, 79,
		80, 5, 111, 0, 0, 80, 84, 5, 114, 0, 0, 81, 82, 5, 124, 0, 0, 82, 84, 5,
		124, 0, 0, 83, 77, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		84, 10, 1, 0, 0, 0, 85, 86, 5, 78, 0, 0, 86, 87, 5, 79, 0, 0, 87, 92, 5,
		84, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 111, 0, 0, 90, 92, 5, 116, 0,
		0, 91, 85, 1, 0, 0, 0, 91, 88, 1, 0, 0, 0, 92, 12, 1, 0, 0, 0, 93, 94,
		5, 84, 0, 0, 94, 95, 5, 82, 0, 0, 95, 96, 5, 85, 0, 0, 96, 102, 5, 69,
		0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5, 114, 0, 0, 99, 100, 5, 117, 0, 0,
		100, 102, 5, 101, 0, 0, 101, 93, 1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 102,
		14, 1, 0, 0, 0, 103, 104, 5, 70, 0, 0, 104, 105, 5, 65, 0, 0, 105, 106,
		5, 76, 0, 0, 106, 107, 5, 83, 0, 0, 107, 114, 5, 69, 0, 0, 108, 109, 5,
		102, 0, 0, 109, 110, 5, 97, 0, 0, 110, 111, 5, 108, 0, 0, 111, 112, 5,
		115, 0, 0, 112, 114, 5, 101, 0, 0, 113, 103, 1, 0, 0, 0, 113, 108, 1, 0,
		0, 0, 114, 16, 1, 0, 0, 0, 115, 116, 5, 67, 0, 0, 116, 117, 5, 79, 0, 0,
		117, 118, 5, 78, 0, 0, 118, 119, 5, 84, 0, 0, 119, 120, 5, 65, 0, 0, 120,
		121, 5, 73, 0, 0, 121, 122, 5, 78, 0, 0, 122, 123, 5, 83, 0, 0, 123, 124,
		5, 95, 0, 0, 124, 125, 5, 65, 0, 0, 125, 126, 5, 76, 0, 0, 126, 140, 5,
		76, 0, 0, 127, 128, 5, 99, 0, 0, 128, 129, 5, 111, 0, 0, 129, 130, 5, 110,
		0, 0, 130, 131, 5, 116, 0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 105,
		0, 0, 133, 134, 5, 110, 0, 0, 134, 135, 5, 115, 0, 0, 135, 136, 5, 95,
		0, 0, 136, 137, 5, 97, 0, 0, 137, 138, 5, 108, 0, 0, 138, 140, 5, 108,
		0, 0, 139, 115, 1, 0, 0, 0, 139, 127, 1, 0, 0, 0, 140, 18, 1, 0, 0, 0,
		141, 142, 5, 67, 0, 0, 142, 143, 5, 79, 0, 0, 143, 144, 5, 78, 0, 0, 144,
		145, 5, 84, 0, 0, 145, 146, 5, 65, 0, 0, 146, 147, 5, 73, 0, 0, 147, 148,
		5, 78, 0, 0, 148, 149, 5, 83, 0, 0, 149, 150, 5, 95, 0, 0, 150, 151, 5,
		65, 0, 0, 151, 152, 5, 78, 0, 0, 152, 166, 5, 89, 0, 0, 153, 154, 5, 99,
		0, 0, 154, 155, 5, 111, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157, 5, 116,
		0, 0, 157, 158, 5, 97, 0, 0, 158, 159, 5, 105, 0, 0, 159, 160, 5, 110,
		0, 0, 160, 161, 5, 115, 0, 0, 161, 162, 5, 95, 0, 0, 162, 163, 5, 97, 0,
		0, 163, 164, 5, 110, 0, 0, 164, 166, 5, 121, 0, 0, 165, 141, 1, 0, 0, 0,
		165, 153, 1, 0, 0, 0, 166, 20, 1, 0, 0, 0, 167, 168, 5, 33, 0, 0, 168,
		169, 5, 61, 0, 0, 169, 22, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171, 24,
		1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 174, 5, 61, 0, 0, 174, 26, 1, 0,
		0, 0, 175, 176, 5, 60, 0, 0, 176, 28, 1, 0, 0, 0, 177, 178, 5, 60, 0, 0,
		178, 179, 5, 61, 0, 0, 179, 30, 1, 0, 0, 0, 180, 181, 5, 61, 0, 0, 181,
		32, 1, 0, 0, 0, 182, 183, 5, 40, 0, 0, 183, 34, 1, 0, 0, 0, 184, 185, 5,
		41, 0, 0, 185, 36, 1, 0, 0, 0, 186, 188, 7, 0, 0, 0, 187, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		191, 1, 0, 0, 0, 191, 193, 5, 46, 0, 0, 192, 194, 7, 0, 0, 0, 193, 192,
		1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 38, 1, 0, 0, 0, 197, 200, 7, 0, 0, 0, 198, 199, 5, 46, 0, 0,
		199, 201, 3, 41, 20, 0, 200, 198, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202,
		200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 40, 1, 0, 0, 0, 204, 206, 7,
		0, 0, 0, 205, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 205, 1, 0, 0,
		0, 207, 208, 1, 0, 0, 0, 208, 42, 1, 0, 0, 0, 209, 211, 7, 1, 0, 0, 210,
		209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213,
		1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 6, 21, 0, 0, 215, 44, 1, 0,
		0, 0, 216, 221, 3, 47, 23, 0, 217, 221, 7, 2, 0, 0, 218, 221, 3, 49, 24,
		0, 219, 221, 3, 51, 25, 0, 220, 216, 1, 0, 0, 0, 220, 217, 1, 0, 0, 0,
		220, 218, 1, 0, 0, 0, 220, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 46, 1, 0, 0, 0, 224, 225, 7,
		3, 0, 0, 225, 48, 1, 0, 0, 0, 226, 230, 5, 39, 0, 0, 227, 229, 9, 0, 0,
		0, 228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 234,
		5, 39, 0, 0, 234, 50, 1, 0, 0, 0, 235, 239, 5, 34, 0, 0, 236, 238, 9, 0,
		0, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0,
		239, 237, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242,
		243, 5, 34, 0, 0, 243, 52, 1, 0, 0, 0, 19, 0, 59, 65, 75, 83, 91, 101,
		113, 139, 165, 189, 195, 202, 207, 212, 220, 222, 230, 239, 1, 6, 0, 0,
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
	BooleanExpressionLexerNE           = 11
	BooleanExpressionLexerGT           = 12
	BooleanExpressionLexerGE           = 13
	BooleanExpressionLexerLT           = 14
	BooleanExpressionLexerLE           = 15
	BooleanExpressionLexerEQ           = 16
	BooleanExpressionLexerLPAREN       = 17
	BooleanExpressionLexerRPAREN       = 18
	BooleanExpressionLexerDECIMAL      = 19
	BooleanExpressionLexerAPP_VERSION  = 20
	BooleanExpressionLexerINTEGER      = 21
	BooleanExpressionLexerWS           = 22
	BooleanExpressionLexerWORD         = 23
	BooleanExpressionLexerALPHANUMERIC = 24
	BooleanExpressionLexerSQ           = 25
	BooleanExpressionLexerDQ           = 26
)
