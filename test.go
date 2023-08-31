package assert

// test is the interface used for tests.
type test interface {
	Errorf(format string, args ...any)
	FailNow()
}
