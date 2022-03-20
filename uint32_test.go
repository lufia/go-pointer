package pointer

import (
	"testing"
)

func TestUint32(t *testing.T) {
	tests := []uint32{
		0, 1, ^uint32(0),
	}
	for _, v := range tests {
		var p *uint32 = Uint32(v)
		switch {
		case p == nil:
			t.Errorf("Uint32(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint32(%v) = %v; want %v", v, *p, v)
		}
	}
}
