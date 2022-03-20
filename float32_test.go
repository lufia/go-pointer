package pointer

import (
	"testing"
)

func TestFloat32(t *testing.T) {
	tests := []float32{
		0., 72.40, 6.67428e-11,
	}
	for _, v := range tests {
		var p *float32 = Float32(v)
		switch {
		case p == nil:
			t.Errorf("Float32(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Float32(%v) = %v; want %v", v, *p, v)
		}
	}
}
