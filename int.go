package pointer

// Int returns a pointer to int that is initialized with v.
func Int(v int) *int {
	return &v
}

// IntSlice returns a slice of pointer to int.
func IntSlice(a ...int) []*int {
	p := make([]*int, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
