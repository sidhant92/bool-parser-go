package application

import (
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"github.com/sidhant92/bool-parser-go/pkg/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

var evaluator = NewBooleanExpressionEvaluator(parser.New())

func TestUnaryExpressionWithBoolean(t *testing.T) {
	data := map[string]interface{}{
		"is_allowed": true,
	}
	res, err := evaluator.Evaluate("not false", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestUnaryExpressionWithField(t *testing.T) {
	data := map[string]interface{}{
		"is_allowed": true,
	}
	res, err := evaluator.Evaluate("not is_allowed", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestInvalidUnaryOperation(t *testing.T) {
	data := map[string]interface{}{
		"is_allowed": 123,
	}
	_, err := evaluator.Evaluate("not is_allowed", data)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errors2.INVALID_UNARY_OPERAND)
}

func TestSimpleTrueCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"name": "abc",
	}
	res, err := evaluator.Evaluate("name = 'abc'", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestSimpleTrueCorrectExpressions(t *testing.T) {
	data := map[string]interface{}{
		"name": "abc-",
	}
	res, err := evaluator.Evaluate("name = \"abc-\"", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestInvalidDataType(t *testing.T) {
	data := map[string]interface{}{
		"name": "abc-",
	}
	_, err := evaluator.Evaluate("name > 123", data)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errors2.INVALID_DATA_TYPE)
}

func TestSimpleFalseIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"name": "def",
	}
	res, err := evaluator.Evaluate("name = 'abc'", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNumericEqualCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age = 24", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNumericEqualIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 26,
	}
	res, err := evaluator.Evaluate("age = 24", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNumericGreaterThanCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age > 20", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNumericGreaterThanCorrectExpressionField(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
		"b":   20,
	}
	res, err := evaluator.Evaluate("age > b", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNestedField(t *testing.T) {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"age": 24,
		},
	}
	res, err := evaluator.Evaluate("person.age > 20", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestTwoNestedField(t *testing.T) {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"details": map[string]interface{}{
				"age": 24,
			},
		},
	}
	res, err := evaluator.Evaluate("person.details.age > 20", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestMissingNestedField(t *testing.T) {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"age": 24,
		},
	}
	res, err := evaluator.Evaluate("person.agee > 20", data)
	assert.ErrorIs(t, err, errors2.KEY_DATA_NOT_PRESENT)
	assert.False(t, res)
}

func TestNumericGreaterThanIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 26,
	}
	res, err := evaluator.Evaluate("age > 27", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNumericGreaterThanEqualCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age >= 20", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNumericGreaterThanEqualIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 26,
	}
	res, err := evaluator.Evaluate("age >= 27", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNumericLessThanEqualCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age < 30", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNumericLessThanEqualIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 26,
	}
	res, err := evaluator.Evaluate("age <= 20", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNumericNotEqualCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age != 30", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNumericNotEqualIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 26,
	}
	res, err := evaluator.Evaluate("age != 26", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestSimpleNotStringExpression(t *testing.T) {
	data := map[string]interface{}{
		"name": "abc",
	}
	res, err := evaluator.Evaluate("NOT (name = 'abc')", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestSimpleNotNumericExpression(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("NOT (age = 24)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestComplexAndCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
	}
	res, err := evaluator.Evaluate("name = 'sid' AND age = 25", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestComplexAndIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
	}
	res, err := evaluator.Evaluate("name = 'sid' AND age = 23", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestComplexORCorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
	}
	res, err := evaluator.Evaluate("name = 'sid' OR age = 23", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestComplexORIncorrectExpression(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sidh",
	}
	res, err := evaluator.Evaluate("name = 'sid' OR age = 23", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestCorrectComplexExpressionWithParenthesis(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
		"num":  45,
	}
	res, err := evaluator.Evaluate("name = 'sid' AND (age = 25 OR num = 44)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNegativeInClauseForIntegers(t *testing.T) {
	data := map[string]interface{}{
		"age": 25,
	}
	res, err := evaluator.Evaluate("age in (26,56,34)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestPositiveInClauseForIntegers(t *testing.T) {
	data := map[string]interface{}{
		"age": 25,
	}
	res, err := evaluator.Evaluate("age in (26.2,25,34)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestPositiveInClauseForIntegersField(t *testing.T) {
	data := map[string]interface{}{
		"age": 25,
		"b":   25,
	}
	res, err := evaluator.Evaluate("age in (26.2,27,34,b)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestPositiveInClauseForDecimals(t *testing.T) {
	data := map[string]interface{}{
		"num": 25.3,
	}
	res, err := evaluator.Evaluate("num in (26.3,25.3,34.4)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestPositiveNotInClauseForDecimals(t *testing.T) {
	data := map[string]interface{}{
		"num": 25.3,
	}
	res, err := evaluator.Evaluate("num not in (26.3,26.3,34.4)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNegativeInClauseForStrings(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
	}
	res, err := evaluator.Evaluate("name in ('tes', 'abc')", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestPositiveInClauseForStrings(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
	}
	res, err := evaluator.Evaluate("name in ('abc', 'test')", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestCorrectComplexExpressionWithParenthesis1(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
		"num":  45,
	}
	res, err := evaluator.Evaluate("name = 'sidh' OR (age = 25 AND num = 45)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestIncorrectComplexExpressionWithParenthesis(t *testing.T) {
	data := map[string]interface{}{
		"age":  25,
		"name": "sid",
		"num":  45,
	}
	res, err := evaluator.Evaluate("name = 'sid' AND (age = 23 OR num = 44)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestWrongDataType(t *testing.T) {
	data := map[string]interface{}{
		"age": 24,
	}
	res, err := evaluator.Evaluate("age = 'dsf'", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestWrongDataType1(t *testing.T) {
	data := map[string]interface{}{
		"age": "sf",
	}
	_, err := evaluator.Evaluate("age = 24", data)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errors2.INVALID_DATA_TYPE)
}

func TestKeyMissing(t *testing.T) {
	data := map[string]interface{}{
		"agee": 34,
	}
	res, err := evaluator.Evaluate("age = 24", data)
	assert.ErrorIs(t, err, errors2.KEY_DATA_NOT_PRESENT)
	assert.False(t, res)
}

func TestAppVersionGraterThan(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.0.6",
	}
	res, err := evaluator.Evaluate("app_version > 1.0.5", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestAppVersionGraterThan1(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.0.6.15",
	}
	res, err := evaluator.Evaluate("app_version > 1.0.6.14", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestAppVersionGraterThan2(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.54",
	}
	res, err := evaluator.Evaluate("app_version > 1.53", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestAppVersionGraterThanEqualTo(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.0.6",
	}
	res, err := evaluator.Evaluate("app_version >= 1.0.6", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestAppVersionLessThan(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.5.9",
	}
	res, err := evaluator.Evaluate("app_version < 1.5.10", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestAppVersionLessThanEqualTo(t *testing.T) {
	data := map[string]interface{}{
		"app_version": "1.5.9",
	}
	res, err := evaluator.Evaluate("app_version <= 1.5.9", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestStringEqualityWithQuotes(t *testing.T) {
	data := map[string]interface{}{
		"name": "sidhant aggarwal",
	}
	res, err := evaluator.Evaluate("name = 'sidhant aggarwal'", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestDefaultFieldTrue(t *testing.T) {
	data := map[string]interface{}{
		"age": 19,
	}
	res, err := evaluator.Evaluate(">= 18 AND < 20", data, "age")
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestDefaultFieldFalse(t *testing.T) {
	data := map[string]interface{}{
		"age": 17,
	}
	res, err := evaluator.Evaluate(">= 18 AND < 20", data, "age")
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNegativeComparison(t *testing.T) {
	data := map[string]interface{}{
		"a": -6,
	}
	res, err := evaluator.Evaluate("a > -10 AND a < -2", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestContainsAnyTrueCondition(t *testing.T) {
	data := map[string]interface{}{
		"age": []int32{1, 2, 3},
	}
	res, err := evaluator.Evaluate("age contains_any (2)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestContainsAnyFalseCondition(t *testing.T) {
	data := map[string]interface{}{
		"age": []int32{1, 2, 3},
	}
	res, err := evaluator.Evaluate("age contains_any (12)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestContainsAllTrueCondition(t *testing.T) {
	data := map[string]interface{}{
		"age": []int32{1, 2, 3},
	}
	res, err := evaluator.Evaluate("age contains_all (1,2)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestContainsAllFalseCondition(t *testing.T) {
	data := map[string]interface{}{
		"age": []int32{1, 2, 3},
	}
	res, err := evaluator.Evaluate("age contains_all (2,5)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestContainsAllFailureCondition(t *testing.T) {
	data := map[string]interface{}{
		"age": "sid",
	}
	_, err := evaluator.Evaluate("age contains_all (2,5)", data)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errors2.INVALID_DATA_TYPE)
}

func TestComparisonWithArithmeticFunction(t *testing.T) {
	data := map[string]interface{}{
		"age": 20,
	}
	res, err := evaluator.Evaluate("age > min (18, 25, 30)", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestComparisonWithArithmeticVariableFunction(t *testing.T) {
	data := map[string]interface{}{
		"age":     20,
		"numbers": []int{1, 1, 2, 30},
	}
	res, err := evaluator.Evaluate("age > max (numbers)", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNullCheck(t *testing.T) {
	data := map[string]interface{}{
		"a":     2.7,
	}
	res, err := evaluator.Evaluate("b = null", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestNullCheck1(t *testing.T) {
	data := map[string]interface{}{
		"a":     2.7,
	}
	res, err := evaluator.Evaluate("a = null", data)
	assert.Nil(t, err)
	assert.False(t, res)
}

func TestNotNullCheck(t *testing.T) {
	data := map[string]interface{}{
		"a":     2.7,
	}
	res, err := evaluator.Evaluate("a != null", data)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestBooleanNullCheck(t *testing.T) {
	data := map[string]interface{}{
		"a":     3,
	}
	res, err := evaluator.Evaluate("b = null && a > 2", data)
	assert.Nil(t, err)
	assert.True(t, res)
}
