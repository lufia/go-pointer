package pointer

// Rune returns a pointer to rune that is initialized with v.
func Rune(v rune) *rune {
	return &v
}

// RuneSlice returns a slice of pointer to rune.
func RuneSlice(a ...rune) []*rune {
	p := make([]*rune, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
