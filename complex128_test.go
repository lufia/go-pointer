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
		switch {
		case p == nil:
			t.Errorf("Complex128(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Complex128(%v) = %v; want %v", v, *p, v)
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
