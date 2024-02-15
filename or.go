package validator

type OrValidator struct {
    AbstractValidator
    validators []Validator
}

func (o *OrValidator) Validate() error {
    for _, v := range o.validators {
        v.setValue(o.ref)
        if err := v.Validate(); err == nil {
            return nil
        }
    }
    return o.Error(o.errMsg)
}

func (o *OrValidator) Msg(msg string) Validator {
    o.errMsg = msg
    return o
}

func Or(validators ...Validator) *OrValidator {
    return &OrValidator{
        validators: validators,
    }
}
