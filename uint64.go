package pointer

// Uint64 returns a pointer to uint64 that is initialized with v.
func Uint64(v uint64) *uint64 {
	return &v
}
