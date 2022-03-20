package pointer

// Int32 returns a pointer to int32 that is initialized with v.
func Int32(v int32) *int32 {
	return &v
}

// Int32Slice returns a slice of pointer to int32.
func Int32Slice(a ...int32) []*int32 {
	p := make([]*int32, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
