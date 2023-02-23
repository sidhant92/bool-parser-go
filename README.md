# bool-parser-go
A Boolean Expression Parser for Go

The library can help parse complex and nested boolean expressions.
The expressions are in SQL-like syntax, where you can use boolean operators and parentheses to combine individual expressions.

An expression can be as simple as `name = Sidhant`.
A Complex expression is formed by combining these small expressions by logical operators and giving precedence using parenthesis

### Examples
#### Textual Equality

Format: `${attributeName} = ${value}`

Example: `name = john`

#### Numeric Comparisons

Format: `${attributeName} ${operator} ${value}`

Example: `price > 12.99`

The ${value} must be numeric. Supported operators are `<`, `<=`, `=`, `!=`, `>=` and `>`, with the same semantics as in virtually all programming languages.

#### Numeric Range

Format: `${attributeName} ${lowerBound} TO ${upperBound}`

Example: `price 5.99 TO 100`

`${lowerBound}` and `${upperBound}` must be numeric. Both are inclusive.

#### Boolean operators

Example:

`price < 10 AND (category:Book OR NOT category:Ebook)`

Individual filters can be combined via boolean operators. The following operators are supported:

* `OR`: must match any of the combined conditions (disjunction)
* `AND`: must match all of the combined conditions (conjunction)
* `NOT`: negate a filter

Parentheses, `(` and `)`, can be used for grouping.

#### Usage Notes
* Phrases that includes quotes, like `content = "It's a wonderful day"`
* Phrases that includes quotes, like `attribute = 'She said "Hello World"'`
* For nested keys in data map you can use the dot notation, like `person.age`

## Usage
```
go get github.com/sidhant92/bool-parser-go
```


Code
```
package main

import (
	"fmt"
	"github.com/sidhant92/bool-parser-go/pkg/application"
	"github.com/sidhant92/bool-parser-go/pkg/parser"
)

func main() {
	antlrParser := parser.New()
	evaluator := application.NewBooleanExpressionEvaluator(antlrParser)
	data := map[string]interface{}{
		"age": 25,
	}
	res, err := evaluator.Evaluate("age > 20", data)

	if err == nil {
		// true
		fmt.Println(res)
	}
}

```


## Applications

### Boolean Expression Evaluator

The library can be used to evaluate a boolean expression.

The following Data Types are supported:
1. String
2. Integer
3. Long
4. Decimal
5. Boolean
6. Semantic Version

Usage examples:

Simple Numerical Comparison
```
data := map[string]interface{}{
	"age": 26,
}
res, err := evaluator.Evaluate("age >= 27", data)
assert.Nil(t, err)
assert.False(t, res)
```
Boolean Comparison
```
data := map[string]interface{}{
	"age":  25,
	"name": "sid",
}
res, err := evaluator.Evaluate("name = sid AND age = 25", data)
assert.Nil(t, err)
assert.True(t, res)
```
Nested Boolean Comparison
```
data := map[string]interface{}{
	"age":  25,
	"name": "sid",
	"num":  45,
}
res, err := evaluator.Evaluate("name = sid AND (age = 25 OR num = 44)", data)
assert.Nil(t, err)
assert.True(t, res)
```
Semantic Version Comparison
```
data := map[string]interface{}{
	"app_version": "1.5.9",
}
res, err := evaluator.Evaluate("app_version < 1.5.10", data)
assert.Nil(t, err)
assert.True(t, res)
```

The return type is `res, err`. Failure means that parsing has failed and any fallback can be used.


[For a complete list of examples please check out the test file](pkg/application/boolean_expression_evaluator_test.go)