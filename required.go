package validator

import "reflect"

type RequiredValidator struct {
    errorMessage
    value reflect.Value
}

func (r *RequiredValidator) Validate() error {
    if isEmpty(r.value) {
        return r.errMsg("value is required")
    }
    return nil
}

func (r *RequiredValidator) Msg(msg string) *RequiredValidator {
    r.setMsg(msg)
    return r
}

func (r *RequiredValidator) SetValue(value reflect.Value) {
    r.value = value
}

func Required() *RequiredValidator {
    return &RequiredValidator{}
}
