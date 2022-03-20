package pointer

import (
	"testing"
)

func TestUint(t *testing.T) {
	tests := []uint{
		0, 1, ^uint(0),
	}
	for _, v := range tests {
		var p *uint = Uint(v)
		switch {
		case p == nil:
			t.Errorf("Uint(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestUintSlice(t *testing.T) {
	data := []uint{
		0, 1, ^uint(0),
	}
	a := UintSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("UintSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("UintSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
