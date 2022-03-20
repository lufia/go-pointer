//go:build go1.18
// +build go1.18

package pointer

// New returns a pointer to T that is initialized with v.
func New[T any](v T) *T {
	return &v
}

// Slice returns a slice of pointer to T.
func Slice[T any](a ...T) []*T {
	p := make([]*T, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Value returns a value referenced by p. If p is nil, it will returns zero value of T.
func Value[T any](p *T) T {
	if p != nil {
		return *p
	}
	var v T
	return v
}
