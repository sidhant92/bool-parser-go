// Code generated from /Users/sid/Desktop/filter1/BooleanExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"parse", "expression", "comparator", "wordlist", "arrayOperators", "numericTypes",
		"types", "binary", "bool",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 1, 3, 1, 41, 8, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 46, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 61, 8, 1, 10, 1, 12, 1, 64, 9,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 70, 8, 3, 10, 3, 12, 3, 73, 9, 3, 1, 3,
		1, 3, 5, 3, 77, 8, 3, 10, 3, 12, 3, 80, 9, 3, 1, 3, 1, 3, 5, 3, 84, 8,
		3, 10, 3, 12, 3, 87, 9, 3, 1, 3, 1, 3, 5, 3, 91, 8, 3, 10, 3, 12, 3, 94,
		9, 3, 5, 3, 96, 8, 3, 10, 3, 12, 3, 99, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 113, 8, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 0, 1, 2, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 5, 1,
		0, 11, 16, 1, 0, 9, 10, 2, 0, 19, 19, 21, 21, 1, 0, 4, 5, 1, 0, 7, 8, 130,
		0, 18, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 65, 1, 0, 0, 0, 6, 67, 1, 0, 0,
		0, 8, 102, 1, 0, 0, 0, 10, 104, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14, 114,
		1, 0, 0, 0, 16, 116, 1, 0, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20, 5, 0, 0, 1,
		20, 1, 1, 0, 0, 0, 21, 22, 6, 1, -1, 0, 22, 23, 5, 17, 0, 0, 23, 24, 3,
		2, 1, 0, 24, 25, 5, 18, 0, 0, 25, 51, 1, 0, 0, 0, 26, 27, 5, 6, 0, 0, 27,
		51, 3, 2, 1, 7, 28, 51, 3, 12, 6, 0, 29, 31, 5, 23, 0, 0, 30, 29, 1, 0,
		0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 3, 10, 5, 0, 33,
		34, 5, 3, 0, 0, 34, 35, 3, 10, 5, 0, 35, 51, 1, 0, 0, 0, 36, 38, 5, 23,
		0, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 41,
		5, 6, 0, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		42, 43, 5, 2, 0, 0, 43, 51, 3, 6, 3, 0, 44, 46, 5, 23, 0, 0, 45, 44, 1,
		0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 3, 8, 4, 0, 48,
		49, 3, 6, 3, 0, 49, 51, 1, 0, 0, 0, 50, 21, 1, 0, 0, 0, 50, 26, 1, 0, 0,
		0, 50, 28, 1, 0, 0, 0, 50, 30, 1, 0, 0, 0, 50, 37, 1, 0, 0, 0, 50, 45,
		1, 0, 0, 0, 51, 62, 1, 0, 0, 0, 52, 53, 10, 6, 0, 0, 53, 54, 3, 4, 2, 0,
		54, 55, 3, 2, 1, 7, 55, 61, 1, 0, 0, 0, 56, 57, 10, 5, 0, 0, 57, 58, 3,
		14, 7, 0, 58, 59, 3, 2, 1, 6, 59, 61, 1, 0, 0, 0, 60, 52, 1, 0, 0, 0, 60,
		56, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0,
		0, 63, 3, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 7, 0, 0, 0, 66, 5, 1,
		0, 0, 0, 67, 71, 5, 17, 0, 0, 68, 70, 5, 22, 0, 0, 69, 68, 1, 0, 0, 0,
		70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 78, 3, 12, 6, 0, 75, 77, 5, 22, 0, 0,
		76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1,
		0, 0, 0, 79, 97, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 85, 5, 1, 0, 0, 82,
		84, 5, 22, 0, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0,
		0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 92,
		3, 12, 6, 0, 89, 91, 5, 22, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92,
		1, 0, 0, 0, 95, 81, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 101,
		5, 18, 0, 0, 101, 7, 1, 0, 0, 0, 102, 103, 7, 1, 0, 0, 103, 9, 1, 0, 0,
		0, 104, 105, 7, 2, 0, 0, 105, 11, 1, 0, 0, 0, 106, 113, 5, 21, 0, 0, 107,
		113, 5, 19, 0, 0, 108, 113, 5, 20, 0, 0, 109, 113, 3, 16, 8, 0, 110, 113,
		5, 23, 0, 0, 111, 113, 1, 0, 0, 0, 112, 106, 1, 0, 0, 0, 112, 107, 1, 0,
		0, 0, 112, 108, 1, 0, 0, 0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0,
		112, 111, 1, 0, 0, 0, 113, 13, 1, 0, 0, 0, 114, 115, 7, 3, 0, 0, 115, 15,
		1, 0, 0, 0, 116, 117, 7, 4, 0, 0, 117, 17, 1, 0, 0, 0, 13, 30, 37, 40,
		45, 50, 60, 62, 71, 78, 85, 92, 97, 112,
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
	BooleanExpressionParserNE           = 11
	BooleanExpressionParserGT           = 12
	BooleanExpressionParserGE           = 13
	BooleanExpressionParserLT           = 14
	BooleanExpressionParserLE           = 15
	BooleanExpressionParserEQ           = 16
	BooleanExpressionParserLPAREN       = 17
	BooleanExpressionParserRPAREN       = 18
	BooleanExpressionParserDECIMAL      = 19
	BooleanExpressionParserAPP_VERSION  = 20
	BooleanExpressionParserINTEGER      = 21
	BooleanExpressionParserWS           = 22
	BooleanExpressionParserWORD         = 23
	BooleanExpressionParserALPHANUMERIC = 24
	BooleanExpressionParserSQ           = 25
	BooleanExpressionParserDQ           = 26
)

// BooleanExpressionParser rules.
const (
	BooleanExpressionParserRULE_parse          = 0
	BooleanExpressionParserRULE_expression     = 1
	BooleanExpressionParserRULE_comparator     = 2
	BooleanExpressionParserRULE_wordlist       = 3
	BooleanExpressionParserRULE_arrayOperators = 4
	BooleanExpressionParserRULE_numericTypes   = 5
	BooleanExpressionParserRULE_types          = 6
	BooleanExpressionParserRULE_binary         = 7
	BooleanExpressionParserRULE_bool           = 8
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
		p.SetState(18)
		p.expression(0)
	}
	{
		p.SetState(19)
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

func (s *InExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
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

func (s *ArrayExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
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

func (s *ToExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(22)
			p.Match(BooleanExpressionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.expression(0)
		}
		{
			p.SetState(24)
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
			p.SetState(26)
			p.Match(BooleanExpressionParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.expression(7)
		}

	case 3:
		localctx = NewTypesExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(28)
			p.Types()
		}

	case 4:
		localctx = NewToExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BooleanExpressionParserWORD {
			{
				p.SetState(29)

				var _m = p.Match(BooleanExpressionParserWORD)

				localctx.(*ToExpressionContext).field = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(32)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).lower = _x
		}
		{
			p.SetState(33)
			p.Match(BooleanExpressionParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).upper = _x
		}

	case 5:
		localctx = NewInExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BooleanExpressionParserWORD {
			{
				p.SetState(36)

				var _m = p.Match(BooleanExpressionParserWORD)

				localctx.(*InExpressionContext).field = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BooleanExpressionParserNOT {
			{
				p.SetState(39)

				var _m = p.Match(BooleanExpressionParserNOT)

				localctx.(*InExpressionContext).not = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(42)
			p.Match(BooleanExpressionParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)

			var _x = p.Wordlist()

			localctx.(*InExpressionContext).data = _x
		}

	case 6:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == BooleanExpressionParserWORD {
			{
				p.SetState(44)

				var _m = p.Match(BooleanExpressionParserWORD)

				localctx.(*ArrayExpressionContext).field = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(47)

			var _x = p.ArrayOperators()

			localctx.(*ArrayExpressionContext).op = _x
		}
		{
			p.SetState(48)

			var _x = p.Wordlist()

			localctx.(*ArrayExpressionContext).data = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(62)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ComparatorExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(53)

					var _x = p.Comparator()

					localctx.(*ComparatorExpressionContext).op = _x
				}
				{
					p.SetState(54)

					var _x = p.expression(7)

					localctx.(*ComparatorExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(57)

					var _x = p.Binary()

					localctx.(*BinaryExpressionContext).op = _x
				}
				{
					p.SetState(58)

					var _x = p.expression(6)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&129024) != 0) {
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
	GetFirst() ITypesContext

	// GetRest returns the rest rule contexts.
	GetRest() ITypesContext

	// SetFirst sets the first rule contexts.
	SetFirst(ITypesContext)

	// SetRest sets the rest rule contexts.
	SetRest(ITypesContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllTypes() []ITypesContext
	Types(i int) ITypesContext
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsWordlistContext differentiates from other interfaces.
	IsWordlistContext()
}

type WordlistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	first  ITypesContext
	rest   ITypesContext
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

func (s *WordlistContext) GetFirst() ITypesContext { return s.first }

func (s *WordlistContext) GetRest() ITypesContext { return s.rest }

func (s *WordlistContext) SetFirst(v ITypesContext) { s.first = v }

func (s *WordlistContext) SetRest(v ITypesContext) { s.rest = v }

func (s *WordlistContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserLPAREN, 0)
}

func (s *WordlistContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserRPAREN, 0)
}

func (s *WordlistContext) AllTypes() []ITypesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypesContext); ok {
			len++
		}
	}

	tst := make([]ITypesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypesContext); ok {
			tst[i] = t.(ITypesContext)
			i++
		}
	}

	return tst
}

func (s *WordlistContext) Types(i int) ITypesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
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

	return t.(ITypesContext)
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
	p.EnterRule(localctx, 6, BooleanExpressionParserRULE_wordlist)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(BooleanExpressionParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(68)
				p.Match(BooleanExpressionParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(74)

		var _x = p.Types()

		localctx.(*WordlistContext).first = _x
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserWS {
		{
			p.SetState(75)
			p.Match(BooleanExpressionParserWS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserT__0 {
		{
			p.SetState(81)
			p.Match(BooleanExpressionParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(82)
					p.Match(BooleanExpressionParserWS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(88)

			var _x = p.Types()

			localctx.(*WordlistContext).rest = _x
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BooleanExpressionParserWS {
			{
				p.SetState(89)
				p.Match(BooleanExpressionParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
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
	p.EnterRule(localctx, 8, BooleanExpressionParserRULE_arrayOperators)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
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
	p.EnterRule(localctx, 10, BooleanExpressionParserRULE_numericTypes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
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
	p.EnterRule(localctx, 12, BooleanExpressionParserRULE_types)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(BooleanExpressionParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(BooleanExpressionParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(BooleanExpressionParserAPP_VERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(110)
			p.Match(BooleanExpressionParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)

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
	p.EnterRule(localctx, 14, BooleanExpressionParserRULE_binary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
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
	p.EnterRule(localctx, 16, BooleanExpressionParserRULE_bool)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
