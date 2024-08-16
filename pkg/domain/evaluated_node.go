package domain

import "github.com/sidhant92/bool-parser-go/pkg/constant"

type EvaluatedNode struct {
	Value    interface{}
	DataType constant.DataType
}
