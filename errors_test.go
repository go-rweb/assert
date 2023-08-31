package assert_test

import (
	"testing"
)

// noop implements the Test interface with noop functions so you can test the failure cases.
type noop struct{}

// Errorf does nothing.
func (t *noop) Errorf(format string, args ...any) {}

// FailNow does nothing.
func (t *noop) FailNow() {}

// fail creates a new test that is expected to fail.
func fail(t *testing.T) *noop {
	return &noop{}
}
