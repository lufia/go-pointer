// Package pointer provides functions that returns pointer to basic type.
package pointer

// New returns a pointer to T that is initialized with v.
func New[T any](v T) *T {
	return &v
}
