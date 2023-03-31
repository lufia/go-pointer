package pointer

import (
	"fmt"
	"testing"
)

func TestUint(t *testing.T) {
	tests := []uint{
		0, 1, ^uint(0),
	}
	for _, v := range tests {
		var p *uint = Uint(v)
		if p == nil {
			t.Errorf("Uint(%v) = nil; want %v", v, v)
		}
		if n := UintValue(p); n != v {
			t.Errorf("Uint(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUintSlice(t *testing.T) {
	data := []uint{
		0, 1, ^uint(0),
	}
	a := UintSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("UintSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("UintSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUintValue_Nil(t *testing.T) {
	var zero uint
	v := UintValue(nil)
	if v != zero {
		t.Errorf("UintValue(nil) = %v; want %v", v, zero)
	}
}

func TestEqualUint(t *testing.T) {
	tests := []struct {
		v1, v2 *uint
		eq     bool
	}{
		{Uint(0), Uint(1), false},
		{Uint(0), Uint(0), true},
		{Uint(1), Uint(1), true},
		{nil, nil, true},
		{nil, Uint(0), false},
		{Uint(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualUint(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualUint(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestUintFormatterFormat(t *testing.T) {
	tests := []struct {
		p *uint
		s string
	}{
		{Uint(0), fmt.Sprintf("%v", 0)},
		{Uint(1), fmt.Sprintf("%v", 1)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewUintFormatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
