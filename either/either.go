package either

import "fmt"

/* ============ Either type ============ */

// Either represents a disjoint union of the left and right types.
type Either[L, R any] struct {
	left    L
	right   R
	isRight bool
}

// String is used only to properly print an Either.
func (e Either[L, R]) String() string {
	if e.isRight {
		return fmt.Sprintf("Right %v", e.right)
	}
	return fmt.Sprintf("Left %v", e.left)
}

/* ============ Constructors ============ */

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

/* ============ Basic functions ============ */

// Copy returns a value copy of the given Either.
func Copy[L, R any](e Either[L, R]) Either[L, R] {
	return e
}

// Equal tests deep equality of the two Eithers.
func Equal[L, R comparable](other Either[L, R]) func(e Either[L, R]) bool {
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

// Converge (Haskell either) converts an Either to another type by mapping either a Left or Right.
func Converge[L, R, T any](lFn func(L) T) func(rFn func(R) T) func(e Either[L, R]) T {
	return func(rFn func(R) T) func(Either[L, R]) T {
		return func(e Either[L, R]) T {
			if IsLeft(e) {
				return lFn(UnwrapLeft(e))
			}
			return rFn(Unwrap(e))
		}
	}
}

// Lefts extracts from a slice of Either all the Left elements. All the Left elements are extracted in order.
func Lefts[L, R any](es []Either[L, R]) []L {
	out := []L{}
	for _, e := range es {
		if IsLeft(e) {
			out = append(out, UnwrapLeft(e))
		}
	}
	return out
}

// Rights extracts from a slice of Either all the Right elements. All the Right elements are extracted in order.
func Rights[L, R any](es []Either[L, R]) []R {
	out := []R{}
	for _, e := range es {
		if IsRight(e) {
			out = append(out, Unwrap(e))
		}
	}
	return out
}

// IsLeft returns true if the Either is an instance of Left.
func IsLeft[L, R any](e Either[L, R]) bool {
	return !e.isRight
}

// IsLeftAnd returns true if the Either is an instance of Left
// and its value satisfies the predicate.
func IsLeftAnd[L, R any](predicate func(L) bool) func(e Either[L, R]) bool {
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
func IsRightAnd[L, R any](predicate func(R) bool) func(e Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		return IsRight(e) && predicate(e.right)
	}
}

/* ============ Maps ============ */

// Map returns a Right with a function applied to the right value (if Right)
// or forwards the left value (if Left).
func Map[L, R1, R2 any](fn func(R1) R2) func(e Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsLeft(e) {
			return Left[L, R2](e.left)
		}
		return Right[L](fn(e.right))
	}
}

// Bind returns the result of the function applied to the right value (if Right),
// or forwards the left value (if Left).
func Bind[L, R1, R2 any](fn func(R1) Either[L, R2]) func(e Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsLeft(e) {
			return Left[L, R2](e.left)
		}
		return fn(e.right)
	}
}

// Flatten returns the nested Either (if Right), or
// forwards the left value (if Left).
func Flatten[L, R any](e Either[L, Either[L, R]]) Either[L, R] {
	if IsLeft(e) {
		return Left[L, R](e.left)
	}
	return e.right
}

// MapOr returns the value of applying a function to the right value (if Right)
// or returns the provided default (if Left).
func MapOr[L, R1, R2 any](fallback R2) func(fn func(R1) R2) func(e Either[L, R1]) R2 {
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
func MapOrElse[L, R1, R2 any](fallbackFn func() R2) func(fn func(R1) R2) func(e Either[L, R1]) R2 {
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
func MapLeft[L1, L2, R any](fn func(L1) L2) func(e Either[L1, R]) Either[L2, R] {
	return func(e Either[L1, R]) Either[L2, R] {
		if IsRight(e) {
			return Right[L2](e.right)
		}
		return Left[L2, R](fn(e.left))
	}
}

// And returns other (if Right) or forwards the left value (if Left).
func And[L, R1, R2 any](other Either[L, R2]) func(e Either[L, R1]) Either[L, R2] {
	return func(e Either[L, R1]) Either[L, R2] {
		if IsRight(e) {
			return other
		}
		return Left[L, R2](e.left)
	}
}

// Or returns the given Either (if Right), or other (if Left).
func Or[L, R any](other Either[L, R]) func(e Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			return e
		}
		return other
	}
}

// OrElse returns the given Either (if Right), or computes the
// return value from the closure.
func OrElse[L, R any](fallbackFn func(L) Either[L, R]) func(e Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			return e
		}
		return fallbackFn(e.left)
	}
}

/* ============ Debugging ============ */

// Inspect calls the provided closure with the contained right
// value (if Right) and returns the unchanged Either.
func Inspect[L, R any](fn func(R)) func(e Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsRight(e) {
			fn(e.right)
		}
		return e
	}
}

// InspectLeft calls the provided closure with the contained right
// value (if Right) and returns the unchanged Either.
func InspectLeft[L, R any](fn func(L)) func(e Either[L, R]) Either[L, R] {
	return func(e Either[L, R]) Either[L, R] {
		if IsLeft(e) {
			fn(e.left)
		}
		return e
	}
}

/* ============ Value extraction ============ */

// FromLeft returns the contents of a Left-value or a default value otherwise.
func FromLeft[L, R any](fallback L) func(e Either[L, R]) L {
	return UnwrapLeftOr[L, R](fallback)
}

// FromRight returns the contents of a Right-value or a default value otherwise.
func FromRight[L, R any](fallback R) func(e Either[L, R]) R {
	return UnwrapOr[L](fallback)
}

// Expect returns the right value (if Right), or panics
// with the given message (if Left).
func Expect[L, R any](msg string) func(e Either[L, R]) R {
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
func UnwrapOr[L, R any](fallback R) func(e Either[L, R]) R {
	return func(e Either[L, R]) R {
		if IsLeft(e) {
			return fallback
		}
		return e.right
	}
}

// UnwrapOrElse returns the right value (if Right), or
// the computed default (if Left).
func UnwrapOrElse[L, R any](fallbackFn func() R) func(e Either[L, R]) R {
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
func ExpectLeft[L, R any](msg string) func(e Either[L, R]) L {
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

// UnwrapLeftOr returns the left value (if Left), or the fallback otherwise.
func UnwrapLeftOr[L, R any](fallback L) func(e Either[L, R]) L {
	return func(e Either[L, R]) L {
		if IsLeft(e) {
			return UnwrapLeft(e)
		}
		return fallback
	}
}

/* ============ Query ============ */

// Contains returns true if and only if the Either is
// a Right and its value is equivalent to the target.
func Contains[L any, R comparable](target R) func(e Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		if IsLeft(e) {
			return false
		}
		return e.right == target
	}
}

// ContainsLeft returns true if and only if the Either is
// a Left and its value is equivalent to the target.
func ContainsLeft[L comparable, R any](target L) func(e Either[L, R]) bool {
	return func(e Either[L, R]) bool {
		if IsRight(e) {
			return false
		}
		return e.left == target
	}
}

/* ========= Functor definitions ========== */

// Fmap transforms the Right value if Right using the given function,
// otherwise forwards the Left value.
func Fmap[L, R0, R any](fn func(R0) R) func(e Either[L, R0]) Either[L, R] {
	return Map[L](fn)
}

// ConstMap replaces the Right value if Right, otherwise forwards the Left value.
func ConstMap[L, R0, R any](value R) func(e Either[L, R0]) Either[L, R] {
	return And[L, R0](Right[L](value))
}

// FmapLeft transforms the Left value if Left using the given function,
// otherwise forwards the Right value.
func FmapLeft[L0, R, L any](fn func(L0) L) func(e Either[L0, R]) Either[L, R] {
	return MapLeft[L0, L, R](fn)
}

// ConstMap replaces the Left value if Left, otherwise forwards the Right value.
func ConstMapLeft[L0, R, L any](value L) func(e Either[L0, R]) Either[L, R] {
	return func(e Either[L0, R]) Either[L, R] {
		if IsRight(e) {
			return Right[L](e.right)
		}
		return Left[L, R](value)
	}
}
