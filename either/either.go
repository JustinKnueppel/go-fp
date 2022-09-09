package either

// Either represents a disjoint union of the left and right types.
type Either[L, R any] struct {
	left    L
	right   R
	isRight bool
}

// Left returns an Either containing a left value.
func Left[L, R any](val L) Either[L, R] {
	var r R
	return Either[L, R]{
		left:    val,
		right:   r,
		isRight: false,
	}
}

// Right returns an Either containing a right value.
func Right[L, R any](val R) Either[L, R] {
	var l L
	return Either[L, R]{
		left:    l,
		right:   val,
		isRight: true,
	}
}

// IsLeft returns true if the Either is an instance of Left.
func IsLeft[L, R any](e Either[L, R]) bool {
	return !e.isRight
}

// IsLeftAnd returns true if the Either is an instance of Left
// and its value satisfies the predicate.
func IsLeftAnd[L, R any](predicate func(L) bool) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		return IsLeft(e) && predicate(e.left)
	}
}

// IsRight returns true if the Either is an instance of Right.
func IsRight[L, R any](e Either[L, R]) bool {
	return e.isRight
}

// IsRightAnd returns true if the Either is an instance of Right
// and its value satisfies the predicate.
func IsRightAnd[L, R any](predicate func(R) bool) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		return IsRight(e) && predicate(e.right)
	}
}

// Map returns a Right with a function applied to the right value (if Right)
// or forwards the left value (if Left).
func Map[L, R1, R2 any](fn func(R1) R2) func(Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsLeft(e) {
			return Left[L, R2](e.left)
		}
		return Right[L](fn(e.right))
	}
}

// Bind returns the result of the function applied to the right value (if Right),
// or forwards the left value (if Left).
func Bind[L, R1, R2 any](fn func(R1) Either[L, R2]) func(Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsLeft(e) {
			return Left[L, R2](e.left)
		}
		return fn(e.right)
	}
}

// MapOr returns the value of applying a function to the right value (if Right)
// or returns the provided default (if Left).
func MapOr[L, R1, R2 any](fallback R2) func(func(R1) R2) func(Either[L, R1]) R2 {
	return func(fn func(R1) R2) func(Either[L, R1]) R2 {
		return func(e Either[L, R1]) R2 {
			if IsLeft(e) {
				return fallback
			}
			return fn(e.right)
		}
	}
}

// MapOrElse returns the value of applying a function to the right value (if Right)
// or returns the computed default (if Left).
func MapOrElse[L, R1, R2 any](fallbackFn func() R2) func(func(R1) R2) func(Either[L, R1]) R2 {
	return func(fn func(R1) R2) func(Either[L, R1]) R2 {
		return func(e Either[L, R1]) R2 {
			if IsLeft(e) {
				return fallbackFn()
			}
			return fn(e.right)
		}
	}
}

// MapLeft returns a Left with a function applied to the left value (if Left)
// or forwards the right value (if Right).
func MapLeft[L1, L2, R any](fn func(L1) L2) func(Either[L1, R]) Either[L2, R] {
	return func(e Either[L1, R]) Either[L2, R] {
		if IsRight(e) {
			return Right[L2](e.right)
		}
		return Left[L2, R](fn(e.left))
	}
}

// Inspect calls the provided closure with the contained right
// value (if Right) and returns the unchanged Either.
func Inspect[L, R any](fn func(R)) func(Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			fn(e.right)
		}
		return e
	}
}

// InspectLeft calls the provided closure with the contained right
// value (if Right) and returns the unchanged Either.
func InspectLeft[L, R any](fn func(L)) func(Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsLeft(e) {
			fn(e.left)
		}
		return e
	}
}

// Expect returns the right value (if Right), or panics
// with the given message (if Left).
func Expect[L, R any](msg string) func(Either[L, R]) R {
	return func(e Either[L, R]) R {
		if IsLeft(e) {
			panic(msg)
		}
		return e.right
	}
}

// Unwrap returns the right value (if Right), or panics.
func Unwrap[L, R any](e Either[L, R]) R {
	if IsLeft(e) {
		panic("Unwrap called on Left")
	}
	return e.right
}

// UnwrapOr returns the right value (if Right), or
// the provided default (if Left).
func UnwrapOr[L, R any](fallback R) func(Either[L, R]) R {
	return func(e Either[L, R]) R {
		if IsLeft(e) {
			return fallback
		}
		return e.right
	}
}

// UnwrapOrElse returns the right value (if Right), or
// the computed default (if Left).
func UnwrapOrElse[L, R any](fallbackFn func() R) func(Either[L, R]) R {
	return func(e Either[L, R]) R {
		if IsLeft(e) {
			return fallbackFn()
		}
		return e.right
	}
}

// UnwrapOrDefault returns the right value (if Right), or
// the zero value of the right type (if Left).
func UnwrapOrDefault[L, R any](e Either[L, R]) R {
	if IsLeft(e) {
		var r R
		return r
	}
	return e.right
}

// ExpectLeft returns the left value (if Left), or panics
// with the given message (if Right).
func ExpectLeft[L, R any](msg string) func(Either[L, R]) L {
	return func(e Either[L, R]) L {
		if IsRight(e) {
			panic(msg)
		}
		return e.left
	}
}

// UnwrapLeft returns the left value (if Left), or panics.
func UnwrapLeft[L, R any](e Either[L, R]) L {
	if IsRight(e) {
		panic("UnwrapLeft called on Right")
	}
	return e.left
}

// And returns other (if Right) or forwards the left value (if Left).
func And[L, R1, R2 any](other Either[L, R2]) func(Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsRight(e) {
			return other
		}
		return Left[L, R2](e.left)
	}
}

// Or returns the given Either (if Right), or other (if Left).
func Or[L, R any](other Either[L, R]) func(Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			return e
		}
		return other
	}
}

// OrElse returns the given Either (if Right), or computes the
// return value from the closure.
func OrElse[L, R any](fn func(L) Either[L, R]) func(Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			return e
		}
		return fn(e.left)
	}
}

// Contains returns true if and only if the Either is
// a Right and its value is equivalent to the target.
func Contains[L any, R comparable](target R) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		if IsLeft(e) {
			return false
		}
		return e.right == target
	}
}

// ContainsLeft returns true if and only if the Either is
// a Left and its value is equivalent to the target.
func ContainsLeft[L comparable, R any](target L) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		if IsRight(e) {
			return false
		}
		return e.left == target
	}
}

// Copy returns a value copy of the given Either.
func Copy[L, R any](e Either[L, R]) Either[L, R] {
	return e
}

// Flatten returns the nested Either (if Right), or
// forwards the left value (if Left).
func Flatten[L, R any](e Either[L, Either[L, R]]) Either[L, R] {
	if IsLeft(e) {
		return Left[L, R](e.left)
	}
	return e.right
}

// Equal tests deep equality of the two Eithers.
func Equal[L, R comparable](other Either[L, R]) func(Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		if IsRight(e) && IsRight(other) {
			return e.right == other.right
		}
		if IsLeft(e) && IsLeft(other) {
			return e.left == other.left
		}
		return false
	}
}
