package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"reflect"
)

type ContainsAnyOperator struct {
}

func (e *ContainsAnyOperator) Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error) {
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
		if res {
			return true, nil
		}
	}
	return false, nil
}

func (e *ContainsAnyOperator) GetSymbol() string {
	return "CONTAINS_ANY"
}

func (e *ContainsAnyOperator) GetOperator() constant.Operator {
	return constant.CONTAINS_ANY
}

func (e *ContainsAnyOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (e *ContainsAnyOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewContainsAnyOperator() AbstractOperator {
	return &ContainsAnyOperator{}
}
