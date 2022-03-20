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

func TestFloat32Slice(t *testing.T) {
	data := []float32{
		0., 72.40, 6.67428e-11,
	}
	a := Float32Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Float32Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Float32Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
