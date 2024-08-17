package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type ArithmeticFunctionNode struct {
	FunctionType constant.FunctionType
	Items        []interface{} // array of nodes
}

func (c ArithmeticFunctionNode) GetNodeType() constant.NodeType {
	return constant.ARITHMETIC_FUNCTION
}
