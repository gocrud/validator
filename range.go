package validator

import (
	"fmt"
	"reflect"
)

type RangeValidator struct {
	AbstractValidator
	min int64
	max int64
}

func (r *RangeValidator) Msg(msg string) Validator {
	r.errMsg = msg
	return r
}

func (r *RangeValidator) Validate() error {
	var value int64
	switch r.ref.Kind() {
	case reflect.String:
		value = int64(len([]rune(r.ref.String())))
	case reflect.Slice, reflect.Array, reflect.Map:
		value = int64(r.ref.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = int64(int(r.ref.Int()))
	default:
		return r.Error(fmt.Sprintf("Range(%d,%d): value must be a string/slice/array/map/int/int8/int16/int32/int64", r.min, r.max))
	}
	if r.min > r.max {
		panic("Range: min must be less than max")
	}
	if value >= r.min && value <= r.max {
		return nil
	}
	return r.Error(fmt.Sprintf("value must be between %d and %d", r.min, r.max))
}

func Range(min, max int64) *RangeValidator {
	return &RangeValidator{
		min: min,
		max: max,
	}
}
