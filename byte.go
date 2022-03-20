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
