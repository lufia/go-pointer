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
		if p == nil {
			t.Errorf("Int(%v) = nil; want %v", v, v)
		}
		if n := IntValue(p); n != v {
			t.Errorf("Int(%v) = %v; want %v", v, n, v)
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

func TestIntValue_Nil(t *testing.T) {
	var zero int
	v := IntValue(nil)
	if v != zero {
		t.Errorf("IntValue(nil) = %v; want %v", v, zero)
	}
}

func TestEqualInt(t *testing.T) {
	tests := []struct {
		v1, v2 *int
		eq     bool
	}{
		{Int(-(^0 >> 1) - 1), Int(-1), false},
		{Int(-(^0 >> 1) - 1), Int(-(^0 >> 1) - 1), true},
		{Int(-1), Int(-1), true},
		{nil, nil, true},
		{nil, Int(-(^0 >> 1) - 1), false},
		{Int(-(^0 >> 1) - 1), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualInt(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualInt(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
