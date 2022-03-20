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
		if p == nil {
			t.Errorf("Uint(%v) = nil; want %v", v, v)
		}
		if n := UintValue(p); n != v {
			t.Errorf("Uint(%v) = %v; want %v", v, n, v)
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

func TestUintValue_Nil(t *testing.T) {
	var zero uint
	v := UintValue(nil)
	if v != zero {
		t.Errorf("UintValue(nil) = %v; want %v", v, zero)
	}
}
