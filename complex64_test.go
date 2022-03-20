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
