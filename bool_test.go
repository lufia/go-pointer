package pointer

import (
	"fmt"
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

func TestEqualBool(t *testing.T) {
	tests := []struct {
		v1, v2 *bool
		eq     bool
	}{
		{Bool(true), Bool(false), false},
		{Bool(true), Bool(true), true},
		{Bool(false), Bool(false), true},
		{nil, nil, true},
		{nil, Bool(true), false},
		{Bool(true), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualBool(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualBool(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestBoolFormatterFormat(t *testing.T) {
	tests := []struct {
		p *bool
		s string
	}{
		{Bool(true), fmt.Sprintf("%v", true)},
		{Bool(false), fmt.Sprintf("%v", false)},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewBoolFormatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
