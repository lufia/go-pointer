package pointer

// Uint returns a pointer to uint that is initialized with v.
func Uint(v uint) *uint {
	return &v
}

// UintSlice returns a slice of pointer to uint.
func UintSlice(a ...uint) []*uint {
	p := make([]*uint, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// UintValue returns a value referenced by p. If p is nil, it will returns zero value of uint.
func UintValue(p *uint) uint {
	if p != nil {
		return *p
	}
	var v uint
	return v
}
