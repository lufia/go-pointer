package pointer

import (
	"testing"
)

func TestNew_int(t *testing.T) {
	tests := []int{
		0, 1, -1, 2,
	}
	for _, v := range tests {
		var p *int = New(v)
		switch {
		case p == nil:
			t.Errorf("New(%d) = nil; want %d", v, v)
		case *p != v:
			t.Errorf("New(%d) = %d; want %d", v, *p, v)
		}
	}
}
