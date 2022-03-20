package pointer

// Int64 returns a pointer to int64 that is initialized with v.
func Int64(v int64) *int64 {
	return &v
}

// Int64Slice returns a slice of pointer to int64.
func Int64Slice(a ...int64) []*int64 {
	p := make([]*int64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
