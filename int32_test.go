package pointer

import (
	"testing"
)

func TestInt32(t *testing.T) {
	tests := []int32{
		-(1 << 31), -1, 0, 1, ^int32(0) >> 1,
	}
	for _, v := range tests {
		var p *int32 = Int32(v)
		if p == nil {
			t.Errorf("Int32(%v) = nil; want %v", v, v)
		}
		if n := Int32Value(p); n != v {
			t.Errorf("Int32(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestInt32Slice(t *testing.T) {
	data := []int32{
		-(1 << 31), -1, 0, 1, ^int32(0) >> 1,
	}
	a := Int32Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int32Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int32Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestInt32Value_Nil(t *testing.T) {
	var zero int32
	v := Int32Value(nil)
	if v != zero {
		t.Errorf("Int32Value(nil) = %v; want %v", v, zero)
	}
}
