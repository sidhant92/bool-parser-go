package containerdatatype

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type ListContainerDataType struct {
}

func (s *ListContainerDataType) GetContainerDataType() constant.ContainerDataType {
	return constant.LIST
}

func (s *ListContainerDataType) IsValid(dataType constant.DataType, value interface{}) bool {
	if !util.IsSlice(value) {
		return false
	}
	var slice = util.GetSliceFromInterface(value)
	for _, value := range slice {
		if !datatype.GetDataType(dataType).IsValid(value) {
			return false
		}
	}
	return true
}

func NewListContainerDataType() ContainerDataType {
	return &ListContainerDataType{}
}
