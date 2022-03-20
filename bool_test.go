package pointer

import (
	"testing"
)

func TestBool(t *testing.T) {
	tests := []bool{
		true, false,
	}
	for _, v := range tests {
		var p *bool = Bool(v)
		if p == nil {
			t.Errorf("Bool(%v) = nil; want %v", v, v)
		}
		if n := BoolValue(p); n != v {
			t.Errorf("Bool(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestBoolSlice(t *testing.T) {
	data := []bool{
		true, false,
	}
	a := BoolSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("BoolSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("BoolSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestBoolValue_Nil(t *testing.T) {
	var zero bool
	v := BoolValue(nil)
	if v != zero {
		t.Errorf("BoolValue(nil) = %v; want %v", v, zero)
	}
}
