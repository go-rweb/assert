package assert_test

import (
	"testing"

	"github.com/rohanthewiz/assert"
)

func TestTrue(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}
