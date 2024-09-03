package errors

import (
	"fmt"
	"reflect"
)

type EnumOutOfRangeError error

func NewEnumOutOfRangeError(actualValue interface{}) EnumOutOfRangeError {
	return fmt.Errorf("EnumOutOfRangeError: EnumType=%s, ActualValue=%s", reflect.TypeOf(actualValue).String(), actualValue)
}
