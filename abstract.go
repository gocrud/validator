package validator

import "reflect"

type AbstractValidator struct {
	ref    reflect.Value
	errMsg string
}

func (a *AbstractValidator) Validate() error {
	return &ValidationError{message: a.errMsg}
}

func (a *AbstractValidator) setValue(value reflect.Value) {
	a.ref = value
}

func (a *AbstractValidator) Msg(msg string) *AbstractValidator {
	a.errMsg = msg
	return a
}

func (a *AbstractValidator) Error(msg string) error {
	if a.errMsg == "" {
		a.errMsg = msg
	}
	return &ValidationError{message: a.errMsg}
}

var _ Validator = (*AbstractValidator)(nil)
