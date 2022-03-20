package pointer

// String returns a pointer to string that is initialized with v.
func String(v string) *string {
	return &v
}

// StringSlice returns a slice of pointer to string.
func StringSlice(a ...string) []*string {
	p := make([]*string, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// StringValue returns a value referenced by p. If p is nil, it will returns zero value of string.
func StringValue(p *string) string {
	if p != nil {
		return *p
	}
	var v string
	return v
}
