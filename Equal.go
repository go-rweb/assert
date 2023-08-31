package assert

import (
	"reflect"
)

// Equal asserts that the two parameters are equal.
func Equal[T comparable](t test, a T, b T) {
	if a == b {
		return
	}

	t.Errorf(twoParameters, file(), "Equal", a, b)
	t.FailNow()
}

// NotEqual asserts that the two parameters are not equal.
func NotEqual[T comparable](t test, a T, b T) {
	if a != b {
		return
	}

	t.Errorf(twoParameters, file(), "NotEqual", a, b)
	t.FailNow()
}

// DeepEqual asserts that the two parameters are deeply equal.
func DeepEqual[T any](t test, a T, b T) {
	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf(twoParameters, file(), "DeepEqual", a, b)
	t.FailNow()
}
