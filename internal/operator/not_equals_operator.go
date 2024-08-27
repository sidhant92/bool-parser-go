package operator

import (
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type NotEqualsOperator struct {
}

func (e *NotEqualsOperator) Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error) {
	eq := GetOperator(constant.EQUALS)
	res, err := eq.Evaluate(containerDataType, leftOperand, leftOperandDataType, rightOperands)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func (e *NotEqualsOperator) GetSymbol() string  {
	return "!="
}

func (e *NotEqualsOperator) GetOperator() constant.Operator  {
	return constant.NOT_EQUAL
}

func (e *NotEqualsOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *NotEqualsOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewNotEqualsOperator() AbstractOperator {
	return &NotEqualsOperator{}
}
