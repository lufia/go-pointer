package pointer

import (
	"testing"
)

func TestUint16(t *testing.T) {
	tests := []uint16{
		0, 1, ^uint16(0),
	}
	for _, v := range tests {
		var p *uint16 = Uint16(v)
		if p == nil {
			t.Errorf("Uint16(%v) = nil; want %v", v, v)
		}
		if n := Uint16Value(p); n != v {
			t.Errorf("Uint16(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestUint16Slice(t *testing.T) {
	data := []uint16{
		0, 1, ^uint16(0),
	}
	a := Uint16Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Uint16Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Uint16Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestUint16Value_Nil(t *testing.T) {
	var zero uint16
	v := Uint16Value(nil)
	if v != zero {
		t.Errorf("Uint16Value(nil) = %v; want %v", v, zero)
	}
}
