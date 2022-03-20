package pointer

import (
	"testing"
)

func TestInt64(t *testing.T) {
	tests := []int64{
		-(1 << 63), -1, 0, 1, ^int64(0) >> 1,
	}
	for _, v := range tests {
		var p *int64 = Int64(v)
		switch {
		case p == nil:
			t.Errorf("Int64(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int64(%v) = %v; want %v", v, *p, v)
		}
	}
}
