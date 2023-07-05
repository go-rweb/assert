# assert

A minimal & stateless assert package for writing tests.

## API

```go
assert.Nil(t, nil)
assert.True(t, true)
assert.Equal(t, "Hello", "Hello")
assert.DeepEqual(t, "Hello", "Hello")
assert.Contains(t, "Hello", "ello")
```
