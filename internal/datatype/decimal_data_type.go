package datatype

import (
	"fmt"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"reflect"
)

type DecimalDataType struct {
}

func (s *DecimalDataType) GetDataType() constant.DataType {
	return constant.DECIMAL
}

func (s *DecimalDataType) IsValid(value interface{}) bool {
	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return true
	}
	_, ok := value.(float64)
	if ok {
		return true
	}
	_, err := util.ConvertValue(fmt.Sprintf("%v",value), constant.DECIMAL)
	if err != nil {
		return false
	}
	return true
}

func (s *DecimalDataType) GetValue(value interface{}) interface{} {
	val, ok := value.(float64)
	if ok {
		return val
	}
	res, _ := util.ConvertValue(fmt.Sprintf("%v",value), constant.DECIMAL)
	return res
}

func (s *DecimalDataType) Compare(left interface{}, right interface{}, validated bool) (int, error) {
	if !validated {
		leftValid := s.IsValid(left)
		rightValid := s.IsValid(right)
		if !leftValid || !rightValid {
			return 0, errors.INVALID_DATA_TYPE
		}
	}
	leftValue := s.GetValue(left).(float64)
	rightValue := s.GetValue(right).(float64)
	if leftValue == rightValue {
		return 0, nil
	}
	if leftValue < rightValue {
		return -1, nil
	}
	return 1, nil
}

func NewDecimalDataType() DataType {
	return &DecimalDataType{}
}
