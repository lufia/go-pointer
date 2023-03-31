package pointer

import (
	"fmt"
	"testing"
)

func TestInt32(t *testing.T) {
	tests := []int32{
		-(1 << 31), -1, 0, 1, ^int32(0) >> 1,
	}
	for _, v := range tests {
		var p *int32 = Int32(v)
		if p == nil {
			t.Errorf("Int32(%v) = nil; want %v", v, v)
		}
		if n := Int32Value(p); n != v {
			t.Errorf("Int32(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestInt32Slice(t *testing.T) {
	data := []int32{
		-(1 << 31), -1, 0, 1, ^int32(0) >> 1,
	}
	a := Int32Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int32Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int32Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestInt32Value_Nil(t *testing.T) {
	var zero int32
	v := Int32Value(nil)
	if v != zero {
		t.Errorf("Int32Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualInt32(t *testing.T) {
	tests := []struct {
		v1, v2 *int32
		eq     bool
	}{
		{Int32(-(1 << 31)), Int32(-1), false},
		{Int32(-(1 << 31)), Int32(-(1 << 31)), true},
		{Int32(-1), Int32(-1), true},
		{nil, nil, true},
		{nil, Int32(-(1 << 31)), false},
		{Int32(-(1 << 31)), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualInt32(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualInt32(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestInt32FormatterFormat(t *testing.T) {
	tests := []struct {
		p *int32
		s string
	}{
		{Int32(-(1 << 31)), fmt.Sprintf("%v", -(1 << 31))},
		{Int32(-1), fmt.Sprintf("%v", -1)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewInt32Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
