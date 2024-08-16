package logical

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type Node interface {
	GetNodeType() constant.NodeType
}
