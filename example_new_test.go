//go:build go1.18
// +build go1.18

package pointer_test

import (
	"fmt"

	"github.com/lufia/go-pointer"
)

func Example_go1_18() {
	s := pointer.New("hello")
	fmt.Println(*s)
	// Output: hello
}
