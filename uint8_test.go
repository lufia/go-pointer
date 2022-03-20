package pointer

import (
	"testing"
)

func TestUint8(t *testing.T) {
	tests := []uint8{
		0, 1, ^uint8(0),
	}
	for _, v := range tests {
		var p *uint8 = Uint8(v)
		switch {
		case p == nil:
			t.Errorf("Uint8(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint8(%v) = %v; want %v", v, *p, v)
		}
	}
}
