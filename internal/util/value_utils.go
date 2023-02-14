package util

import (
	"github.com/hashicorp/go-version"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"strconv"
	"strings"
)

func ConvertValue(value string, dataType constant.DataType) interface{} {
	switch dataType {
	case constant.INTEGER:
		res, _ := strconv.Atoi(value)
		return res
	case constant.LONG:
		res, _ := strconv.ParseInt(value, 10, 64)
		return res
	case constant.DECIMAL:
		res, _ := strconv.ParseFloat(value, 64)
		return res
	case constant.BOOLEAN:
		res, _ := strconv.ParseBool(value)
		return res
	case constant.APP_VERSION:
		res, _ := version.NewVersion(value)
		return res
	default:
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			return value[1 : len(value)-1]
		}
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			return value[1 : len(value)-1]
		}
		return value
	}
}

func IsValidInteger(value string) bool {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return true
}

func GetNumericDataType(value string) constant.DataType {
	if IsValidInteger(value) {
		return constant.INTEGER
	}
	return constant.LONG
}
