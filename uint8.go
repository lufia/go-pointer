package pointer

// Uint8 returns a pointer to uint8 that is initialized with v.
func Uint8(v uint8) *uint8 {
	return &v
}
