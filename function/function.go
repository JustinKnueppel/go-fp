package function

import "github.com/JustinKnueppel/go-fp/tuple"

// Const always returns the first argument, disregarding the second.
func Const[T, U any](t T) func(U) T {
	return func(_ U) T {
		return t
	}
}

// Id is the identity function and will return its argument.
func Id[T any](t T) T {
	return t
}

// Inspect calls the given function with the provided value
// and returns the unchanged value.
func Inspect[T any](fn func(T)) func(T) T {
	return func(t T) T {
		fn(t)
		return t
	}
}

// On transforms a binary operation on type B into a binary operation on type A
// by applying a transformation from A to B before operating.
func On[A, B, C any](fn func(B) func(B) C) func(transform func(A) B) func(A) func(A) C {
	return func(transform func(A) B) func(A) func(A) C {
		return func(a1 A) func(A) C {
			return func(a2 A) C {
				return fn(transform(a1))(transform(a2))
			}
		}
	}
}

// Tupled converts a binary function into a unary function with a pair as the argument.
func Tupled[A, B, C any](fn func(A) func(B) C) func(tuple.Pair[A, B]) C {
	return func(t tuple.Pair[A, B]) C {
		return fn(tuple.Fst(t))(tuple.Snd(t))
	}
}

// Untupled converts a unary function which takes a tuple into a curried binary function.
func Untupled[A, B, C any](fn func(tuple.Pair[A, B]) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return fn(tuple.NewPair[A, B](a)(b))
		}
	}
}
