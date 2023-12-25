package validator

type IP struct {
	Message string
}

func (v *IP) Validate(value interface{}) error {
	ref, err := isString(value)
	if err != nil {
		return err
	}
	if !isIP(ref.String()) {
		var err = "invalid ip"
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}
