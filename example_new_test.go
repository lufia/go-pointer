//go:build go1.18
// +build go1.18

package pointer_test

import (
	"fmt"

	"github.com/lufia/go-pointer"
)

func Example_go1_18() {
	s := pointer.New("hello")
	fmt.Println(pointer.Value(s))
	// Output: hello
}

func ExampleNew() {
	p := pointer.New(10)
	fmt.Println(pointer.Value(p))
	// Output: 10
}

func ExampleSlice() {
	a := pointer.Slice(10, 20)
	for _, p := range a {
		fmt.Println(pointer.Value(p))
	}
	// Output:
	// 10
	// 20
}
