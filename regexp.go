package validator

import "regexp"

type RegexValidator struct {
	AbstractValidator
	rex string
}

func (r *RegexValidator) Validate() error {
	if !isString(r.ref) {
		return r.Error("Regex: value must be a string")
	}
	match, err := regexp.MatchString(r.rex, r.ref.String())
	if err != nil {
		return r.Error(err.Error())
	}
	if !match {
		return r.Error(r.errMsg)
	}
	return nil
}

func (r *RegexValidator) Msg(msg string) Validator {
	r.errMsg = msg
	return r
}

func Regex(rex string) *RegexValidator {
	v := RegexValidator{}
	v.rex = rex
	return &v
}
