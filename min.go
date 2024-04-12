package validator

import (
	"fmt"
	"reflect"
)

type MinValidator struct {
	AbstractValidator
	min int64
}

func (m *MinValidator) Validate() error {
	var value int64
	switch m.ref.Kind() {
	case reflect.String:
		value = int64(len([]rune(m.ref.String())))
	case reflect.Slice, reflect.Array, reflect.Map:
		value = int64(m.ref.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = int64(int(m.ref.Int()))
	default:
		return m.Error(fmt.Sprintf("Min(%d): value must be a string/slice/array/map/int/int8/int16/int32/int64", m.min))
	}

	if value >= m.min {
		return nil
	}
	return m.Error(fmt.Sprintf("value must be greater than or equal to %d", m.min))
}

func (m *MinValidator) Msg(msg string) Validator {
	m.errMsg = msg
	return m
}

func Min(min int64) *MinValidator {
	v := MinValidator{}
	v.min = min
	return &v
}
