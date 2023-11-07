package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/sidhant92/bool-parser-go/internal/parser/antlr/lib"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"reflect"
)

type myListener struct {
	Nodes []domain.Node

	LastToken antlr.Token

	Result domain.Node

	OperatorService *service.OperatorService

	DefaultField string
}

func New(defaultField string) *myListener {
	l := myListener{
		Nodes:           []domain.Node{},
		LastToken:       nil,
		Result:          nil,
		OperatorService: service.NewOperatorService(),
		DefaultField:    defaultField,
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
	l.ValidateField(c.GetField(), c.GetText())
	field := l.GetField(c.GetField().GetText())

	pairs := l.GetArrayElements(c.GetData().GetChildren())
	var inNode = domain.InNode{
		Field: field,
		Items: pairs,
	}
	if c.NOT() == nil {
		l.push(inNode)
	} else {
		var booleanNode = domain.BooleanNode{
			Left:     inNode,
			Right:    nil,
			Operator: constant.NOT,
		}
		l.push(booleanNode)
	}
}

func (l *myListener) ExitArrayExpression(ctx *lib.ArrayExpressionContext) {
	l.ValidateField(ctx.GetField(), ctx.GetText())
	field := l.GetField(ctx.GetField().GetText())
	pairs := l.GetArrayElements(ctx.GetData().GetChildren())
	operator := l.OperatorService.GetOperatorFromSymbol(ctx.GetOp().GetText())
	l.push(domain.ArrayNode{
		Field:    field,
		Operator: operator,
		Items:    pairs,
	})
}

func (l *myListener) ExitToExpression(c *lib.ToExpressionContext) {
	l.ValidateField(c.GetField(), c.GetText())
	field := l.GetField(c.GetField().GetText())
	lowerDataType := GetDataType(c.GetLower().GetStart())
	lowerValue, _ := util.ConvertValue(c.GetLower().GetStart().GetText(), lowerDataType)
	upperDataType := GetDataType(c.GetUpper().GetStart())
	upperValue, _ := util.ConvertValue(c.GetUpper().GetText(), upperDataType)
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
		value, _ := util.ConvertValue(l.LastToken.GetText(), dataType)
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
	field := l.GetField(c.GetLeft().GetText())
	dataType := GetDataType(c.GetRight().GetStart())
	value, _ := util.ConvertValue(c.GetRight().GetText(), dataType)
	operator := l.OperatorService.GetOperatorFromSymbol(c.GetOp().GetText())
	l.push(domain.ComparisonNode{
		Field:    field,
		Value:    value,
		Operator: operator,
		DataType: dataType,
	})
}

func (l *myListener) EnterArrayExpression(ctx *lib.ArrayExpressionContext) {}

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

func (l myListener) EnterArrayOperators(ctx *lib.ArrayOperatorsContext) {}

func (l myListener) ExitArrayOperators(ctx *lib.ArrayOperatorsContext) {}

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
		return constant.VERSION
	case lib.BooleanExpressionLexerTRUE:
		return constant.BOOLEAN
	case lib.BooleanExpressionLexerFALSE:
		return constant.BOOLEAN
	default:
		return constant.STRING
	}
}

func (l myListener) ValidateField(token antlr.Token, text string) {
	if token == nil || (len(token.GetText()) == 0 && len(text) == 0) {
		panic("Invalid Expression, field and default value both are empty")
	}
}

func (l myListener) GetField(field string) string {
	if len(field) == 0 {
		return l.DefaultField
	}
	return field
}

func (l myListener) GetArrayElements(trees []antlr.Tree) []domain.Pair {
	typesContextFilter := func(tree antlr.Tree) bool { return reflect.TypeOf(tree) == reflect.TypeOf(&lib.TypesContext{}) }
	var typesContextChildren = util.Filter(trees, typesContextFilter)
	var pairs []domain.Pair

	for _, child := range typesContextChildren {
		dataType := GetDataType(child.(*lib.TypesContext).GetStart())
		value, _ := util.ConvertValue(child.(*lib.TypesContext).GetText(), dataType)
		pairs = append(pairs, domain.Pair{
			DataType: dataType,
			Value:    value,
		})
	}

	return pairs
}
