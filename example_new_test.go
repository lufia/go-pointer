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

func Example_New() {
	p := pointer.New(10)
	fmt.Println(*p)
	// Output: 10
}

func Example_Slice() {
	a := pointer.Slice(10, 20)
	for _, p := range a {
		fmt.Println(*p)
	}
	// Output:
	// 10
	// 20
}
