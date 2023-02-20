package constant

type NodeType string

const (
	BOOLEAN_NODE    NodeType = "BOOLEAN_NODE"
	COMPARISON_NODE    NodeType = "COMPARISON_NODE"
	NUMERIC_RANGE_NODE NodeType = "NUMERIC_RANGE_NODE"
	IN_NODE            NodeType = "IN_NODE"
	UNARY_NODE NodeType = "UNARY_NODE"
)
