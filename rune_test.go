package pointer

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	tests := []rune{
		0, 'a', 'ä', '本', '\t', '\007',
	}
	for _, v := range tests {
		var p *rune = Rune(v)
		if p == nil {
			t.Errorf("Rune(%v) = nil; want %v", v, v)
		}
		if n := RuneValue(p); n != v {
			t.Errorf("Rune(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestRuneSlice(t *testing.T) {
	data := []rune{
		0, 'a', 'ä', '本', '\t', '\007',
	}
	a := RuneSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("RuneSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("RuneSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestRuneValue_Nil(t *testing.T) {
	var zero rune
	v := RuneValue(nil)
	if v != zero {
		t.Errorf("RuneValue(nil) = %v; want %v", v, zero)
	}
}

func TestEqualRune(t *testing.T) {
	tests := []struct {
		v1, v2 *rune
		eq     bool
	}{
		{Rune(0), Rune('a'), false},
		{Rune(0), Rune(0), true},
		{Rune('a'), Rune('a'), true},
		{nil, nil, true},
		{nil, Rune(0), false},
		{Rune(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualRune(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualRune(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestRuneFormatterFormat(t *testing.T) {
	tests := []struct {
		p *rune
		s string
	}{
		{Rune(0), fmt.Sprintf("%v", 0)},
		{Rune('a'), fmt.Sprintf("%v", 'a')},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewRuneFormatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
