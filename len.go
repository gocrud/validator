package validator

import (
	"fmt"
	"reflect"
)

type StrLen struct {
	Min     int
	Max     int
	Message string
}

func (v *StrLen) Validate(value interface{}) error {
	ref, err := isString(value)
	if err != nil {
		return err
	}
	r := []rune(ref.String())
	if len(r) < v.Min || len(r) > v.Max {
		var err = fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type SliceLen struct {
	Min     int
	Max     int
	Message string
}

func (v *SliceLen) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Slice {
		return Error{Message: "value must be a slice"}
	}
	if ref.Len() < v.Min || ref.Len() > v.Max {
		var err = fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}

type MapLen struct {
	Min int
	Max int
}

func (v *MapLen) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Map {
		return Error{Message: "value must be a map"}
	}
	if ref.Len() < v.Min || ref.Len() > v.Max {
		return &Error{Message: fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)}
	}
	return nil
}

type ArrayLen struct {
	Min     int
	Max     int
	Message string
}

func (v *ArrayLen) Validate(value interface{}) error {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Array {
		return Error{Message: "value must be a array"}
	}
	if ref.Len() < v.Min || ref.Len() > v.Max {
		var err = fmt.Sprintf("value length must be between %d and %d", v.Min, v.Max)
		if v.Message != "" {
			err = v.Message
		}
		return &Error{Message: err}
	}
	return nil
}
