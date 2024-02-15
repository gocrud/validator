package validator

type PrefixValidator struct {
    AbstractValidator
    prefix string
}

func (p *PrefixValidator) Validate() error {
    if !isString(p.ref) {
        return p.Error("Prefix: value must be a string")
    }

    if !hasPrefix(p.ref.String(), p.prefix) {
        return p.Error(p.errMsg)
    }

    return nil
}

func (p *PrefixValidator) Msg(msg string) Validator {
    p.errMsg = msg
    return p
}

func Prefix(prefix string) *PrefixValidator {
    v := PrefixValidator{}
    v.prefix = prefix
    return &v
}
