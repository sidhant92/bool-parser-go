package constant

type Operator string

const (
	EQUALS             Operator = "EQUALS"
	GREATER_THAN       Operator = "GREATER_THAN"
	GREATER_THAN_EQUAL Operator = "GREATER_THAN_EQUAL"
	LESS_THAN          Operator = "LESS_THAN"
	LESS_THAN_EQUAL    Operator = "LESS_THAN_EQUAL"
	NOT_EQUAL          Operator = "NOT_EQUAL"
	IN                 Operator = "IN"

	ADD      Operator = "ADD"
	SUBTRACT Operator = "SUBTRACT"
	MULTIPLY Operator = "MULTIPLY"
	DIVIDE   Operator = "DIVIDE"
	MODULUS  Operator = "MODULUS"
	EXPONENT Operator = "EXPONENT"
	UNARY    Operator = "UNARY"

	CONTAINS_ANY Operator = "CONTAINS_ANY"
	CONTAINS_ALL Operator = "CONTAINS_ALL"
)

var EqualityOperators = []Operator {EQUALS, NOT_EQUAL}
