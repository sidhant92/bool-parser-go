package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type ArrayNode struct {
	Field    string
	Operator constant.Operator
	Items    []interface{}
}

func (c ArrayNode) GetNodeType() constant.NodeType {
	return constant.ARRAY_NODE
}
