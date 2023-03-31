package pointer

import "fmt"

// Int8 returns a pointer to int8 that is initialized with v.
func Int8(v int8) *int8 {
	return &v
}

// Int8Slice returns a slice of pointer to int8.
func Int8Slice(a ...int8) []*int8 {
	p := make([]*int8, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Int8Value returns a value referenced by p. If p is nil, it will returns zero value of int8.
func Int8Value(p *int8) int8 {
	if p != nil {
		return *p
	}
	var v int8
	return v
}

// EqualInt8 reports whether p1 and p2 represent the same value.
func EqualInt8(p1, p2 *int8) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Int8Formatter implements fmt.Formatter of a pointer to int8.
type Int8Formatter struct {
	p *int8
}

// NewInt8Formatter returns the formatter of a pointer to int8.
func NewInt8Formatter(p *int8) *Int8Formatter {
	return &Int8Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Int8Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
