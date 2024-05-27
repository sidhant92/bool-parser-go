package application

import (
	"errors"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"github.com/sidhant92/bool-parser-go/pkg/domain/arithmetic"
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"github.com/sidhant92/bool-parser-go/pkg/parser"
	"reflect"
)

type BooleanExpressionEvaluator struct {
	Parser                        parser.Parser
	OperatorService               *service.OperatorService
	ArithmeticExpressionEvaluator ArithmeticExpressionEvaluator
}

func (b *BooleanExpressionEvaluator) Evaluate(expression string, data map[string]interface{}, defaultField ...string) (bool, error) {
	node, err := b.Parser.Parse(expression, defaultField...)
	if err != nil {
		return false, err
	}
	return b.evaluateNode(node, data)
}

func (b *BooleanExpressionEvaluator) evaluateNode(node domain.Node, data map[string]interface{}) (bool, error) {
	switch node.GetNodeType() {
	case constant.COMPARISON_NODE:
		return b.evaluateComparisonNode(node.(domain.ComparisonNode), data)
	case constant.NUMERIC_RANGE_NODE:
		return b.evaluateNumericRangeNode(node.(domain.NumericRangeNode), data)
	case constant.IN_NODE:
		return b.evaluateInNode(node.(domain.InNode), data)
	case constant.UNARY_NODE:
		return b.evaluateUnaryNode(node.(domain.UnaryNode), data)
	case constant.BOOLEAN_NODE:
		return b.evaluateBooleanNode(node.(domain.BooleanNode), data)
	case constant.ARRAY_NODE:
		return b.evaluateArrayNode(node.(domain.ArrayNode), data)
	default:
		return false, errors.New("unknown node")
	}
}

func (b *BooleanExpressionEvaluator) evaluateComparisonNode(node domain.ComparisonNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	value := node.Value
	arithmeticType := reflect.TypeOf(new(arithmetic.ArithmeticBaseNode)).Elem()
	if reflect.TypeOf(node.Value).Implements(arithmeticType) {
		res, _ := b.ArithmeticExpressionEvaluator.evaluateNode(value.(domain.Node), data)
		value = res
	}
	return b.OperatorService.Evaluate(node.Operator, constant.PRIMITIVE, node.DataType, fieldData, value)
}

func (b *BooleanExpressionEvaluator) evaluateNumericRangeNode(node domain.NumericRangeNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	leftRes, err := b.OperatorService.Evaluate(constant.GREATER_THAN_EQUAL, constant.PRIMITIVE, node.FromDataType, fieldData, node.FromValue)
	if err != nil {
		return false, err
	}
	rightRes, err := b.OperatorService.Evaluate(constant.LESS_THAN_EQUAL, constant.PRIMITIVE, node.ToDataType, fieldData, node.ToValue)
	if err != nil {
		return false, err
	}
	return leftRes && rightRes, nil
}

func (b *BooleanExpressionEvaluator) evaluateInNode(node domain.InNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	dataType := node.Items[0].DataType
	var values []interface{}
	for _, item := range node.Items {
		values = append(values, item.Value)
	}
	res, err := b.OperatorService.Evaluate(constant.IN, constant.PRIMITIVE, dataType, fieldData, values...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (b *BooleanExpressionEvaluator) evaluateUnaryNode(node domain.UnaryNode, data map[string]interface{}) (bool, error) {
	if node.DataType == constant.BOOLEAN {
		return node.Value.(bool), nil
	}
	fieldData := util.GetValueFromMap(node.Value.(string), data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	val, ok := fieldData.(bool)
	if ok {
		return val, nil
	}
	return false, errors2.INVALID_UNARY_OPERAND
}

func (b *BooleanExpressionEvaluator) evaluateBooleanNode(node domain.BooleanNode, data map[string]interface{}) (bool, error) {
	switch node.Operator {
	case constant.AND:
		leftRes, err := b.evaluateNode(node.Left, data)
		if err != nil {
			return false, err
		}
		rightRes, err := b.evaluateNode(node.Right, data)
		if err != nil {
			return false, err
		}
		return leftRes && rightRes, nil
	case constant.OR:
		leftRes, err := b.evaluateNode(node.Left, data)
		if err != nil {
			return false, err
		}
		rightRes, err := b.evaluateNode(node.Right, data)
		if err != nil {
			return false, err
		}
		return leftRes || rightRes, nil
	default:
		res, err := b.evaluateNode(node.Left, data)
		if err != nil {
			return false, err
		}
		return !res, nil
	}
}

func (b *BooleanExpressionEvaluator) evaluateArrayNode(node domain.ArrayNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	dataType := node.Items[0].DataType
	var values []interface{}
	for _, item := range node.Items {
		values = append(values, item.Value)
	}
	res, err := b.OperatorService.Evaluate(node.Operator, constant.LIST, dataType, fieldData, values...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func NewBooleanExpressionEvaluator(parser parser.Parser) BooleanExpressionEvaluator {
	return BooleanExpressionEvaluator{parser, service.NewOperatorService(), NewArithmeticExpressionEvaluator(parser)}
}
