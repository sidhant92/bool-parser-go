package datatype

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

var dataTypeMap = map[constant.DataType]DataType{
	constant.STRING:  NewStringDataType(),
	constant.LONG:    NewLongDataType(),
	constant.INTEGER: NewIntDataType(),
	constant.DECIMAL: NewDecimalDataType(),
	constant.BOOLEAN: NewBooleanDataType(),
	constant.VERSION: NewVersionDataType(),
}

func GetDataType(dataType constant.DataType) DataType {
	return dataTypeMap[dataType]
}
