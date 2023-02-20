package application

import (
	"errors"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"github.com/sidhant92/bool-parser-go/pkg/parser"
	"log"
)

type BooleanExpressionEvaluator struct {
	Parser          parser.Parser
	OperatorService *service.OperatorService
}

func (b *BooleanExpressionEvaluator) Evaluate(expression string, data map[string]interface{}) (bool, error) {
	node, err := b.Parser.Parse(expression)
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
	default:
		return false, errors.New("unknown node")
	}
}

func (b *BooleanExpressionEvaluator) evaluateComparisonNode(node domain.ComparisonNode, data map[string]interface{}) (bool, error) {
	if b.checkFieldDataMissing(node.Field, data) {
		return false, nil
	}
	fieldData := data[node.Field]
	return b.OperatorService.Evaluate(node.Operator, node.DataType, fieldData, node.Value)
}

func (b *BooleanExpressionEvaluator) evaluateNumericRangeNode(node domain.NumericRangeNode, data map[string]interface{}) (bool, error) {
	if b.checkFieldDataMissing(node.Field, data) {
		return false, nil
	}
	fieldData := data[node.Field]
	leftRes, err := b.OperatorService.Evaluate(constant.GREATER_THAN_EQUAL, node.FromDataType, fieldData, node.FromValue)
	if err != nil {
		return false, err
	}
	rightRes, err := b.OperatorService.Evaluate(constant.LESS_THAN_EQUAL, node.ToDataType, fieldData, node.ToValue)
	if err != nil {
		return false, err
	}
	return leftRes && rightRes, nil
}

func (b *BooleanExpressionEvaluator) evaluateInNode(node domain.InNode, data map[string]interface{}) (bool, error) {
	if b.checkFieldDataMissing(node.Field, data) {
		return false, nil
	}
	fieldData := data[node.Field]
	dataType := node.Items[0].DataType
	var values []interface{}
	for _, item := range node.Items {
		values = append(values, item.Value)
	}
	res, err := b.OperatorService.Evaluate(constant.IN, dataType, fieldData, values...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (b *BooleanExpressionEvaluator) evaluateUnaryNode(node domain.UnaryNode, data map[string]interface{}) (bool, error) {
	if node.DataType == constant.BOOLEAN {
		return node.Value.(bool), nil
	}
	if b.checkFieldDataMissing(node.Value.(string), data) {
		return false, nil
	}
	fieldData := data[node.Value.(string)]
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

func (b *BooleanExpressionEvaluator) checkFieldDataMissing(field string, data map[string]interface{}) bool {
	if _, ok := data[field]; ok {
		return false
	}
	log.Println("required field data is missing")
	return true
}

func NewBooleanExpressionEvaluator(parser parser.Parser) BooleanExpressionEvaluator {
	return BooleanExpressionEvaluator{parser, service.NewOperatorService()}
}
