package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain/logical"
)

type ArithmeticNode struct {
	Left logical.Node

	Right logical.Node

	Operator constant.Operator
}

func (c ArithmeticNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC
}
