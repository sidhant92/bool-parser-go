package arithmetic

import (
	"github.com/samber/lo"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/internal/util"
	"github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
	"slices"
)

type MinFunction struct {
}

func (m *MinFunction) Evaluate(items []domain.EvaluatedNode) (interface{}, error) {
	dataTypes := lo.Map(items, func(item domain.EvaluatedNode, index int) constant.DataType {
		return item.DataType
	})
	if util.Contains(dataTypes, constant.DECIMAL) {
		decimalDataType := datatype.GetDataType(constant.DECIMAL)
		decimalValues := lo.Map(items, func(item domain.EvaluatedNode, index int) float64 {
			return decimalDataType.GetValue(item.Value).(float64)
		})
		return datatypeutil.CastToWholeIfPossible(slices.Min(decimalValues)), nil
	}
	if util.Contains(dataTypes, constant.LONG) {
		longDataType := datatype.GetDataType(constant.LONG)
		longValues := lo.Map(items, func(item domain.EvaluatedNode, index int) int64 {
			return longDataType.GetValue(item.Value).(int64)
		})
		return datatypeutil.CastToIntIfPossible(slices.Min(longValues)), nil
	}
	intDataType := datatype.GetDataType(constant.INTEGER)
	intValues := lo.Map(items, func(item domain.EvaluatedNode, index int) int {
		return intDataType.GetValue(item.Value).(int)
	})
	return slices.Min(intValues), nil
}

func (m *MinFunction) GetFunctionType() constant.FunctionType {
	return constant.MIN
}

func (m *MinFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *MinFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewMinFunction() AbstractFunction {
	return &MinFunction{}
}
