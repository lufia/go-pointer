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

func TestFloat64Slice(t *testing.T) {
	data := []float64{
		0., 72.40, 6.67428e-11,
	}
	a := Float64Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Float64Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Float64Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
