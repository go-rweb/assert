package assert

// True asserts that the given parameter is true.
func True(t test, a bool) {
	Equal(t, a, true)
}

// False asserts that the given parameter is false.
func False(t test, a bool) {
	Equal(t, a, false)
}
