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

// Uint32Value returns a value referenced by p. If p is nil, it will returns zero value of uint32.
func Uint32Value(p *uint32) uint32 {
	if p != nil {
		return *p
	}
	var v uint32
	return v
}
