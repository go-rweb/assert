package assert_test

import (
	"testing"

	"git.akyoto.dev/go/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}
