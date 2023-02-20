package util

import (
	"github.com/hashicorp/go-version"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"strconv"
	"strings"
)

func ConvertValue(value string, dataType constant.DataType) (interface{}, error) {
	switch dataType {
	case constant.INTEGER:
		res, err := strconv.Atoi(value)
		return res, err
	case constant.LONG:
		res, err := strconv.ParseInt(value, 10, 64)
		return res, err
	case constant.DECIMAL:
		res, err := strconv.ParseFloat(value, 64)
		return res, err
	case constant.BOOLEAN:
		res, err := strconv.ParseBool(value)
		return res, err
	case constant.VERSION:
		res, err := version.NewVersion(value)
		return res, err
	default:
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			return value[1 : len(value)-1], nil
		}
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			return value[1 : len(value)-1], nil
		}
		return value, nil
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
