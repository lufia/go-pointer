package pointer

// Float64 returns a pointer to float64 that is initialized with v.
func Float64(v float64) *float64 {
	return &v
}

// Float64Slice returns a slice of pointer to float64.
func Float64Slice(a ...float64) []*float64 {
	p := make([]*float64, len(a))
	for i, v := range a {
		v := v
		p[i] = &v
	}
	return p
}
