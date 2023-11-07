package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type EqualsOperator struct {
}

func (e *EqualsOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0], validated)
	if err != nil {
		return false, err
	}
	return res == 0, nil
}

func (e *EqualsOperator) GetSymbol() string {
	return "="
}

func (e *EqualsOperator) GetOperator() constant.Operator {
	return constant.EQUALS
}

func (e *EqualsOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *EqualsOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewEqualsOperator() AbstractOperator {
	return &EqualsOperator{}
}
