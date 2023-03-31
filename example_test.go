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

func ExampleEqualInt() {
	p1 := pointer.New(10)
	p2 := pointer.New(10)
	fmt.Println(pointer.Equal(p1, p2))
	fmt.Println(pointer.Equal(p1, nil))
	// Output:
	// true
	// false
}

func ExampleIntFormatter_Format() {
	fmt.Printf("%v\n", pointer.NewIntFormatter(pointer.Int(10)))
	fmt.Printf("%v\n", pointer.NewStringFormatter(nil))
	// Output:
	// 10
	// <nil>
}
