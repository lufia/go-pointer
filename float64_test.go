package pointer

import (
	"fmt"
	"testing"
)

func TestFloat64(t *testing.T) {
	tests := []float64{
		0., 72.40, 6.67428e-11,
	}
	for _, v := range tests {
		var p *float64 = Float64(v)
		if p == nil {
			t.Errorf("Float64(%v) = nil; want %v", v, v)
		}
		if n := Float64Value(p); n != v {
			t.Errorf("Float64(%v) = %v; want %v", v, n, v)
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

func TestFloat64Value_Nil(t *testing.T) {
	var zero float64
	v := Float64Value(nil)
	if v != zero {
		t.Errorf("Float64Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualFloat64(t *testing.T) {
	tests := []struct {
		v1, v2 *float64
		eq     bool
	}{
		{Float64(0.), Float64(72.40), false},
		{Float64(0.), Float64(0.), true},
		{Float64(72.40), Float64(72.40), true},
		{nil, nil, true},
		{nil, Float64(0.), false},
		{Float64(0.), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualFloat64(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualFloat64(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestFloat64FormatterFormat(t *testing.T) {
	tests := []struct {
		p *float64
		s string
	}{
		{Float64(0.), fmt.Sprintf("%v", 0.)},
		{Float64(72.40), fmt.Sprintf("%v", 72.40)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewFloat64Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
