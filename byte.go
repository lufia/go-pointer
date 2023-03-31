package pointer

// Byte returns a pointer to byte that is initialized with v.
func Byte(v byte) *byte {
	return &v
}

// ByteSlice returns a slice of pointer to byte.
func ByteSlice(a ...byte) []*byte {
	p := make([]*byte, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}

// ByteValue returns a value referenced by p. If p is nil, it will returns zero value of byte.
func ByteValue(p *byte) byte {
	if p != nil {
		return *p
	}
	var v byte
	return v
}

// EqualByte reports whether p1 and p2 represent the same value.
func EqualByte(p1, p2 *byte) bool {
	if p1 == nil || p2 == nil {
		return p1 == p2
	}
	return *p1 == *p2
}
