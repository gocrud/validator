package validator

type IntRangeValidator struct {
	Min     int
	Max     int
	Message string
}

func (v IntRangeValidator) Validate(value interface{}) error {
	ref, ok := value.(int)
	if !ok {
		return &Error{Message: "invalid type, expected int"}
	}
	if ref < v.Min || ref > v.Max {
		var err = "value out of range"
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type IntRangeOptions func(validator *IntRangeValidator)

func IntRangeMsg(msg string) IntRangeOptions {
	return func(validator *IntRangeValidator) {
		validator.Message = msg
	}
}

func IntRange(min int, max int, opts ...IntRangeOptions) IntRangeValidator {
	v := IntRangeValidator{Min: min, Max: max}
	for _, f := range opts {
		f(&v)
	}
	return v
}
