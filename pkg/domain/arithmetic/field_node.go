package arithmetic

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type FieldNode struct {
	Field string
}

func (f FieldNode) GetNodeType() constant.NodeType {
	return constant.FIELD_NODE
}
