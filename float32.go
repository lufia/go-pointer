package pointer

// Float32 returns a pointer to float32 that is initialized with v.
func Float32(v float32) *float32 {
	return &v
}

// Float32Slice returns a slice of pointer to float32.
func Float32Slice(a ...float32) []*float32 {
	p := make([]*float32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
