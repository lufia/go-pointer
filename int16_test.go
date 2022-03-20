package pointer

import (
	"testing"
)

func TestInt16(t *testing.T) {
	tests := []int16{
		-(1 << 15), -1, 0, 1, ^int16(0) >> 1,
	}
	for _, v := range tests {
		var p *int16 = Int16(v)
		switch {
		case p == nil:
			t.Errorf("Int16(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int16(%v) = %v; want %v", v, *p, v)
		}
	}
}
