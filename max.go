package validator

import (
	"fmt"
	"reflect"
)

type MaxValidator struct {
	max int64
	AbstractValidator
}

func (m *MaxValidator) Validate() error {
	var value int64
	switch m.ref.Kind() {
	case reflect.String:
		value = int64(len([]rune(m.ref.String())))
	case reflect.Slice, reflect.Array, reflect.Map:
		value = int64(m.ref.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = int64(int(m.ref.Int()))
	default:
		return m.Error(fmt.Sprintf("Max(%d): value must be a string/slice/array/map/int/int8/int16/int32/int64", m.max))
	}
	if value <= m.max {
		return nil
	}
	return m.Error(fmt.Sprintf("value must be less than or equal to %d", m.max))
}

func (m *MaxValidator) Msg(msg string) Validator {
	m.errMsg = msg
	return m
}

func Max(max int64) *MaxValidator {
	return &MaxValidator{
		max: max,
	}
}
