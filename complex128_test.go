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
