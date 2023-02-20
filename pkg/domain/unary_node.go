package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type UnaryNode struct {
	DataType constant.DataType
	Value    interface{}
}

func (c UnaryNode) GetNodeType() constant.NodeType {
	return constant.UNARY_NODE
}
