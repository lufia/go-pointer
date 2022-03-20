package pointer

import (
	"testing"
)

func TestInt8(t *testing.T) {
	tests := []int8{
		-(1 << 7), 0, 1, ^int8(0) >> 1,
	}
	for _, v := range tests {
		var p *int8 = Int8(v)
		switch {
		case p == nil:
			t.Errorf("Int8(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int8(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestInt8Slice(t *testing.T) {
	data := []int8{
		-(1 << 7), 0, 1, ^int8(0) >> 1,
	}
	a := Int8Slice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("Int8Slice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("Int8Slice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
