package pointer

import "fmt"

// String returns a pointer to string that is initialized with v.
func String(v string) *string {
	return &v
}

// StringSlice returns a slice of pointer to string.
func StringSlice(a ...string) []*string {
	p := make([]*string, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// StringValue returns a value referenced by p. If p is nil, it will returns zero value of string.
func StringValue(p *string) string {
	if p != nil {
		return *p
	}
	var v string
	return v
}

// EqualString reports whether p1 and p2 represent the same value.
func EqualString(p1, p2 *string) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}

// StringFormatter implements fmt.Formatter of a pointer to string.
type StringFormatter struct {
	p *string
}

// NewStringFormatter returns the formatter of a pointer to string.
func NewStringFormatter(p *string) *StringFormatter {
	return &StringFormatter{p}
}

// Format implements the fmt.Formatter interface.
func (p StringFormatter) Format(f fmt.State, c rune) {
	if p.p == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.p)
}
