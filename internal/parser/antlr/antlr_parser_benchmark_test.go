package parser

import (
	"github.com/sidhant92/bool-parser-go/pkg/domain/logical"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression1()
	}
}

func expression1() logical.Node {
	res, _ := parser.Parse("b>0 AND z IN ('c1', 'c2')")
	return res
}
