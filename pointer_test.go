//go:build go1.18
// +build go1.18

package pointer

import (
	"reflect"
	"testing"
)

func TestNew_int(t *testing.T) {
	tests := []int{
		0, 1, -1, 2,
	}
	for _, v := range tests {
		var p *int = New(v)
		if p == nil {
			t.Errorf("New(%d) = nil; want %d", v, v)
		}
		if n := Value(p); n != v {
			t.Errorf("New(%d) = %d; want %d", v, n, v)
		}
	}
}

func TestSlice_int(t *testing.T) {
	tests := [][]int{
		{},
		{0, 1, 2},
	}
	for _, tt := range tests {
		var a []*int = Slice(tt...)
		x := deref(a)
		if !reflect.DeepEqual(x, tt) {
			t.Errorf("deref(Slice(%v)) = %v; want %v", tt, x, tt)
		}
	}
}

func deref[T any](a []*T) []T {
	p := make([]T, len(a))
	for i, v := range a {
		p[i] = *v
	}
	return p
}

func TestValue_Nil_int(t *testing.T) {
	v := Value[int](nil)
	if v != 0 {
		t.Errorf("Value[int](nil) = %v; want 0", v)
	}
}

func TestEqual_int(t *testing.T) {
	tests := []struct {
		v1, v2 *int
		eq     bool
	}{
		{New(0), New(1), false},
		{New(0), New(0), true},
		{New(-1), New(-1), true},
		{nil, nil, true},
		{nil, New(0), false},
		{New(0), nil, false},
	}
	for _, tt := range tests {
		if eq := Equal(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("Equal(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}
