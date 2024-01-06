package validator

import (
	"fmt"
	"reflect"
)

type LengthValidator struct {
	errorMessage
	min   int
	max   int
	value reflect.Value
}

func (l *LengthValidator) Validate() error {
	length, err := l.len()
	if err != nil {
		return err
	}
	if l.min != 0 && length < l.min {
		return l.errMsg(fmt.Sprintf("value length must be greater than %d", l.min))
	}
	if l.max != 0 && length > l.max {
		return l.errMsg(fmt.Sprintf("value length must be less than %d", l.max))
	}
	return nil
}

func (l *LengthValidator) SetValue(value reflect.Value) {
	l.value = value
}

func (l *LengthValidator) Msg(msg string) *LengthValidator {
	l.setMsg(msg)
	return l
}

func (l *LengthValidator) len() (int, error) {
	switch l.value.Kind() {
	case reflect.String:
		return len([]rune(l.value.String())), nil
	case reflect.Slice, reflect.Array, reflect.Map:
		return l.value.Len(), nil
	default:
		return 0, Error{Message: fmt.Sprintf("Length(%d,%d): value must be a string/slice/array/map", l.min, l.max)}
	}
}

func Length(min, max int) *LengthValidator {
	return &LengthValidator{
		min: min,
		max: max,
	}
}
