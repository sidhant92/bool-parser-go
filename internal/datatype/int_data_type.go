package datatype

import (
	"fmt"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"reflect"
)

type IntDataType struct {
}

func (s *IntDataType) GetDataType() constant.DataType {
	return constant.INTEGER
}

func (s *IntDataType) IsValid(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Int {
		return true
	}
	_, ok := value.(int)
	if ok {
		return true
	}
	_, err := util.ConvertValue(fmt.Sprintf("%v",value), constant.INTEGER)
	if err != nil {
		return false
	}
	return true
}

func (s *IntDataType) GetValue(value interface{}) interface{} {
	val, ok := value.(int)
	if ok {
		return val
	}
	res, _ := util.ConvertValue(fmt.Sprintf("%v",value), constant.INTEGER)
	return res
}

func (s *IntDataType) Compare(left interface{}, right interface{}) (int, error) {
	leftValid := s.IsValid(left)
	rightValid := s.IsValid(right)
	if !leftValid || !rightValid {
		return 0, errors.INVALID_DATA_TYPE
	}
	leftValue := s.GetValue(left).(int)
	rightValue := s.GetValue(right).(int)
	if leftValue == rightValue {
		return 0, nil
	}
	if leftValue < rightValue {
		return -1, nil
	}
	return 1, nil
}

func NewIntDataType() DataType {
	return &IntDataType{}
}
