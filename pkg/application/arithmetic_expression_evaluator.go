package application

import (
	"errors"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/service"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"github.com/sidhant92/bool-parser-go/pkg/domain/arithmetic"
	"github.com/sidhant92/bool-parser-go/pkg/parser"
	"reflect"
)

type ArithmeticLeafNodeResponse struct {
	Value    interface{}
	DataType constant.DataType
}

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

func (b *ArithmeticExpressionEvaluator) evaluateNode(node domain.Node, data map[string]interface{}) (interface{}, error) {
	switch node.GetNodeType() {
	case constant.ARITHMETIC:
		return b.evaluateArithmeticNode(node.(*arithmetic.ArithmeticNode), data)
	case constant.ARITHMETIC_LEAF:
		return b.evaluateArithmeticLeafNode(node.(*arithmetic.ArithmeticLeafNode), data)
	case constant.ARITHMETIC_FUNCTION:
		return b.evaluateArithmeticFunctionNode(node.(*arithmetic.ArithmeticFunctionNode), data)
	case constant.ARITHMETIC_UNARY:
		return b.evaluateArithmeticUnaryNode(node.(*arithmetic.ArithmeticUnaryNode), data)
	case constant.UNARY_NODE:
		return b.evaluateUnaryNode(node.(*domain.UnaryNode), data)
	default:
		return nil, errors.New("unknown node")
	}
}

func (b *ArithmeticExpressionEvaluator) evaluateUnaryNode(node *domain.UnaryNode, data map[string]interface{}) (interface{}, error) {
	if node.DataType == constant.STRING {
		fieldData := util.GetValueFromMap(node.Value.(string), data)
		if fieldData != nil {
			return fieldData, nil
		}
	}
	return node.Value, nil
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticLeafNode(node *arithmetic.ArithmeticLeafNode, data map[string]interface{}) (*ArithmeticLeafNodeResponse, error) {
	stringValue := b.StringDataType.GetValue(node.Operand).(string)
	fetchedValue := util.GetValueFromMap(stringValue, data)
	if fetchedValue != nil {
		return &ArithmeticLeafNodeResponse{
			Value:    fetchedValue,
			DataType: util.GetDataType(fetchedValue),
		}, nil
	}
	return &ArithmeticLeafNodeResponse{
		Value:    node.Operand,
		DataType: node.DataType,
	}, nil
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticUnaryNode(node *arithmetic.ArithmeticUnaryNode, data map[string]interface{}) (interface{}, error) {
	resolvedValue, _ := b.evaluateNode(node.Operand, data)
	if (reflect.TypeOf(resolvedValue) == reflect.TypeOf(&ArithmeticLeafNodeResponse{})) {
		arithmeticLeafNodeResponse := resolvedValue.(*ArithmeticLeafNodeResponse)
		return b.OperatorService.EvaluateArithmeticExpression(arithmeticLeafNodeResponse.Value, arithmeticLeafNodeResponse.DataType, nil, constant.STRING, constant.UNARY, constant.PRIMITIVE)
	}
	dataType := util.GetDataType(resolvedValue)
	return b.OperatorService.EvaluateArithmeticExpression(resolvedValue, dataType, nil, constant.STRING, constant.UNARY, constant.PRIMITIVE)
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticFunctionNode(node *arithmetic.ArithmeticFunctionNode, data map[string]interface{}) (interface{}, error) {
	var resolvedValues []ArithmeticLeafNodeResponse
	for _, item := range node.Items {
		res, _ := b.evaluateArithmeticLeafNode(&item, data)
		resolvedValues = append(resolvedValues, *res)
	}
	var flattenedValues []domain.EvaluatedNode
	for _, item := range resolvedValues {
		if util.IsSlice(item.Value) {
			data := util.GetSliceFromInterface(item.Value)
			for _, val := range data {
				flattenedValues = append(flattenedValues, domain.EvaluatedNode{
					Value:    val,
					DataType: util.GetDataType(val),
				})
			}
		} else {
			flattenedValues = append(flattenedValues, domain.EvaluatedNode{
				Value:    item.Value,
				DataType: item.DataType,
			})
		}
	}
	return b.FunctionEvaluatorService.Evaluate(node.FunctionType, flattenedValues)
}

func (b *ArithmeticExpressionEvaluator) evaluateArithmeticNode(node *arithmetic.ArithmeticNode, data map[string]interface{}) (interface{}, error) {
	leftValue, _ := b.evaluateNode(node.Left, data)
	rightValue, _ := b.evaluateNode(node.Right, data)
	if (reflect.TypeOf(leftValue) == reflect.TypeOf(&ArithmeticLeafNodeResponse{}) && reflect.TypeOf(rightValue) == reflect.TypeOf(&ArithmeticLeafNodeResponse{})) {
		leftLeaf := leftValue.(*ArithmeticLeafNodeResponse)
		rightLeaf := rightValue.(*ArithmeticLeafNodeResponse)
		return b.OperatorService.EvaluateArithmeticExpression(leftLeaf.Value, leftLeaf.DataType, rightLeaf.Value, rightLeaf.DataType, node.Operator, constant.PRIMITIVE)
	} else if (reflect.TypeOf(leftValue) == reflect.TypeOf(&ArithmeticLeafNodeResponse{})) {
		leftLeaf := leftValue.(*ArithmeticLeafNodeResponse)
		rightDataType := util.GetDataType(rightValue)
		return b.OperatorService.EvaluateArithmeticExpression(leftLeaf.Value, leftLeaf.DataType, rightValue, rightDataType, node.Operator, constant.PRIMITIVE)
	} else if (reflect.TypeOf(rightValue) == reflect.TypeOf(&ArithmeticLeafNodeResponse{})) {
		rightLeaf := rightValue.(*ArithmeticLeafNodeResponse)
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
