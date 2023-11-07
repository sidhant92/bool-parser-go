package operator

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type AbstractOperator interface {
	Evaluate(containerDataType constant.ContainerDataType, dataType constant.DataType, validated bool, left interface{}, right ...interface{}) (bool, error)

	GetSymbol() string

	GetOperator() constant.Operator

	GetAllowedContainerTypes() []constant.ContainerDataType

	GetAllowedDataTypes() []constant.DataType
}
