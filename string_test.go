package pointer

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []string{
		"", "hello, world!\n", "日本語",
	}
	for _, v := range tests {
		var p *string = String(v)
		switch {
		case p == nil:
			t.Errorf("String(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("String(%v) = %v; want %v", v, *p, v)
		}
	}
}
