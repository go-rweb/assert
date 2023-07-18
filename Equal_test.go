package assert_test

import (
	"testing"

	"git.akyoto.dev/go/assert"
)

type T struct{ A int }

func TestEqual(t *testing.T) {
	assert.Equal(t, 0, 0)
	assert.Equal(t, "Hello", "Hello")
	assert.Equal(t, T{A: 10}, T{A: 10})
}

func TestNotEqual(t *testing.T) {
	assert.NotEqual(t, 0, 1)
	assert.NotEqual(t, "Hello", "World")
	assert.NotEqual(t, &T{A: 10}, &T{A: 10})
	assert.NotEqual(t, T{A: 10}, T{A: 20})
}

func TestDeepEqual(t *testing.T) {
	assert.DeepEqual(t, 0, 0)
	assert.DeepEqual(t, "Hello", "Hello")
	assert.DeepEqual(t, []byte("Hello"), []byte("Hello"))
	assert.DeepEqual(t, T{A: 10}, T{A: 10})
	assert.DeepEqual(t, &T{A: 10}, &T{A: 10})
}
