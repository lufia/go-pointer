package pointer

import "fmt"

// Uint returns a pointer to uint that is initialized with v.
func Uint(v uint) *uint {
	return &v
}

// UintSlice returns a slice of pointer to uint.
func UintSlice(a ...uint) []*uint {
	p := make([]*uint, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// UintValue returns a value referenced by p. If p is nil, it will returns zero value of uint.
func UintValue(p *uint) uint {
	if p != nil {
		return *p
	}
	var v uint
	return v
}

// EqualUint reports whether p1 and p2 represent the same value.
func EqualUint(p1, p2 *uint) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// UintFormatter implements fmt.Formatter of a pointer to uint.
type UintFormatter struct {
	p *uint
}

// NewUintFormatter returns the formatter of a pointer to uint.
func NewUintFormatter(p *uint) *UintFormatter {
	return &UintFormatter{p}
}

// Format implements the fmt.Formatter interface.
func (p UintFormatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
