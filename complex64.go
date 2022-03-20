package pointer

// Complex64 returns a pointer to complex64 that is initialized with v.
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex64Slice returns a slice of pointer to complex64.
func Complex64Slice(a ...complex64) []*complex64 {
	p := make([]*complex64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
