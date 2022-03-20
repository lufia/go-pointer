package pointer

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []int{
		-(^0 >> 1) - 1, -1, 0, 1, (^0) >> 1,
	}
	for _, v := range tests {
		var p *int = Int(v)
		switch {
		case p == nil:
			t.Errorf("Int(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Int(%v) = %v; want %v", v, *p, v)
		}
	}
}

func TestIntSlice(t *testing.T) {
	data := []int{
		-(^0 >> 1) - 1, -1, 0, 1, (^0) >> 1,
	}
	a := IntSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("IntSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("IntSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}
