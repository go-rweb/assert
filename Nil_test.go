
package assert_test

import (
	"testing"

	"git.akyoto.dev/go/assert"
)

func TestNil(t *testing.T) {
	var nilPointer *T
	var nilSlice []byte

	assert.Nil(t, nil)
	assert.Nil(t, nilPointer)
	assert.Nil(t, nilSlice)
}

func TestNotNil(t *testing.T) {
	assert.NotNil(t, 0)
	assert.NotNil(t, "Hello")
	assert.NotNil(t, []byte{})
	assert.NotNil(t, T{})
	assert.NotNil(t, &T{})
}
