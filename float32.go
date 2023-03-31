package pointer

import "fmt"

// Float32 returns a pointer to float32 that is initialized with v.
func Float32(v float32) *float32 {
	return &v
}

// Float32Slice returns a slice of pointer to float32.
func Float32Slice(a ...float32) []*float32 {
	p := make([]*float32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Float32Value returns a value referenced by p. If p is nil, it will returns zero value of float32.
func Float32Value(p *float32) float32 {
	if p != nil {
		return *p
	}
	var v float32
	return v
}

// EqualFloat32 reports whether p1 and p2 represent the same value.
func EqualFloat32(p1, p2 *float32) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Float32Formatter implements fmt.Formatter of a pointer to float32.
type Float32Formatter struct {
	p *float32
}

// NewFloat32Formatter returns the formatter of a pointer to float32.
func NewFloat32Formatter(p *float32) *Float32Formatter {
	return &Float32Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Float32Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
