package assert_test

import (
	"testing"

	"git.akyoto.dev/go/assert"
)

func TestContains(t *testing.T) {
	assert.Contains(t, "Hello", "H")
	assert.Contains(t, "Hello", "Hello")
	assert.Contains(t, []string{"Hello", "World"}, "Hello")
	assert.Contains(t, []int{1, 2, 3}, 2)
	assert.Contains(t, []int{1, 2, 3}, []int{})
	assert.Contains(t, []int{1, 2, 3}, []int{1, 2})
	assert.Contains(t, []byte{'H', 'e', 'l', 'l', 'o'}, byte('e'))
	assert.Contains(t, []byte{'H', 'e', 'l', 'l', 'o'}, []byte{'e', 'l'})
	assert.Contains(t, map[string]int{"Hello": 1, "World": 2}, "Hello")
}

func TestNotContains(t *testing.T) {
	assert.NotContains(t, "Hello", "h")
	assert.NotContains(t, "Hello", "hello")
	assert.NotContains(t, []string{"Hello", "World"}, "hello")
	assert.NotContains(t, []int{1, 2, 3}, 4)
	assert.NotContains(t, []int{1, 2, 3}, []int{2, 1})
	assert.NotContains(t, []int{1, 2, 3}, []int{1, 2, 3, 4})
	assert.NotContains(t, []byte{'H', 'e', 'l', 'l', 'o'}, byte('a'))
	assert.NotContains(t, []byte{'H', 'e', 'l', 'l', 'o'}, []byte{'l', 'e'})
	assert.NotContains(t, map[string]int{"Hello": 1, "World": 2}, "hello")
}

func TestFailContains(t *testing.T) {
	assert.Contains(fail(t), "Hello", "h")
}

func TestFailNotContains(t *testing.T) {
	assert.NotContains(fail(t), "Hello", "H")
}
