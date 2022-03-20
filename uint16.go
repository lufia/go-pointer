package pointer

// Uint16 returns a pointer to uint16 that is initialized with v.
func Uint16(v uint16) *uint16 {
	return &v
}
