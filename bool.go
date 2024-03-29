package pointer

import "fmt"

// Bool returns a pointer to bool that is initialized with v.
func Bool(v bool) *bool {
	return &v
}

// BoolSlice returns a slice of pointer to bool.
func BoolSlice(a ...bool) []*bool {
	p := make([]*bool, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// BoolValue returns a value referenced by p. If p is nil, it will returns zero value of bool.
func BoolValue(p *bool) bool {
	if p != nil {
		return *p
	}
	var v bool
	return v
}

// EqualBool reports whether p1 and p2 represent the same value.
func EqualBool(p1, p2 *bool) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// BoolFormatter implements fmt.Formatter of a pointer to bool.
type BoolFormatter struct {
	p *bool
}

// NewBoolFormatter returns the formatter of a pointer to bool.
func NewBoolFormatter(p *bool) *BoolFormatter {
	return &BoolFormatter{p}
}

// Format implements the fmt.Formatter interface.
func (p BoolFormatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
