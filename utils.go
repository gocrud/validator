package validator

import (
	"net"
	"reflect"
)

func getPointerValue(value interface{}) reflect.Value {
	ref := reflect.ValueOf(value)
	if ref.Kind() != reflect.Ptr {
		return ref
	}
	return ref.Elem()
}

func isString(value interface{}) (reflect.Value, error) {
	ref := getPointerValue(value)
	if ref.Kind() != reflect.String {
		return ref, Error{Message: "value must be a string"}
	}
	return ref, nil
}

func isIP(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}

// 所有类型空值判断
func isEmpty(value interface{}) bool {
	ref := getPointerValue(value)
	switch ref.Kind() {
	case reflect.String:
		return ref.String() == ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return ref.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ref.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ref.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return ref.Float() == 0
	default:
		return false
	}
}
