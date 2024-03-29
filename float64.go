package pointer

import "fmt"

// Float64 returns a pointer to float64 that is initialized with v.
func Float64(v float64) *float64 {
	return &v
}

// Float64Slice returns a slice of pointer to float64.
func Float64Slice(a ...float64) []*float64 {
	p := make([]*float64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Float64Value returns a value referenced by p. If p is nil, it will returns zero value of float64.
func Float64Value(p *float64) float64 {
	if p != nil {
		return *p
	}
	var v float64
	return v
}

// EqualFloat64 reports whether p1 and p2 represent the same value.
func EqualFloat64(p1, p2 *float64) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Float64Formatter implements fmt.Formatter of a pointer to float64.
type Float64Formatter struct {
	p *float64
}

// NewFloat64Formatter returns the formatter of a pointer to float64.
func NewFloat64Formatter(p *float64) *Float64Formatter {
	return &Float64Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Float64Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
