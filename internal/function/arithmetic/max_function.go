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

type MaxFunction struct {
}

func (m *MaxFunction) Evaluate(items []domain.EvaluatedNode) (interface{}, error) {
	dataTypes := lo.Map(items, func(item domain.EvaluatedNode, index int) constant.DataType {
		return item.DataType
	})
	if util.Contains(dataTypes, constant.DECIMAL) {
		decimalDataType := datatype.GetDataType(constant.DECIMAL)
		decimalValues := lo.Map(items, func(item domain.EvaluatedNode, index int) float64 {
			return decimalDataType.GetValue(item.Value).(float64)
		})
		return datatypeutil.CastToWholeIfPossible(slices.Max(decimalValues)), nil
	}
	if util.Contains(dataTypes, constant.LONG) {
		longDataType := datatype.GetDataType(constant.LONG)
		longValues := lo.Map(items, func(item domain.EvaluatedNode, index int) int64 {
			return longDataType.GetValue(item.Value).(int64)
		})
		return datatypeutil.CastToIntIfPossible(slices.Max(longValues)), nil
	}
	intDataType := datatype.GetDataType(constant.INTEGER)
	intValues := lo.Map(items, func(item domain.EvaluatedNode, index int) int {
		return intDataType.GetValue(item.Value).(int)
	})
	return slices.Max(intValues), nil
}

func (m *MaxFunction) GetFunctionType() constant.FunctionType {
	return constant.MAX
}

func (m *MaxFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *MaxFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewMaxFunction() AbstractFunction {
	return &MaxFunction{}
}
