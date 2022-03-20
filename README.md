# go-pointer

Utilities that returns a pointer to the value.

[![GoDev][godev-image]][godev-url]
[![Actions Status][actions-image]][actions-url]
[![Coverage Status][coveralls-image]][coveralls-url]

## DESCRIPTION

**go-pointer** provides functions returned pointer to basic types in Go such as  *Int*, *Uint32* and *String*. When your Go version is 1.18 or higher, it provides generic *New* function too.

Rarely, when you need a slice of pointer to any basic types, you can use **Slice**-suffixed functions corresponded to the type.

### EXAMPLE

```go
import (
	"fmt"

	"github.com/lufia/go-pointer"
)

func main() {
	p := pointer.Int(10)
	fmt.Printf("p = %d\n", *p)

	for _, p := range pointer.StringSlice("hello", "world") {
		fmt.Printf("p = %s\n", *p)
	}

	// Go 1.18 or higher
	q := pointer.New(20)
	fmt.Printf("q = %d\n", *q)
	for _, p := range pointer.Slice(10, 20) {
		fmt.Printf("p = %d\n", *p)
	}
}
```

[godev-image]: https://pkg.go.dev/badge/github.com/lufia/go-pointer
[godev-url]: https://pkg.go.dev/github.com/lufia/go-pointer
[actions-image]: https://github.com/lufia/go-pointer/workflows/Test/badge.svg?branch=master
[actions-url]: https://github.com/lufia/go-pointer/actions?workflow=Test
[coveralls-image]: https://coveralls.io/repos/github/lufia/go-pointer/badge.svg
[coveralls-url]: https://coveralls.io/github/lufia/go-pointer
