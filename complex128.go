package pointer

// Complex128 returns a pointer to complex128 that is initialized with v.
func Complex128(v complex128) *complex128 {
	return &v
}
