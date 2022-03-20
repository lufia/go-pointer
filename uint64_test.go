package pointer

import (
	"testing"
)

func TestUint64(t *testing.T) {
	tests := []uint64{
		0, 1, ^uint64(0),
	}
	for _, v := range tests {
		var p *uint64 = Uint64(v)
		switch {
		case p == nil:
			t.Errorf("Uint64(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint64(%v) = %v; want %v", v, *p, v)
		}
	}
}
