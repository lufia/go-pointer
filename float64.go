package pointer

// Float64 returns a pointer to float64 that is initialized with v.
func Float64(v float64) *float64 {
	return &v
}
