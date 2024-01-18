package validator

import (
	"fmt"
	"reflect"
)

type LengthValidator struct {
	min int
	max int
	AbstractValidator
}

func (l *LengthValidator) Validate() error {
	length, err := l.len()
	if err != nil {
		return err
	}
	if l.min != 0 && length < l.min {
		return l.Error(fmt.Sprintf("value length must be greater than %d", l.min))
	}
	if l.max != 0 && length > l.max {
		return l.Error(fmt.Sprintf("value length must be less than %d", l.max))
	}
	return nil
}

func (l *LengthValidator) len() (int, error) {
	switch l.ref.Kind() {
	case reflect.String:
		return len([]rune(l.ref.String())), nil
	case reflect.Slice, reflect.Array, reflect.Map:
		return l.ref.Len(), nil
	default:
		return 0, l.Error(fmt.Sprintf("Length(%d,%d): value must be a string/slice/array/map", l.min, l.max))
	}
}

func Length(min, max int) *LengthValidator {
	return &LengthValidator{
		min: min,
		max: max,
	}
}
