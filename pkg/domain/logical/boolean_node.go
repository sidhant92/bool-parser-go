package logical

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type BooleanNode struct {
	Left     Node
	Right    Node
	Operator constant.LogicalOperatorType
}

func (c BooleanNode) GetNodeType() constant.NodeType {
	return constant.BOOLEAN_NODE
}
