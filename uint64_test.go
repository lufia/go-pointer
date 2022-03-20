package pointer

import (
	"testing"
)

func TestUint64(t *testing.T) {
	tests := []uint64{
		0, 1, ^uint64(0),
	}
	for _, v := range tests {
		var p *uint64 = Uint64(v)
		switch {
		case p == nil:
			t.Errorf("Uint64(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Uint64(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestUint64Slice(t *testing.T) {
	data := []uint64{
		0, 1, ^uint64(0),
	}
	a := Uint64Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint64Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint64Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
