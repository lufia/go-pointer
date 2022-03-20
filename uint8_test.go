package pointer

import (
	"testing"
)

func TestUint8(t *testing.T) {
	tests := []uint8{
		0, 1, ^uint8(0),
	}
	for _, v := range tests {
		var p *uint8 = Uint8(v)
		switch {
		case p == nil:
			t.Errorf("Uint8(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint8(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestUint8Slice(t *testing.T) {
	data := []uint8{
		0, 1, ^uint8(0),
	}
	a := Uint8Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint8Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint8Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
