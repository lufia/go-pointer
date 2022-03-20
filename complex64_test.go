package pointer

import (
	"testing"
)

func TestComplex64(t *testing.T) {
	tests := []complex64{
		0i, 123i, 2.71828i, 6.67428 - 11i,
	}
	for _, v := range tests {
		var p *complex64 = Complex64(v)
		switch {
		case p == nil:
			t.Errorf("Complex64(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Complex64(%v) = %v; want %v", v, *p, v)
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
