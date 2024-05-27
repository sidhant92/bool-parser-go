package arithmetic

import (
	"github.com/montanaflynn/stats"
	"github.com/samber/lo"
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	datatypeutil "github.com/sidhant92/bool-parser-go/internal/util/data_type"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"github.com/sidhant92/bool-parser-go/pkg/domain"
)

type ModeFunction struct {
}

func (m *ModeFunction) Evaluate(items []domain.Field) (interface{}, error) {
	decimalDataType := datatype.GetDataType(constant.DECIMAL)
	decimalValues := lo.Map(items, func(item domain.Field, index int) float64 {
		return decimalDataType.GetValue(item.Value).(float64)
	})
	res, _ := stats.Mode(decimalValues)
	if len(res) == 0 {
		return datatypeutil.CastToWholeIfPossible(decimalValues[0]), nil
	}
	return datatypeutil.CastToWholeIfPossible(res[0]), nil
}

func (m *ModeFunction) GetFunctionType() constant.FunctionType {
	return constant.MODE
}

func (m *ModeFunction) GetAllowedContainerTypes() []constant.ContainerDataType {
	return []constant.ContainerDataType{constant.LIST}
}

func (m *ModeFunction) GetAllowedDataTypes() []constant.DataType {
	return []constant.DataType{constant.INTEGER, constant.LONG, constant.DECIMAL}
}

func NewModeFunction() AbstractFunction {
	return &ModeFunction{}
}
