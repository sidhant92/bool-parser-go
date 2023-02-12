package parser

import (
	"bool-parser-go/internal/parser/antlr/lib"
	"bool-parser-go/internal/service"
	"bool-parser-go/internal/util"
	"bool-parser-go/pkg/constant"
	"bool-parser-go/pkg/domain"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"reflect"
)

type myListener struct {
	Nodes []domain.Node

	LastToken antlr.Token

	Result domain.Node
}

func New() *myListener {
	l := myListener{
		Nodes:     []domain.Node{},
		LastToken: nil,
		Result:    nil,
	}
	return &l
}

func (l *myListener) push(node domain.Node) {
	l.Nodes = append(l.Nodes, node)
}

func (l *myListener) size() int {
	return len(l.Nodes)
}

func (l *myListener) pop() domain.Node {
	if l.size() < 1 {
		panic("Invalid Expression")
	}
	result := l.Nodes[l.size()-1]
	l.Nodes = l.Nodes[:l.size()-1]
	return result
}

func (l *myListener) GetResult() domain.Node {
	return l.Result
}

func (l myListener) VisitTerminal(node antlr.TerminalNode) {
}

func (l myListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (l myListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (l myListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (l myListener) EnterParse(c *lib.ParseContext) {
}

func (l myListener) EnterBinaryExpression(c *lib.BinaryExpressionContext) {
}

func (l myListener) EnterTypesExpression(c *lib.TypesExpressionContext) {
}

func (l myListener) EnterInExpression(c *lib.InExpressionContext) {
}

func (l myListener) EnterToExpression(c *lib.ToExpressionContext) {
}

func (l myListener) EnterNotExpression(c *lib.NotExpressionContext) {
}

func (l myListener) EnterComparatorExpression(c *lib.ComparatorExpressionContext) {
}

func (l myListener) EnterParentExpression(c *lib.ParentExpressionContext) {
}

func (l myListener) EnterComparator(c *lib.ComparatorContext) {
}

func (l myListener) EnterWordlist(c *lib.WordlistContext) {
}

func (l myListener) EnterNumericTypes(c *lib.NumericTypesContext) {
}

func (l myListener) EnterTypes(c *lib.TypesContext) {
}

func (l myListener) EnterBinary(c *lib.BinaryContext) {
}

func (l myListener) EnterBool(c *lib.BoolContext) {
}

func (l *myListener) ExitParse(c *lib.ParseContext) {
	if l.Result == nil && l.size() == 1 {
		l.Result = l.pop()
	} else {
		panic("Error parsing expression for the string " + c.GetText())
	}
}

func (l *myListener) ExitBinaryExpression(c *lib.BinaryExpressionContext) {
	if l.size() < 2 {
		panic("Error parsing binary expression for the string " + c.GetText())
	}
	firstNode := l.pop()
	secondNode := l.pop()
	operator := GetLogicalOperator(c.GetOp().GetStart())
	l.push(domain.BooleanNode{
		Left:     secondNode,
		Right:    firstNode,
		Operator: operator,
	})
}

func (l *myListener) ExitTypesExpression(c *lib.TypesExpressionContext) {
	l.LastToken = c.GetStart()
}

func (l *myListener) ExitInExpression(c *lib.InExpressionContext) {
	field := c.GetField().GetText()

	typesContextFilter := func(tree antlr.Tree) bool { return reflect.TypeOf(tree) == reflect.TypeOf(&lib.TypesContext{}) }
	var typesContextChildren = util.Filter(c.GetData().GetChildren(), typesContextFilter)
	var pairs []domain.Pair

	for _, child := range typesContextChildren {
		dataType := GetDataType(child.(*lib.TypesContext).GetStart())
		value := util.ConvertValue(child.(*lib.TypesContext).GetText(), dataType)
		pairs = append(pairs, domain.Pair{
			DataType: dataType,
			Value:    value,
		})
	}

	l.push(domain.InNode{
		Field: field,
		Items: pairs,
	})
}

func (l *myListener) ExitToExpression(c *lib.ToExpressionContext) {
	field := c.GetField().GetText()
	lowerDataType := GetDataType(c.GetLower().GetStart())
	lowerValue := util.ConvertValue(c.GetLower().GetStart().GetText(), lowerDataType)
	upperDataType := GetDataType(c.GetUpper().GetStart())
	upperValue := util.ConvertValue(c.GetUpper().GetText(), upperDataType)
	l.push(domain.NumericRangeNode{
		Field:        field,
		FromValue:    lowerValue,
		ToValue:      upperValue,
		FromDataType: lowerDataType,
		ToDataType:   upperDataType,
	})
}

func (l *myListener) ExitNotExpression(c *lib.NotExpressionContext) {
	if l.size() < 1 {
		if l.LastToken == nil {
			panic("Error parsing not expression for the string " + c.GetText())
		}
		dataType := GetDataType(l.LastToken)
		value := util.ConvertValue(l.LastToken.GetText(), dataType)
		l.push(domain.UnaryNode{
			DataType: dataType,
			Value:    value,
		})
	}
	boolNode := domain.BooleanNode{
		Left:     l.pop(),
		Right:    nil,
		Operator: constant.NOT,
	}
	l.push(boolNode)
}

func (l *myListener) ExitComparatorExpression(c *lib.ComparatorExpressionContext) {
	field := c.GetLeft().GetText()
	dataType := GetDataType(c.GetRight().GetStart())
	value := util.ConvertValue(c.GetRight().GetText(), dataType)
	operator := service.GetOperatorFromSymbol(c.GetOp().GetText())
	l.push(domain.ComparisonNode{
		Field:    field,
		Value:    value,
		Operator: operator,
		DataType: dataType,
	})
}

func (l myListener) ExitParentExpression(c *lib.ParentExpressionContext) {
}

func (l myListener) ExitComparator(c *lib.ComparatorContext) {
}

func (l myListener) ExitWordlist(c *lib.WordlistContext) {
}

func (l myListener) ExitNumericTypes(c *lib.NumericTypesContext) {
}

func (l myListener) ExitTypes(c *lib.TypesContext) {
}

func (l myListener) ExitBinary(c *lib.BinaryContext) {
}

func (l myListener) ExitBool(c *lib.BoolContext) {
}

func GetLogicalOperator(token antlr.Token) constant.LogicalOperatorType {
	switch token.GetTokenType() {
	case lib.BooleanExpressionLexerAND:
		return constant.AND
	case lib.BooleanExpressionLexerOR:
		return constant.OR
	default:
		return constant.NOT
	}
}

func GetDataType(token antlr.Token) constant.DataType {
	switch token.GetTokenType() {
	case lib.BooleanExpressionLexerDECIMAL:
		return constant.DECIMAL
	case lib.BooleanExpressionLexerINTEGER:
		return util.GetNumericDataType(token.GetText())
	case lib.BooleanExpressionLexerAPP_VERSION:
		return constant.APP_VERSION
	case lib.BooleanExpressionLexerTRUE:
		return constant.BOOLEAN
	case lib.BooleanExpressionLexerFALSE:
		return constant.BOOLEAN
	default:
		return constant.STRING
	}
}
