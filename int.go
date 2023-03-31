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

// IntValue returns a value referenced by p. If p is nil, it will returns zero value of int.
func IntValue(p *int) int {
	if p != nil {
		return *p
	}
	var v int
	return v
}

// EqualInt reports whether p1 and p2 represent the same value.
func EqualInt(p1, p2 *int) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}
