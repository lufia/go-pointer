package pointer

import (
	"fmt"
	"testing"
)

func TestComplex64(t *testing.T) {
	tests := []complex64{
		0i, 123i, 2.71828i, 6.67428 - 11i,
	}
	for _, v := range tests {
		var p *complex64 = Complex64(v)
		if p == nil {
			t.Errorf("Complex64(%v) = nil; want %v", v, v)
		}
		if n := Complex64Value(p); n != v {
			t.Errorf("Complex64(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestComplex64Slice(t *testing.T) {
	data := []complex64{
		0i, 123i, 2.71828i, 6.67428 - 11i,
	}
	a := Complex64Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Complex64Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Complex64Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestComplex64Value_Nil(t *testing.T) {
	var zero complex64
	v := Complex64Value(nil)
	if v != zero {
		t.Errorf("Complex64Value(nil) = %v; want %v", v, zero)
	}
}

func TestEqualComplex64(t *testing.T) {
	tests := []struct {
		v1, v2 *complex64
		eq     bool
	}{
		{Complex64(0i), Complex64(123i), false},
		{Complex64(0i), Complex64(0i), true},
		{Complex64(123i), Complex64(123i), true},
		{nil, nil, true},
		{nil, Complex64(0i), false},
		{Complex64(0i), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualComplex64(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualComplex64(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestComplex64FormatterFormat(t *testing.T) {
	tests := []struct {
		p *complex64
		s string
	}{
		{Complex64(0i), fmt.Sprintf("%v", 0i)},
		{Complex64(123i), fmt.Sprintf("%v", 123i)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewComplex64Formatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
