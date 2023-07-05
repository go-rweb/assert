package assert

import (
	"runtime/debug"
	"strings"
	"testing"
)

// file returns the first line containing "_test.go" in the debug stack.
func file(t testing.TB) string {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")

	for _, line := range lines {
		if strings.Contains(line, "_test.go") {
			return strings.TrimSpace(line)
		}
	}

	return ""
}
