package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type ArithmeticNode struct {
	Left domain.Node

	Right domain.Node

	Operator constant.Operator
}

func (c ArithmeticNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC
}
