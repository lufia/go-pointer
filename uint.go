package pointer

// Uint returns a pointer to uint that is initialized with v.
func Uint(v uint) *uint {
	return &v
}
