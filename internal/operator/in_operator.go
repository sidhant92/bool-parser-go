package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type InOperator struct {
}

func (e *InOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	for _, value := range right {
		res, err := GetOperator(constant.EQUALS).Evaluate(containerDataType, dataType, validated, left, value)
		if err != nil {
			return false, err
		}
		if res {
			return true, nil
		}
	}
	return false, nil
}

func (e *InOperator) GetSymbol() string {
	return "IN"
}

func (e *InOperator) GetOperator() constant.Operator {
	return constant.IN
}

func (e *InOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *InOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewInOperator() AbstractOperator {
	return &InOperator{}
}
