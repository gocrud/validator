package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

type EmailValidator struct {
	errorMessage
	Rex string
	ref reflect.Value
}

func (e *EmailValidator) Msg(msg string) *EmailValidator {
	e.setMsg(msg)
	return e
}

func (e *EmailValidator) SetValue(value reflect.Value) {
	e.ref = value
}

func (e *EmailValidator) Regex(rex string) *EmailValidator {
	e.Rex = rex
	return e
}

func (e *EmailValidator) Validate() error {
	if err := e.isString(); err != nil {
		return err
	}
	if e.Rex == "" {
		e.Rex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	}
	match, err := regexp.MatchString(e.Rex, e.ref.String())
	if err != nil {
		return Error{Message: err.Error()}
	}
	if !match {
		return e.errMsg("invalid email address")
	}
	return nil
}

func (e *EmailValidator) isString() error {
	if e.ref.Kind() == reflect.String {
		return nil
	}
	return e.errMsg(fmt.Sprintf("Email(): value must be a string"))
}

func Email() *EmailValidator {
	v := EmailValidator{}
	return &v
}
