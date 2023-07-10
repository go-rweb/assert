package assert

import (
	"reflect"
	"testing"
)

// Equal asserts that the two parameters are equal.
func Equal[T comparable](t testing.TB, a T, b T) {
	if a == b {
		return
	}

	t.Errorf(formatTwoParameters, file(), "Equal", a, b)
	t.FailNow()
}

// NotEqual asserts that the two parameters are not equal.
func NotEqual[T comparable](t testing.TB, a T, b T) {
	if a != b {
		return
	}

	t.Errorf(formatTwoParameters, file(), "NotEqual", a, b)
	t.FailNow()
}

// DeepEqual asserts that the two parameters are deeply equal.
func DeepEqual[T comparable](t testing.TB, a T, b T) {
	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf(formatTwoParameters, file(), "DeepEqual", a, b)
	t.FailNow()
}
