package pointer

import (
	"testing"
)

func TestByte(t *testing.T) {
	tests := []byte{
		0, 'a', '\007', '\t',
	}
	for _, v := range tests {
		var p *byte = Byte(v)
		switch {
		case p == nil:
			t.Errorf("Byte(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Byte(%v) = %v; want %v", v, *p, v)
		}
	}
}
