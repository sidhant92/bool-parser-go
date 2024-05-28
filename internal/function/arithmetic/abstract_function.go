package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type AbstractFunction interface {
	Evaluate(fields []domain.EvaluatedNode) (interface{}, error)

	GetFunctionType() constant.FunctionType

	GetAllowedContainerTypes() []constant.ContainerDataType

	GetAllowedDataTypes() []constant.DataType
}
