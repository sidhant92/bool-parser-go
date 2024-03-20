package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"math"
)

type ExponentOperator struct {
}

func (e *ExponentOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	left := decimalDataType.GetValue(leftOperand).(float64)
	right := decimalDataType.GetValue(rightOperand).(float64)
	res, _ := util.ConvertValue(math.Pow(left, right), constant.INTEGER)
	return res, nil
}

func (e *ExponentOperator) GetSymbol() string {
	return "^"
}

func (e *ExponentOperator) GetOperator() constant.Operator {
	return constant.EXPONENT
}

func (e *ExponentOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *ExponentOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG}
}

func NewExponentOperator() AbstractOperator {
	return &ExponentOperator{}
}
