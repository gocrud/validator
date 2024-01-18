package validator

import (
	"fmt"
	"reflect"
)

type IntRangeValidator struct {
	AbstractValidator
	min int64
	max int64
}

func (i *IntRangeValidator) Validate() error {
	if !isInt(i.ref) {
		return i.Error("IntRange: value must be an int")
	}
	value := i.ref.Int()
	if i.min != 0 && value < i.min {
		return i.Error(fmt.Sprintf("value must be greater than %d", i.min))
	}
	if i.max != 0 && value > i.max {
		return i.Error(fmt.Sprintf("value must be less than %d", i.max))
	}
	return nil
}

func (i *IntRangeValidator) SetValue(value reflect.Value) {
	i.ref = value
}

func IntRange(min, max int64) *IntRangeValidator {
	return &IntRangeValidator{
		min: min,
		max: max,
	}
}
