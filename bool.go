package pointer

// Bool returns a pointer to bool that is initialized with v.
func Bool(v bool) *bool {
	return &v
}
