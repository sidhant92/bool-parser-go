package arithmetic

import (
	"github.com/montanaflynn/stats"
	"github.com/samber/lo"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	datatypeutil "github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type MedianFunction struct {
}

func (m *MedianFunction) Evaluate(items []domain.Field) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	decimalValues := lo.Map(items, func(item domain.Field, index int) float64 {
		return decimalDataType.GetValue(item.Value).(float64)
	})
	res, _ := stats.Median(decimalValues)
	return datatypeutil.CastToWholeIfPossible(res), nil
}

func (m *MedianFunction) GetFunctionType() constant.FunctionType {
	return constant.MEDIAN
}

func (m *MedianFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *MedianFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewMedianFunction() AbstractFunction {
	return &MedianFunction{}
}
