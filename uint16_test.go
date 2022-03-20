package pointer

import (
	"testing"
)

func TestUint16(t *testing.T) {
	tests := []uint16{
		0, 1, ^uint16(0),
	}
	for _, v := range tests {
		var p *uint16 = Uint16(v)
		switch {
		case p == nil:
			t.Errorf("Uint16(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint16(%v) = %v; want %v", v, *p, v)
		}
	}
}
