package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type NotEqualsOperator struct {
}

func (e *NotEqualsOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	eq := GetOperator(constant.EQUALS)
	res, err := eq.Evaluate(dataType, left, right[0])
	if err != nil {
		return false, err
	}
	return !res, nil
}

func (e *NotEqualsOperator) GetSymbol() string  {
	return "!="
}

func (e *NotEqualsOperator) GetOperator() constant.Operator  {
	return constant.NOT_EQUAL
}

func NewNotEqualsOperator() AbstractOperator {
	return &NotEqualsOperator{}
}
