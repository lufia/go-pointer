package pointer

// Int returns a pointer to int that is initialized with v.
func Int(v int) *int {
	return &v
}
