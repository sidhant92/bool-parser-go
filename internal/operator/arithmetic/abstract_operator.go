package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type AbstractOperator interface {
	Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightRightData constant.DataType) (interface{}, error)

	GetSymbol() string

	GetOperator() constant.Operator

	GetAllowedContainerTypes() []constant.ContainerDataType

	GetAllowedDataTypes() []constant.DataType
}
