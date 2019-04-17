// Package pointer provides functions that returns pointer to basic type.
package pointer

// Int returns pointer to int that is initialized with n.
func Int(n int) *int {
	return &n
}
