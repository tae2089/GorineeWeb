package main

import (
	"unsafe"
)

// GetString gets the content of a string as a []byte without copying
// #nosec G103
func GetString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func InterfaceConvertIntType(v interface{}) int {
	return v.(int)
}

func InterfaceConvertStringType(v interface{}) string {
	return v.(string)
}

func InterfaceConvertFloat32Type(v interface{}) float32 {
	return v.(float32)
}
func InterfaceConvertFloat64Type(v interface{}) float64 {
	return v.(float64)
}
