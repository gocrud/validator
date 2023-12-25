package validator

import (
	"fmt"
	"reflect"
)

type SliceLenValidator struct {
	Min     int
	Max     int
	Message string
}

type SliceLenOptions func(validator *SliceLenValidator)

func SliceLenMsg(msg string) SliceLenOptions {
	return func(validator *SliceLenValidator) {
		validator.Message = msg
	}
}

func SliceLen(min, max int, opts ...SliceLenOptions) SliceLenValidator {
	v := SliceLenValidator{
		Min: min,
		Max: max,
	}
	for _, f := range opts {
		f(&v)
	}
	return v
}

func (v SliceLenValidator) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Slice {
		return Error{Message: "value must be a slice"}
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
