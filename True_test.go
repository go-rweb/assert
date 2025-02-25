package assert_test

import (
	"testing"

	"git.urbach.dev/go/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}
