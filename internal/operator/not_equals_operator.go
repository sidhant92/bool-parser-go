package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type NotEqualsOperator struct {
}

func (e *NotEqualsOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	eq := GetOperator(constant.EQUALS)
	res, err := eq.Evaluate(containerDataType, dataType, validated, left, right[0])
	if err != nil {
		return false, err
	}
	return !res, nil
}

func (e *NotEqualsOperator) GetSymbol() string  {
	return "!="
}

func (e *NotEqualsOperator) GetOperator() constant.Operator  {
	return constant.NOT_EQUAL
}

func (e *NotEqualsOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *NotEqualsOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewNotEqualsOperator() AbstractOperator {
	return &NotEqualsOperator{}
}
