package pointer

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	tests := []float64{
		0., 72.40, 6.67428e-11,
	}
	for _, v := range tests {
		var p *float64 = Float64(v)
		switch {
		case p == nil:
			t.Errorf("Float64(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Float64(%v) = %v; want %v", v, *p, v)
		}
	}
}
