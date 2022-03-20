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
		switch {
		case p == nil:
			t.Errorf("Int32(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int32(%v) = %v; want %v", v, *p, v)
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
