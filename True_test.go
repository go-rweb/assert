package assert_test

import (
	"testing"

	"github.com/go-rweb/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}
