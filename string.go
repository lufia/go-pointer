package pointer

// String returns a pointer to string that is initialized with v.
func String(v string) *string {
	return &v
}
