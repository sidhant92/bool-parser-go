package constant

type NodeType string

const (
	BOOLEAN_NODE       NodeType = "BOOLEAN_NODE"
	COMPARISON_NODE    NodeType = "COMPARISON_NODE"
	NUMERIC_RANGE_NODE NodeType = "NUMERIC_RANGE_NODE"
	IN_NODE            NodeType = "IN_NODE"
	ARRAY_NODE         NodeType = "ARRAY_NODE"
	UNARY_NODE         NodeType = "UNARY_NODE"
	ARITHMETIC         NodeType = "ARITHMETIC"
	ARITHMETIC_LEAF    NodeType = "ARITHMETIC_LEAF"
	ARITHMETIC_UNARY   NodeType = "ARITHMETIC_UNARY"
	STRING_NODE        NodeType = "STRING_NODE"
)
