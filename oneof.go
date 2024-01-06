package validator

import "reflect"

type OneOfValidator struct {
	errorMessage
	values []interface{}
	ref    reflect.Value
}

func (o *OneOfValidator) Validate() error {
	if !o.isInSlice() {
		return o.errMsg("value is not in slice")
	}
	return nil
}

func (o *OneOfValidator) SetValue(value reflect.Value) {
	o.ref = value
}

func (o *OneOfValidator) isInSlice() bool {
	for _, v := range o.values {
		if v == o.ref.Interface() {
			return true
		}
	}
	return false
}

func (o *OneOfValidator) Msg(msg string) *OneOfValidator {
	o.setMsg(msg)
	return o
}

func OneOf(values ...interface{}) *OneOfValidator {
	return &OneOfValidator{
		values: values,
	}
}
