package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type LessThanOperator struct {
}

func (e *LessThanOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0], validated)
	if err != nil {
		return false, err
	}
	return res < 0, nil
}

func (e *LessThanOperator) GetSymbol() string {
	return "<"
}

func (e *LessThanOperator) GetOperator() constant.Operator {
	return constant.LESS_THAN
}

func (e *LessThanOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *LessThanOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION}
}

func NewLessThanOperator() AbstractOperator {
	return &LessThanOperator{}
}
