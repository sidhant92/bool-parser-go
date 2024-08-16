package application

import (
	"errors"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"github.com/sidhant92/bool-parser-go/pkg/domain/arithmetic"
	"github.com/sidhant92/bool-parser-go/pkg/domain/logical"
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

func (b *BooleanExpressionEvaluator) evaluateNode(node logical.Node, data map[string]interface{}) (bool, error) {
	switch node.GetNodeType() {
	case constant.COMPARISON_NODE:
		return b.evaluateComparisonNode(node.(logical.ComparisonNode), data)
	case constant.NUMERIC_RANGE_NODE:
		return b.evaluateNumericRangeNode(node.(logical.NumericRangeNode), data)
	case constant.IN_NODE:
		return b.evaluateInNode(node.(logical.InNode), data)
	case constant.UNARY_NODE:
		return b.evaluateUnaryNode(node.(arithmetic.UnaryNode), data)
	case constant.BOOLEAN_NODE:
		return b.evaluateBooleanNode(node.(logical.BooleanNode), data)
	case constant.ARRAY_NODE:
		return b.evaluateArrayNode(node.(logical.ArrayNode), data)
	default:
		return false, errors.New("unknown node")
	}
}

func comparisonNodeNullCheck(node logical.ComparisonNode) bool  {
	return node.IsNullCheck() && node.Value.(*arithmetic.FieldNode).IsNull()
}

func checkForNullValues(fieldData interface{}, node logical.ComparisonNode) (interface{}, error) {
	if comparisonNodeNullCheck(node) {
		if fieldData == nil {
			return "null", nil
		} else {
			return fieldData, nil
		}
	} else {
		if fieldData == nil {
			return nil, errors2.KEY_DATA_NOT_PRESENT
		} else {
			return fieldData, nil
		}
	}
}

func (b *BooleanExpressionEvaluator) evaluateComparisonNode(node logical.ComparisonNode, data map[string]interface{}) (bool, error) {
	fieldDataOptional := util.GetValueFromMap(node.Field, data)
	fieldData, err := checkForNullValues(fieldDataOptional, node)
	if err != nil {
		return false, err
	}
	var value interface{}
	value = node.Value
	var dataType = node.DataType
	arithmeticType := reflect.TypeOf(new(arithmetic.ArithmeticBaseNode)).Elem()
	if comparisonNodeNullCheck(node) {
		value = "null"
	} else if reflect.TypeOf(node.Value).Implements(arithmeticType) {
		res, _ := b.ArithmeticExpressionEvaluator.evaluateNode(value.(logical.Node), data)
		value = res
		dataType = util.GetDataType(value)
	}
	return b.OperatorService.Evaluate(node.Operator, constant.PRIMITIVE, dataType, fieldData, value)
}

func (b *BooleanExpressionEvaluator) resolveArrayElements(items []interface{}, data map[string]interface{}) []domain.EvaluatedNode {
	var resolvedValues []interface{}
	for _, item := range util.GetSliceFromInterface(items) {
		arithmeticBaseNode := reflect.TypeOf(new(arithmetic.ArithmeticBaseNode)).Elem()
		if reflect.TypeOf(item).Implements(arithmeticBaseNode) {
			res, _ := b.ArithmeticExpressionEvaluator.evaluateNode(item.(logical.Node), data)
			resolvedValues = append(resolvedValues, res)
		} else {
			res, _ := b.evaluateNode(item.(logical.Node), data)
			resolvedValues = append(resolvedValues, res)
		}
	}
	flattenedValues := util.MapToEvaluatedNodes(resolvedValues)
	return flattenedValues
}

func (b *BooleanExpressionEvaluator) evaluateNumericRangeNode(node logical.NumericRangeNode, data map[string]interface{}) (bool, error) {
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

func (b *BooleanExpressionEvaluator) evaluateInNode(node logical.InNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	items := b.resolveArrayElements(node.Items, data)
	dataType := items[0].DataType
	var values []interface{}
	for _, item := range items {
		values = append(values, item.Value)
	}
	res, err := b.OperatorService.Evaluate(constant.IN, constant.PRIMITIVE, dataType, fieldData, values...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (b *BooleanExpressionEvaluator) evaluateUnaryNode(node arithmetic.UnaryNode, data map[string]interface{}) (bool, error) {
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

func (b *BooleanExpressionEvaluator) evaluateBooleanNode(node logical.BooleanNode, data map[string]interface{}) (bool, error) {
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

func (b *BooleanExpressionEvaluator) evaluateArrayNode(node logical.ArrayNode, data map[string]interface{}) (bool, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData == nil {
		return false, errors2.KEY_DATA_NOT_PRESENT
	}
	items := b.resolveArrayElements(node.Items, data)
	var values []interface{}
	for _, item := range items {
		values = append(values, item.Value)
	}
	dataType := items[0].DataType
	res, err := b.OperatorService.Evaluate(node.Operator, constant.LIST, dataType, fieldData, values...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func NewBooleanExpressionEvaluator(parser parser.Parser) BooleanExpressionEvaluator {
	return BooleanExpressionEvaluator{parser, service.NewOperatorService(), NewArithmeticExpressionEvaluator(parser)}
}
