package pointer

import (
	"testing"
)

func TestInt32(t *testing.T) {
	tests := []int32{
		-(1 << 31), -1, 0, 1, ^int32(0) >> 1,
	}
	for _, v := range tests {
		var p *int32 = Int32(v)
		switch {
		case p == nil:
			t.Errorf("Int32(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int32(%v) = %v; want %v", v, *p, v)
		}
	}
}
