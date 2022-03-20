package pointer

import (
	"testing"
)

func TestUint(t *testing.T) {
	tests := []uint{
		0, 1, ^uint(0),
	}
	for _, v := range tests {
		var p *uint = Uint(v)
		switch {
		case p == nil:
			t.Errorf("Uint(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint(%v) = %v; want %v", v, *p, v)
		}
	}
}
