package pointer_test

import (
	"fmt"

	"github.com/lufia/go-pointer"
)

func Example() {
	v := pointer.Int(10)
	fmt.Println(*v)
	s := pointer.String("hello")
	fmt.Println(*s)
	// Output:
	// 10
	// hello
}

func Example_IntSlice() {
	a := pointer.IntSlice(10, 20)
	for _, p := range a {
		fmt.Println(*p)
	}
	// Output:
	// 10
	// 20
}
