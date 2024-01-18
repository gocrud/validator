package validator

type RequiredValidator struct {
	AbstractValidator
}

func (r *RequiredValidator) Validate() error {
	if isEmpty(r.ref) {
		return r.Error("value is required")
	}
	return nil
}

func Required() *RequiredValidator {
	return &RequiredValidator{}
}
