package service

import (
	"github.com/sidhant92/bool-parser-go/internal/containerdatatype"
	arithmetic "github.com/sidhant92/bool-parser-go/internal/function"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"log"
	"slices"
	"strings"
)

type FunctionEvaluatorService struct {
	symbolMap map[string]constant.FunctionType
}

func (f *FunctionEvaluatorService) GetFunctionFromSymbol(symbol string) constant.FunctionType {
	if val, ok := f.symbolMap[strings.ToLower(symbol)]; ok {
		return val
	}
	panic("Unknown Function " + symbol)
}

func (f *FunctionEvaluatorService) Evaluate(functionType constant.FunctionType, items []domain.Field) (interface{}, error) {
	var abstractFunction = arithmetic.GetArithmeticFunction(functionType)
	if len(items) < 1 {
		log.Printf("Empty items now allowed")
		return nil, errors2.INVALID_EXPRESSION
	}
	containerDataType := constant.LIST
	if len(items) == 1 {
		containerDataType = constant.PRIMITIVE
	}
	if !slices.Contains(abstractFunction.GetAllowedContainerTypes(), containerDataType) {
		log.Printf("Invalid container type %s for the the function %s", containerDataType, string(functionType))
		return false, errors2.INVALID_CONTAINER_DATA_TYPE
	}

	for _, item := range items {
		if !slices.Contains(abstractFunction.GetAllowedDataTypes(), item.DataType) {
			log.Printf("Invalid data type %s for the the function %s", item.DataType, functionType)
			return false, errors2.INVALID_DATA_TYPE
		}
		if !containerdatatype.GetContainerDataType(constant.PRIMITIVE).IsValid(item.DataType, item.Value) {
			log.Printf("validation failed for the function %s for the the operand %v", functionType, item.Value)
			return false, errors2.INVALID_DATA_TYPE
		}
	}
	res, err := abstractFunction.Evaluate(items)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewFunctionEvaluatorService() *FunctionEvaluatorService {
	allArithmeticFunctionTypes := arithmetic.GetAllArithmeticValues()
	symbolMap := map[string]constant.FunctionType{}
	for _, fn := range allArithmeticFunctionTypes {
		symbolMap[strings.ToLower(string(fn.GetFunctionType()))] = fn.GetFunctionType()
	}
	return &FunctionEvaluatorService{symbolMap}
}
