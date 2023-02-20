package datatype

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
	"reflect"
)

type VersionDataType struct {
}

func (s *VersionDataType) GetDataType() constant.DataType {
	return constant.VERSION
}

func (s *VersionDataType) IsValid(value interface{}) bool {
	if reflect.TypeOf(value).Name() == "version.Version" {
		return true
	}
	_, err := util.ConvertValue(fmt.Sprintf("%v", value), constant.VERSION)
	if err != nil {
		return false
	}
	return true
}

func (s *VersionDataType) GetValue(value interface{}) interface{} {
	val, ok := value.(version.Version)
	if ok {
		return val
	}
	res, _ := util.ConvertValue(fmt.Sprintf("%v", value), constant.VERSION)
	return res
}

func (s *VersionDataType) Compare(left interface{}, right interface{}) (int, error) {
	leftValid := s.IsValid(left)
	rightValid := s.IsValid(right)
	if !leftValid || !rightValid {
		return 0, errors.INVALID_DATA_TYPE
	}
	leftValue := s.GetValue(left).(*version.Version)
	rightValue := s.GetValue(right).(*version.Version)
	return leftValue.Compare(rightValue), nil
}

func NewVersionDataType() DataType {
	return &VersionDataType{}
}
