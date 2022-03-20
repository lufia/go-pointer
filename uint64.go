package pointer

// Uint64 returns a pointer to uint64 that is initialized with v.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64Slice returns a slice of pointer to uint64.
func Uint64Slice(a ...uint64) []*uint64 {
	p := make([]*uint64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
