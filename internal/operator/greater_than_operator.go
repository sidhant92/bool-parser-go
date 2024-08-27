package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type GreaterThanOperator struct {
}

func (e *GreaterThanOperator) Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error) {
	err := Validate(leftOperand, leftOperandDataType, rightOperands[0].Value, rightOperands[0].DataType)
	if err != nil {
		return false, err
	}
	comparableDataType := GetComparableDataType(leftOperandDataType, rightOperands[0].DataType)
	dt := datatype.GetDataType(comparableDataType)
	res, err := dt.Compare(leftOperand, rightOperands[0].Value, true)
	if err != nil {
		return false, err
	}
	return res > 0, nil
}

func (e *GreaterThanOperator) GetSymbol() string {
	return ">"
}

func (e *GreaterThanOperator) GetOperator() constant.Operator {
	return constant.GREATER_THAN
}

func (e *GreaterThanOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *GreaterThanOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.VERSION}
}

func NewGreaterThanOperator() AbstractOperator {
	return &GreaterThanOperator{}
}
