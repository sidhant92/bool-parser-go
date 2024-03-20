package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type ArithmeticUnaryNode struct {
	Operand domain.Node
}

func (c ArithmeticUnaryNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC_UNARY
}
