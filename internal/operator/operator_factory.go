package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/operator/arithmetic"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

var operatorMap = map[constant.Operator]AbstractOperator{
	constant.EQUALS:             NewEqualsOperator(),
	constant.NOT_EQUAL:          NewNotEqualsOperator(),
	constant.GREATER_THAN:       NewGreaterThanOperator(),
	constant.GREATER_THAN_EQUAL: NewGreaterThanEqualOperator(),
	constant.LESS_THAN:          NewLessThanOperator(),
	constant.LESS_THAN_EQUAL:    NewLessThanEqualOperator(),
	constant.IN:                 NewInOperator(),
	constant.CONTAINS_ANY:       NewContainsAnyOperator(),
	constant.CONTAINS_ALL:       NewContainsAllOperator(),
}

var arithmeticOperatorMap = map[constant.Operator]arithmetic.AbstractOperator{
	constant.ADD:      arithmetic.NewAddOperator(),
	constant.SUBTRACT: arithmetic.NewSubtractOperator(),
	constant.DIVIDE:   arithmetic.NewDivideOperator(),
	constant.MULTIPLY: arithmetic.NewMultiplyOperator(),
	constant.EXPONENT: arithmetic.NewExponentOperator(),
	constant.MODULUS:  arithmetic.NewModulusOperator(),
	constant.UNARY:    arithmetic.NewUnaryOperator(),
}

func GetOperator(operator constant.Operator) AbstractOperator {
	return operatorMap[operator]
}

func GetArithmeticOperator(operator constant.Operator) arithmetic.AbstractOperator {
	return arithmeticOperatorMap[operator]
}

func GetAllValues() []AbstractOperator {
	var operatorList []AbstractOperator
	for _, v := range operatorMap {
		operatorList = append(operatorList, v)
	}
	return operatorList
}

func GetAllArithmeticValues() []arithmetic.AbstractOperator {
	var operatorList []arithmetic.AbstractOperator
	for _, v := range arithmeticOperatorMap {
		operatorList = append(operatorList, v)
	}
	return operatorList
}
