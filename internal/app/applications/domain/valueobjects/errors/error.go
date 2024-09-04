package errors

import (
	"fmt"
	"reflect"
)

type EnumOutOfRangeError error

func NewEnumOutOfRangeError(actualValue interface{}) EnumOutOfRangeError {
	return NewEnumStrOutOfRangeError(reflect.TypeOf(actualValue).String(), actualValue)
}

func NewEnumStrOutOfRangeError(enumType string, actualValue interface{}) EnumOutOfRangeError {
	return fmt.Errorf("EnumOutOfRangeError: EnumType=%s, ActualValue=%s", enumType, actualValue)
}
