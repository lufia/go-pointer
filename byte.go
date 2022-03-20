package pointer

// Byte returns a pointer to byte that is initialized with v.
func Byte(v byte) *byte {
	return &v
}
