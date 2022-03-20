package pointer

// Int16 returns a pointer to int16 that is initialized with v.
func Int16(v int16) *int16 {
	return &v
}

// Int16Slice returns a slice of pointer to int16.
func Int16Slice(a ...int16) []*int16 {
	p := make([]*int16, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
