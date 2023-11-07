package containerdatatype

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"golang.org/x/exp/maps"
)

var containerDataTypeMap = map[constant.ContainerDataType]ContainerDataType{
	constant.PRIMITIVE: NewPrimitiveContainerDataType(),
	constant.LIST:      NewListContainerDataType(),
}

func GetContainerDataType(containerDataType constant.ContainerDataType) ContainerDataType {
	return containerDataTypeMap[containerDataType]
}

func GetAllDataTypes() []ContainerDataType {
	return maps.Values(containerDataTypeMap)
}
