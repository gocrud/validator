package validator

import "regexp"

type EmailValidator struct {
	Rex     string
	Message string
}

func (v EmailValidator) Validate(value interface{}) error {
	ref, err := isString(value)
	if err != nil {
		return err
	}
	if v.Rex == "" {
		v.Rex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	}
	match, err := regexp.MatchString(v.Rex, ref.String())
	if err != nil {
		return Error{Message: err.Error()}
	}
	if !match {
		var err = "invalid email"
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type EmailOptions func(validator *EmailValidator)

func EmailMsg(msg string) EmailOptions {
	return func(validator *EmailValidator) {
		validator.Message = msg
	}
}

func EmailRex(rex string) EmailOptions {
	return func(validator *EmailValidator) {
		validator.Rex = rex
	}
}

func Email(opts ...EmailOptions) EmailValidator {
	v := EmailValidator{}
	for _, f := range opts {
		f(&v)
	}
	return v
}
