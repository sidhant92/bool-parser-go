package errors

import "errors"

var INVALID_DATA_TYPE = errors.New("invalid data type")

var INVALID_CONTAINER_DATA_TYPE = errors.New("invalid container data type")

var INVALID_UNARY_OPERAND = errors.New("invalid unary operand")

var KEY_DATA_NOT_PRESENT = errors.New("key data not present")

var INVALID_EXPRESSION = errors.New("invalid expression")

var INCOMPATIBLE_DATA_TYPE = errors.New("incompatible data type")
