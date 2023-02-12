package domain

import "bool-parser-go/pkg/constant"

type ComparisonNode struct {
	Field    string
	Value    interface{}
	Operator constant.Operator
	DataType constant.DataType
}

func (c ComparisonNode) GetNodeType() constant.NodeType {
	return constant.COMPARISON_NODE
}
