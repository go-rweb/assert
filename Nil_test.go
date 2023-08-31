package assert_test

import (
	"testing"

	"git.akyoto.dev/go/assert"
)

func TestNil(t *testing.T) {
	var (
		nilPointer   *T
		nilInterface any
		nilSlice     []byte
		nilMap       map[byte]byte
		nilChannel   chan byte
		nilFunction  func()
	)

	assert.Nil(t, nil)
	assert.Nil(t, nilPointer)
	assert.Nil(t, nilInterface)
	assert.Nil(t, nilSlice)
	assert.Nil(t, nilMap)
	assert.Nil(t, nilChannel)
	assert.Nil(t, nilFunction)
}

func TestNotNil(t *testing.T) {
	assert.NotNil(t, 0)
	assert.NotNil(t, "Hello")
	assert.NotNil(t, T{})
	assert.NotNil(t, &T{})
	assert.NotNil(t, make([]byte, 0))
	assert.NotNil(t, make(map[byte]byte))
	assert.NotNil(t, make(chan byte))
	assert.NotNil(t, TestNotNil)
}

func TestFailNil(t *testing.T) {
	assert.Nil(fail(t), 0)
}

func TestFailNotNil(t *testing.T) {
	assert.NotNil(fail(t), nil)
}
