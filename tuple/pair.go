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
func NewPair[T, U any](t T) func(u U) Pair[T, U] {
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

// MapLeft returns a new Pair with the function applied to the left element.
func MapLeft[T1, U, T2 any](fn func(T1) T2) func(p Pair[T1, U]) Pair[T2, U] {
	return func(p Pair[T1, U]) Pair[T2, U] {
		return NewPair[T2, U](fn(p.fst))(p.snd)
	}
}

// MapRight returns a new Pair with the function applied to the right element.
func MapRight[T, U1, U2 any](fn func(U1) U2) func(p Pair[T, U1]) Pair[T, U2] {
	return func(p Pair[T, U1]) Pair[T, U2] {
		return NewPair[T, U2](p.fst)(fn(p.snd))
	}
}

// Pattern returns each value of the pair using Go multiple return.
// This allows pattern matching like syntax with x, y := Pattern(pair).
func Pattern[T, U any](pair Pair[T, U]) (T, U) {
	return pair.fst, pair.snd
}

/* ========= Functor definitions ========= */

// Fmap applies the given function to the first element of the pair.
func Fmap[T0, U, T any](fn func(T0) T) func(p Pair[T0, U]) Pair[T, U] {
	return MapLeft[T0, U](fn)
}

// ConstMap replaces the first element with the given value.
func ConstMap[T0, U, T any](value T) func(p Pair[T0, U]) Pair[T, U] {
	return func(p Pair[T0, U]) Pair[T, U] {
		return NewPair[T, U](value)(Snd(p))
	}
}

// FmapRight applies the given function to the second element of the pair.
func FmapRight[T, U0, U any](fn func(U0) U) func(Pair[T, U0]) Pair[T, U] {
	return MapRight[T](fn)
}

// ConstMapRight replaces the second element with the given value.
func ConstMapRight[T, U0, U any](value U) func(Pair[T, U0]) Pair[T, U] {
	return func(p Pair[T, U0]) Pair[T, U] {
		return NewPair[T, U](Fst(p))(value)
	}
}
