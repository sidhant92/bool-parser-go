package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type GreaterThanEqualOperator struct {
}

func (e *GreaterThanEqualOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0])
	if err != nil {
		return false, err
	}
	return res >= 0, nil
}

func (e *GreaterThanEqualOperator) GetSymbol() string  {
	return ">="
}

func (e *GreaterThanEqualOperator) GetOperator() constant.Operator  {
	return constant.GREATER_THAN_EQUAL
}

func NewGreaterThanEqualOperator() AbstractOperator {
	return &GreaterThanEqualOperator{}
}
