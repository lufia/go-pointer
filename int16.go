package pointer

// Int16 returns a pointer to int16 that is initialized with v.
func Int16(v int16) *int16 {
	return &v
}
