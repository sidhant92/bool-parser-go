package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type LenFunction struct {
}

func (m *LenFunction) Evaluate(items []domain.Field) (interface{}, error) {
	if len(items) == 1 && items[0].DataType == constant.STRING {
		return len(items[0].Value.(string)), nil
	}
	return len(items), nil
}

func (m *LenFunction) GetFunctionType() constant.FunctionType {
	return constant.LEN
}

func (m *LenFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE, constant.LIST}
}

func (m *LenFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION, constant.BOOLEAN}
}

func NewLenFunction() AbstractFunction {
	return &LenFunction{}
}
