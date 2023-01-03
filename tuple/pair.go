package tuple

import "fmt"

type Pair[T, U any] struct {
	fst T
	snd U
}

// String is used only for properly displaying pairs.
func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v %v)", Fst(p), Snd(p))
}

// NewPair creates a Pair from the two arguments.
func NewPair[T, U any](t T) func(U) Pair[T, U] {
	return func(u U) Pair[T, U] {
		return Pair[T, U]{t, u}
	}
}

// Fst returns the first element of the pair.
func Fst[T, U any](p Pair[T, U]) T {
	return p.fst
}

// Snd returns the first element of the pair.
func Snd[T, U any](p Pair[T, U]) U {
	return p.snd
}
