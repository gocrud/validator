package validator

import "reflect"

type OptionalValidator struct {
	validators []Validator
	value      reflect.Value
}

func (o *OptionalValidator) Validate() error {
	if isEmpty(o.value) {
		return nil
	}
	for _, v := range o.validators {
		v.SetValue(o.value)
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (o *OptionalValidator) SetValue(value reflect.Value) {
	o.value = value
}

func Optional(validators ...Validator) *OptionalValidator {
	return &OptionalValidator{
		validators: validators,
	}
}
