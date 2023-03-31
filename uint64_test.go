package pointer

import (
	"testing"
)

func TestUint64(t *testing.T) {
	tests := []uint64{
		0, 1, ^uint64(0),
	}
	for _, v := range tests {
		var p *uint64 = Uint64(v)
		if p == nil {
			t.Errorf("Uint64(%v) = nil; want %v", v, v)
		}
		if n := Uint64Value(p); n != v {
			t.Errorf("Uint64(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUint64Slice(t *testing.T) {
	data := []uint64{
		0, 1, ^uint64(0),
	}
	a := Uint64Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint64Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint64Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUint64Value_Nil(t *testing.T) {
	var zero uint64
	v := Uint64Value(nil)
	if v != zero {
		t.Errorf("Uint64Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualUint64(t *testing.T) {
	tests := []struct {
		v1, v2 *uint64
		eq     bool
	}{
		{Uint64(0), Uint64(1), false},
		{Uint64(0), Uint64(0), true},
		{Uint64(1), Uint64(1), true},
		{nil, nil, true},
		{nil, Uint64(0), false},
		{Uint64(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualUint64(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualUint64(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
