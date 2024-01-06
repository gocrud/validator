package validator

import (
	"fmt"
	"reflect"
)

type IntRangeValidator struct {
	message errorMessage
	min     int64
	max     int64
	ref     reflect.Value
}

func (i *IntRangeValidator) Validate() error {
	if err := i.isInt(); err != nil {
		return err
	}
	value := i.ref.Int()
	if i.min != 0 && value < i.min {
		return i.message.errMsg(fmt.Sprintf("value must be greater than %d", i.min))
	}
	if i.max != 0 && value > i.max {
		return i.message.errMsg(fmt.Sprintf("value must be less than %d", i.max))
	}
	return nil
}

func (i *IntRangeValidator) isInt() error {
	if i.ref.Kind() >= reflect.Int && i.ref.Kind() <= reflect.Int64 {
		return nil
	}
	return i.message.errMsg(fmt.Sprintf("IntRange(%d,%d): value must be an int", i.min, i.max))
}

func (i *IntRangeValidator) SetValue(value reflect.Value) {
	i.ref = value
}

func (i *IntRangeValidator) Msg(msg string) *IntRangeValidator {
	i.message.setMsg(msg)
	return i
}

func IntRange(min, max int64) *IntRangeValidator {
	return &IntRangeValidator{
		min: min,
		max: max,
	}
}
