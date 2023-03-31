package pointer

import (
	"fmt"
	"testing"
)

func TestFormat_int(t *testing.T) {
	tests := []struct {
		p *int
		s string
	}{
		{Int(0), "0"},
		{Int(-1), "-1"},
		{nil, "<nil>"},
	}
	for _, tt := range tests {
		p := NewFormatter(tt.p)
		s := fmt.Sprintf("%v", p)
		if s != tt.s {
			t.Errorf("{%+v}.Format() = %q; want %q", p, s, tt.s)
		}
	}
}
