package application

import (
	"errors"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
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

type ArithmeticExpressionEvaluator struct {
	Parser                   parser.Parser
	OperatorService          *service.OperatorService
	FunctionEvaluatorService *service.FunctionEvaluatorService
	StringDataType           datatype.DataType
}

func (b *ArithmeticExpressionEvaluator) Evaluate(expression string, data map[string]interface{}) (interface{}, error) {
	node, err := b.Parser.Parse(expression)
	if err != nil {
		return false, err
	}
	return b.evaluateNode(node, data)
}

func (b *ArithmeticExpressionEvaluator) evaluateNode(node logical.Node, data map[string]interface{}) (interface{}, error) {
	switch node.GetNodeType() {
	case constant.ARITHMETIC:
		return b.evaluateArithmeticNode(node.(*arithmetic.ArithmeticNode), data)
	case constant.ARITHMETIC_FUNCTION:
		return b.evaluateArithmeticFunctionNode(node.(*arithmetic.ArithmeticFunctionNode), data)
	case constant.UNARY_NODE:
		return b.evaluateUnaryNode(node.(*arithmetic.UnaryNode), data)
	case constant.FIELD_NODE:
		return b.evaluateFieldNode(node.(*arithmetic.FieldNode), data)
	default:
		return nil, errors.New("unknown node")
	}
}

func (b *ArithmeticExpressionEvaluator) evaluateFieldNode(node *arithmetic.FieldNode, data map[string]interface{}) (interface{}, error) {
	fieldData := util.GetValueFromMap(node.Field, data)
	if fieldData != nil {
		return fieldData, nil
	}
	return nil, errors2.KEY_DATA_NOT_PRESENT
}

func (b *ArithmeticExpressionEvaluator) evaluateUnaryNode(node *arithmetic.UnaryNode, data map[string]interface{}) (interface{}, error) {
	return node.Value, nil
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticFunctionNode(node *arithmetic.ArithmeticFunctionNode, data map[string]interface{}) (interface{}, error) {
	var resolvedValues []interface{}
	for _, item := range util.GetSliceFromInterface(node.Items) {

		res, _ := b.evaluateNode(item.(logical.Node), data)
		resolvedValues = append(resolvedValues, res)
	}
	flattenedValues := util.MapToEvaluatedNodes(resolvedValues)
	return b.FunctionEvaluatorService.Evaluate(node.FunctionType, flattenedValues)
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticNode(node *arithmetic.ArithmeticNode, data map[string]interface{}) (interface{}, error) {
	leftValue, _ := b.evaluateNode(node.Left, data)
	if node.Operator == constant.UNARY {
		if reflect.TypeOf(leftValue) == reflect.TypeOf(&domain.EvaluatedNode{}) {
			evaluatedNode := leftValue.(domain.EvaluatedNode)
			return b.OperatorService.EvaluateArithmeticExpression(evaluatedNode.Value, evaluatedNode.DataType, nil, constant.STRING, node.Operator, constant.PRIMITIVE)
		} else {
			dataType := util.GetDataType(leftValue)
			return b.OperatorService.EvaluateArithmeticExpression(leftValue, dataType, nil, constant.STRING, node.Operator, constant.PRIMITIVE)
		}
	}
	rightValue, _ := b.evaluateNode(node.Right, data)
	if (reflect.TypeOf(leftValue) == reflect.TypeOf(&domain.EvaluatedNode{}) && reflect.TypeOf(rightValue) == reflect.TypeOf(&domain.EvaluatedNode{})) {
		leftLeaf := leftValue.(*domain.EvaluatedNode)
		rightLeaf := rightValue.(*domain.EvaluatedNode)
		return b.OperatorService.EvaluateArithmeticExpression(leftLeaf.Value, leftLeaf.DataType, rightLeaf.Value, rightLeaf.DataType, node.Operator, constant.PRIMITIVE)
	} else if (reflect.TypeOf(leftValue) == reflect.TypeOf(&domain.EvaluatedNode{})) {
		leftLeaf := leftValue.(*domain.EvaluatedNode)
		rightDataType := util.GetDataType(rightValue)
		return b.OperatorService.EvaluateArithmeticExpression(leftLeaf.Value, leftLeaf.DataType, rightValue, rightDataType, node.Operator, constant.PRIMITIVE)
	} else if (reflect.TypeOf(rightValue) == reflect.TypeOf(&domain.EvaluatedNode{})) {
		rightLeaf := rightValue.(*domain.EvaluatedNode)
		leftDataType := util.GetDataType(leftValue)
		return b.OperatorService.EvaluateArithmeticExpression(leftValue, leftDataType, rightLeaf.Value, rightLeaf.DataType, node.Operator, constant.PRIMITIVE)
	} else {
		leftDataType := util.GetDataType(leftValue)
		rightDataType := util.GetDataType(rightValue)
		return b.OperatorService.EvaluateArithmeticExpression(leftValue, leftDataType, rightValue, rightDataType, node.Operator, constant.PRIMITIVE)
	}
}

func NewArithmeticExpressionEvaluator(parser parser.Parser) ArithmeticExpressionEvaluator {
	return ArithmeticExpressionEvaluator{parser, service.NewOperatorService(), service.NewFunctionEvaluatorService(), datatype.NewStringDataType()}
}
