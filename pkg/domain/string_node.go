package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type StringNode struct {
	Field string
}

func (c StringNode) GetNodeType() constant.NodeType {
	return constant.STRING_NODE
}
