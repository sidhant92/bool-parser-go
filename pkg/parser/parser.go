package parser

import (
	parser "bool-parser-go/internal/parser/antlr"
	"bool-parser-go/pkg/domain"
)

type Parser interface {
	Parse(input string) (domain.Node, error)
}

func New() Parser {
	return parser.Default()
}

func NewWithCache(size int) Parser {
	return parser.Cached(size)
}
