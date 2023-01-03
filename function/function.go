package function

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
func On[A, B, C any](fn func(B) func(B) C) func(func(A) B) func(A) func(A) C {
	return func(transform func(A) B) func(A) func(A) C {
		return func(a1 A) func(A) C {
			return func(a2 A) C {
				return fn(transform(a1))(transform(a2))
			}
		}
	}
}
