package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type LessThanEqualOperator struct {
}

func (e *LessThanEqualOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0])
	if err != nil {
		return false, err
	}
	return res <= 0, nil
}

func (e *LessThanEqualOperator) GetSymbol() string  {
	return "<="
}

func (e *LessThanEqualOperator) GetOperator() constant.Operator  {
	return constant.LESS_THAN_EQUAL
}

func NewLessThanEqualOperator() AbstractOperator {
	return &LessThanEqualOperator{}
}
