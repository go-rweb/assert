package assert

import (
	"reflect"
	"strings"
)

// Contains asserts that a contains b.
func Contains(t test, a any, b any) {
	if contains(a, b) {
		return
	}

	t.Errorf(twoParameters, file(), "Contains", a, b)
	t.FailNow()
}

// NotContains asserts that a doesn't contain b.
func NotContains(t test, a any, b any) {
	if !contains(a, b) {
		return
	}

	t.Errorf(twoParameters, file(), "NotContains", a, b)
	t.FailNow()
}

// contains returns whether container contains the given the element.
// It works with strings, maps and slices.
func contains(container any, element any) bool {
	containerValue := reflect.ValueOf(container)

	switch containerValue.Kind() {
	case reflect.String:
		elementValue := reflect.ValueOf(element)
		return strings.Contains(containerValue.String(), elementValue.String())

	case reflect.Map:
		keys := containerValue.MapKeys()

		for _, key := range keys {
			if key.Interface() == element {
				return true
			}
		}

	case reflect.Slice:
		elementValue := reflect.ValueOf(element)

		if elementValue.Kind() == reflect.Slice {
			elementLength := elementValue.Len()

			if elementLength == 0 {
				return true
			}

			if elementLength > containerValue.Len() {
				return false
			}

			matchingElements := 0

			for i := 0; i < containerValue.Len(); i++ {
				if containerValue.Index(i).Interface() == elementValue.Index(matchingElements).Interface() {
					matchingElements++
				} else {
					matchingElements = 0
				}

				if matchingElements == elementLength {
					return true
				}
			}

			return false
		}

		for i := 0; i < containerValue.Len(); i++ {
			if containerValue.Index(i).Interface() == element {
				return true
			}
		}
	}

	return false
}
