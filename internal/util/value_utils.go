package util

import (
	"github.com/hashicorp/go-version"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	"log"
	"strconv"
	"strings"
)

func GetValueFromMap(key string, data map[string]interface{}) interface{} {
	keys := strings.Split(key, ".")
	size := len(keys)
	value, ok := data[keys[0]]
	if size == 1 {
		if ok {
			return value
		} else {
			return nil
		}
	}
	nestedMap, ok := value.(map[string]interface{})
	if ok {
		return GetValueFromMap(keys[1], nestedMap)
	}
	log.Printf("could not find key %s for the data %s", key, data)
	return nil
}

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
