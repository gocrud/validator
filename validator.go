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

type Validators struct {
	Value      interface{}
	Validators []Validator
}

// ValidateAll validates all values
func ValidateAll(validators ...Validators) error {
	for _, v := range validators {
		if err := Validate(v.Value, v.Validators...); err != nil {
			return err
		}
	}
	return nil
}
