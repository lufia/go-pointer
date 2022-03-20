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

func TestStringSlice(t *testing.T) {
	data := []string{
		"", "hello, world!\n", "日本語",
	}
	a := StringSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("StringSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("StringSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
