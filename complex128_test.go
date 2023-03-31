package pointer

import (
	"testing"
)

func TestComplex128(t *testing.T) {
	tests := []complex128{
		0i, 123i, 2.71828i, 6.67428 - 11i,
	}
	for _, v := range tests {
		var p *complex128 = Complex128(v)
		if p == nil {
			t.Errorf("Complex128(%v) = nil; want %v", v, v)
		}
		if n := Complex128Value(p); n != v {
			t.Errorf("Complex128(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestComplex128Slice(t *testing.T) {
	data := []complex128{
		0i, 123i, 2.71828i, 6.67428 - 11i,
	}
	a := Complex128Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Complex128Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Complex128Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestComplex128Value_Nil(t *testing.T) {
	var zero complex128
	v := Complex128Value(nil)
	if v != zero {
		t.Errorf("Complex128Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualComplex128(t *testing.T) {
	tests := []struct {
		v1, v2 *complex128
		eq     bool
	}{
		{Complex128(0i), Complex128(123i), false},
		{Complex128(0i), Complex128(0i), true},
		{Complex128(123i), Complex128(123i), true},
		{nil, nil, true},
		{nil, Complex128(0i), false},
		{Complex128(0i), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualComplex128(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualComplex128(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
