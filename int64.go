package pointer

import "fmt"

// Int64 returns a pointer to int64 that is initialized with v.
func Int64(v int64) *int64 {
	return &v
}

// Int64Slice returns a slice of pointer to int64.
func Int64Slice(a ...int64) []*int64 {
	p := make([]*int64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Int64Value returns a value referenced by p. If p is nil, it will returns zero value of int64.
func Int64Value(p *int64) int64 {
	if p != nil {
		return *p
	}
	var v int64
	return v
}

// EqualInt64 reports whether p1 and p2 represent the same value.
func EqualInt64(p1, p2 *int64) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// Int64Formatter implements fmt.Formatter of a pointer to int64.
type Int64Formatter struct {
	p *int64
}

// NewInt64Formatter returns the formatter of a pointer to int64.
func NewInt64Formatter(p *int64) *Int64Formatter {
	return &Int64Formatter{p}
}

// Format implements the fmt.Formatter interface.
func (p Int64Formatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
