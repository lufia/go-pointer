package pointer

import (
	"fmt"
	"testing"
)

func TestInt16(t *testing.T) {
	tests := []int16{
		-(1 << 15), -1, 0, 1, ^int16(0) >> 1,
	}
	for _, v := range tests {
		var p *int16 = Int16(v)
		if p == nil {
			t.Errorf("Int16(%v) = nil; want %v", v, v)
		}
		if n := Int16Value(p); n != v {
			t.Errorf("Int16(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestInt16Slice(t *testing.T) {
	data := []int16{
		-(1 << 15), -1, 0, 1, ^int16(0) >> 1,
	}
	a := Int16Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int16Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int16Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestInt16Value_Nil(t *testing.T) {
	var zero int16
	v := Int16Value(nil)
	if v != zero {
		t.Errorf("Int16Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualInt16(t *testing.T) {
	tests := []struct {
		v1, v2 *int16
		eq     bool
	}{
		{Int16(-(1 << 15)), Int16(-1), false},
		{Int16(-(1 << 15)), Int16(-(1 << 15)), true},
		{Int16(-1), Int16(-1), true},
		{nil, nil, true},
		{nil, Int16(-(1 << 15)), false},
		{Int16(-(1 << 15)), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualInt16(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualInt16(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestInt16FormatterFormat(t *testing.T) {
	tests := []struct {
		p *int16
		s string
	}{
		{Int16(-(1 << 15)), fmt.Sprintf("%v", -(1 << 15))},
		{Int16(-1), fmt.Sprintf("%v", -1)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewInt16Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
