package validator

import (
	"reflect"
)

type Validator interface {
	Validate() error
	Msg(msg string) Validator
	setValue(value reflect.Value)
}

func Validate(value interface{}, validators ...Validator) error {
	ref := getPointerValue(value)
	for _, v := range validators {
		v.setValue(ref)
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
