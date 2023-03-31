package pointer

import (
	"fmt"
	"testing"
)

func TestFloat32(t *testing.T) {
	tests := []float32{
		0., 72.40, 6.67428e-11,
	}
	for _, v := range tests {
		var p *float32 = Float32(v)
		if p == nil {
			t.Errorf("Float32(%v) = nil; want %v", v, v)
		}
		if n := Float32Value(p); n != v {
			t.Errorf("Float32(%v) = %v; want %v", v, n, v)
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

func TestFloat32Value_Nil(t *testing.T) {
	var zero float32
	v := Float32Value(nil)
	if v != zero {
		t.Errorf("Float32Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualFloat32(t *testing.T) {
	tests := []struct {
		v1, v2 *float32
		eq     bool
	}{
		{Float32(0.), Float32(72.40), false},
		{Float32(0.), Float32(0.), true},
		{Float32(72.40), Float32(72.40), true},
		{nil, nil, true},
		{nil, Float32(0.), false},
		{Float32(0.), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualFloat32(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualFloat32(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestFloat32FormatterFormat(t *testing.T) {
	tests := []struct {
		p *float32
		s string
	}{
		{Float32(0.), fmt.Sprintf("%v", 0.)},
		{Float32(72.40), fmt.Sprintf("%v", 72.40)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewFloat32Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
