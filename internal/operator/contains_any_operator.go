package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"reflect"
)

type ContainsAnyOperator struct {
}

func (e *ContainsAnyOperator) Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error) {
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
		if res {
			return true, nil
		}
	}
	return false, nil
}

func (e *ContainsAnyOperator) GetSymbol() string {
	return "CONTAINS_ANY"
}

func (e *ContainsAnyOperator) GetOperator() constant.Operator {
	return constant.CONTAINS_ANY
}

func (e *ContainsAnyOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (e *ContainsAnyOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewContainsAnyOperator() AbstractOperator {
	return &ContainsAnyOperator{}
}
