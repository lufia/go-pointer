package pointer

// Int32 returns a pointer to int32 that is initialized with v.
func Int32(v int32) *int32 {
	return &v
}
