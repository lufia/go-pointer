package pointer

// Uint32 returns a pointer to uint32 that is initialized with v.
func Uint32(v uint32) *uint32 {
	return &v
}
