package validator

import "fmt"

type StringLengthValidator struct {
	Min     int
	Max     int
	Message string
}

func (v StringLengthValidator) Validate(value interface{}) error {
	ref, err := isString(value)
	if err != nil {
		return err
	}
	r := []rune(ref.String())
	if len(r) < v.Min || len(r) > v.Max {
		var err = fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type StrLenOptions func(validator *StringLengthValidator)

func StrLenMsg(msg string) StrLenOptions {
	return func(validator *StringLengthValidator) {
		validator.Message = msg
	}
}

func StrLen(min, max int, opts ...StrLenOptions) StringLengthValidator {
	v := StringLengthValidator{
		Min: min,
		Max: max,
	}
	for _, f := range opts {
		f(&v)
	}
	return v
}
