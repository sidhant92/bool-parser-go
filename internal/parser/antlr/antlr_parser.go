package parser

import (
	"bool-parser-go/internal/parser/antlr/lib"
	"bool-parser-go/pkg/domain"
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	lru "github.com/hashicorp/golang-lru/v2"
	"log"
)

type ANTLRParser struct {
	UseCache bool
	cache    *lru.Cache[string, domain.Node]
}

func Default() *ANTLRParser {
	return &ANTLRParser{
		UseCache: false,
		cache:    nil,
	}
}

func Cached(size int) *ANTLRParser {
	cache, _ := lru.New[string, domain.Node](size)
	return &ANTLRParser{
		UseCache: true,
		cache:    cache,
	}
}

func (p *ANTLRParser) Parse(input string) (res domain.Node, err error) {
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
				err = errors.New("unknown panic")
			}
		}
	}()

	if p.UseCache {
		val, ok := p.cache.Get(input)
		if ok {
			return val, nil
		}
	}

	inputStream := antlr.NewInputStream(input)
	lexer := lib.NewBooleanExpressionLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	antlrParser := lib.NewBooleanExpressionParser(stream)
	listener := New()
	antlr.ParseTreeWalkerDefault.Walk(listener, antlrParser.Parse())
	if p.UseCache {
		p.cache.Add(input, listener.GetResult())
	}
	return listener.GetResult(), nil
}
