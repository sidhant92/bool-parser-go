package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type GreaterThanOperator struct {
}

func (e *GreaterThanOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0])
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

func NewGreaterThanOperator() AbstractOperator {
	return &GreaterThanOperator{}
}
