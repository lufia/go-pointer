package pointer

import "fmt"

// Uint32 returns a pointer to uint32 that is initialized with v.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32Slice returns a slice of pointer to uint32.
func Uint32Slice(a ...uint32) []*uint32 {
	p := make([]*uint32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Uint32Value returns a value referenced by p. If p is nil, it will returns zero value of uint32.
func Uint32Value(p *uint32) uint32 {
	if p != nil {
		return *p
	}
	var v uint32
	return v
}

// EqualUint32 reports whether p1 and p2 represent the same value.
func EqualUint32(p1, p2 *uint32) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Uint32Formatter implements fmt.Formatter of a pointer to uint32.
type Uint32Formatter struct {
	p *uint32
}

// NewUint32Formatter returns the formatter of a pointer to uint32.
func NewUint32Formatter(p *uint32) *Uint32Formatter {
	return &Uint32Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Uint32Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
