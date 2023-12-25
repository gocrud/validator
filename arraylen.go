package validator

import (
	"fmt"
	"reflect"
)

type ArrayLenthValidator struct {
	Min     int
	Max     int
	Message string
}

type ArrayLenOptions func(validator *ArrayLenthValidator)

func ArrayLenMsg(msg string) ArrayLenOptions {
	return func(validator *ArrayLenthValidator) {
		validator.Message = msg
	}
}

func ArrayLen(min, max int, opts ...ArrayLenOptions) ArrayLenthValidator {
	v := ArrayLenthValidator{
		Min: min,
		Max: max,
	}
	for _, f := range opts {
		f(&v)
	}
	return v
}

func (v ArrayLenthValidator) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Array {
		return Error{Message: "value must be a array"}
	}
	if ref.Len() < v.Min || ref.Len() > v.Max {
		var err = fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}
