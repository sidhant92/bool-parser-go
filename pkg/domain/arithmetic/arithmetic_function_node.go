package arithmetic

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type ArithmeticFunctionNode struct {
	FunctionType constant.FunctionType
	Items        []ArithmeticLeafNode
}

func (c ArithmeticFunctionNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC_FUNCTION
}
