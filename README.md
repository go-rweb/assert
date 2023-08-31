# assert

A minimal & stateless assert package for writing tests.

## Features

- Simple interface
- Tiny codebase
- Zero dependencies

## Installation

```shell
go get git.akyoto.dev/go/assert
```

## Usage

```go
assert.Nil(t, nil)
assert.True(t, true)
assert.Equal(t, "Hello", "Hello")
assert.DeepEqual(t, "Hello", "Hello")
assert.Contains(t, "Hello", "ello")
```

## Tests

```
PASS: TestContains
PASS: TestNotContains
PASS: TestFailContains
PASS: TestFailNotContains
PASS: TestEqual
PASS: TestNotEqual
PASS: TestDeepEqual
PASS: TestFailEqual
PASS: TestFailNotEqual
PASS: TestFailDeepEqual
PASS: TestNil
PASS: TestNotNil
PASS: TestFailNil
PASS: TestFailNotNil
PASS: TestTrue
coverage: 100.0% of statements
```

## License

Please see the [license documentation](https://akyoto.dev/license).

## Copyright

© 2019 Eduard Urbach
