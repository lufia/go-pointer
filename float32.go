package pointer

// Float32 returns a pointer to float32 that is initialized with v.
func Float32(v float32) *float32 {
	return &v
}
