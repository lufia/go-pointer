package pointer

import "fmt"

// Complex128 returns a pointer to complex128 that is initialized with v.
func Complex128(v complex128) *complex128 {
	return &v
}

// Complex128Slice returns a slice of pointer to complex128.
func Complex128Slice(a ...complex128) []*complex128 {
	p := make([]*complex128, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Complex128Value returns a value referenced by p. If p is nil, it will returns zero value of complex128.
func Complex128Value(p *complex128) complex128 {
	if p != nil {
		return *p
	}
	var v complex128
	return v
}

// EqualComplex128 reports whether p1 and p2 represent the same value.
func EqualComplex128(p1, p2 *complex128) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Complex128Formatter implements fmt.Formatter of a pointer to complex128.
type Complex128Formatter struct {
	p *complex128
}

// NewComplex128Formatter returns the formatter of a pointer to complex128.
func NewComplex128Formatter(p *complex128) *Complex128Formatter {
	return &Complex128Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Complex128Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
