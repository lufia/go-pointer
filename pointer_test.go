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
		switch {
		case p == nil:
			t.Errorf("New(%d) = nil; want %d", v, v)
		case *p != v:
			t.Errorf("New(%d) = %d; want %d", v, *p, v)
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
