package pointer

import (
	"testing"
)

func TestInt8(t *testing.T) {
	tests := []int8{
		-(1 << 7), 0, 1, ^int8(0) >> 1,
	}
	for _, v := range tests {
		var p *int8 = Int8(v)
		switch {
		case p == nil:
			t.Errorf("Int8(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int8(%v) = %v; want %v", v, *p, v)
		}
	}
}
