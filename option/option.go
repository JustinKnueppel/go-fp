package option

import "fmt"

/* ============ Option type ============ */

// Option represents the presence of a value with Some, or
// the absence of value with None.
type Option[T any] struct {
	data     T
	has_data bool
}

// String is used only for correctly printing an Option.
func (o Option[T]) String() string {
	if IsNone(o) {
		return "None"
	}
	return fmt.Sprintf("Some %v", o.data)
}

/* ============ Constructors ============ */

// Some returns an Option that contains the value.
func Some[T any](val T) Option[T] {
	return Option[T]{
		data:     val,
		has_data: true,
	}
}

// None returns an Option with no value.
func None[T any]() Option[T] {
	var t T
	return Option[T]{
		data:     t,
		has_data: false,
	}
}

/* ============ Basic functions ============ */

// Copy returns a value copy of the option.
func Copy[T any](o Option[T]) Option[T] {
	return o
}

// Equal returns true if either both options are None,
// or if both are Some with the same value.
func Equal[T comparable](optb Option[T]) func(opta Option[T]) bool {
	return func(o Option[T]) bool {
		if IsNone(o) && IsNone(optb) {
			return true
		}
		if IsNone(o) || IsNone(optb) {
			return false
		}
		return o.data == optb.data
	}
}

// IsSome returns true if the option is Some, or false if None.
func IsSome[T any](o Option[T]) bool {
	return o.has_data
}

// IsSomeAnd returns true if the option is a Some value
// and the value inside of it matches a predicate.
func IsSomeAnd[T any](predicate func(T) bool) func(o Option[T]) bool {
	return func(o Option[T]) bool {
		return o.has_data && predicate(o.data)
	}
}

// IsNone returns true if the option is None, or false if Some.
func IsNone[T any](o Option[T]) bool {
	return !o.has_data
}

/* ============ Query ============ */

// Contains returns true if the option is
// a Some value containing the given value.
func Contains[T comparable](x T) func(o Option[T]) bool {
	return func(o Option[T]) bool {
		if IsNone(o) {
			return false
		}
		return o.data == x
	}
}

/* ============ Extract value ============ */

// Expect returns the contained Some value unsafely.
// Panics with the given message if None.
func Expect[T any](msg string) func(o Option[T]) T {
	return func(o Option[T]) T {
		if IsNone(o) {
			panic(msg)
		}
		return o.data
	}
}

// Unwrap returns the contained Some value unsafely.
// Panics if None.
func Unwrap[T any](o Option[T]) T {
	if IsNone(o) {
		panic("Unwrap called on None")
	}
	return o.data
}

// UnwrapOr returns the contained Some value or
// a provided default.
func UnwrapOr[T any](fallback T) func(o Option[T]) T {
	return func(o Option[T]) T {
		if IsNone(o) {
			return fallback
		}
		return o.data
	}
}

// UnwrapOrElse returns the contained Some value or
// computes it from a closure.
func UnwrapOrElse[T any](fallbackFn func() T) func(o Option[T]) T {
	return func(o Option[T]) T {
		if IsNone(o) {
			return fallbackFn()
		}
		return o.data
	}
}

// UnwrapOrDefault returns the contained Some value or
// the zero value of type T.
func UnwrapOrDefault[T any](o Option[T]) T {
	if IsNone(o) {
		var t T
		return t
	}
	return o.data
}

/* ============ Map ============ */

// Map maps an Option[T] to an Option[U] by applying a function
// to the contained value if it exists.
func Map[T any, U any](f func(T) U) func(o Option[T]) Option[U] {
	return func(o Option[T]) Option[U] {
		if IsNone(o) {
			return None[U]()
		}
		return Some(f(o.data))
	}
}

// Bind returns None if the option is None,
// otherwise calls f with the wrapped value and returns the result.
func Bind[T any, U any](f func(T) Option[U]) func(o Option[T]) Option[U] {
	return func(o Option[T]) Option[U] {
		if IsNone(o) {
			return None[U]()
		}
		return f(o.data)
	}
}

// Flatten converts from Option[Option[T]] to Option[T].
func Flatten[T any](o Option[Option[T]]) Option[T] {
	if IsNone(o) {
		return None[T]()
	}
	return o.data
}

// MapOr returns the provided default result (if None),
// or applies a function to the contained value (if Some).
func MapOr[T any, U any](fallback U) func(fn func(T) U) func(o Option[T]) U {
	return func(f func(T) U) func(Option[T]) U {
		return func(o Option[T]) U {
			if IsNone(o) {
				return fallback
			}
			return f(o.data)
		}
	}
}

// MapOrElse computes a default function result (if None),
// or applies a different function to the contained value (if Some).
func MapOrElse[T any, U any](fallbackFn func() U) func(fn func(T) U) func(o Option[T]) U {
	return func(f func(T) U) func(Option[T]) U {
		return func(o Option[T]) U {
			if IsNone(o) {
				return fallbackFn()
			}
			return f(o.data)
		}
	}
}

// And returns None if the option is None, otherwise returns optb.
func And[T any, U any](optB Option[U]) func(o Option[T]) Option[U] {
	return func(o Option[T]) Option[U] {
		if IsNone(o) {
			return None[U]()
		}
		return optB
	}
}

// Or returns the option if it contains a value,
// otherwise returns optB.
func Or[T any](optB Option[T]) func(o Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsSome(o) {
			return o
		}
		return optB
	}
}

// OrElse returns the option if it contains a value,
// otherwise computes the fallback.
func OrElse[T any](fallbackFn func() Option[T]) func(o Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsSome(o) {
			return o
		}
		return fallbackFn()
	}
}

// Xor returns Some if exactly one of self, optB is Some,
// otherwise returns None.
func Xor[T any](optB Option[T]) func(o Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsSome(o) && IsNone(optB) {
			return o
		}
		if IsNone(o) && IsSome(optB) {
			return optB
		}
		return None[T]()
	}
}

/* ============ Filter ============ */

// Filter returns None if the option is None,
// otherwise calls predicate with the wrapped value and returns:
// - Some(t) if predicate returns true (where t is the wrapped value), and
// - None if predicate returns false.
func Filter[T any](f func(T) bool) func(o Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsNone(o) || !f(o.data) {
			return None[T]()
		}
		return o
	}
}

/* ============ Slice operations ============ */

// MapOption maps a function over a slice, throwing out any Nones and unwrapping all Somes.
func MapOption[T, U any](fn func(T) Option[U]) func(ts []T) []U {
	return func(ts []T) []U {
		out := []U{}
		for _, t := range ts {
			optVal := fn(t)
			if IsSome(optVal) {
				out = append(out, Unwrap(optVal))
			}
		}
		return out
	}
}

// SliceToOption returns None if the slice is empty, otherwise returns Some with the first element.
func SliceToOption[T any](ts []T) Option[T] {
	if len(ts) == 0 {
		return None[T]()
	}
	return Some(ts[0])
}

// OptionToSlice returns a singleton slice if
func OptionToSlice[T any](o Option[T]) []T {
	if IsNone(o) {
		return []T{}
	}
	return []T{Unwrap(o)}
}

// CatOptions takes a list of options and returns a list of the Some values.
func CatOptions[T any](opts []Option[T]) []T {
	out := []T{}
	for _, o := range opts {
		if IsSome(o) {
			out = append(out, Unwrap(o))
		}
	}
	return out
}

/* ============ Debug ============ */

// Inspect calls the provided closure with the contained value
// if it exists and returns the unchanged Option.
func Inspect[T any](f func(T)) func(o Option[T]) Option[T] {
	return func(o Option[T]) Option[T] {
		if IsSome(o) {
			f(o.data)
		}
		return o
	}
}

/* ============= Functor definitions ============= */

// Fmap maps the value of the option to the new value if it exists.
func Fmap[T, U any](fn func(T) U) func(o Option[T]) Option[U] {
	return Map(fn)
}

// ConstMap converts the option into Some(value) if Some, otherwise continues to be None.
func ConstMap[T, U any](value U) func(o Option[T]) Option[U] {
	return And[T](Some(value))
}
