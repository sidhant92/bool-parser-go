package containerdatatype

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type PrimitiveContainerDataType struct {
}

func (s *PrimitiveContainerDataType) GetContainerDataType() constant.ContainerDataType {
	return constant.PRIMITIVE
}

func (s *PrimitiveContainerDataType) IsValid(dataType constant.DataType, value interface{}) bool {
	if util.IsSlice(value) {
		return false
	}
	return datatype.GetDataType(dataType).IsValid(value)
}

func NewPrimitiveContainerDataType() ContainerDataType {
	return &PrimitiveContainerDataType{}
}
