package service

import (
	op "github.com/sidhant92/bool-parser-go/internal/operator"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

type OperatorService struct {
	symbolMap map[string]constant.Operator
}

func (o *OperatorService) GetOperatorFromSymbol(symbol string) constant.Operator {
	if val, ok := o.symbolMap[symbol]; ok {
		return val
	}
	panic("Unknown Operator " + symbol)
}

func (o *OperatorService) Evaluate(operator constant.Operator, dataType constant.DataType, leftOperand interface{}, rightOperand ...interface{}) (bool, error) {
	res, err := op.GetOperator(operator).Evaluate(dataType, leftOperand, rightOperand...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func NewOperatorService() *OperatorService {
	allOperators := op.GetAllValues()
	symbolMap := map[string]constant.Operator{}
	for _, op := range allOperators {
		symbolMap[op.GetSymbol()] = op.GetOperator()
	}
	return &OperatorService{symbolMap}
}
