package validator

import (
	"net"
	"reflect"
)

func isPtr(value interface{}) reflect.Value {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Ptr {
		return ref
	}
	return ref.Elem()
}

func isString(value interface{}) (reflect.Value, error) {
	ref := isPtr(value)
	if ref.Kind() != reflect.String {
		return ref, Error{Message: "value must be a string"}
	}
	return ref, nil
}

func isNumber(value interface{}) (reflect.Value, error) {
	ref := isPtr(value)
	switch ref.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ref, nil
	default:
		return ref, Error{Message: "value must be a number"}
	}
	return ref, nil
}

func isIP(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
