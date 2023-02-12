package parser

import (
	"bool-parser-go/internal/parser/antlr/lib"
	"bool-parser-go/pkg/domain"
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"log"
)

type ANTLRParser struct {
}

func (p ANTLRParser) Parse(input string) (res domain.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occurred:", r)
			res = nil
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	inputStream := antlr.NewInputStream(input)
	lexer := lib.NewBooleanExpressionLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	antlrParser := lib.NewBooleanExpressionParser(stream)
	listener := New()
	antlr.ParseTreeWalkerDefault.Walk(listener, antlrParser.Parse())
	return listener.GetResult(), nil
}
