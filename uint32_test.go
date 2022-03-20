package pointer

import (
	"testing"
)

func TestUint32(t *testing.T) {
	tests := []uint32{
		0, 1, ^uint32(0),
	}
	for _, v := range tests {
		var p *uint32 = Uint32(v)
		switch {
		case p == nil:
			t.Errorf("Uint32(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint32(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestUint32Slice(t *testing.T) {
	data := []uint32{
		0, 1, ^uint32(0),
	}
	a := Uint32Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint32Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint32Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
