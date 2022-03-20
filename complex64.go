package pointer

// Complex64 returns a pointer to complex64 that is initialized with v.
func Complex64(v complex64) *complex64 {
	return &v
}
