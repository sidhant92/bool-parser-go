package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type SumFunction struct {
}

func (m *SumFunction) Evaluate(items []domain.Field) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	total := 0.0
	for _, item := range items {
		total += decimalDataType.GetValue(item.Value).(float64)
	}
	return datatypeutil.CastToWholeIfPossible(total), nil
}

func (m *SumFunction) GetFunctionType() constant.FunctionType {
	return constant.SUM
}

func (m *SumFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *SumFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewSumFunction() AbstractFunction {
	return &SumFunction{}
}
