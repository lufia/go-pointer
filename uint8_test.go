package pointer

import (
	"fmt"
	"testing"
)

func TestUint8(t *testing.T) {
	tests := []uint8{
		0, 1, ^uint8(0),
	}
	for _, v := range tests {
		var p *uint8 = Uint8(v)
		if p == nil {
			t.Errorf("Uint8(%v) = nil; want %v", v, v)
		}
		if n := Uint8Value(p); n != v {
			t.Errorf("Uint8(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUint8Slice(t *testing.T) {
	data := []uint8{
		0, 1, ^uint8(0),
	}
	a := Uint8Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint8Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint8Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUint8Value_Nil(t *testing.T) {
	var zero uint8
	v := Uint8Value(nil)
	if v != zero {
		t.Errorf("Uint8Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualUint8(t *testing.T) {
	tests := []struct {
		v1, v2 *uint8
		eq     bool
	}{
		{Uint8(0), Uint8(1), false},
		{Uint8(0), Uint8(0), true},
		{Uint8(1), Uint8(1), true},
		{nil, nil, true},
		{nil, Uint8(0), false},
		{Uint8(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualUint8(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualUint8(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestUint8FormatterFormat(t *testing.T) {
	tests := []struct {
		p *uint8
		s string
	}{
		{Uint8(0), fmt.Sprintf("%v", 0)},
		{Uint8(1), fmt.Sprintf("%v", 1)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewUint8Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
