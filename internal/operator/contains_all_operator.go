package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"reflect"
)

type ContainsAllOperator struct {
}

func (e *ContainsAllOperator) Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error) {
	var leftArray []domain.EvaluatedNode
	rv := reflect.ValueOf(leftOperand)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			leftArray = append(leftArray, domain.EvaluatedNode{
				Value:    rv.Index(i).Interface(),
				DataType: leftOperandDataType,
			})
		}
	}
	for _, rightOperand := range rightOperands {
		res, err := GetOperator(constant.IN).Evaluate(constant.PRIMITIVE, rightOperand.Value, rightOperand.DataType, leftArray)
		if err != nil {
			return false, err
		}
		if !res {
			return false, nil
		}
	}
	return true, nil
}

func (e *ContainsAllOperator) GetSymbol() string {
	return "CONTAINS_ALL"
}

func (e *ContainsAllOperator) GetOperator() constant.Operator {
	return constant.CONTAINS_ALL
}

func (e *ContainsAllOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (e *ContainsAllOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewContainsAllOperator() AbstractOperator {
	return &ContainsAllOperator{}
}
