package parser

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"github.com/sidhant92/bool-parser-go/pkg/domain/arithmetic"
	"github.com/stretchr/testify/assert"
	"testing"
)

var parser = Default()

func TestUnaryTokenBoolean(t *testing.T) {
	res, _ := parser.Parse("not false")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	booleanNode := res.(domain.BooleanNode)
	assert.Equal(t, booleanNode.Operator, constant.NOT)
	assert.Equal(t, booleanNode.Left.GetNodeType(), constant.UNARY_NODE)
	unaryNode := booleanNode.Left.(domain.UnaryNode)
	assert.Equal(t, unaryNode.DataType, constant.BOOLEAN)
	assert.Equal(t, unaryNode.Value, false)
}

func TestUnaryTokenString(t *testing.T) {
	res, _ := parser.Parse("not abc")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	booleanNode := res.(domain.BooleanNode)
	assert.Equal(t, booleanNode.Operator, constant.NOT)
	assert.Equal(t, booleanNode.Left.GetNodeType(), constant.UNARY_NODE)
	unaryNode := booleanNode.Left.(domain.UnaryNode)
	assert.Equal(t, unaryNode.DataType, constant.STRING)
	assert.Equal(t, unaryNode.Value, "abc")
}

func TestSingleStringToken(t *testing.T) {
	res, _ := parser.Parse("name = test")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.STRING)
	assert.Equal(t, res.(domain.ComparisonNode).Value, "test")
}

func TestSingleStringTokenWithSingleQuotes(t *testing.T) {
	res, _ := parser.Parse("name = \"te'st\"")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.STRING)
	assert.Equal(t, res.(domain.ComparisonNode).Value, "te'st")
}

func TestSingleStringTokenWithDoubleQuotes(t *testing.T) {
	res, _ := parser.Parse("name = 'te\"st'")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.STRING)
	assert.Equal(t, res.(domain.ComparisonNode).Value, "te\"st")
}

func TestSingleStringTokenWithSpace(t *testing.T) {
	res, _ := parser.Parse("name = \"first second\"")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.STRING)
	assert.Equal(t, res.(domain.ComparisonNode).Value, "first second")
}

func TestSingleIntToken(t *testing.T) {
	res, _ := parser.Parse("age = 44")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.ComparisonNode).Value, 44)
}

func TestSingleLongToken(t *testing.T) {
	res, _ := parser.Parse("age=1611473334456113434")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.ComparisonNode).Value, 1611473334456113434)
}

func TestSingleDecimalToken(t *testing.T) {
	res, _ := parser.Parse("age=44.34")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.ComparisonNode).DataType, constant.DECIMAL)
	assert.Equal(t, res.(domain.ComparisonNode).Value, 44.34)
}

func TestSingleIntRangeToken(t *testing.T) {
	res, _ := parser.Parse("age 18 TO 44")
	assert.Equal(t, res.GetNodeType(), constant.NUMERIC_RANGE_NODE)
	assert.Equal(t, res.(domain.NumericRangeNode).Field, "age")
	assert.Equal(t, res.(domain.NumericRangeNode).FromValue, 18)
	assert.Equal(t, res.(domain.NumericRangeNode).FromDataType, constant.INTEGER)
	assert.Equal(t, res.(domain.NumericRangeNode).ToValue, 44)
	assert.Equal(t, res.(domain.NumericRangeNode).ToDataType, constant.INTEGER)
}

func TestSingleDecimalRangeToken(t *testing.T) {
	res, _ := parser.Parse("age 18.4 TO 44.2")
	assert.Equal(t, res.GetNodeType(), constant.NUMERIC_RANGE_NODE)
	assert.Equal(t, res.(domain.NumericRangeNode).Field, "age")
	assert.Equal(t, res.(domain.NumericRangeNode).FromValue, 18.4)
	assert.Equal(t, res.(domain.NumericRangeNode).FromDataType, constant.DECIMAL)
	assert.Equal(t, res.(domain.NumericRangeNode).ToValue, 44.2)
	assert.Equal(t, res.(domain.NumericRangeNode).ToDataType, constant.DECIMAL)
}

func TestGreaterThan(t *testing.T) {
	res, _ := parser.Parse("age > 18")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	assert.Equal(t, res.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.ComparisonNode).Operator, constant.GREATER_THAN)
	assert.Equal(t, res.(domain.ComparisonNode).Value, 18)
}

func TestSimpleOrCondition(t *testing.T) {
	res, _ := parser.Parse("name = test OR age=33")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.OR)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, "test")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, 33)
}

func TestSimpleAndCondition(t *testing.T) {
	res, _ := parser.Parse("name = test AND age=333")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.AND)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, "test")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, 333)
}

func TestSimpleNotCondition(t *testing.T) {
	res, _ := parser.Parse("NOT (name = test)")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.NOT)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, "test")
	assert.Nil(t, res.(domain.BooleanNode).Right)
}

func TestNestedAndCondition(t *testing.T) {
	res, _ := parser.Parse("name = test OR (age=33 AND city = dummy)")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.OR)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, "test")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Operator, constant.AND)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, 33)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "city")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, "dummy")
}

func TestNestedAndCondition1(t *testing.T) {
	res, _ := parser.Parse("(agel=44 AND cityl = abc) OR (ager=33 AND cityr = dummy)")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.OR)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Operator, constant.AND)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "agel")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, 44)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "cityl")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, "abc")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Operator, constant.AND)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "ager")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, 33)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "cityr")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, "dummy")
}

func TestNestedOrCondition(t *testing.T) {
	res, _ := parser.Parse("name = test AND (age=33 OR city = dummy)")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.AND)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "name")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, "test")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Operator, constant.OR)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Field, "age")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Left.(domain.ComparisonNode).Value, 33)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Field, "city")
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Operator, constant.EQUALS)
	assert.Equal(t, res.(domain.BooleanNode).Right.(domain.BooleanNode).Right.(domain.ComparisonNode).Value, "dummy")
}

func TestIntegerList(t *testing.T) {
	res, _ := parser.Parse("age IN (12,45)")
	assert.Equal(t, res.GetNodeType(), constant.IN_NODE)
	assert.Equal(t, res.(domain.InNode).Field, "age")
	assert.Equal(t, res.(domain.InNode).Items[0].DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.InNode).Items[0].Value, 12)
	assert.Equal(t, res.(domain.InNode).Items[1].DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.InNode).Items[1].Value, 45)
}

func TestIntegerNotInList(t *testing.T) {
	res, _ := parser.Parse("age not IN (12,45)")
	assert.Equal(t, res.GetNodeType(), constant.BOOLEAN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Operator, constant.NOT)
	assert.Nil(t, res.(domain.BooleanNode).Right)
	assert.Equal(t, res.(domain.BooleanNode).Left.GetNodeType(), constant.IN_NODE)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.InNode).Field, "age")
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.InNode).Items[0].DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.InNode).Items[0].Value, 12)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.InNode).Items[1].DataType, constant.INTEGER)
	assert.Equal(t, res.(domain.BooleanNode).Left.(domain.InNode).Items[1].Value, 45)
}

func TestStringList(t *testing.T) {
	res, _ := parser.Parse("name IN (abc, def, 'abc def')")
	assert.Equal(t, res.GetNodeType(), constant.IN_NODE)
	assert.Equal(t, res.(domain.InNode).Field, "name")
	assert.Equal(t, res.(domain.InNode).Items[0].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[0].Value, "abc")
	assert.Equal(t, res.(domain.InNode).Items[1].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[1].Value, "def")
	assert.Equal(t, res.(domain.InNode).Items[2].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[2].Value, "abc def")
}

func TestStringList1(t *testing.T) {
	res, _ := parser.Parse("name IN (abc, def, 'abc, def'")
	assert.Equal(t, res.GetNodeType(), constant.IN_NODE)
	assert.Equal(t, res.(domain.InNode).Field, "name")
	assert.Equal(t, res.(domain.InNode).Items[0].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[0].Value, "abc")
	assert.Equal(t, res.(domain.InNode).Items[1].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[1].Value, "def")
	assert.Equal(t, res.(domain.InNode).Items[2].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[2].Value, "abc, def")
}

func TestStringList2(t *testing.T) {
	res, _ := parser.Parse("name IN (abc, def, 'ab\"c')")
	assert.Equal(t, res.GetNodeType(), constant.IN_NODE)
	assert.Equal(t, res.(domain.InNode).Field, "name")
	assert.Equal(t, res.(domain.InNode).Items[0].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[0].Value, "abc")
	assert.Equal(t, res.(domain.InNode).Items[1].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[1].Value, "def")
	assert.Equal(t, res.(domain.InNode).Items[2].DataType, constant.STRING)
	assert.Equal(t, res.(domain.InNode).Items[2].Value, "ab\"c")
}

func TestInvalidExpression(t *testing.T) {
	_, err := parser.Parse("a")
	assert.Nil(t, err)
}

func TestInvalidNotExpression(t *testing.T) {
	_, err := parser.Parse("not a > 5")
	assert.NotNil(t, err)
}

func TestContainsAny(t *testing.T) {
	res, _ := parser.Parse("a contains_any (1,2,3)")
	assert.Equal(t, res.GetNodeType(), constant.ARRAY_NODE)
	assert.Equal(t, res.(domain.ArrayNode).Field, "a")
	assert.Equal(t, res.(domain.ArrayNode).Operator, constant.CONTAINS_ANY)
	assert.Equal(t, len(res.(domain.ArrayNode).Items), 3)
	assert.Equal(t, res.(domain.ArrayNode).Items[0].Value, 1)
}

func TestContainsAll(t *testing.T) {
	res, _ := parser.Parse("a CONTAINS_ALL (\"a\", \"b\")")
	assert.Equal(t, res.GetNodeType(), constant.ARRAY_NODE)
	assert.Equal(t, res.(domain.ArrayNode).Field, "a")
	assert.Equal(t, res.(domain.ArrayNode).Operator, constant.CONTAINS_ALL)
	assert.Equal(t, len(res.(domain.ArrayNode).Items), 2)
	assert.Equal(t, res.(domain.ArrayNode).Items[0].Value, "a")
}

func TestAddOperatorString(t *testing.T) {
	res, _ := parser.Parse("a + b")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC)
	arithmeticNode := res.(*arithmetic.ArithmeticNode)
	left := arithmeticNode.Left.(*domain.UnaryNode)
	right := arithmeticNode.Right.(*domain.UnaryNode)
	assert.Equal(t, left.Value, "a")
	assert.Equal(t, left.DataType, constant.STRING)
	assert.Equal(t, right.Value, "b")
	assert.Equal(t, right.DataType, constant.STRING)
	assert.Equal(t, arithmeticNode.Operator, constant.ADD)
}

func TestAddOperatorInt(t *testing.T) {
	res, _ := parser.Parse("20 + 5")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC)
	arithmeticNode := res.(*arithmetic.ArithmeticNode)
	left := arithmeticNode.Left.(*domain.UnaryNode)
	right := arithmeticNode.Right.(*domain.UnaryNode)
	assert.Equal(t, left.Value, 20)
	assert.Equal(t, left.DataType, constant.INTEGER)
	assert.Equal(t, right.Value, 5)
	assert.Equal(t, right.DataType, constant.INTEGER)
	assert.Equal(t, arithmeticNode.Operator, constant.ADD)
}

func TestAddOperatorDecimal(t *testing.T) {
	res, _ := parser.Parse("20.5 + 5")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC)
	arithmeticNode := res.(*arithmetic.ArithmeticNode)
	left := arithmeticNode.Left.(*domain.UnaryNode)
	right := arithmeticNode.Right.(*domain.UnaryNode)
	assert.Equal(t, left.Value, 20.5)
	assert.Equal(t, left.DataType, constant.DECIMAL)
	assert.Equal(t, right.Value, 5)
	assert.Equal(t, right.DataType, constant.INTEGER)
	assert.Equal(t, arithmeticNode.Operator, constant.ADD)
}

func TestAddOperatorBool(t *testing.T) {
	res, _ := parser.Parse("false + 5")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC)
	arithmeticNode := res.(*arithmetic.ArithmeticNode)
	left := arithmeticNode.Left.(*domain.UnaryNode)
	right := arithmeticNode.Right.(*domain.UnaryNode)
	assert.Equal(t, left.Value, false)
	assert.Equal(t, left.DataType, constant.BOOLEAN)
	assert.Equal(t, right.Value, 5)
	assert.Equal(t, right.DataType, constant.INTEGER)
	assert.Equal(t, arithmeticNode.Operator, constant.ADD)
}

func TestComparisonWithArithmetic(t *testing.T) {
	res, _ := parser.Parse("a > (10 + 20)")
	assert.Equal(t, res.GetNodeType(), constant.COMPARISON_NODE)
	comparisonNode := res.(domain.ComparisonNode)
	assert.Equal(t, comparisonNode.Field, "a")
	arithmeticNode := comparisonNode.Value.(*arithmetic.ArithmeticNode)
	left := arithmeticNode.Left.(*domain.UnaryNode)
	right := arithmeticNode.Right.(*domain.UnaryNode)
	assert.Equal(t, arithmeticNode.Operator, constant.ADD)
	assert.Equal(t, left.Value, 10)
	assert.Equal(t, left.DataType, constant.INTEGER)
	assert.Equal(t, right.Value, 20)
	assert.Equal(t, right.DataType, constant.INTEGER)
}

func TestArithmeticFunction(t *testing.T) {
	res, _ := parser.Parse("min (1,2,3)")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC_FUNCTION)
	arithmeticFunctionNode := res.(*arithmetic.ArithmeticFunctionNode)
	assert.Equal(t, arithmeticFunctionNode.FunctionType, constant.MIN)
	assert.Equal(t, len(arithmeticFunctionNode.Items), 3)
	assert.Equal(t, arithmeticFunctionNode.Items[0].(*domain.UnaryNode).DataType, constant.INTEGER)
	assert.Equal(t, arithmeticFunctionNode.Items[0].(*domain.UnaryNode).Value, 1)
	assert.Equal(t, arithmeticFunctionNode.Items[1].(*domain.UnaryNode).DataType, constant.INTEGER)
	assert.Equal(t, arithmeticFunctionNode.Items[1].(*domain.UnaryNode).Value, 2)
	assert.Equal(t, arithmeticFunctionNode.Items[2].(*domain.UnaryNode).DataType, constant.INTEGER)
	assert.Equal(t, arithmeticFunctionNode.Items[2].(*domain.UnaryNode).Value, 3)
}

func TestArithmeticFunctionWithSubstitution(t *testing.T) {
	res, _ := parser.Parse("min(abc)")
	assert.Equal(t, res.GetNodeType(), constant.ARITHMETIC_FUNCTION)
	arithmeticFunctionNode := res.(*arithmetic.ArithmeticFunctionNode)
	assert.Equal(t, arithmeticFunctionNode.FunctionType, constant.MIN)
	assert.Equal(t, len(arithmeticFunctionNode.Items), 1)
	assert.Equal(t, arithmeticFunctionNode.Items[0].(*domain.UnaryNode).DataType, constant.STRING)
	assert.Equal(t, arithmeticFunctionNode.Items[0].(*domain.UnaryNode).Value, "abc")
}

func TestArithmeticFunctionWithError(t *testing.T) {
	res, err := parser.Parse("min abc")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}
