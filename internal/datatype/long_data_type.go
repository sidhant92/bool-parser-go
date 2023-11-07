package datatype

import (
	"fmt"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"reflect"
)

type LongDataType struct {
}

func (s *LongDataType) GetDataType() constant.DataType {
	return constant.LONG
}

func (s *LongDataType) IsValid(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Int64 {
		return true
	}
	_, ok := value.(int64)
	if ok {
		return true
	}
	_, err := util.ConvertValue(fmt.Sprintf("%v",value), constant.LONG)
	if err != nil {
		return false
	}
	return true
}

func (s *LongDataType) GetValue(value interface{}) interface{} {
	val, ok := value.(int64)
	if ok {
		return val
	}
	res, _ := util.ConvertValue(fmt.Sprintf("%v",value), constant.LONG)
	return res
}

func (s *LongDataType) Compare(left interface{}, right interface{}, validated bool) (int, error) {
	if !validated {
		leftValid := s.IsValid(left)
		rightValid := s.IsValid(right)
		if !leftValid || !rightValid {
			return 0, errors.INVALID_DATA_TYPE
		}
	}
	leftValue := s.GetValue(left).(int64)
	rightValue := s.GetValue(right).(int64)
	if leftValue == rightValue {
		return 0, nil
	}
	if leftValue < rightValue {
		return -1, nil
	}
	return 1, nil
}

func NewLongDataType() DataType {
	return &LongDataType{}
}
