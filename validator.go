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
	Field      interface{}
	Validators []Validator
}

// ValidateMultiple validates multiple values with multiple validators
func ValidateMultiple(validators ...Group) error {
	for _, v := range validators {
		if err := Validate(v.Field, v.Validators...); err != nil {
			return err
		}
	}
	return nil
}
