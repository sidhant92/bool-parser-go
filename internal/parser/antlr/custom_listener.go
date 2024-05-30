package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/sidhant92/bool-parser-go/internal/parser/antlr/lib"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain/arithmetic"
	"github.com/sidhant92/bool-parser-go/pkg/domain/logical"
	"reflect"
)

type CustomListener struct {
	BaseBooleanExpressionListener,

	Nodes []logical.Node

	LastToken antlr.Token

	Result logical.Node

	OperatorService *service.OperatorService

	FunctionEvaluatorService *service.FunctionEvaluatorService

	TokenCount uint8

	ErrorCount int

	DefaultField string
}

func New(defaultField string, errors int) *CustomListener {
	l := CustomListener{
		Nodes:                    []logical.Node{},
		LastToken:                nil,
		Result:                   nil,
		OperatorService:          service.NewOperatorService(),
		FunctionEvaluatorService: service.NewFunctionEvaluatorService(),
		TokenCount:               0,
		ErrorCount:               errors,
		DefaultField:             defaultField,
	}
	return &l
}

func (l *CustomListener) push(node logical.Node) {
	l.Nodes = append(l.Nodes, node)
}

func (l *CustomListener) size() int {
	return len(l.Nodes)
}

func (l *CustomListener) pop() logical.Node {
	if l.size() < 1 {
		panic("Invalid Expression")
	}
	result := l.Nodes[l.size()-1]
	l.Nodes = l.Nodes[:l.size()-1]
	return result
}

func (l *CustomListener) GetResult() logical.Node {
	return l.Result
}

func (l *CustomListener) GetUnaryNode(ctx *lib.TypesExpressionContext) *arithmetic.UnaryNode {
	dataType := GetDataType(ctx.GetStart())
	operand, _ := util.ConvertValue(ctx.GetText(), dataType)
	return &arithmetic.UnaryNode{
		Value:    operand,
		DataType: dataType,
	}
}

func (l *CustomListener) ExitParse(c *lib.ParseContext) {
	if l.Result == nil && l.size() == 1 {
		l.Result = l.pop()
	} else if l.Result == nil && l.size() == 2 {
		firstNode := l.pop()
		secondNode := l.pop()
		if firstNode.GetNodeType() == constant.ARITHMETIC && secondNode.GetNodeType() == constant.ARITHMETIC && secondNode.(*arithmetic.ArithmeticNode).Right == nil {
			l.Result = &arithmetic.ArithmeticNode{Left: firstNode, Operator: constant.UNARY}
		}
		if secondNode.GetNodeType() == constant.ARITHMETIC && firstNode.GetNodeType() == constant.ARITHMETIC && firstNode.(*arithmetic.ArithmeticNode).Right == nil {
			l.Result = &arithmetic.ArithmeticNode{Left: secondNode, Operator: constant.UNARY}
		}
	}
	if (l.Result == nil && l.TokenCount == 1 && reflect.TypeOf(l.LastToken) == reflect.TypeOf(&antlr.CommonToken{})) {
		field, _ := util.ConvertValue(l.LastToken.GetText(), constant.STRING)
		l.Result = &arithmetic.UnaryNode{
			DataType: constant.STRING,
			Value:    field,
		}
	}
	if l.Result == nil {
		panic("Error parsing expression for the string " + c.GetText())
	}
}

func (l *CustomListener) ExitBinaryExpression(c *lib.BinaryExpressionContext) {
	if l.size() < 2 {
		panic("Error parsing binary expression for the string " + c.GetText())
	}
	firstNode := l.pop()
	secondNode := l.pop()
	operator := GetLogicalOperator(c.GetOp().GetStart())
	l.push(logical.BooleanNode{
		Left:     secondNode,
		Right:    firstNode,
		Operator: operator,
	})
}

func (l *CustomListener) ExitTypesExpression(c *lib.TypesExpressionContext) {
	l.LastToken = c.GetStart()
}

func (l *CustomListener) ExitInExpression(c *lib.InExpressionContext) {
	l.ValidateField(c.GetField(), c.GetText())
	field := l.GetField(c.GetField().GetText())

	items := l.GetArrayElements(c.GetData().GetChildren())
	var itemsMapped []interface{}
	for _, i := range items {
		itemsMapped = append(itemsMapped, i)
	}
	var inNode = logical.InNode{
		Field: field,
		Items: itemsMapped,
	}
	if c.NOT() == nil {
		l.push(inNode)
	} else {
		var booleanNode = logical.BooleanNode{
			Left:     inNode,
			Right:    nil,
			Operator: constant.NOT,
		}
		l.push(booleanNode)
	}
}

func (l *CustomListener) ExitArrayExpression(ctx *lib.ArrayExpressionContext) {
	l.ValidateField(ctx.GetField(), ctx.GetText())
	field := l.GetField(ctx.GetField().GetText())
	items := l.GetArrayElements(ctx.GetData().GetChildren())
	var itemsMapped []interface{}
	for _, i := range items {
		itemsMapped = append(itemsMapped, i)
	}
	operator := l.OperatorService.GetOperatorFromSymbol(ctx.GetOp().GetText())
	l.push(logical.ArrayNode{
		Field:    field,
		Operator: operator,
		Items:    itemsMapped,
	})
}

func (l *CustomListener) ExitToExpression(c *lib.ToExpressionContext) {
	l.ValidateField(c.GetField(), c.GetText())
	field := l.GetField(c.GetField().GetText())
	lowerDataType := GetDataType(c.GetLower().GetStart())
	lowerValue, _ := util.ConvertValue(c.GetLower().GetStart().GetText(), lowerDataType)
	upperDataType := GetDataType(c.GetUpper().GetStart())
	upperValue, _ := util.ConvertValue(c.GetUpper().GetText(), upperDataType)
	l.push(logical.NumericRangeNode{
		Field:        field,
		FromValue:    lowerValue,
		ToValue:      upperValue,
		FromDataType: lowerDataType,
		ToDataType:   upperDataType,
	})
}

func (l *CustomListener) ExitNotExpression(c *lib.NotExpressionContext) {
	if l.size() < 1 {
		if l.LastToken == nil {
			panic("Error parsing not expression for the string " + c.GetText())
		}
		dataType := GetDataType(l.LastToken)
		value, _ := util.ConvertValue(l.LastToken.GetText(), dataType)
		l.push(arithmetic.UnaryNode{
			DataType: dataType,
			Value:    value,
		})
	}
	boolNode := logical.BooleanNode{
		Left:     l.pop(),
		Right:    nil,
		Operator: constant.NOT,
	}
	l.push(boolNode)
}

func (l *CustomListener) ExitComparatorExpression(c *lib.ComparatorExpressionContext) {
	field := l.GetField(c.GetLeft().GetText())
	operator := l.OperatorService.GetOperatorFromSymbol(c.GetOp().GetText())
	if (reflect.TypeOf(c.GetRight()) != reflect.TypeOf(&lib.TypesExpressionContext{}) && l.size() != 0) {
		value := l.pop()
		l.push(logical.ComparisonNode{
			Field:    field,
			Value:    value,
			Operator: operator,
			DataType: constant.INTEGER,
		})
	} else {
		dataType := GetDataType(c.GetRight().GetStart())
		value, _ := util.ConvertValue(c.GetRight().GetText(), dataType)
		l.push(logical.ComparisonNode{
			Field:    field,
			Value:    value,
			Operator: operator,
			DataType: dataType,
		})
	}
}

func (l *CustomListener) ExitUnaryArithmeticExpression(ctx *lib.UnaryArithmeticExpressionContext) {
	dataType := GetDataType(ctx.GetExp().GetStart())
	operand, _ := util.ConvertValue(ctx.GetExp().GetText(), dataType)
	leafNode := &arithmetic.UnaryNode{
		Value:    operand,
		DataType: dataType,
	}
	node := &arithmetic.ArithmeticNode{Left: leafNode, Operator: constant.UNARY}
	l.push(node)
}

func (l *CustomListener) ExitArithmeticFunctionExpression(ctx *lib.ArithmeticFunctionExpressionContext) {
	if l.ErrorCount > 0 {
		panic("Invalid Expression")
	}
	functionType := l.FunctionEvaluatorService.GetFunctionFromSymbol(ctx.GetLeft().GetText())
	items := l.GetArrayElements(ctx.GetData().GetChildren())
	var itemsMapped []interface{}
	for _, i := range items {
		itemsMapped = append(itemsMapped, i)
	}
	node := &arithmetic.ArithmeticFunctionNode{FunctionType: functionType, Items: itemsMapped}
	l.push(node)
}

func (l *CustomListener) ExitArithmeticExpression(ctx *lib.ArithmeticExpressionContext) {
	operator := l.OperatorService.GetOperatorFromSymbol(ctx.GetOp().GetText())
	if (reflect.TypeOf(ctx.GetLeft()) == reflect.TypeOf(&lib.TypesExpressionContext{}) && reflect.TypeOf(ctx.GetRight()) == reflect.TypeOf(&lib.TypesExpressionContext{})) {
		left := l.GetUnaryNode(ctx.GetLeft().(*lib.TypesExpressionContext))
		right := l.GetUnaryNode(ctx.GetRight().(*lib.TypesExpressionContext))
		node := &arithmetic.ArithmeticNode{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
		l.push(node)
	} else if (reflect.TypeOf(ctx.GetLeft()) == reflect.TypeOf(&lib.TypesExpressionContext{})) {
		left := l.GetUnaryNode(ctx.GetLeft().(*lib.TypesExpressionContext))
		right := l.pop()
		node := &arithmetic.ArithmeticNode{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
		l.push(node)
	} else if (reflect.TypeOf(ctx.GetRight()) == reflect.TypeOf(&lib.TypesExpressionContext{})) {
		right := l.GetUnaryNode(ctx.GetRight().(*lib.TypesExpressionContext))
		left := l.pop()
		node := &arithmetic.ArithmeticNode{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
		l.push(node)
	} else {
		if l.size() < 2 {
			panic("Error parsing expression for the string " + ctx.GetText())
		}
		right := l.pop()
		left := l.pop()
		node := &arithmetic.ArithmeticNode{
			Left:     left,
			Right:    right,
			Operator: operator,
		}
		l.push(node)
	}
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
		return constant.VERSION
	case lib.BooleanExpressionLexerTRUE:
		return constant.BOOLEAN
	case lib.BooleanExpressionLexerFALSE:
		return constant.BOOLEAN
	default:
		return constant.STRING
	}
}

func (l *CustomListener) ValidateField(token antlr.Token, text string) {
	if token == nil || (len(token.GetText()) == 0 && len(text) == 0) {
		panic("Invalid Expression, field and default value both are empty")
	}
}

func (l *CustomListener) ExitTypes(ctx *lib.TypesContext) {
	l.TokenCount++
	l.LastToken = ctx.GetStart()
}

func (l *CustomListener) GetField(field string) string {
	if len(field) == 0 {
		return l.DefaultField
	}
	return field
}

func (l *CustomListener) GetArrayElements(trees []antlr.Tree) []*arithmetic.UnaryNode {
	typesContextFilter := func(tree antlr.Tree) bool { return reflect.TypeOf(tree) == reflect.TypeOf(&lib.TypesContext{}) }
	var typesContextChildren = util.Filter(trees, typesContextFilter)
	var items []*arithmetic.UnaryNode

	for _, child := range typesContextChildren {
		dataType := GetDataType(child.(*lib.TypesContext).GetStart())
		value, _ := util.ConvertValue(child.(*lib.TypesContext).GetText(), dataType)
		items = append(items, &arithmetic.UnaryNode{
			DataType: dataType,
			Value:    value,
		})
	}

	return items
}

// VisitTerminal is called when a terminal node is visited.
func (s *CustomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CustomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CustomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CustomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *CustomListener) EnterParse(ctx *lib.ParseContext) {}

// EnterUnaryArithmeticExpression is called when production unaryArithmeticExpression is entered.
func (s *CustomListener) EnterUnaryArithmeticExpression(ctx *lib.UnaryArithmeticExpressionContext) {
}

func (s *CustomListener) EnterArithmeticFunctionExpression(ctx *lib.ArithmeticFunctionExpressionContext) {
}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *CustomListener) EnterBinaryExpression(ctx *lib.BinaryExpressionContext) {}

// EnterTypesExpression is called when production typesExpression is entered.
func (s *CustomListener) EnterTypesExpression(ctx *lib.TypesExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *CustomListener) EnterInExpression(ctx *lib.InExpressionContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *CustomListener) EnterArrayExpression(ctx *lib.ArrayExpressionContext) {}

// EnterToExpression is called when production toExpression is entered.
func (s *CustomListener) EnterToExpression(ctx *lib.ToExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *CustomListener) EnterNotExpression(ctx *lib.NotExpressionContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *CustomListener) EnterArithmeticExpression(ctx *lib.ArithmeticExpressionContext) {}

// EnterComparatorExpression is called when production comparatorExpression is entered.
func (s *CustomListener) EnterComparatorExpression(ctx *lib.ComparatorExpressionContext) {}

// EnterParentExpression is called when production parentExpression is entered.
func (s *CustomListener) EnterParentExpression(ctx *lib.ParentExpressionContext) {}

// ExitParentExpression is called when production parentExpression is exited.
func (s *CustomListener) ExitParentExpression(ctx *lib.ParentExpressionContext) {}

// EnterComparator is called when production comparator is entered.
func (s *CustomListener) EnterComparator(ctx *lib.ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *CustomListener) ExitComparator(ctx *lib.ComparatorContext) {}

// EnterArithmeticOperator is called when production arithmeticOperator is entered.
func (s *CustomListener) EnterArithmeticOperator(ctx *lib.ArithmeticOperatorContext) {}

// ExitArithmeticOperator is called when production arithmeticOperator is exited.
func (s *CustomListener) ExitArithmeticOperator(ctx *lib.ArithmeticOperatorContext) {}

// EnterWordlist is called when production wordlist is entered.
func (s *CustomListener) EnterWordlist(ctx *lib.WordlistContext) {}

// ExitWordlist is called when production wordlist is exited.
func (s *CustomListener) ExitWordlist(ctx *lib.WordlistContext) {}

// EnterArrayOperators is called when production arrayOperators is entered.
func (s *CustomListener) EnterArrayOperators(ctx *lib.ArrayOperatorsContext) {}

// ExitArrayOperators is called when production arrayOperators is exited.
func (s *CustomListener) ExitArrayOperators(ctx *lib.ArrayOperatorsContext) {}

// EnterNumericTypes is called when production numericTypes is entered.
func (s *CustomListener) EnterNumericTypes(ctx *lib.NumericTypesContext) {}

// ExitNumericTypes is called when production numericTypes is exited.
func (s *CustomListener) ExitNumericTypes(ctx *lib.NumericTypesContext) {}

// EnterTypes is called when production types is entered.
func (s *CustomListener) EnterTypes(ctx *lib.TypesContext) {}

// EnterBinary is called when production binary is entered.
func (s *CustomListener) EnterBinary(ctx *lib.BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *CustomListener) ExitBinary(ctx *lib.BinaryContext) {}

// EnterBool is called when production bool is entered.
func (s *CustomListener) EnterBool(ctx *lib.BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *CustomListener) ExitBool(ctx *lib.BoolContext) {}

// EnterArrayArithmeticFunction is called when production arrayArithmeticFunction is entered.
func (s *CustomListener) EnterArrayArithmeticFunction(ctx *lib.ArrayArithmeticFunctionContext) {
}

// ExitArrayArithmeticFunction is called when production arrayArithmeticFunction is exited.
func (s *CustomListener) ExitArrayArithmeticFunction(ctx *lib.ArrayArithmeticFunctionContext) {
}
