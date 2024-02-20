package operator

import "github.com/sidhant92/bool-parser-go/pkg/constant"

var operatorMap = map[constant.Operator]AbstractOperator{
	constant.EQUALS: NewEqualsOperator(),
	constant.NOT_EQUAL: NewNotEqualsOperator(),
	constant.GREATER_THAN: NewGreaterThanOperator(),
	constant.GREATER_THAN_EQUAL: NewGreaterThanEqualOperator(),
	constant.LESS_THAN: NewLessThanOperator(),
	constant.LESS_THAN_EQUAL: NewLessThanEqualOperator(),
	constant.IN: NewInOperator(),
	constant.CONTAINS_ANY: NewContainsAnyOperator(),
	constant.CONTAINS_ALL: NewContainsAllOperator(),
}

func GetOperator(operator constant.Operator) AbstractOperator {
	return operatorMap[operator]
}

func GetAllValues() []AbstractOperator {
	var operatorList []AbstractOperator
	for _,v := range operatorMap {
		operatorList = append(operatorList, v)
	}
	return operatorList
}
