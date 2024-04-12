package validator

type InValidator struct {
	values []interface{}
	AbstractValidator
}

func (o *InValidator) Validate() error {
	if !o.isInSlice() {
		return o.Error("value is not in slice")
	}
	return nil
}

func (o *InValidator) Msg(msg string) Validator {
	o.errMsg = msg
	return o
}

func (o *InValidator) isInSlice() bool {
	for _, v := range o.values {
		if v == o.ref.Interface() {
			return true
		}
	}
	return false
}

func In(values ...interface{}) *InValidator {
	return &InValidator{
		values: values,
	}
}
