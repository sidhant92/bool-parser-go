package parser

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/sidhant92/bool-parser-go/internal/parser/antlr/lib"
	"github.com/sidhant92/bool-parser-go/pkg/domain/logical"
	"log"
)

type ANTLRParser struct {
	UseCache bool
	cache    *lru.Cache[string, logical.Node]
}

func Default() *ANTLRParser {
	return &ANTLRParser{
		UseCache: false,
		cache:    nil,
	}
}

func Cached(size int) *ANTLRParser {
	cache, _ := lru.New[string, logical.Node](size)
	return &ANTLRParser{
		UseCache: true,
		cache:    cache,
	}
}

func (p *ANTLRParser) Parse(input string, defaultField ...string) (res logical.Node, err error) {
	cacheKey := input
	if len(defaultField) > 0 {
		cacheKey = cacheKey + "_" + defaultField[0]
	}
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
		val, ok := p.cache.Get(cacheKey)
		if ok {
			return val, nil
		}
	}

	errorListener := NewCustomErrorListener()
	inputStream := antlr.NewInputStream(input)
	lexer := lib.NewBooleanExpressionLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	antlrParser := lib.NewBooleanExpressionParser(stream)
	antlrParser.RemoveErrorListeners()
	antlrParser.AddErrorListener(errorListener)
	var df = ""
	if len(defaultField) > 0 {
		df = defaultField[0]
	}
	parseResult := antlrParser.Parse()
	listener := New(df, errorListener.errors)
	antlr.ParseTreeWalkerDefault.Walk(listener, parseResult)

	if p.UseCache {
		p.cache.Add(cacheKey, listener.GetResult())
	}
	return listener.GetResult(), nil
}
