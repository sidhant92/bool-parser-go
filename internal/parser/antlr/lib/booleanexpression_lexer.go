// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lib

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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

var booleanexpressionlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func booleanexpressionlexerLexerInit() {
	staticData := &booleanexpressionlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "", "", "", "", "", "", "", "'!='", "'>'", "'>='", "'<'",
		"'<='", "'='", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "NE", "GT",
		"GE", "LT", "LE", "EQ", "LPAREN", "RPAREN", "DECIMAL", "APP_VERSION",
		"INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ", "DQ",
	}
	staticData.ruleNames = []string{
		"T__0", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "NE", "GT",
		"GE", "LT", "LE", "EQ", "LPAREN", "RPAREN", "DECIMAL", "APP_VERSION",
		"INTEGER", "WS", "WORD", "ALPHANUMERIC", "SQ", "DQ",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 24, 188, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 62, 8, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 72, 8, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 80, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 88, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 98,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		110, 8, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		4, 16, 132, 8, 16, 11, 16, 12, 16, 133, 1, 16, 1, 16, 4, 16, 138, 8, 16,
		11, 16, 12, 16, 139, 1, 17, 1, 17, 1, 17, 4, 17, 145, 8, 17, 11, 17, 12,
		17, 146, 1, 18, 4, 18, 150, 8, 18, 11, 18, 12, 18, 151, 1, 19, 4, 19, 155,
		8, 19, 11, 19, 12, 19, 156, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 4,
		20, 165, 8, 20, 11, 20, 12, 20, 166, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22,
		173, 8, 22, 10, 22, 12, 22, 176, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 5,
		23, 182, 8, 23, 10, 23, 12, 23, 185, 9, 23, 1, 23, 1, 23, 2, 174, 183,
		0, 24, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 1, 0, 4, 1, 0, 48, 57, 3, 0, 9,
		10, 12, 13, 32, 32, 2, 0, 45, 46, 95, 95, 3, 0, 48, 57, 65, 90, 97, 122,
		207, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 1, 49, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 61, 1, 0,
		0, 0, 7, 71, 1, 0, 0, 0, 9, 79, 1, 0, 0, 0, 11, 87, 1, 0, 0, 0, 13, 97,
		1, 0, 0, 0, 15, 109, 1, 0, 0, 0, 17, 111, 1, 0, 0, 0, 19, 114, 1, 0, 0,
		0, 21, 116, 1, 0, 0, 0, 23, 119, 1, 0, 0, 0, 25, 121, 1, 0, 0, 0, 27, 124,
		1, 0, 0, 0, 29, 126, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 131, 1, 0, 0,
		0, 35, 141, 1, 0, 0, 0, 37, 149, 1, 0, 0, 0, 39, 154, 1, 0, 0, 0, 41, 164,
		1, 0, 0, 0, 43, 168, 1, 0, 0, 0, 45, 170, 1, 0, 0, 0, 47, 179, 1, 0, 0,
		0, 49, 50, 5, 44, 0, 0, 50, 2, 1, 0, 0, 0, 51, 52, 5, 73, 0, 0, 52, 56,
		5, 78, 0, 0, 53, 54, 5, 105, 0, 0, 54, 56, 5, 110, 0, 0, 55, 51, 1, 0,
		0, 0, 55, 53, 1, 0, 0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 84, 0, 0, 58, 62,
		5, 79, 0, 0, 59, 60, 5, 116, 0, 0, 60, 62, 5, 111, 0, 0, 61, 57, 1, 0,
		0, 0, 61, 59, 1, 0, 0, 0, 62, 6, 1, 0, 0, 0, 63, 64, 5, 65, 0, 0, 64, 65,
		5, 78, 0, 0, 65, 72, 5, 68, 0, 0, 66, 67, 5, 97, 0, 0, 67, 68, 5, 110,
		0, 0, 68, 72, 5, 100, 0, 0, 69, 70, 5, 38, 0, 0, 70, 72, 5, 38, 0, 0, 71,
		63, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 8, 1, 0, 0,
		0, 73, 74, 5, 79, 0, 0, 74, 80, 5, 82, 0, 0, 75, 76, 5, 111, 0, 0, 76,
		80, 5, 114, 0, 0, 77, 78, 5, 124, 0, 0, 78, 80, 5, 124, 0, 0, 79, 73, 1,
		0, 0, 0, 79, 75, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 10, 1, 0, 0, 0, 81,
		82, 5, 78, 0, 0, 82, 83, 5, 79, 0, 0, 83, 88, 5, 84, 0, 0, 84, 85, 5, 110,
		0, 0, 85, 86, 5, 111, 0, 0, 86, 88, 5, 116, 0, 0, 87, 81, 1, 0, 0, 0, 87,
		84, 1, 0, 0, 0, 88, 12, 1, 0, 0, 0, 89, 90, 5, 84, 0, 0, 90, 91, 5, 82,
		0, 0, 91, 92, 5, 85, 0, 0, 92, 98, 5, 69, 0, 0, 93, 94, 5, 116, 0, 0, 94,
		95, 5, 114, 0, 0, 95, 96, 5, 117, 0, 0, 96, 98, 5, 101, 0, 0, 97, 89, 1,
		0, 0, 0, 97, 93, 1, 0, 0, 0, 98, 14, 1, 0, 0, 0, 99, 100, 5, 70, 0, 0,
		100, 101, 5, 65, 0, 0, 101, 102, 5, 76, 0, 0, 102, 103, 5, 83, 0, 0, 103,
		110, 5, 69, 0, 0, 104, 105, 5, 102, 0, 0, 105, 106, 5, 97, 0, 0, 106, 107,
		5, 108, 0, 0, 107, 108, 5, 115, 0, 0, 108, 110, 5, 101, 0, 0, 109, 99,
		1, 0, 0, 0, 109, 104, 1, 0, 0, 0, 110, 16, 1, 0, 0, 0, 111, 112, 5, 33,
		0, 0, 112, 113, 5, 61, 0, 0, 113, 18, 1, 0, 0, 0, 114, 115, 5, 62, 0, 0,
		115, 20, 1, 0, 0, 0, 116, 117, 5, 62, 0, 0, 117, 118, 5, 61, 0, 0, 118,
		22, 1, 0, 0, 0, 119, 120, 5, 60, 0, 0, 120, 24, 1, 0, 0, 0, 121, 122, 5,
		60, 0, 0, 122, 123, 5, 61, 0, 0, 123, 26, 1, 0, 0, 0, 124, 125, 5, 61,
		0, 0, 125, 28, 1, 0, 0, 0, 126, 127, 5, 40, 0, 0, 127, 30, 1, 0, 0, 0,
		128, 129, 5, 41, 0, 0, 129, 32, 1, 0, 0, 0, 130, 132, 7, 0, 0, 0, 131,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 5, 46, 0, 0, 136, 138, 7, 0,
		0, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 34, 1, 0, 0, 0, 141, 144, 7, 0, 0, 0, 142, 143,
		5, 46, 0, 0, 143, 145, 3, 37, 18, 0, 144, 142, 1, 0, 0, 0, 145, 146, 1,
		0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 36, 1, 0, 0,
		0, 148, 150, 7, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 38, 1, 0, 0, 0, 153, 155, 7,
		1, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0,
		0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 6, 19, 0, 0, 159,
		40, 1, 0, 0, 0, 160, 165, 3, 43, 21, 0, 161, 165, 7, 2, 0, 0, 162, 165,
		3, 45, 22, 0, 163, 165, 3, 47, 23, 0, 164, 160, 1, 0, 0, 0, 164, 161, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 42, 1, 0, 0, 0, 168,
		169, 7, 3, 0, 0, 169, 44, 1, 0, 0, 0, 170, 174, 5, 39, 0, 0, 171, 173,
		9, 0, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0,
		177, 178, 5, 39, 0, 0, 178, 46, 1, 0, 0, 0, 179, 183, 5, 34, 0, 0, 180,
		182, 9, 0, 0, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 186, 1, 0, 0, 0, 185, 183, 1, 0,
		0, 0, 186, 187, 5, 34, 0, 0, 187, 48, 1, 0, 0, 0, 17, 0, 55, 61, 71, 79,
		87, 97, 109, 133, 139, 146, 151, 156, 164, 166, 174, 183, 1, 6, 0, 0,
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
	staticData := &booleanexpressionlexerLexerStaticData
	staticData.once.Do(booleanexpressionlexerLexerInit)
}

// NewBooleanExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBooleanExpressionLexer(input antlr.CharStream) *BooleanExpressionLexer {
	BooleanExpressionLexerInit()
	l := new(BooleanExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &booleanexpressionlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
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
	BooleanExpressionLexerNE           = 9
	BooleanExpressionLexerGT           = 10
	BooleanExpressionLexerGE           = 11
	BooleanExpressionLexerLT           = 12
	BooleanExpressionLexerLE           = 13
	BooleanExpressionLexerEQ           = 14
	BooleanExpressionLexerLPAREN       = 15
	BooleanExpressionLexerRPAREN       = 16
	BooleanExpressionLexerDECIMAL      = 17
	BooleanExpressionLexerAPP_VERSION  = 18
	BooleanExpressionLexerINTEGER      = 19
	BooleanExpressionLexerWS           = 20
	BooleanExpressionLexerWORD         = 21
	BooleanExpressionLexerALPHANUMERIC = 22
	BooleanExpressionLexerSQ           = 23
	BooleanExpressionLexerDQ           = 24
)
