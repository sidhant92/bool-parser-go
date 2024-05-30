package util

import (
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"reflect"
)

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func GetSliceFromInterface(value interface{}) []interface{} {
	var slice []interface{}
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			slice = append(slice, rv.Index(i).Interface())
		}
	}
	return slice
}

func MapToEvaluatedNodes(items []interface{}) []domain.EvaluatedNode {
	var flattenedValues []domain.EvaluatedNode
	for _, item := range items {
		if (reflect.TypeOf(item) == reflect.TypeOf(&domain.EvaluatedNode{})) {
			evaluatedNode := item.(domain.EvaluatedNode)
			if IsSlice(evaluatedNode.Value) {
				data := GetSliceFromInterface(evaluatedNode.Value)
				for _, val := range data {
					flattenedValues = append(flattenedValues, domain.EvaluatedNode{
						Value:    val,
						DataType: GetDataType(val),
					})
				}
			} else {
				flattenedValues = append(flattenedValues, evaluatedNode)
			}
		}
		if IsSlice(item) {
			data := GetSliceFromInterface(item)
			for _, val := range data {
				flattenedValues = append(flattenedValues, domain.EvaluatedNode{
					Value:    val,
					DataType: GetDataType(val),
				})
			}
		} else {
			flattenedValues = append(flattenedValues, domain.EvaluatedNode{
				Value:    item,
				DataType: GetDataType(item),
			})
		}
	}
	return flattenedValues
}
