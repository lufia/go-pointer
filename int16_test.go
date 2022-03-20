package pointer

import (
	"testing"
)

func TestInt16(t *testing.T) {
	tests := []int16{
		-(1 << 15), -1, 0, 1, ^int16(0) >> 1,
	}
	for _, v := range tests {
		var p *int16 = Int16(v)
		if p == nil {
			t.Errorf("Int16(%v) = nil; want %v", v, v)
		}
		if n := Int16Value(p); n != v {
			t.Errorf("Int16(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestInt16Slice(t *testing.T) {
	data := []int16{
		-(1 << 15), -1, 0, 1, ^int16(0) >> 1,
	}
	a := Int16Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int16Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int16Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestInt16Value_Nil(t *testing.T) {
	var zero int16
	v := Int16Value(nil)
	if v != zero {
		t.Errorf("Int16Value(nil) = %v; want %v", v, zero)
	}
}
