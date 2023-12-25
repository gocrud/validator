package validator

type OptionalValidator struct {
	validators []Validator
}

func (o *OptionalValidator) Validate(value interface{}) error {
	if isEmpty(value) {
		return nil
	}
	return Validate(value, o.validators...)
}

func Optional(validators ...Validator) Validator {
	return &OptionalValidator{
		validators: validators,
	}
}
