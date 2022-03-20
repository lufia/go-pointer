package pointer

// Int64 returns a pointer to int64 that is initialized with v.
func Int64(v int64) *int64 {
	return &v
}
