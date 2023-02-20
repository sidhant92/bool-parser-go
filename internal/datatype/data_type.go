package datatype

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type DataType interface {
	IsValid(value interface{}) bool

	GetDataType() constant.DataType

	GetValue(value interface{}) any

	Compare(left interface{}, right interface{}) (int, error)
}
