package service

import (
	"github.com/sidhant92/bool-parser-go/internal/containerdatatype"
	op "github.com/sidhant92/bool-parser-go/internal/operator"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"log"
	"slices"
	"strings"
)

type OperatorService struct {
	symbolMap map[string]constant.Operator
}

func (o *OperatorService) GetOperatorFromSymbol(symbol string) constant.Operator {
	if val, ok := o.symbolMap[strings.ToLower(symbol)]; ok {
		return val
	}
	panic("Unknown Operator " + symbol)
}

func (o *OperatorService) Evaluate(operator constant.Operator, containerDataType constant.ContainerDataType, dataType constant.DataType, leftOperand interface{}, rightOperand ...interface{}) (bool, error) {
	var abstractOperator = op.GetOperator(operator)
	if !slices.Contains(abstractOperator.GetAllowedContainerTypes(), containerDataType) {
		log.Printf("Invalid container type %s for the the operator %s", containerDataType, operator)
		return false, errors2.INVALID_CONTAINER_DATA_TYPE
	}
	if !slices.Contains(abstractOperator.GetAllowedDataTypes(), dataType) {
		log.Printf("Invalid data type %s for the the operator %s", dataType, operator)
		return false, errors2.INVALID_DATA_TYPE
	}
	if !containerdatatype.GetContainerDataType(containerDataType).IsValid(dataType, leftOperand) {
		log.Printf("validation failed for the operator %s for the the operand %v", operator, leftOperand)
		return false, errors2.INVALID_DATA_TYPE
	}
	res, err := abstractOperator.Evaluate(containerDataType, dataType, true, leftOperand, rightOperand...)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (o *OperatorService) EvaluateArithmeticExpression(leftOperand interface{}, leftOperandDataType constant.DataType, rightOperand interface{}, rightOperandDataType constant.DataType, operator constant.Operator, containerDataType constant.ContainerDataType) (interface{}, error) {
	var abstractOperator = op.GetArithmeticOperator(operator)
	if !slices.Contains(abstractOperator.GetAllowedContainerTypes(), containerDataType) {
		log.Printf("Invalid container type %s for the the operator %s", containerDataType, operator)
		return false, errors2.INVALID_CONTAINER_DATA_TYPE
	}
	if !slices.Contains(abstractOperator.GetAllowedDataTypes(), leftOperandDataType) {
		log.Printf("Invalid left operand data type %s for the the operator %s", leftOperandDataType, operator)
		return false, errors2.INVALID_CONTAINER_DATA_TYPE
	}
	if rightOperand != nil && !slices.Contains(abstractOperator.GetAllowedDataTypes(), rightOperandDataType) {
		log.Printf("Invalid right operand data type %s for the the operator %s", leftOperandDataType, operator)
		return false, errors2.INVALID_CONTAINER_DATA_TYPE
	}
	if !containerdatatype.GetContainerDataType(containerDataType).IsValid(leftOperandDataType, leftOperand) {
		log.Printf("validation failed for the operator %s for the the operand %v", operator, leftOperand)
		return false, errors2.INVALID_DATA_TYPE
	}
	if rightOperand != nil && !containerdatatype.GetContainerDataType(containerDataType).IsValid(rightOperandDataType, rightOperand) {
		log.Printf("validation failed for the operator %s for the the operand %v", operator, rightOperand)
		return false, errors2.INVALID_DATA_TYPE
	}
	res, err := abstractOperator.Evaluate(leftOperand, leftOperandDataType, rightOperand, rightOperandDataType)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewOperatorService() *OperatorService {
	allOperators := op.GetAllValues()
	allArithmeticOperators := op.GetAllArithmeticValues()
	symbolMap := map[string]constant.Operator{}
	for _, op := range allOperators {
		symbolMap[strings.ToLower(op.GetSymbol())] = op.GetOperator()
	}
	for _, op := range allArithmeticOperators {
		symbolMap[strings.ToLower(op.GetSymbol())] = op.GetOperator()
	}
	return &OperatorService{symbolMap}
}
