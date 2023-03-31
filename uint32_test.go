package pointer

import (
	"fmt"
	"testing"
)

func TestUint32(t *testing.T) {
	tests := []uint32{
		0, 1, ^uint32(0),
	}
	for _, v := range tests {
		var p *uint32 = Uint32(v)
		if p == nil {
			t.Errorf("Uint32(%v) = nil; want %v", v, v)
		}
		if n := Uint32Value(p); n != v {
			t.Errorf("Uint32(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUint32Slice(t *testing.T) {
	data := []uint32{
		0, 1, ^uint32(0),
	}
	a := Uint32Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint32Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint32Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUint32Value_Nil(t *testing.T) {
	var zero uint32
	v := Uint32Value(nil)
	if v != zero {
		t.Errorf("Uint32Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualUint32(t *testing.T) {
	tests := []struct {
		v1, v2 *uint32
		eq     bool
	}{
		{Uint32(0), Uint32(1), false},
		{Uint32(0), Uint32(0), true},
		{Uint32(1), Uint32(1), true},
		{nil, nil, true},
		{nil, Uint32(0), false},
		{Uint32(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualUint32(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualUint32(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestUint32FormatterFormat(t *testing.T) {
	tests := []struct {
		p *uint32
		s string
	}{
		{Uint32(0), fmt.Sprintf("%v", 0)},
		{Uint32(1), fmt.Sprintf("%v", 1)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewUint32Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
