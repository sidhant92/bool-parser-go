package application

import (
	"github.com/sidhant92/bool-parser-go/pkg/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

var arithmeticEvaluator = NewArithmeticExpressionEvaluator(parser.New())

func TestSimpleAddOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("5 + 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 10)
}

func TestSimpleSubtractOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("5 - 4", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 1)
}

func TestSimpleMultiplyOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("5 * 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 25)
}

func TestSimpleDivideOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("5 / 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 1)
}

func TestSimpleDivideOperationWithoutVariable1(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("5 / 4", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 1.25)
}

func TestSimpleModulusOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("10 % 4", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 2)
}

func TestSimpleExponentOperationWithoutVariable(t *testing.T) {
	data := map[string]interface{}{}
	res, err := arithmeticEvaluator.Evaluate("2 ^ 4", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 16)
}

func TestSimpleAddOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a + 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 15)
}

func TestSimpleSubtractOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a - 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 5)
}

func TestSimpleSubtractOperationWithTwoVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
		"b": 5,
	}
	res, err := arithmeticEvaluator.Evaluate("a-b", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 5)
}

func TestStringEvaluation(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
		"b": 5,
	}
	res, err := arithmeticEvaluator.Evaluate("\"a-b\"", data)
	assert.Nil(t, err)
	assert.Equal(t, res, "a-b")
}

func TestSimpleMultiplyOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a * 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 50)
}

func TestSimpleDivideOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a / 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 2)
}

func TestSimpleModulusOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a % 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 0)
}

func TestSimpleExponentOperationWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("a ^ 5", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 100000)
}

func TestStringConcatenation(t *testing.T) {
	data := map[string]interface{}{
		"a": "first",
		"b": "second",
	}
	res, err := arithmeticEvaluator.Evaluate("a + \" \" + b", data)
	assert.Nil(t, err)
	assert.Equal(t, res, "first second")
}

func TestExpressionWithParenthesisNoVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("((5 * 2) + 10) * 2 + (1 + 3 * (10 / 2))", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 56)
}

func TestExpressionWithParenthesisWithVariable(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("((5 * 2) + a) * 2 + (1 + 3 * (a / 2))", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 56)
}

func TestExpressionWithComparison(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("(a > (4 + 1))", data)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestExpressionWithUnaryNode(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("- a", data)
	assert.Nil(t, err)
	assert.Equal(t, res, -10)
}

func TestExpressionWithNestedUnaryNode(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("- (a + 5)", data)
	assert.Nil(t, err)
	assert.Equal(t, res, -15)
}

func TestExpressionWithNestedUnaryNode1(t *testing.T) {
	data := map[string]interface{}{
		"a": "test",
	}
	res, err := arithmeticEvaluator.Evaluate("a", data)
	assert.Nil(t, err)
	assert.Equal(t, res, "test")
}

func TestUnaryNode(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("-a", data)
	assert.Nil(t, err)
	assert.Equal(t, res, -10)
}

func TestDoubleUnaryNode(t *testing.T) {
	data := map[string]interface{}{
		"a": 10,
	}
	res, err := arithmeticEvaluator.Evaluate("--a", data)
	assert.Nil(t, err)
	assert.Equal(t, res, 10)
}
