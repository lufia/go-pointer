package pointer

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []int{
		0, 1, -1, 2,
	}
	for _, v := range tests {
		var p *int = Int(v)
		switch {
		case p == nil:
			t.Errorf("Int(%d) = nil; want %d", v, v)
		case *p != v:
			t.Errorf("Int(%d) = %d; want %d", v, *p, v)
		}
	}
}
