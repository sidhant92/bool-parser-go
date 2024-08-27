package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type AbstractOperator interface {
	Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error)

	GetSymbol() string

	GetOperator() constant.Operator

	GetAllowedContainerTypes() []constant.ContainerDataType

	GetAllowedDataTypes() []constant.DataType
}
