package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors "github.com/sidhant92/bool-parser-go/pkg/error"
)

type MultiplyOperator struct {
}

func (e *MultiplyOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	if leftDataType == constant.DECIMAL || rightDataType == constant.DECIMAL {
		decimalDataType := datatype.GetDataType(constant.DECIMAL)
		res := decimalDataType.GetValue(leftOperand).(float64) * decimalDataType.GetValue(rightOperand).(float64)
		return res, nil
	}
	if leftDataType == constant.LONG || rightDataType == constant.LONG {
		longDataType := datatype.GetDataType(constant.LONG)
		res := longDataType.GetValue(leftOperand).(int64) * longDataType.GetValue(rightOperand).(int64)
		return res, nil
	}
	if leftDataType == constant.INTEGER || rightDataType == constant.INTEGER {
		intDataType := datatype.GetDataType(constant.INTEGER)
		res := intDataType.GetValue(leftOperand).(int) * intDataType.GetValue(rightOperand).(int)
		return res, nil
	}
	return nil, errors.INVALID_DATA_TYPE
}

func (e *MultiplyOperator) GetSymbol() string {
	return "*"
}

func (e *MultiplyOperator) GetOperator() constant.Operator {
	return constant.MULTIPLY
}

func (e *MultiplyOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *MultiplyOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewMultiplyOperator() AbstractOperator {
	return &MultiplyOperator{}
}
