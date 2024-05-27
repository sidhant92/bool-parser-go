package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type IntFunction struct {
}

func (m *IntFunction) Evaluate(items []domain.Field) (interface{}, error) {
	datatype := items[0].DataType
	value := items[0].Value
	if datatype == constant.DECIMAL {
		return int(value.(float64)), nil
	}
	if datatype == constant.LONG {
		return int(value.(int64)), nil
	}
	return value.(int), nil
}

func (m *IntFunction) GetFunctionType() constant.FunctionType {
	return constant.INT
}

func (m *IntFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (m *IntFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewIntFunction() AbstractFunction {
	return &IntFunction{}
}
