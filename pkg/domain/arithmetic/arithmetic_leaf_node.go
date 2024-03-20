package arithmetic

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type ArithmeticLeafNode struct {
	Operand  interface{}
	DataType constant.DataType
}

func (c ArithmeticLeafNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC_LEAF
}
