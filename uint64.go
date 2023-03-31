package pointer

import "fmt"

// Uint64 returns a pointer to uint64 that is initialized with v.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64Slice returns a slice of pointer to uint64.
func Uint64Slice(a ...uint64) []*uint64 {
	p := make([]*uint64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Uint64Value returns a value referenced by p. If p is nil, it will returns zero value of uint64.
func Uint64Value(p *uint64) uint64 {
	if p != nil {
		return *p
	}
	var v uint64
	return v
}

// EqualUint64 reports whether p1 and p2 represent the same value.
func EqualUint64(p1, p2 *uint64) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Uint64Formatter implements fmt.Formatter of a pointer to uint64.
type Uint64Formatter struct {
	p *uint64
}

// NewUint64Formatter returns the formatter of a pointer to uint64.
func NewUint64Formatter(p *uint64) *Uint64Formatter {
	return &Uint64Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Uint64Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
