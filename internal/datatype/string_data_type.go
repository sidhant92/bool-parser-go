package datatype

import (
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"strings"
)

type StringDataType struct {
}

func (s *StringDataType) GetDataType() constant.DataType {
	return constant.STRING
}

func (s *StringDataType) IsValid(value interface{}) bool {
	return true
}

func (s *StringDataType) GetValue(value interface{}) interface{} {
	return util.ToString(value)
}

func (s *StringDataType) Compare(left interface{}, right interface{}, validated bool) (int, error) {
	if !validated {
		leftValid := s.IsValid(left)
		rightValid := s.IsValid(right)
		if !leftValid || !rightValid {
			return 0, errors.INVALID_DATA_TYPE
		}
	}
	leftValue := s.GetValue(left).(string)
	rightValue := s.GetValue(right).(string)
	return strings.Compare(leftValue, rightValue), nil
}

func NewStringDataType() DataType {
	return &StringDataType{}
}
