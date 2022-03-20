package pointer

// Rune returns a pointer to rune that is initialized with v.
func Rune(v rune) *rune {
	return &v
}
