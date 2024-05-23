package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type ModulusOperator struct {
}

func (e *ModulusOperator) Evaluate(leftOperand interface{}, leftDataType constant.DataType, rightOperand interface{}, rightDataType constant.DataType) (interface{}, error) {
	if leftDataType == constant.LONG || rightDataType == constant.LONG {
		longDataType := datatype.GetDataType(constant.LONG)
		res := longDataType.GetValue(leftOperand).(int64) % longDataType.GetValue(rightOperand).(int64)
		return res, nil
	}
	intDataType := datatype.GetDataType(constant.INTEGER)
	res := intDataType.GetValue(leftOperand).(int) % intDataType.GetValue(rightOperand).(int)
	return res, nil
}

func (e *ModulusOperator) GetSymbol() string {
	return "%"
}

func (e *ModulusOperator) GetOperator() constant.Operator {
	return constant.MODULUS
}

func (e *ModulusOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *ModulusOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG}
}

func NewModulusOperator() AbstractOperator {
	return &ModulusOperator{}
}
