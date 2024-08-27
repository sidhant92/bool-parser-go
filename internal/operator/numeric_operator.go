package operator

import (
	"github.com/sidhant92/bool-parser-go/internal/datatype"
	"github.com/sidhant92/bool-parser-go/pkg/constant"
	errors2 "github.com/sidhant92/bool-parser-go/pkg/error"
	"log"
)

func GetComparableDataType(leftOperandDataType constant.DataType, rightOperandDataType constant.DataType) constant.DataType {
	if (constant.DataTypeAttributes[leftOperandDataType].Numeric && !constant.DataTypeAttributes[rightOperandDataType].Numeric) || (!constant.DataTypeAttributes[leftOperandDataType].Numeric && constant.DataTypeAttributes[rightOperandDataType].Numeric) {
		if constant.DataTypeAttributes[leftOperandDataType].Numeric {
			return leftOperandDataType
		}
		return rightOperandDataType
	}
	if constant.DataTypeAttributes[leftOperandDataType].Priority > constant.DataTypeAttributes[rightOperandDataType].Priority {
		return leftOperandDataType
	}
	return rightOperandDataType
}

func Validate(leftOperand interface{}, leftOperandDataType constant.DataType, rightOperand interface{}, rightOperandDataType constant.DataType) error {
	if constant.DataTypeAttributes[leftOperandDataType].Numeric && constant.DataTypeAttributes[rightOperandDataType].Numeric {
		return nil
	}
	if !constant.DataTypeAttributes[leftOperandDataType].Numeric && !constant.DataTypeAttributes[rightOperandDataType].Numeric {
		return nil
	}
	rightDataType := datatype.GetDataType(rightOperandDataType)
	if !constant.DataTypeAttributes[leftOperandDataType].Numeric && !rightDataType.IsValid(leftOperand) {
		log.Printf("Incompatible data types %s and %s", leftOperandDataType, rightOperandDataType)
		return errors2.INCOMPATIBLE_DATA_TYPE
	}
	leftDataType := datatype.GetDataType(leftOperandDataType)
	if !constant.DataTypeAttributes[rightOperandDataType].Numeric && !leftDataType.IsValid(rightOperand) {
		log.Printf("Incompatible data types %s and %s", leftOperandDataType, rightOperandDataType)
		return errors2.INCOMPATIBLE_DATA_TYPE
	}
	return nil
}
