package validator

type OneOfValidator struct {
	values []interface{}
	AbstractValidator
}

func (o *OneOfValidator) Validate() error {
	if !o.isInSlice() {
		return o.Error("value is not in slice")
	}
	return nil
}

func (o *OneOfValidator) isInSlice() bool {
	for _, v := range o.values {
		if v == o.ref.Interface() {
			return true
		}
	}
	return false
}

func OneOf(values ...interface{}) *OneOfValidator {
	return &OneOfValidator{
		values: values,
	}
}
