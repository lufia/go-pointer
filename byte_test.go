package pointer

import (
	"testing"
)

func TestByte(t *testing.T) {
	tests := []byte{
		0, 'a', '\007', '\t',
	}
	for _, v := range tests {
		var p *byte = Byte(v)
		switch {
		case p == nil:
			t.Errorf("Byte(%v) = nil; want %v", v, v)
		case *p != v:
			t.Errorf("Byte(%v) = %v; want %v", v, *p, v)
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
