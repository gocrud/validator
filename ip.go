package validator

type IpValidator struct {
	AbstractValidator
}

func (i *IpValidator) Validate() error {
	if !isString(i.ref) {
		return i.Error("Ip: value must be a string")
	}
	if !isIP(i.ref.String()) {
		return i.Error("invalid ip address")
	}
	return nil
}

func (i *IpValidator) Msg(msg string) Validator {
	i.errMsg = msg
	return i
}

func IP() *IpValidator {
	return &IpValidator{}
}
