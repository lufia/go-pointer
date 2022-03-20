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
