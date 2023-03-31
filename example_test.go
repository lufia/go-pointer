package pointer_test

import (
	"fmt"

	"github.com/lufia/go-pointer"
)

func Example() {
	v := pointer.Int(10)
	fmt.Println(pointer.IntValue(v))
	s := pointer.String("hello")
	fmt.Println(pointer.StringValue(s))
	// Output:
	// 10
	// hello
}

func ExampleIntSlice() {
	a := pointer.IntSlice(10, 20)
	for _, p := range a {
		fmt.Println(pointer.IntValue(p))
	}
	// Output:
	// 10
	// 20
}
