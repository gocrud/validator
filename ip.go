package validator

type IpValidator struct {
	Message string
}

func (v IpValidator) Validate(value interface{}) error {
	ref, err := isString(value)
	if err != nil {
		return err
	}
	if !isIP(ref.String()) {
		var err = "invalid ip"
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type IpOptions func(validator *IpValidator)

func IpMsg(msg string) IpOptions {
	return func(validator *IpValidator) {
		validator.Message = msg
	}
}

func IP(opts ...IpOptions) IpValidator {
	v := IpValidator{}
	for _, f := range opts {
		f(&v)
	}
	return v
}
