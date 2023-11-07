package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type LessThanEqualOperator struct {
}

func (e *LessThanEqualOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0], validated)
	if err != nil {
		return false, err
	}
	return res <= 0, nil
}

func (e *LessThanEqualOperator) GetSymbol() string {
	return "<="
}

func (e *LessThanEqualOperator) GetOperator() constant.Operator {
	return constant.LESS_THAN_EQUAL
}

func (e *LessThanEqualOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *LessThanEqualOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION}
}

func NewLessThanEqualOperator() AbstractOperator {
	return &LessThanEqualOperator{}
}
