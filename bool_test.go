package pointer

import (
	"testing"
)

func TestBool(t *testing.T) {
	tests := []bool{
		true, false,
	}
	for _, v := range tests {
		var p *bool = Bool(v)
		switch {
		case p == nil:
			t.Errorf("Bool(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Bool(%v) = %v; want %v", v, *p, v)
		}
	}
}
