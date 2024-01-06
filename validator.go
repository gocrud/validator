package validator

import (
	"reflect"
)

type Validator interface {
    Validate() error
    SetValue(value reflect.Value)
}

func Validate(value interface{}, validators ...Validator) error {
    ref := getPointerValue(value)
    for _, v := range validators {
        v.SetValue(ref)
        if err := v.Validate(); err != nil {
            return err
        }
    }
    return nil
}
