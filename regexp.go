package validator

import "regexp"

type RegexpValidator struct {
	AbstractValidator
	rex string
}

func (r *RegexpValidator) Validate() error {
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

func (r *RegexpValidator) Msg(msg string) Validator {
	r.errMsg = msg
	return r
}

func Regexp(rex string) *RegexpValidator {
	v := RegexpValidator{}
	v.rex = rex
	return &v
}
