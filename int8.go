package pointer

// Int8 returns a pointer to int8 that is initialized with v.
func Int8(v int8) *int8 {
	return &v
}
