package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type EqualsOperator struct {
}

func (e *EqualsOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	dt := datatype.GetDataType(dataType)
	res, err := dt.Compare(left, right[0])
	if err != nil {
		return false, err
	}
	return res == 0, nil
}

func (e *EqualsOperator) GetSymbol() string  {
	return "="
}

func (e *EqualsOperator) GetOperator() constant.Operator  {
	return constant.EQUALS
}

func NewEqualsOperator() AbstractOperator {
	return &EqualsOperator{}
}
