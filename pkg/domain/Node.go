package domain

import "bool-parser-go/pkg/constant"

type Node interface {
	GetNodeType() constant.NodeType
}
