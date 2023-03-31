package pointer

import (
	"testing"
)

func TestUint16(t *testing.T) {
	tests := []uint16{
		0, 1, ^uint16(0),
	}
	for _, v := range tests {
		var p *uint16 = Uint16(v)
		if p == nil {
			t.Errorf("Uint16(%v) = nil; want %v", v, v)
		}
		if n := Uint16Value(p); n != v {
			t.Errorf("Uint16(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUint16Slice(t *testing.T) {
	data := []uint16{
		0, 1, ^uint16(0),
	}
	a := Uint16Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint16Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint16Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUint16Value_Nil(t *testing.T) {
	var zero uint16
	v := Uint16Value(nil)
	if v != zero {
		t.Errorf("Uint16Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualUint16(t *testing.T) {
	tests := []struct {
		v1, v2 *uint16
		eq     bool
	}{
		{Uint16(0), Uint16(1), false},
		{Uint16(0), Uint16(0), true},
		{Uint16(1), Uint16(1), true},
		{nil, nil, true},
		{nil, Uint16(0), false},
		{Uint16(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualUint16(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualUint16(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
