package arithmetic

import (
	arithmetic2 "github.com/sidhant92/bool-parser-go/internal/function/arithmetic"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
)

var functionMap = map[constant.FunctionType]arithmetic2.AbstractFunction{
	constant.MIN: arithmetic2.NewMinFunction(),
	constant.MAX: arithmetic2.NewMaxFunction(),
	constant.AVG: arithmetic2.NewAvgFunction(),
	constant.SUM: arithmetic2.NewSumFunction(),
	constant.MEAN: arithmetic2.NewMeanFunction(),
	constant.MODE: arithmetic2.NewModeFunction(),
	constant.MEDIAN: arithmetic2.NewMedianFunction(),
	constant.INT: arithmetic2.NewIntFunction(),
	constant.LEN: arithmetic2.NewLenFunction(),
}

func GetArithmeticFunction(functionType constant.FunctionType) arithmetic2.AbstractFunction {
	return functionMap[functionType]
}

func GetAllArithmeticValues() []arithmetic2.AbstractFunction {
	var functionList []arithmetic2.AbstractFunction
	for _, v := range functionMap {
		functionList = append(functionList, v)
	}
	return functionList
}
