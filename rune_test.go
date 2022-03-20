package pointer

import (
	"testing"
)

func TestRune(t *testing.T) {
	tests := []rune{
		0, 'a', 'ä', '本', '\t', '\007',
	}
	for _, v := range tests {
		var p *rune = Rune(v)
		switch {
		case p == nil:
			t.Errorf("Rune(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Rune(%v) = %v; want %v", v, *p, v)
		}
	}
}
