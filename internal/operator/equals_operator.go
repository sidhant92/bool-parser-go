package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type EqualsOperator struct {
}

func (e *EqualsOperator) Evaluate(containerDataType constant.ContainerDataType, leftOperand interface{}, leftOperandDataType constant.DataType, rightOperands []domain.EvaluatedNode) (bool, error) {
	comparableDataType := getComparableDataType(leftOperandDataType, rightOperands[0].DataType)
	dt := datatype.GetDataType(comparableDataType)
	res, err := dt.Compare(leftOperand, rightOperands[0].Value, true)
	if err != nil {
		return false, err
	}
	return res == 0, nil
}

func getComparableDataType(leftOperandDataType constant.DataType, rightOperandDataType constant.DataType) constant.DataType  {
	if constant.DataTypeAttributes[leftOperandDataType].Priority > constant.DataTypeAttributes[rightOperandDataType].Priority {
		return leftOperandDataType
	}
	return rightOperandDataType
}

func (e *EqualsOperator) GetSymbol() string {
	return "="
}

func (e *EqualsOperator) GetOperator() constant.Operator {
	return constant.EQUALS
}

func (e *EqualsOperator) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.PRIMITIVE}
}

func (e *EqualsOperator) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.STRING, constant.INTEGER, constant.LONG, constant.DECIMAL, constant.BOOLEAN, constant.VERSION}
}

func NewEqualsOperator() AbstractOperator {
	return &EqualsOperator{}
}
