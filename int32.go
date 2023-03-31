package pointer

import "fmt"

// Int32 returns a pointer to int32 that is initialized with v.
func Int32(v int32) *int32 {
	return &v
}

// Int32Slice returns a slice of pointer to int32.
func Int32Slice(a ...int32) []*int32 {
	p := make([]*int32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Int32Value returns a value referenced by p. If p is nil, it will returns zero value of int32.
func Int32Value(p *int32) int32 {
	if p != nil {
		return *p
	}
	var v int32
	return v
}

// EqualInt32 reports whether p1 and p2 represent the same value.
func EqualInt32(p1, p2 *int32) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Int32Formatter implements fmt.Formatter of a pointer to int32.
type Int32Formatter struct {
	p *int32
}

// NewInt32Formatter returns the formatter of a pointer to int32.
func NewInt32Formatter(p *int32) *Int32Formatter {
	return &Int32Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Int32Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
