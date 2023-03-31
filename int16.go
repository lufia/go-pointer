package pointer

import "fmt"

// Int16 returns a pointer to int16 that is initialized with v.
func Int16(v int16) *int16 {
	return &v
}

// Int16Slice returns a slice of pointer to int16.
func Int16Slice(a ...int16) []*int16 {
	p := make([]*int16, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Int16Value returns a value referenced by p. If p is nil, it will returns zero value of int16.
func Int16Value(p *int16) int16 {
	if p != nil {
		return *p
	}
	var v int16
	return v
}

// EqualInt16 reports whether p1 and p2 represent the same value.
func EqualInt16(p1, p2 *int16) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Int16Formatter implements fmt.Formatter of a pointer to int16.
type Int16Formatter struct {
	p *int16
}

// NewInt16Formatter returns the formatter of a pointer to int16.
func NewInt16Formatter(p *int16) *Int16Formatter {
	return &Int16Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Int16Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
