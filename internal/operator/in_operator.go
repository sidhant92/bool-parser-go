package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type InOperator struct {
}

func (e *InOperator) Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error) {
	for _, value := range right {
		res, err := GetOperator(constant.EQUALS).Evaluate(dataType, left, value)
		if err != nil {
			return false, err
		}
		if res {
			return true, nil
		}
	}
	return false, nil
}

func (e *InOperator) GetSymbol() string {
	return "IN"
}

func (e *InOperator) GetOperator() constant.Operator {
	return constant.IN
}

func NewInOperator() AbstractOperator {
	return &InOperator{}
}
