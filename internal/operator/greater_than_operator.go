package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type GreaterThanOperator struct {
}

func (e *GreaterThanOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0], validated)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

func (e *GreaterThanOperator) GetSymbol() string {
	return ">"
}

func (e *GreaterThanOperator) GetOperator() constant.Operator {
	return constant.GREATER_THAN
}

func (e *GreaterThanOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *GreaterThanOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION}
}

func NewGreaterThanOperator() AbstractOperator {
	return &GreaterThanOperator{}
}
