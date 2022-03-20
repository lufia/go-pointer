package pointer

// Uint32 returns a pointer to uint32 that is initialized with v.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32Slice returns a slice of pointer to uint32.
func Uint32Slice(a ...uint32) []*uint32 {
	p := make([]*uint32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
