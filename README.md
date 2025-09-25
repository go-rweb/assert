# assert

A minimal & stateless assert package for writing tests.

Forked from git.akyoto.dev/go/assert
because `go mod tidy` failures here

## Features

- Simple interface
- Tiny codebase
- Zero dependencies

## Installation

```shell
go get github.com/go-rweb/assert
```

## Usage

```go
assert.Nil(t, nil)
assert.True(t, true)
assert.Equal(t, "Hello", "Hello")
assert.DeepEqual(t, "Hello", "Hello")
assert.Contains(t, "Hello", "ello")
```
