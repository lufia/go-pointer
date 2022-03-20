package pointer

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []int{
		-(^0 >> 1) - 1, -1, 0, 1, (^0) >> 1,
	}
	for _, v := range tests {
		var p *int = Int(v)
		switch {
		case p == nil:
			t.Errorf("Int(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int(%v) = %v; want %v", v, *p, v)
		}
	}
}
