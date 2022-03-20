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
		if p == nil {
			t.Errorf("Uint8(%v) = nil; want %v", v, v)
		}
		if n := Uint8Value(p); n != v {
			t.Errorf("Uint8(%v) = %v; want %v", v, n, v)
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

func TestUint8Value_Nil(t *testing.T) {
	var zero uint8
	v := Uint8Value(nil)
	if v != zero {
		t.Errorf("Uint8Value(nil) = %v; want %v", v, zero)
	}
}
