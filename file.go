package assert

import (
	"runtime/debug"
	"strings"
)

// file returns the first line containing "_test.go" in the debug stack.
func file() string {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")

	for _, line := range lines {
		if strings.Contains(line, "_test.go") {
			space := strings.LastIndex(line, " ")

			if space != -1 {
				line = line[:space]
			}

			return strings.TrimSpace(line)
		}
	}

	return ""
}
