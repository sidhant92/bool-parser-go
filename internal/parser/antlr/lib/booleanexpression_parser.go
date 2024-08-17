// Code generated from /Users/sid/Desktop/filter2/BooleanExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package lib // BooleanExpression

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BooleanExpressionParser struct {
	*antlr.BaseParser
}

var BooleanExpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func booleanexpressionParserInit() {
	staticData := &BooleanExpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'!='", "'>'",
		"'>='", "'<'", "'<='", "'='", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "IN", "TO", "AND", "OR", "NOT", "TRUE", "FALSE", "CONTAINS_ALL",
		"CONTAINS_ANY", "MIN", "MAX", "AVG", "SUM", "MEAN", "MODE", "MEDIAN",
		"LEN", "INT", "ADD", "SUBTRACT", "MULTIPLY", "DIVIDE", "MODULUS", "EXPONENT",
		"NE", "GT", "GE", "LT", "LE", "EQ", "LPAREN", "RPAREN", "DECIMAL", "APP_VERSION",
		"INTEGER", "WS", "WORD", "SQSTR", "DQSTR", "FIELD", "ALPHANUMERIC",
		"SQ", "DQ",
	}
	staticData.RuleNames = []string{
		"parse", "expression", "comparator", "arithmeticOperator", "arithmeticFunction",
		"wordlist", "arrayOperators", "numericTypes", "types", "binary", "bool",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 145, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 46, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 54, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 82, 8, 1, 10, 1, 12, 1, 85, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 5, 5, 95, 8, 5, 10, 5, 12, 5, 98, 9, 5, 1, 5, 1, 5,
		5, 5, 102, 8, 5, 10, 5, 12, 5, 105, 9, 5, 1, 5, 1, 5, 5, 5, 109, 8, 5,
		10, 5, 12, 5, 112, 9, 5, 1, 5, 1, 5, 5, 5, 116, 8, 5, 10, 5, 12, 5, 119,
		9, 5, 5, 5, 121, 8, 5, 10, 5, 12, 5, 124, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 139, 8, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 1, 2, 11, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 0, 7, 1, 0, 26, 31, 1, 0, 20, 25, 1, 0, 11, 19, 1, 0, 9, 10,
		2, 0, 34, 34, 36, 36, 1, 0, 4, 5, 1, 0, 7, 8, 160, 0, 22, 1, 0, 0, 0, 2,
		53, 1, 0, 0, 0, 4, 86, 1, 0, 0, 0, 6, 88, 1, 0, 0, 0, 8, 90, 1, 0, 0, 0,
		10, 92, 1, 0, 0, 0, 12, 127, 1, 0, 0, 0, 14, 129, 1, 0, 0, 0, 16, 138,
		1, 0, 0, 0, 18, 140, 1, 0, 0, 0, 20, 142, 1, 0, 0, 0, 22, 23, 3, 2, 1,
		0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 6, 1, -1, 0, 26, 27,
		5, 32, 0, 0, 27, 28, 3, 2, 1, 0, 28, 29, 5, 33, 0, 0, 29, 54, 1, 0, 0,
		0, 30, 31, 5, 6, 0, 0, 31, 54, 3, 2, 1, 15, 32, 33, 5, 21, 0, 0, 33, 54,
		3, 2, 1, 13, 34, 35, 3, 8, 4, 0, 35, 36, 3, 10, 5, 0, 36, 54, 1, 0, 0,
		0, 37, 54, 3, 16, 8, 0, 38, 39, 5, 41, 0, 0, 39, 40, 3, 14, 7, 0, 40, 41,
		5, 3, 0, 0, 41, 42, 3, 14, 7, 0, 42, 54, 1, 0, 0, 0, 43, 45, 5, 41, 0,
		0, 44, 46, 5, 6, 0, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47,
		1, 0, 0, 0, 47, 48, 5, 2, 0, 0, 48, 54, 3, 10, 5, 0, 49, 50, 5, 41, 0,
		0, 50, 51, 3, 12, 6, 0, 51, 52, 3, 10, 5, 0, 52, 54, 1, 0, 0, 0, 53, 25,
		1, 0, 0, 0, 53, 30, 1, 0, 0, 0, 53, 32, 1, 0, 0, 0, 53, 34, 1, 0, 0, 0,
		53, 37, 1, 0, 0, 0, 53, 38, 1, 0, 0, 0, 53, 43, 1, 0, 0, 0, 53, 49, 1,
		0, 0, 0, 54, 83, 1, 0, 0, 0, 55, 56, 10, 14, 0, 0, 56, 57, 3, 4, 2, 0,
		57, 58, 3, 2, 1, 15, 58, 82, 1, 0, 0, 0, 59, 60, 10, 12, 0, 0, 60, 61,
		5, 25, 0, 0, 61, 82, 3, 2, 1, 13, 62, 63, 10, 11, 0, 0, 63, 64, 5, 23,
		0, 0, 64, 82, 3, 2, 1, 12, 65, 66, 10, 10, 0, 0, 66, 67, 5, 22, 0, 0, 67,
		82, 3, 2, 1, 11, 68, 69, 10, 9, 0, 0, 69, 70, 5, 24, 0, 0, 70, 82, 3, 2,
		1, 10, 71, 72, 10, 8, 0, 0, 72, 73, 5, 20, 0, 0, 73, 82, 3, 2, 1, 9, 74,
		75, 10, 7, 0, 0, 75, 76, 5, 21, 0, 0, 76, 82, 3, 2, 1, 8, 77, 78, 10, 5,
		0, 0, 78, 79, 3, 18, 9, 0, 79, 80, 3, 2, 1, 6, 80, 82, 1, 0, 0, 0, 81,
		55, 1, 0, 0, 0, 81, 59, 1, 0, 0, 0, 81, 62, 1, 0, 0, 0, 81, 65, 1, 0, 0,
		0, 81, 68, 1, 0, 0, 0, 81, 71, 1, 0, 0, 0, 81, 74, 1, 0, 0, 0, 81, 77,
		1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0,
		84, 3, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 7, 0, 0, 0, 87, 5, 1, 0,
		0, 0, 88, 89, 7, 1, 0, 0, 89, 7, 1, 0, 0, 0, 90, 91, 7, 2, 0, 0, 91, 9,
		1, 0, 0, 0, 92, 96, 5, 32, 0, 0, 93, 95, 5, 37, 0, 0, 94, 93, 1, 0, 0,
		0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99,
		1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 103, 3, 2, 1, 0, 100, 102, 5, 37, 0,
		0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 122, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 110,
		5, 1, 0, 0, 107, 109, 5, 37, 0, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0,
		0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0,
		112, 110, 1, 0, 0, 0, 113, 117, 3, 2, 1, 0, 114, 116, 5, 37, 0, 0, 115,
		114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 106, 1, 0,
		0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0,
		123, 125, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 5, 33, 0, 0, 126,
		11, 1, 0, 0, 0, 127, 128, 7, 3, 0, 0, 128, 13, 1, 0, 0, 0, 129, 130, 7,
		4, 0, 0, 130, 15, 1, 0, 0, 0, 131, 139, 5, 36, 0, 0, 132, 139, 5, 34, 0,
		0, 133, 139, 5, 35, 0, 0, 134, 139, 3, 20, 10, 0, 135, 139, 5, 38, 0, 0,
		136, 139, 5, 41, 0, 0, 137, 139, 1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 138,
		132, 1, 0, 0, 0, 138, 133, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 135,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 17, 1, 0,
		0, 0, 140, 141, 7, 5, 0, 0, 141, 19, 1, 0, 0, 0, 142, 143, 7, 6, 0, 0,
		143, 21, 1, 0, 0, 0, 10, 45, 53, 81, 83, 96, 103, 110, 117, 122, 138,
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

// BooleanExpressionParserInit initializes any static state used to implement BooleanExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBooleanExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BooleanExpressionParserInit() {
	staticData := &BooleanExpressionParserStaticData
	staticData.once.Do(booleanexpressionParserInit)
}

// NewBooleanExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBooleanExpressionParser(input antlr.TokenStream) *BooleanExpressionParser {
	BooleanExpressionParserInit()
	this := new(BooleanExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BooleanExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BooleanExpression.g4"

	return this
}

// BooleanExpressionParser tokens.
const (
	BooleanExpressionParserEOF          = antlr.TokenEOF
	BooleanExpressionParserT__0         = 1
	BooleanExpressionParserIN           = 2
	BooleanExpressionParserTO           = 3
	BooleanExpressionParserAND          = 4
	BooleanExpressionParserOR           = 5
	BooleanExpressionParserNOT          = 6
	BooleanExpressionParserTRUE         = 7
	BooleanExpressionParserFALSE        = 8
	BooleanExpressionParserCONTAINS_ALL = 9
	BooleanExpressionParserCONTAINS_ANY = 10
	BooleanExpressionParserMIN          = 11
	BooleanExpressionParserMAX          = 12
	BooleanExpressionParserAVG          = 13
	BooleanExpressionParserSUM          = 14
	BooleanExpressionParserMEAN         = 15
	BooleanExpressionParserMODE         = 16
	BooleanExpressionParserMEDIAN       = 17
	BooleanExpressionParserLEN          = 18
	BooleanExpressionParserINT          = 19
	BooleanExpressionParserADD          = 20
	BooleanExpressionParserSUBTRACT     = 21
	BooleanExpressionParserMULTIPLY     = 22
	BooleanExpressionParserDIVIDE       = 23
	BooleanExpressionParserMODULUS      = 24
	BooleanExpressionParserEXPONENT     = 25
	BooleanExpressionParserNE           = 26
	BooleanExpressionParserGT           = 27
	BooleanExpressionParserGE           = 28
	BooleanExpressionParserLT           = 29
	BooleanExpressionParserLE           = 30
	BooleanExpressionParserEQ           = 31
	BooleanExpressionParserLPAREN       = 32
	BooleanExpressionParserRPAREN       = 33
	BooleanExpressionParserDECIMAL      = 34
	BooleanExpressionParserAPP_VERSION  = 35
	BooleanExpressionParserINTEGER      = 36
	BooleanExpressionParserWS           = 37
	BooleanExpressionParserWORD         = 38
	BooleanExpressionParserSQSTR        = 39
	BooleanExpressionParserDQSTR        = 40
	BooleanExpressionParserFIELD        = 41
	BooleanExpressionParserALPHANUMERIC = 42
	BooleanExpressionParserSQ           = 43
	BooleanExpressionParserDQ           = 44
)

// BooleanExpressionParser rules.
const (
	BooleanExpressionParserRULE_parse              = 0
	BooleanExpressionParserRULE_expression         = 1
	BooleanExpressionParserRULE_comparator         = 2
	BooleanExpressionParserRULE_arithmeticOperator = 3
	BooleanExpressionParserRULE_arithmeticFunction = 4
	BooleanExpressionParserRULE_wordlist           = 5
	BooleanExpressionParserRULE_arrayOperators     = 6
	BooleanExpressionParserRULE_numericTypes       = 7
	BooleanExpressionParserRULE_types              = 8
	BooleanExpressionParserRULE_binary             = 9
	BooleanExpressionParserRULE_bool               = 10
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *BooleanExpressionParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BooleanExpressionParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.expression(0)
	}
	{
		p.SetState(23)
		p.Match(BooleanExpressionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryArithmeticExpressionContext struct {
	ExpressionContext
	op  antlr.Token
	exp IExpressionContext
}

func NewUnaryArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryArithmeticExpressionContext {
	var p = new(UnaryArithmeticExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryArithmeticExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryArithmeticExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryArithmeticExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *UnaryArithmeticExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *UnaryArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryArithmeticExpressionContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserSUBTRACT, 0)
}

func (s *UnaryArithmeticExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterUnaryArithmeticExpression(s)
	}
}

func (s *UnaryArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitUnaryArithmeticExpression(s)
	}
}

type BinaryExpressionContext struct {
	ExpressionContext
	left  IExpressionContext
	op    IBinaryContext
	right IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryExpressionContext) GetOp() IBinaryContext { return s.op }

func (s *BinaryExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryExpressionContext) SetOp(v IBinaryContext) { s.op = v }

func (s *BinaryExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) Binary() IBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

type TypesExpressionContext struct {
	ExpressionContext
}

func NewTypesExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypesExpressionContext {
	var p = new(TypesExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TypesExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesExpressionContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TypesExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterTypesExpression(s)
	}
}

func (s *TypesExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitTypesExpression(s)
	}
}

type InExpressionContext struct {
	ExpressionContext
	field antlr.Token
	not   antlr.Token
	data  IWordlistContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetField() antlr.Token { return s.field }

func (s *InExpressionContext) GetNot() antlr.Token { return s.not }

func (s *InExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *InExpressionContext) SetNot(v antlr.Token) { s.not = v }

func (s *InExpressionContext) GetData() IWordlistContext { return s.data }

func (s *InExpressionContext) SetData(v IWordlistContext) { s.data = v }

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserIN, 0)
}

func (s *InExpressionContext) Wordlist() IWordlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordlistContext)
}

func (s *InExpressionContext) FIELD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserFIELD, 0)
}

func (s *InExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserNOT, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type ArrayExpressionContext struct {
	ExpressionContext
	field antlr.Token
	op    IArrayOperatorsContext
	data  IWordlistContext
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayExpressionContext) GetField() antlr.Token { return s.field }

func (s *ArrayExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *ArrayExpressionContext) GetOp() IArrayOperatorsContext { return s.op }

func (s *ArrayExpressionContext) GetData() IWordlistContext { return s.data }

func (s *ArrayExpressionContext) SetOp(v IArrayOperatorsContext) { s.op = v }

func (s *ArrayExpressionContext) SetData(v IWordlistContext) { s.data = v }

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) ArrayOperators() IArrayOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayOperatorsContext)
}

func (s *ArrayExpressionContext) Wordlist() IWordlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordlistContext)
}

func (s *ArrayExpressionContext) FIELD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserFIELD, 0)
}

func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

type ToExpressionContext struct {
	ExpressionContext
	field antlr.Token
	lower INumericTypesContext
	upper INumericTypesContext
}

func NewToExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToExpressionContext {
	var p = new(ToExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ToExpressionContext) GetField() antlr.Token { return s.field }

func (s *ToExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *ToExpressionContext) GetLower() INumericTypesContext { return s.lower }

func (s *ToExpressionContext) GetUpper() INumericTypesContext { return s.upper }

func (s *ToExpressionContext) SetLower(v INumericTypesContext) { s.lower = v }

func (s *ToExpressionContext) SetUpper(v INumericTypesContext) { s.upper = v }

func (s *ToExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToExpressionContext) TO() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserTO, 0)
}

func (s *ToExpressionContext) AllNumericTypes() []INumericTypesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericTypesContext); ok {
			len++
		}
	}

	tst := make([]INumericTypesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericTypesContext); ok {
			tst[i] = t.(INumericTypesContext)
			i++
		}
	}

	return tst
}

func (s *ToExpressionContext) NumericTypes(i int) INumericTypesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericTypesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericTypesContext)
}

func (s *ToExpressionContext) FIELD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserFIELD, 0)
}

func (s *ToExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterToExpression(s)
	}
}

func (s *ToExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitToExpression(s)
	}
}

type ArithmeticFunctionExpressionContext struct {
	ExpressionContext
	left IArithmeticFunctionContext
	data IWordlistContext
}

func NewArithmeticFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticFunctionExpressionContext {
	var p = new(ArithmeticFunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArithmeticFunctionExpressionContext) GetLeft() IArithmeticFunctionContext { return s.left }

func (s *ArithmeticFunctionExpressionContext) GetData() IWordlistContext { return s.data }

func (s *ArithmeticFunctionExpressionContext) SetLeft(v IArithmeticFunctionContext) { s.left = v }

func (s *ArithmeticFunctionExpressionContext) SetData(v IWordlistContext) { s.data = v }

func (s *ArithmeticFunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticFunctionExpressionContext) ArithmeticFunction() IArithmeticFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticFunctionContext)
}

func (s *ArithmeticFunctionExpressionContext) Wordlist() IWordlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordlistContext)
}

func (s *ArithmeticFunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArithmeticFunctionExpression(s)
	}
}

func (s *ArithmeticFunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArithmeticFunctionExpression(s)
	}
}

type NotExpressionContext struct {
	ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserNOT, 0)
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type ArithmeticExpressionContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticExpressionContext {
	var p = new(ArithmeticExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArithmeticExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ArithmeticExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ArithmeticExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ArithmeticExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArithmeticExpressionContext) EXPONENT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserEXPONENT, 0)
}

func (s *ArithmeticExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserDIVIDE, 0)
}

func (s *ArithmeticExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMULTIPLY, 0)
}

func (s *ArithmeticExpressionContext) MODULUS() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMODULUS, 0)
}

func (s *ArithmeticExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserADD, 0)
}

func (s *ArithmeticExpressionContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserSUBTRACT, 0)
}

func (s *ArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArithmeticExpression(s)
	}
}

func (s *ArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArithmeticExpression(s)
	}
}

type ComparatorExpressionContext struct {
	ExpressionContext
	left  IExpressionContext
	op    IComparatorContext
	right IExpressionContext
}

func NewComparatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorExpressionContext {
	var p = new(ComparatorExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparatorExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *ComparatorExpressionContext) GetOp() IComparatorContext { return s.op }

func (s *ComparatorExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *ComparatorExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *ComparatorExpressionContext) SetOp(v IComparatorContext) { s.op = v }

func (s *ComparatorExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *ComparatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparatorExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparatorExpressionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ComparatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterComparatorExpression(s)
	}
}

func (s *ComparatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitComparatorExpression(s)
	}
}

type ParentExpressionContext struct {
	ExpressionContext
}

func NewParentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentExpressionContext {
	var p = new(ParentExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLPAREN, 0)
}

func (s *ParentExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParentExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserRPAREN, 0)
}

func (s *ParentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterParentExpression(s)
	}
}

func (s *ParentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitParentExpression(s)
	}
}

func (p *BooleanExpressionParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *BooleanExpressionParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, BooleanExpressionParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(26)
			p.Match(BooleanExpressionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.expression(0)
		}
		{
			p.SetState(28)
			p.Match(BooleanExpressionParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(BooleanExpressionParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.expression(15)
		}

	case 3:
		localctx = NewUnaryArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)

			var _m = p.Match(BooleanExpressionParserSUBTRACT)

			localctx.(*UnaryArithmeticExpressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)

			var _x = p.expression(13)

			localctx.(*UnaryArithmeticExpressionContext).exp = _x
		}

	case 4:
		localctx = NewArithmeticFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)

			var _x = p.ArithmeticFunction()

			localctx.(*ArithmeticFunctionExpressionContext).left = _x
		}
		{
			p.SetState(35)

			var _x = p.Wordlist()

			localctx.(*ArithmeticFunctionExpressionContext).data = _x
		}

	case 5:
		localctx = NewTypesExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Types()
		}

	case 6:
		localctx = NewToExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(38)

			var _m = p.Match(BooleanExpressionParserFIELD)

			localctx.(*ToExpressionContext).field = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(39)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).lower = _x
		}
		{
			p.SetState(40)
			p.Match(BooleanExpressionParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).upper = _x
		}

	case 7:
		localctx = NewInExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(43)

			var _m = p.Match(BooleanExpressionParserFIELD)

			localctx.(*InExpressionContext).field = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BooleanExpressionParserNOT {
			{
				p.SetState(44)

				var _m = p.Match(BooleanExpressionParserNOT)

				localctx.(*InExpressionContext).not = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(47)
			p.Match(BooleanExpressionParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)

			var _x = p.Wordlist()

			localctx.(*InExpressionContext).data = _x
		}

	case 8:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(49)

			var _m = p.Match(BooleanExpressionParserFIELD)

			localctx.(*ArrayExpressionContext).field = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(50)

			var _x = p.ArrayOperators()

			localctx.(*ArrayExpressionContext).op = _x
		}
		{
			p.SetState(51)

			var _x = p.Wordlist()

			localctx.(*ArrayExpressionContext).data = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ComparatorExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(56)

					var _x = p.Comparator()

					localctx.(*ComparatorExpressionContext).op = _x
				}
				{
					p.SetState(57)

					var _x = p.expression(15)

					localctx.(*ComparatorExpressionContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(60)

					var _m = p.Match(BooleanExpressionParserEXPONENT)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(61)

					var _x = p.expression(13)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(63)

					var _m = p.Match(BooleanExpressionParserDIVIDE)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(64)

					var _x = p.expression(12)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 4:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(66)

					var _m = p.Match(BooleanExpressionParserMULTIPLY)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(67)

					var _x = p.expression(11)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 5:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(69)

					var _m = p.Match(BooleanExpressionParserMODULUS)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)

					var _x = p.expression(10)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 6:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(72)

					var _m = p.Match(BooleanExpressionParserADD)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)

					var _x = p.expression(9)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 7:
				localctx = NewArithmeticExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _m = p.Match(BooleanExpressionParserSUBTRACT)

					localctx.(*ArithmeticExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)

					var _x = p.expression(8)

					localctx.(*ArithmeticExpressionContext).right = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _x = p.Binary()

					localctx.(*BinaryExpressionContext).op = _x
				}
				{
					p.SetState(79)

					var _x = p.expression(6)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserGT, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserGE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLE, 0)
}

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserEQ, 0)
}

func (s *ComparatorContext) NE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserNE, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *BooleanExpressionParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BooleanExpressionParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4227858432) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticOperatorContext is an interface to support dynamic dispatch.
type IArithmeticOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULTIPLY() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUBTRACT() antlr.TerminalNode
	MODULUS() antlr.TerminalNode
	EXPONENT() antlr.TerminalNode

	// IsArithmeticOperatorContext differentiates from other interfaces.
	IsArithmeticOperatorContext()
}

type ArithmeticOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticOperatorContext() *ArithmeticOperatorContext {
	var p = new(ArithmeticOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticOperator
	return p
}

func InitEmptyArithmeticOperatorContext(p *ArithmeticOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticOperator
}

func (*ArithmeticOperatorContext) IsArithmeticOperatorContext() {}

func NewArithmeticOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOperatorContext {
	var p = new(ArithmeticOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticOperator

	return p
}

func (s *ArithmeticOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticOperatorContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMULTIPLY, 0)
}

func (s *ArithmeticOperatorContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserDIVIDE, 0)
}

func (s *ArithmeticOperatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserADD, 0)
}

func (s *ArithmeticOperatorContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserSUBTRACT, 0)
}

func (s *ArithmeticOperatorContext) MODULUS() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMODULUS, 0)
}

func (s *ArithmeticOperatorContext) EXPONENT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserEXPONENT, 0)
}

func (s *ArithmeticOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArithmeticOperator(s)
	}
}

func (s *ArithmeticOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArithmeticOperator(s)
	}
}

func (p *BooleanExpressionParser) ArithmeticOperator() (localctx IArithmeticOperatorContext) {
	localctx = NewArithmeticOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BooleanExpressionParserRULE_arithmeticOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66060288) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticFunctionContext is an interface to support dynamic dispatch.
type IArithmeticFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode
	SUM() antlr.TerminalNode
	AVG() antlr.TerminalNode
	MEAN() antlr.TerminalNode
	MODE() antlr.TerminalNode
	LEN() antlr.TerminalNode
	MEDIAN() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsArithmeticFunctionContext differentiates from other interfaces.
	IsArithmeticFunctionContext()
}

type ArithmeticFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticFunctionContext() *ArithmeticFunctionContext {
	var p = new(ArithmeticFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticFunction
	return p
}

func InitEmptyArithmeticFunctionContext(p *ArithmeticFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticFunction
}

func (*ArithmeticFunctionContext) IsArithmeticFunctionContext() {}

func NewArithmeticFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticFunctionContext {
	var p = new(ArithmeticFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_arithmeticFunction

	return p
}

func (s *ArithmeticFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMIN, 0)
}

func (s *ArithmeticFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMAX, 0)
}

func (s *ArithmeticFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserSUM, 0)
}

func (s *ArithmeticFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserAVG, 0)
}

func (s *ArithmeticFunctionContext) MEAN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMEAN, 0)
}

func (s *ArithmeticFunctionContext) MODE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMODE, 0)
}

func (s *ArithmeticFunctionContext) LEN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLEN, 0)
}

func (s *ArithmeticFunctionContext) MEDIAN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserMEDIAN, 0)
}

func (s *ArithmeticFunctionContext) INT() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserINT, 0)
}

func (s *ArithmeticFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArithmeticFunction(s)
	}
}

func (s *ArithmeticFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArithmeticFunction(s)
	}
}

func (p *BooleanExpressionParser) ArithmeticFunction() (localctx IArithmeticFunctionContext) {
	localctx = NewArithmeticFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BooleanExpressionParserRULE_arithmeticFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1046528) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWordlistContext is an interface to support dynamic dispatch.
type IWordlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IExpressionContext

	// GetRest returns the rest rule contexts.
	GetRest() IExpressionContext

	// SetFirst sets the first rule contexts.
	SetFirst(IExpressionContext)

	// SetRest sets the rest rule contexts.
	SetRest(IExpressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsWordlistContext differentiates from other interfaces.
	IsWordlistContext()
}

type WordlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	first  IExpressionContext
	rest   IExpressionContext
}

func NewEmptyWordlistContext() *WordlistContext {
	var p = new(WordlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_wordlist
	return p
}

func InitEmptyWordlistContext(p *WordlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_wordlist
}

func (*WordlistContext) IsWordlistContext() {}

func NewWordlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordlistContext {
	var p = new(WordlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_wordlist

	return p
}

func (s *WordlistContext) GetParser() antlr.Parser { return s.parser }

func (s *WordlistContext) GetFirst() IExpressionContext { return s.first }

func (s *WordlistContext) GetRest() IExpressionContext { return s.rest }

func (s *WordlistContext) SetFirst(v IExpressionContext) { s.first = v }

func (s *WordlistContext) SetRest(v IExpressionContext) { s.rest = v }

func (s *WordlistContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLPAREN, 0)
}

func (s *WordlistContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserRPAREN, 0)
}

func (s *WordlistContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *WordlistContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WordlistContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(BooleanExpressionParserWS)
}

func (s *WordlistContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWS, i)
}

func (s *WordlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterWordlist(s)
	}
}

func (s *WordlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitWordlist(s)
	}
}

func (p *BooleanExpressionParser) Wordlist() (localctx IWordlistContext) {
	localctx = NewWordlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BooleanExpressionParserRULE_wordlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(BooleanExpressionParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(93)
				p.Match(BooleanExpressionParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(99)

		var _x = p.expression(0)

		localctx.(*WordlistContext).first = _x
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserWS {
		{
			p.SetState(100)
			p.Match(BooleanExpressionParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserT__0 {
		{
			p.SetState(106)
			p.Match(BooleanExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(107)
					p.Match(BooleanExpressionParserWS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _x = p.expression(0)

			localctx.(*WordlistContext).rest = _x
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BooleanExpressionParserWS {
			{
				p.SetState(114)
				p.Match(BooleanExpressionParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(BooleanExpressionParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayOperatorsContext is an interface to support dynamic dispatch.
type IArrayOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTAINS_ANY() antlr.TerminalNode
	CONTAINS_ALL() antlr.TerminalNode

	// IsArrayOperatorsContext differentiates from other interfaces.
	IsArrayOperatorsContext()
}

type ArrayOperatorsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayOperatorsContext() *ArrayOperatorsContext {
	var p = new(ArrayOperatorsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arrayOperators
	return p
}

func InitEmptyArrayOperatorsContext(p *ArrayOperatorsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_arrayOperators
}

func (*ArrayOperatorsContext) IsArrayOperatorsContext() {}

func NewArrayOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayOperatorsContext {
	var p = new(ArrayOperatorsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_arrayOperators

	return p
}

func (s *ArrayOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayOperatorsContext) CONTAINS_ANY() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserCONTAINS_ANY, 0)
}

func (s *ArrayOperatorsContext) CONTAINS_ALL() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserCONTAINS_ALL, 0)
}

func (s *ArrayOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterArrayOperators(s)
	}
}

func (s *ArrayOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitArrayOperators(s)
	}
}

func (p *BooleanExpressionParser) ArrayOperators() (localctx IArrayOperatorsContext) {
	localctx = NewArrayOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BooleanExpressionParserRULE_arrayOperators)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserCONTAINS_ALL || _la == BooleanExpressionParserCONTAINS_ANY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericTypesContext is an interface to support dynamic dispatch.
type INumericTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode

	// IsNumericTypesContext differentiates from other interfaces.
	IsNumericTypesContext()
}

type NumericTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericTypesContext() *NumericTypesContext {
	var p = new(NumericTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_numericTypes
	return p
}

func InitEmptyNumericTypesContext(p *NumericTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_numericTypes
}

func (*NumericTypesContext) IsNumericTypesContext() {}

func NewNumericTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericTypesContext {
	var p = new(NumericTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_numericTypes

	return p
}

func (s *NumericTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericTypesContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserINTEGER, 0)
}

func (s *NumericTypesContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserDECIMAL, 0)
}

func (s *NumericTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterNumericTypes(s)
	}
}

func (s *NumericTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitNumericTypes(s)
	}
}

func (p *BooleanExpressionParser) NumericTypes() (localctx INumericTypesContext) {
	localctx = NewNumericTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BooleanExpressionParserRULE_numericTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserDECIMAL || _la == BooleanExpressionParserINTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	APP_VERSION() antlr.TerminalNode
	Bool_() IBoolContext
	WORD() antlr.TerminalNode
	FIELD() antlr.TerminalNode

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserINTEGER, 0)
}

func (s *TypesContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserDECIMAL, 0)
}

func (s *TypesContext) APP_VERSION() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserAPP_VERSION, 0)
}

func (s *TypesContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *TypesContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
}

func (s *TypesContext) FIELD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserFIELD, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *BooleanExpressionParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BooleanExpressionParserRULE_types)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(BooleanExpressionParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(BooleanExpressionParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(BooleanExpressionParserAPP_VERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.Match(BooleanExpressionParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(136)
			p.Match(BooleanExpressionParserFIELD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_binary
	return p
}

func InitEmptyBinaryContext(p *BinaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_binary
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserOR, 0)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *BooleanExpressionParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BooleanExpressionParserRULE_binary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserAND || _la == BooleanExpressionParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_bool
	return p
}

func InitEmptyBoolContext(p *BoolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_bool
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserTRUE, 0)
}

func (s *BoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserFALSE, 0)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BooleanExpressionListener); ok {
		listenerT.ExitBool(s)
	}
}

func (p *BooleanExpressionParser) Bool_() (localctx IBoolContext) {
	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BooleanExpressionParserRULE_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserTRUE || _la == BooleanExpressionParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *BooleanExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *BooleanExpressionParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
