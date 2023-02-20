package datatype

import (
	"fmt"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"reflect"
)

type BooleanDataType struct {
}

func (s *BooleanDataType) GetDataType() constant.DataType {
	return constant.BOOLEAN
}

func (s *BooleanDataType) IsValid(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Bool {
		return true
	}
	_, ok := value.(bool)
	if ok {
		return true
	}
	_, err := util.ConvertValue(fmt.Sprintf("%v",value), constant.BOOLEAN)
	if err != nil {
		return false
	}
	return true
}

func (s *BooleanDataType) GetValue(value interface{}) interface{} {
	val, ok := value.(bool)
	if ok {
		return val
	}
	res, _ := util.ConvertValue(fmt.Sprintf("%v",value), constant.BOOLEAN)
	return res
}

func (s *BooleanDataType) Compare(left interface{}, right interface{}) (int, error) {
	leftValid := s.IsValid(left)
	rightValid := s.IsValid(right)
	if !leftValid || !rightValid {
		return 0, errors.INVALID_DATA_TYPE
	}
	leftValue := s.GetValue(left).(bool)
	rightValue := s.GetValue(right).(bool)
	if leftValue == rightValue {
		return 0, nil
	}
	return -1, nil
}

func NewBooleanDataType() DataType {
	return &BooleanDataType{}
}
