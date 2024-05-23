package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
)

type UnaryOperator struct {
}

func (e *UnaryOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	if leftDataType == constant.DECIMAL {
		decimalDataType := datatype.GetDataType(constant.DECIMAL)
		res := decimalDataType.GetValue(leftOperand).(float64)
		return -res, nil
	}
	if leftDataType == constant.LONG {
		longDataType := datatype.GetDataType(constant.LONG)
		res := longDataType.GetValue(leftOperand).(int64)
		return -res, nil
	}
	if leftDataType == constant.INTEGER {
		intDataType := datatype.GetDataType(constant.INTEGER)
		res := intDataType.GetValue(leftOperand).(int)
		return -res, nil
	}
	return nil, errors.INVALID_DATA_TYPE
}

func (e *UnaryOperator) GetSymbol() string {
	return "UNARY"
}

func (e *UnaryOperator) GetOperator() constant.Operator {
	return constant.UNARY
}

func (e *UnaryOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *UnaryOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewUnaryOperator() AbstractOperator {
	return &UnaryOperator{}
}
