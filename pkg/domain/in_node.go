package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type InNode struct {
	Field string
	Items []interface{}
}

func (c InNode) GetNodeType() constant.NodeType {
	return constant.IN_NODE
}
