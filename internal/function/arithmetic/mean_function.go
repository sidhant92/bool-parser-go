package arithmetic

import (
	"github.com/montanaflynn/stats"
	"github.com/samber/lo"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	datatypeutil "github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type MeanFunction struct {
}

func (m *MeanFunction) Evaluate(items []domain.EvaluatedNode) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	decimalValues := lo.Map(items, func(item domain.EvaluatedNode, index int) float64 {
		return decimalDataType.GetValue(item.Value).(float64)
	})
	res, _ := stats.Mean(decimalValues)
	return datatypeutil.CastToWholeIfPossible(res), nil
}

func (m *MeanFunction) GetFunctionType() constant.FunctionType {
	return constant.MEAN
}

func (m *MeanFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *MeanFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewMeanFunction() AbstractFunction {
	return &MeanFunction{}
}
