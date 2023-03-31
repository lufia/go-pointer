package pointer

// Uint8 returns a pointer to uint8 that is initialized with v.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint8Slice returns a slice of pointer to uint8.
func Uint8Slice(a ...uint8) []*uint8 {
	p := make([]*uint8, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// Uint8Value returns a value referenced by p. If p is nil, it will returns zero value of uint8.
func Uint8Value(p *uint8) uint8 {
	if p != nil {
		return *p
	}
	var v uint8
	return v
}

// EqualUint8 reports whether p1 and p2 represent the same value.
func EqualUint8(p1, p2 *uint8) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}
