package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"reflect"
)

type ContainsAllOperator struct {
}

func (e *ContainsAllOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
	var leftArray []interface{}
	rv := reflect.ValueOf(left)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			leftArray = append(leftArray, rv.Index(i).Interface())
		}
	}
	for _, value := range right {
		res, err := GetOperator(constant.IN).Evaluate(constant.PRIMITIVE, dataType, validated, value, leftArray...)
		if err != nil {
			return false, err
		}
		if !res {
			return false, nil
		}
	}
	return true, nil
}

func (e *ContainsAllOperator) GetSymbol() string {
	return "CONTAINS_ALL"
}

func (e *ContainsAllOperator) GetOperator() constant.Operator {
	return constant.CONTAINS_ALL
}

func (e *ContainsAllOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (e *ContainsAllOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewContainsAllOperator() AbstractOperator {
	return &ContainsAllOperator{}
}
