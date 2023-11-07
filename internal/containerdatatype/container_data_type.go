package containerdatatype

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type ContainerDataType interface {
	IsValid(dataType constant.DataType, value interface{}) bool

	GetContainerDataType() constant.ContainerDataType
}
