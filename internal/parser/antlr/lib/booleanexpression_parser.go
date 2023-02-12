// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lib // BooleanExpression

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BooleanExpressionParser struct {
	*antlr.BaseParser
}

var booleanexpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func booleanexpressionParserInit() {
	staticData := &booleanexpressionParserStaticData
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
		"parse", "expression", "comparator", "wordlist", "numericTypes", "types",
		"binary", "bool",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 102, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 46, 8, 1, 10, 1, 12, 1, 49, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 55,
		8, 3, 10, 3, 12, 3, 58, 9, 3, 1, 3, 1, 3, 5, 3, 62, 8, 3, 10, 3, 12, 3,
		65, 9, 3, 1, 3, 1, 3, 5, 3, 69, 8, 3, 10, 3, 12, 3, 72, 9, 3, 1, 3, 1,
		3, 5, 3, 76, 8, 3, 10, 3, 12, 3, 79, 9, 3, 5, 3, 81, 8, 3, 10, 3, 12, 3,
		84, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 96, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 1, 2, 8, 0, 2, 4, 6, 8, 10,
		12, 14, 0, 4, 1, 0, 9, 14, 2, 0, 17, 17, 19, 19, 1, 0, 4, 5, 1, 0, 7, 8,
		109, 0, 16, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6, 52, 1,
		0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 95, 1, 0, 0, 0, 12, 97, 1, 0, 0, 0, 14,
		99, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0,
		0, 19, 20, 6, 1, -1, 0, 20, 21, 5, 15, 0, 0, 21, 22, 3, 2, 1, 0, 22, 23,
		5, 16, 0, 0, 23, 36, 1, 0, 0, 0, 24, 25, 5, 6, 0, 0, 25, 36, 3, 2, 1, 6,
		26, 36, 3, 10, 5, 0, 27, 28, 5, 21, 0, 0, 28, 29, 3, 8, 4, 0, 29, 30, 5,
		3, 0, 0, 30, 31, 3, 8, 4, 0, 31, 36, 1, 0, 0, 0, 32, 33, 5, 21, 0, 0, 33,
		34, 5, 2, 0, 0, 34, 36, 3, 6, 3, 0, 35, 19, 1, 0, 0, 0, 35, 24, 1, 0, 0,
		0, 35, 26, 1, 0, 0, 0, 35, 27, 1, 0, 0, 0, 35, 32, 1, 0, 0, 0, 36, 47,
		1, 0, 0, 0, 37, 38, 10, 5, 0, 0, 38, 39, 3, 4, 2, 0, 39, 40, 3, 2, 1, 6,
		40, 46, 1, 0, 0, 0, 41, 42, 10, 4, 0, 0, 42, 43, 3, 12, 6, 0, 43, 44, 3,
		2, 1, 5, 44, 46, 1, 0, 0, 0, 45, 37, 1, 0, 0, 0, 45, 41, 1, 0, 0, 0, 46,
		49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 3, 1, 0, 0,
		0, 49, 47, 1, 0, 0, 0, 50, 51, 7, 0, 0, 0, 51, 5, 1, 0, 0, 0, 52, 56, 5,
		15, 0, 0, 53, 55, 5, 20, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0,
		56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1,
		0, 0, 0, 59, 63, 3, 10, 5, 0, 60, 62, 5, 20, 0, 0, 61, 60, 1, 0, 0, 0,
		62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 82, 1,
		0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 70, 5, 1, 0, 0, 67, 69, 5, 20, 0, 0, 68,
		67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0,
		0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 77, 3, 10, 5, 0, 74, 76,
		5, 20, 0, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		77, 78, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 66, 1,
		0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83,
		85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 16, 0, 0, 86, 7, 1, 0, 0,
		0, 87, 88, 7, 1, 0, 0, 88, 9, 1, 0, 0, 0, 89, 96, 5, 19, 0, 0, 90, 96,
		5, 17, 0, 0, 91, 96, 5, 18, 0, 0, 92, 96, 3, 14, 7, 0, 93, 96, 5, 21, 0,
		0, 94, 96, 1, 0, 0, 0, 95, 89, 1, 0, 0, 0, 95, 90, 1, 0, 0, 0, 95, 91,
		1, 0, 0, 0, 95, 92, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0,
		96, 11, 1, 0, 0, 0, 97, 98, 7, 2, 0, 0, 98, 13, 1, 0, 0, 0, 99, 100, 7,
		3, 0, 0, 100, 15, 1, 0, 0, 0, 9, 35, 45, 47, 56, 63, 70, 77, 82, 95,
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
	staticData := &booleanexpressionParserStaticData
	staticData.once.Do(booleanexpressionParserInit)
}

// NewBooleanExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBooleanExpressionParser(input antlr.TokenStream) *BooleanExpressionParser {
	BooleanExpressionParserInit()
	this := new(BooleanExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &booleanexpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

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
	BooleanExpressionParserNE           = 9
	BooleanExpressionParserGT           = 10
	BooleanExpressionParserGE           = 11
	BooleanExpressionParserLT           = 12
	BooleanExpressionParserLE           = 13
	BooleanExpressionParserEQ           = 14
	BooleanExpressionParserLPAREN       = 15
	BooleanExpressionParserRPAREN       = 16
	BooleanExpressionParserDECIMAL      = 17
	BooleanExpressionParserAPP_VERSION  = 18
	BooleanExpressionParserINTEGER      = 19
	BooleanExpressionParserWS           = 20
	BooleanExpressionParserWORD         = 21
	BooleanExpressionParserALPHANUMERIC = 22
	BooleanExpressionParserSQ           = 23
	BooleanExpressionParserDQ           = 24
)

// BooleanExpressionParser rules.
const (
	BooleanExpressionParserRULE_parse        = 0
	BooleanExpressionParserRULE_expression   = 1
	BooleanExpressionParserRULE_comparator   = 2
	BooleanExpressionParserRULE_wordlist     = 3
	BooleanExpressionParserRULE_numericTypes = 4
	BooleanExpressionParserRULE_types        = 5
	BooleanExpressionParserRULE_binary       = 6
	BooleanExpressionParserRULE_bool         = 7
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BooleanExpressionParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expression(0)
	}
	{
		p.SetState(17)
		p.Match(BooleanExpressionParserEOF)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BooleanExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    IBinaryContext
	right IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewTypesExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypesExpressionContext {
	var p = new(TypesExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
	field antlr.Token
	data  IWordlistContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetField() antlr.Token { return s.field }

func (s *InExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *InExpressionContext) GetData() IWordlistContext { return s.data }

func (s *InExpressionContext) SetData(v IWordlistContext) { s.data = v }

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserIN, 0)
}

func (s *InExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
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

type ToExpressionContext struct {
	*ExpressionContext
	field antlr.Token
	lower INumericTypesContext
	upper INumericTypesContext
}

func NewToExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToExpressionContext {
	var p = new(ToExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

func (s *ToExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(BooleanExpressionParserWORD, 0)
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
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
	left  IExpressionContext
	op    IComparatorContext
	right IExpressionContext
}

func NewComparatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorExpressionContext {
	var p = new(ComparatorExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	*ExpressionContext
}

func NewParentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentExpressionContext {
	var p = new(ParentExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, BooleanExpressionParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(20)
			p.Match(BooleanExpressionParserLPAREN)
		}
		{
			p.SetState(21)
			p.expression(0)
		}
		{
			p.SetState(22)
			p.Match(BooleanExpressionParserRPAREN)
		}

	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(BooleanExpressionParserNOT)
		}
		{
			p.SetState(25)
			p.expression(6)
		}

	case 3:
		localctx = NewTypesExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Types()
		}

	case 4:
		localctx = NewToExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)

			var _m = p.Match(BooleanExpressionParserWORD)

			localctx.(*ToExpressionContext).field = _m
		}
		{
			p.SetState(28)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).lower = _x
		}
		{
			p.SetState(29)
			p.Match(BooleanExpressionParserTO)
		}
		{
			p.SetState(30)

			var _x = p.NumericTypes()

			localctx.(*ToExpressionContext).upper = _x
		}

	case 5:
		localctx = NewInExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)

			var _m = p.Match(BooleanExpressionParserWORD)

			localctx.(*InExpressionContext).field = _m
		}
		{
			p.SetState(33)
			p.Match(BooleanExpressionParserIN)
		}
		{
			p.SetState(34)

			var _x = p.Wordlist()

			localctx.(*InExpressionContext).data = _x
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparatorExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*ComparatorExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(38)

					var _x = p.Comparator()

					localctx.(*ComparatorExpressionContext).op = _x
				}
				{
					p.SetState(39)

					var _x = p.expression(6)

					localctx.(*ComparatorExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, BooleanExpressionParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(42)

					var _x = p.Binary()

					localctx.(*BinaryExpressionContext).op = _x
				}
				{
					p.SetState(43)

					var _x = p.expression(5)

					localctx.(*BinaryExpressionContext).right = _x
				}

			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BooleanExpressionParserRULE_comparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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

	// IsWordlistContext differentiates from other interfaces.
	IsWordlistContext()
}

type WordlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	first  ITypesContext
	rest   ITypesContext
}

func NewEmptyWordlistContext() *WordlistContext {
	var p = new(WordlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_wordlist
	return p
}

func (*WordlistContext) IsWordlistContext() {}

func NewWordlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordlistContext {
	var p = new(WordlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewWordlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BooleanExpressionParserRULE_wordlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(BooleanExpressionParserLPAREN)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(53)
				p.Match(BooleanExpressionParserWS)
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(59)

		var _x = p.Types()

		localctx.(*WordlistContext).first = _x
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserWS {
		{
			p.SetState(60)
			p.Match(BooleanExpressionParserWS)
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BooleanExpressionParserT__0 {
		{
			p.SetState(66)
			p.Match(BooleanExpressionParserT__0)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(67)
					p.Match(BooleanExpressionParserWS)
				}

			}
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		{
			p.SetState(73)

			var _x = p.Types()

			localctx.(*WordlistContext).rest = _x
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == BooleanExpressionParserWS {
			{
				p.SetState(74)
				p.Match(BooleanExpressionParserWS)
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(BooleanExpressionParserRPAREN)
	}

	return localctx
}

// INumericTypesContext is an interface to support dynamic dispatch.
type INumericTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericTypesContext differentiates from other interfaces.
	IsNumericTypesContext()
}

type NumericTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericTypesContext() *NumericTypesContext {
	var p = new(NumericTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_numericTypes
	return p
}

func (*NumericTypesContext) IsNumericTypesContext() {}

func NewNumericTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericTypesContext {
	var p = new(NumericTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewNumericTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BooleanExpressionParserRULE_numericTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserDECIMAL || _la == BooleanExpressionParserINTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BooleanExpressionParserRULE_types)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Match(BooleanExpressionParserINTEGER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(BooleanExpressionParserDECIMAL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(BooleanExpressionParserAPP_VERSION)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)
			p.Match(BooleanExpressionParserWORD)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)

	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BooleanExpressionParserRULE_binary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserAND || _la == BooleanExpressionParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BooleanExpressionParserRULE_bool
	return p
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BooleanExpressionParserRULE_bool)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BooleanExpressionParserTRUE || _la == BooleanExpressionParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
