package service

import "github.com/sidhant92/bool-parser-go/pkg/constant"

var Operators = map[string]constant.Operator{
	"=":  constant.EQUALS,
	">":  constant.GREATER_THAN,
	">=": constant.GREATER_THAN_EQUAL,
	"<":  constant.LESS_THAN,
	"<=": constant.LESS_THAN_EQUAL,
	"!=": constant.NOT_EQUAL,
	"IN": constant.IN,
}

func GetOperatorFromSymbol(symbol string) constant.Operator {
	if val, ok := Operators[symbol]; ok {
		return val
	}
	panic("Unknown Operator " + symbol)
}
