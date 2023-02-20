package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type NumericRangeNode struct {
	Field        string
	FromValue    interface{}
	ToValue      interface{}
	FromDataType constant.DataType
	ToDataType   constant.DataType
}

func (c NumericRangeNode) GetNodeType() constant.NodeType {
	return constant.NUMERIC_RANGE_NODE
}
