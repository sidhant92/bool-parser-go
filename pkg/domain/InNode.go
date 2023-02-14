package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type Pair struct {
	DataType constant.DataType
	Value    interface{}
}

type InNode struct {
	Field string
	Items []Pair
}

func (c InNode) GetNodeType() constant.NodeType {
	return constant.IN_NODE
}
