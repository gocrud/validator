package validator

import (
	"fmt"
	"reflect"
)

type MapLenValidator struct {
	Min int
	Max int
}

type MapLenOptions func(validator *MapLenValidator)

func MapLen(min, max int, opts ...MapLenOptions) MapLenValidator {
	v := MapLenValidator{
		Min: min,
		Max: max,
	}
	for _, f := range opts {
		f(&v)
	}
	return v
}

func (v *MapLenValidator) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Map {
		return Error{Message: "value must be a map"}
	}
	if ref.Len() < v.Min || ref.Len() > v.Max {
		return &Error{Message: fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)}
	}
	return nil
}
