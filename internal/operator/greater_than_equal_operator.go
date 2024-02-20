package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type GreaterThanEqualOperator struct {
}

func (e *GreaterThanEqualOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0], validated)
	if err != nil {
		return false, err
	}
	return res >= 0, nil
}

func (e *GreaterThanEqualOperator) GetSymbol() string {
	return ">="
}

func (e *GreaterThanEqualOperator) GetOperator() constant.Operator {
	return constant.GREATER_THAN_EQUAL
}

func (e *GreaterThanEqualOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *GreaterThanEqualOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION}
}

func NewGreaterThanEqualOperator() AbstractOperator {
	return &GreaterThanEqualOperator{}
}
