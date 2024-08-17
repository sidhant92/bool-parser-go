package datatypeutil

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"math"
)

func CastToWholeIfPossible(res float64) interface{} {
	if res == math.Trunc(res) {
		val := datatype.GetDataType(constant.LONG).GetValue(res).(int64)
		valInt := int(val)
		if val == int64(valInt) {
			return valInt
		}
		return val
	}
	return res
}

func CastToIntIfPossible(val int64) interface{} {
	valInt := int(val)
	if val == int64(valInt) {
		return valInt
	}
	return val
}
