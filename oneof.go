package validator

type OneOfValidator struct {
	Values  []interface{}
	Message string
}

func (v OneOfValidator) Validate(value interface{}) error {
	for _, val := range v.Values {
		if val == value {
			return nil
		}
	}
	var err = "invalid value"
	if v.Message != "" {
		err = v.Message
	}
	return &Error{Message: err}
}

type OneOfOptions func(validator *OneOfValidator)

func OneOfMsg(msg string) OneOfOptions {
	return func(validator *OneOfValidator) {
		validator.Message = msg
	}
}

func OneOf(values []any, opts ...OneOfOptions) OneOfValidator {
	v := OneOfValidator{
		Values: values,
	}
	for _, f := range opts {
		f(&v)
	}
	return v
}
