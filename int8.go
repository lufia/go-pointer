package pointer

// Int8 returns a pointer to int8 that is initialized with v.
func Int8(v int8) *int8 {
	return &v
}

// Int8Slice returns a slice of pointer to int8.
func Int8Slice(a ...int8) []*int8 {
	p := make([]*int8, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
