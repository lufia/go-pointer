package pointer

import (
	"testing"
)

func TestInt64(t *testing.T) {
	tests := []int64{
		-(1 << 63), -1, 0, 1, ^int64(0) >> 1,
	}
	for _, v := range tests {
		var p *int64 = Int64(v)
		if p == nil {
			t.Errorf("Int64(%v) = nil; want %v", v, v)
		}
		if n := Int64Value(p); n != v {
			t.Errorf("Int64(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestInt64Slice(t *testing.T) {
	data := []int64{
		-(1 << 63), -1, 0, 1, ^int64(0) >> 1,
	}
	a := Int64Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int64Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int64Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestInt64Value_Nil(t *testing.T) {
	var zero int64
	v := Int64Value(nil)
	if v != zero {
		t.Errorf("Int64Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualInt64(t *testing.T) {
	tests := []struct {
		v1, v2 *int64
		eq     bool
	}{
		{Int64(-(1 << 63)), Int64(-1), false},
		{Int64(-(1 << 63)), Int64(-(1 << 63)), true},
		{Int64(-1), Int64(-1), true},
		{nil, nil, true},
		{nil, Int64(-(1 << 63)), false},
		{Int64(-(1 << 63)), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualInt64(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualInt64(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
