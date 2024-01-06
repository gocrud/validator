package validator

import "reflect"

type IpValidator struct {
	errorMessage
	Message string
	ref     reflect.Value
}

func (i *IpValidator) Validate() error {
	ref, err := i.isString(i.value)
	if err != nil {
		return err
	}
	if !isIP(ref.String()) {
		return i.errMsg("invalid ip address")
	}
	return nil
}

func (i *IpValidator) SetValue(value reflect.Value) {
	i.ref = value
}

func (i *IpValidator) isString(value string) (reflect.Value, error) {
	if i.ref.Kind() == reflect.String {
		return i.ref, nil
	}
	return i.ref, i.errMsg("Ip(): value must be a string")
}

func (i *IpValidator) Msg(msg string) *IpValidator {
	i.Message = msg
	return i
}

func IP() *IpValidator {
	return &IpValidator{}
}
