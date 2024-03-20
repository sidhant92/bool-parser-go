package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"math"
)

type DivideOperator struct {
}

func (e *DivideOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	left := decimalDataType.GetValue(leftOperand).(float64)
	right := decimalDataType.GetValue(rightOperand).(float64)
	res := left / right
	if res == math.Trunc(res) {
		val := datatype.GetDataType(constant.LONG).GetValue(res).(int64)
		valInt := int(val)
		if val == int64(valInt) {
			return valInt, nil
		}
		return val, nil
	}
	return res, nil
}

func (e *DivideOperator) GetSymbol() string {
	return "/"
}

func (e *DivideOperator) GetOperator() constant.Operator {
	return constant.DIVIDE
}

func (e *DivideOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *DivideOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewDivideOperator() AbstractOperator {
	return &DivideOperator{}
}
