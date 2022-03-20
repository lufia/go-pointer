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

// Int32Value returns a value referenced by p. If p is nil, it will returns zero value of int32.
func Int32Value(p *int32) int32 {
	if p != nil {
		return *p
	}
	var v int32
	return v
}
