package util

import (
	"fmt"
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
		newKey := strings.Join(keys[1:], ".")
		return GetValueFromMap(newKey, nestedMap)
	}
	log.Printf("could not find key %s for the data %s", key, data)
	return nil
}

func ToString(input interface{}) string {
	switch input.(type) {
	case float64, float32:
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", input), "0"), ".")
	default:
		return fmt.Sprintf("%v", input)
	}
}

func ConvertValue(value interface{}, dataType constant.DataType) (interface{}, error) {
	valueString := ToString(value)
	switch dataType {
	case constant.INTEGER:
		res, err := strconv.Atoi(valueString)
		return res, err
	case constant.LONG:
		res, err := strconv.ParseInt(valueString, 10, 64)
		return res, err
	case constant.DECIMAL:
		res, err := strconv.ParseFloat(valueString, 64)
		return res, err
	case constant.BOOLEAN:
		res, err := strconv.ParseBool(valueString)
		return res, err
	case constant.VERSION:
		res, err := version.NewVersion(valueString)
		return res, err
	default:
		if strings.HasPrefix(valueString, "'") && strings.HasSuffix(valueString, "'") {
			return valueString[1 : len(valueString)-1], nil
		}
		if strings.HasPrefix(valueString, "\"") && strings.HasSuffix(valueString, "\"") {
			return valueString[1 : len(valueString)-1], nil
		}
		return value, nil
	}
}

func GetDataType(value interface{}) constant.DataType {
	switch value.(type) {
	case float32:
	case float64:
		return constant.DECIMAL
	case int64:
		return constant.LONG
	case int32:
	case int:
		return constant.INTEGER
	case bool:
		return constant.BOOLEAN
	case *version.Version:
		return constant.VERSION
	default:
		return constant.STRING
	}
	return constant.STRING
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
