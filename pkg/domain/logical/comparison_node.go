package logical

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"slices"
)

type ComparisonNode struct {
	Field    string
	Value    Node
	Operator constant.Operator
	DataType constant.DataType
}

func (c ComparisonNode) GetNodeType() constant.NodeType {
	return constant.COMPARISON_NODE
}

func (c ComparisonNode) IsNullCheck() bool {
	return slices.Contains(constant.EqualityOperators, c.Operator) && c.Value.GetNodeType() == constant.FIELD_NODE
}
