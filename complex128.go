package pointer

// Complex128 returns a pointer to complex128 that is initialized with v.
func Complex128(v complex128) *complex128 {
	return &v
}

// Complex128Slice returns a slice of pointer to complex128.
func Complex128Slice(a ...complex128) []*complex128 {
	p := make([]*complex128, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
