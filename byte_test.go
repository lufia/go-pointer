package pointer

import (
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	tests := []byte{
		0, 'a', '\007', '\t',
	}
	for _, v := range tests {
		var p *byte = Byte(v)
		if p == nil {
			t.Errorf("Byte(%v) = nil; want %v", v, v)
		}
		if n := ByteValue(p); n != v {
			t.Errorf("Byte(%v) = %v; want %v", v, n, v)
		}
	}
}

func TestByteSlice(t *testing.T) {
	data := []byte{
		0, 'a', '\007', '\t',
	}
	a := ByteSlice(data...)
	if len(a) != len(data) {
		t.Errorf("len(slice) = %d; want %d", len(a), len(data))
		return
	}
	for i, p := range a {
		v := data[i]
		switch {
		case p == nil:
			t.Errorf("ByteSlice(%v)[%d] = nil; want %v", data, i, v)
		case *p != v:
			t.Errorf("ByteSlice(%v)[%d] = %v; want %v", data, i, *p, v)
		}
	}
}

func TestByteValue_Nil(t *testing.T) {
	var zero byte
	v := ByteValue(nil)
	if v != zero {
		t.Errorf("ByteValue(nil) = %v; want %v", v, zero)
	}
}

func TestEqualByte(t *testing.T) {
	tests := []struct {
		v1, v2 *byte
		eq     bool
	}{
		{Byte(0), Byte('a'), false},
		{Byte(0), Byte(0), true},
		{Byte('a'), Byte('a'), true},
		{nil, nil, true},
		{nil, Byte(0), false},
		{Byte(0), nil, false},
	}
	for _, tt := range tests {
		if eq := EqualByte(tt.v1, tt.v2); eq != tt.eq {
			t.Errorf("EqualByte(%d, %d) = %t; want %t", tt.v1, tt.v2, eq, tt.eq)
		}
	}
}

func TestByteFormatterFormat(t *testing.T) {
	tests := []struct {
		p *byte
		s string
	}{
		{Byte(0), fmt.Sprintf("%v", 0)},
		{Byte('a'), fmt.Sprintf("%v", 'a')},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewByteFormatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
