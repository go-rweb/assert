package assert

import (
	"runtime/debug"
	"strings"
)

const oneParameter = `
  %s
󰅙  assert.%s
    󰯬  %v`

const twoParameters = `
  %s
󰅙  assert.%s
    󰯬  %v
    󰯯  %v`

// file returns the first line containing "_test.go" in the debug stack.
func file() string {
	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	name := ""

	for _, line := range lines {
		if strings.Contains(line, "_test.go") {
			space := strings.LastIndex(line, " ")

			if space != -1 {
				line = line[:space]
			}

			name = strings.TrimSpace(line)
			break
		}
	}

	return name
}
