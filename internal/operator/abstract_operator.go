package operator

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type AbstractOperator interface {
	Evaluate(dataType constant.DataType, left interface{}, right ...interface{}) (bool, error)

	GetSymbol() string

	GetOperator() constant.Operator
}
