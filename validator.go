package validator

type Validator interface {
	Validate(value interface{}) error
}

func Validate(value interface{}, validators ...Validator) error {
	for _, v := range validators {
		if err := v.Validate(value); err != nil {
			return err
		}
	}
	return nil
}

type Group struct {
	Value      interface{}
	Validators []Validator
}

// ValidateAll validates all values
func ValidateAll(validators ...Group) error {
	for _, v := range validators {
		if err := Validate(v.Value, v.Validators...); err != nil {
			return err
		}
	}
	return nil
}

func Optional(value interface{}, validators ...Validator) error {
	if isEmpty(value) {
		return nil
	}
	return Validate(value, validators...)
}

func OptionalAll(validators ...Group) error {
	for _, v := range validators {
		if err := Optional(v.Value, v.Validators...); err != nil {
			return err
		}
	}
	return nil
}
