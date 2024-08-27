package constant

type DataType string

const (
	STRING  DataType = "STRING"
	INTEGER DataType = "INTEGER"
	LONG    DataType = "LONG"
	DECIMAL DataType = "DECIMAL"
	VERSION DataType = "VERSION"
	BOOLEAN DataType = "BOOL"
)

type DataTypeAttribute struct {
	Priority int
	Numeric bool
}

var DataTypeAttributes = map[DataType]DataTypeAttribute {
	STRING: {
		Priority: 6,
		Numeric:  false,
	},
	INTEGER: {
		Priority: 3,
		Numeric: true,
	},
	LONG: {
		Priority: 4,
		Numeric: true,
	},
	DECIMAL: {
		Priority: 5,
		Numeric: true,
	},
	VERSION: {
		Priority: 2,
		Numeric: true,
	},
	BOOLEAN: {
		Priority: 1,
		Numeric: false,
	},
}
