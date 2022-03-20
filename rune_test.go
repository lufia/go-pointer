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

func TestRuneSlice(t *testing.T) {
	data := []rune{
		0, 'a', 'ä', '本', '\t', '\007',
	}
	a := RuneSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("RuneSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("RuneSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
