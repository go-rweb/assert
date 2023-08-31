package assert

import (
	"reflect"
)

// Nil asserts that the given parameter equals nil.
func Nil(t test, a any) {
	if isNil(a) {
		return
	}

	t.Errorf(oneParameter, file(), "Nil", a)
	t.FailNow()
}

// NotNil asserts that the given parameter does not equal nil.
func NotNil(t test, a any) {
	if !isNil(a) {
		return
	}

	t.Errorf(oneParameter, file(), "NotNil", a)
	t.FailNow()
}

// isNil returns true if the object is nil.
func isNil(object any) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)

	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return value.IsNil()
	}

	return false
}
