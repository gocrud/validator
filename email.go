package validator

import (
	"regexp"
)

type EmailValidator struct {
	AbstractValidator
	Rex string
}

func (e *EmailValidator) Regex(rex string) *EmailValidator {
	e.Rex = rex
	return e
}

func (e *EmailValidator) Msg(msg string) Validator {
	e.errMsg = msg
	return e
}

func (e *EmailValidator) Validate() error {
	if !isString(e.ref) {
		return e.Error("Email: value must be a string")
	}
	if e.Rex == "" {
		e.Rex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	}
	match, err := regexp.MatchString(e.Rex, e.ref.String())
	if err != nil {
		return e.Error(err.Error())
	}
	if !match {
		return e.Error("invalid email address")
	}
	return nil
}

func Email() *EmailValidator {
	v := EmailValidator{}
	return &v
}
