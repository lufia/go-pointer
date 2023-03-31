package pointer

import "fmt"

// Formatter implements fmt.Formatter of a pointer to T.
type Formatter[T any] struct {
	p *T
}

var _ fmt.Formatter = Formatter[int]{}

// NewFormatter returns the formatter of a pointer to T.
func NewFormatter[T any](p *T) Formatter[T] {
	return Formatter[T]{p}
}

// Format implements the fmt.Formatter interface.
func (p Formatter[T]) Format(f fmt.State, c rune) {
	if p.v == nil {
		fmt.Fprintf(f, "<nil>")
		return
	}
	fmt.Fprintf(f, "%"+string(c), *p.v)
}
