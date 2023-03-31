package pointer

// Uint16 returns a pointer to uint16 that is initialized with v.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint16Slice returns a slice of pointer to uint16.
func Uint16Slice(a ...uint16) []*uint16 {
	p := make([]*uint16, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Uint16Value returns a value referenced by p. If p is nil, it will returns zero value of uint16.
func Uint16Value(p *uint16) uint16 {
	if p != nil {
		return *p
	}
	var v uint16
	return v
}

// EqualUint16 reports whether p1 and p2 represent the same value.
func EqualUint16(p1, p2 *uint16) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}
