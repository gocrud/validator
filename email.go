package validator

import "regexp"

type Email struct {
	Rex     string
	Message string
}

func (v *Email) Validate(value interface{}) error {
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
