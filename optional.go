package validator

type OptionalValidator struct {
	validators []Validator
	AbstractValidator
}

func (o *OptionalValidator) Validate() error {
	if isEmpty(o.ref) {
		return nil
	}
	for _, v := range o.validators {
		v.setValue(o.ref)
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func Optional(validators ...Validator) *OptionalValidator {
	return &OptionalValidator{
		validators: validators,
	}
}
