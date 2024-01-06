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

func isIP(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}

// 所有类型空值判断
func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return value.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	default:
		return false
	}
}
