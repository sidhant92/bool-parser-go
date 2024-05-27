package arithmetic

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	datatypeutil "github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type AvgFunction struct {
}

func (m *AvgFunction) Evaluate(items []domain.Field) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	total := 0.0
	for _, item := range items {
		total += decimalDataType.GetValue(item.Value).(float64)
	}
	res := total / float64(len(items))
	return datatypeutil.CastToWholeIfPossible(res), nil
}

func (m *AvgFunction) GetFunctionType() constant.FunctionType {
	return constant.AVG
}

func (m *AvgFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *AvgFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewAvgFunction() AbstractFunction {
	return &AvgFunction{}
}
