package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type AddOperator struct {
}

func (e *AddOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	var stringDataTypes = []constant.DataType{constant.STRING, constant.BOOLEAN, constant.VERSION}
	if util.Contains(stringDataTypes, leftDataType) || util.Contains(stringDataTypes, rightDataType) {
		stringDataType := datatype.GetDataType(constant.STRING)
		res := stringDataType.GetValue(leftOperand).(string) + stringDataType.GetValue(rightOperand).(string)
		return res, nil
	}
	if leftDataType == constant.DECIMAL || rightDataType == constant.DECIMAL {
		decimalDataType := datatype.GetDataType(constant.DECIMAL)
		res := decimalDataType.GetValue(leftOperand).(float64) + decimalDataType.GetValue(rightOperand).(float64)
		return res, nil
	}
	if leftDataType == constant.LONG || rightDataType == constant.LONG {
		longDataType := datatype.GetDataType(constant.LONG)
		res := longDataType.GetValue(leftOperand).(int64) + longDataType.GetValue(rightOperand).(int64)
		return res, nil
	}
	if leftDataType == constant.INTEGER || rightDataType == constant.INTEGER {
		intDataType := datatype.GetDataType(constant.INTEGER)
		res := intDataType.GetValue(leftOperand).(int) + intDataType.GetValue(rightOperand).(int)
		return res, nil
	}
	stringDataType := datatype.GetDataType(constant.STRING)
	res := stringDataType.GetValue(leftOperand).(string) + stringDataType.GetValue(rightOperand).(string)
	return res, nil
}

func (e *AddOperator) GetSymbol() string {
	return "+"
}

func (e *AddOperator) GetOperator() constant.Operator {
	return constant.ADD
}

func (e *AddOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *AddOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION, constant.BOOLEAN}
}

func NewAddOperator() AbstractOperator {
	return &AddOperator{}
}
