package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"strings"
)

type FieldNode struct {
	Field string
}

func (f FieldNode) GetNodeType() constant.NodeType {
	return constant.FIELD_NODE
}

func (f FieldNode) IsNull() bool {
	return strings.ToLower(f.Field) == "null"
}
