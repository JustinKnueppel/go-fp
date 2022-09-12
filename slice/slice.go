package slice

import (
	"sort"

	"github.com/JustinKnueppel/go-fp/option"
)

// Append returns the slice with the given element appended.
func Append[T any](t T) func([]T) []T {
	return func(ts []T) []T {
		return append(ts, t)
	}
}

// AppendSlice returns the slice with all elements in other appended.
func AppendSlice[T any](other []T) func([]T) []T {
	return func(ts []T) []T {
		return append(ts, other...)
	}
}

// At returns None if the index is out of bounds, otherwise
// returns the value at the given index.
func At[T any](index int) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if index < 0 || index >= Length(ts) {
			return option.None[T]()
		}
		return option.Some(ts[index])
	}
}

// Contains returns true if the target exists in the slice.
func Contains[T comparable](target T) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if t == target {
				return true
			}
		}
		return false
	}
}

// Copy returns a value copy of the given slice.
func Copy[T any](slice []T) []T {
	out := []T{}
	out = append(out, slice...)
	return out
}

// Delete removes the first instance of the target if exists
// and returns the resulting slice.
func Delete[T comparable](target T) func([]T) []T {
	return func(ts []T) []T {
		found := false
		out := []T{}
		for _, t := range ts {
			if t == target && !found {
				found = true
				continue
			}
			out = append(out, t)
		}
		return out
	}
}

// DeleteBy removes the first element which satisfies the predicate
// if exists and returns the resulting slice.
func DeleteBy[T any](predicate func(T) bool) func([]T) []T {
	return func(ts []T) []T {
		found := false
		out := []T{}
		for _, t := range ts {
			if predicate(t) && !found {
				found = true
				continue
			}
			out = append(out, t)
		}
		return out
	}
}

// Drop returns the slice without its first n elements.
func Drop[T any](n int) func([]T) []T {
	return func(ts []T) []T {
		out := []T{}
		for i := n; i < Length(ts); i++ {
			out = append(out, ts[i])
		}
		return out
	}
}

// IsEmpty returns true if the slice is empty.
func IsEmpty[T any](ts []T) bool {
	return Length(ts) == 0
}

// Equal returns true if the two slices have the same elements in the same order.
func Equal[T comparable](other []T) func([]T) bool {
	return func(slice []T) bool {
		if Length(slice) != Length(other) {
			return false
		}
		for i, el := range slice {
			if other[i] != el {
				return false
			}
		}
		return true
	}
}

// Every returns true if every element in the slice satisfies the predicate.
func Every[T any](predicate func(T) bool) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if !predicate(t) {
				return false
			}
		}
		return true
	}
}

// Filter returns a new slice with all elements of the slice
// which satisfy the predicate.
func Filter[T any](predicate func(T) bool) func([]T) []T {
	return func(ts []T) []T {
		out := []T{}
		for _, t := range ts {
			if predicate(t) {
				out = append(out, t)
			}
		}
		return out
	}
}

// Find returns the first element to satisfy the predicate,
// or None if no such element exists.
func Find[T any](predicate func(T) bool) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		for _, t := range ts {
			if predicate(t) {
				return option.Some(t)
			}
		}
		return option.None[T]()
	}
}

// FlatMap applies the function to each element of the slice,
// then flattens the slice.
func FlatMap[T, U any](fn func(T) []U) func([]T) []U {
	return func(ts []T) []U {
		out := []U{}
		for _, t := range ts {
			out = append(out, fn(t)...)
		}
		return out
	}
}

// Fold returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func Fold[T any](fn func(acc T, t T) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, t := range ts {
				val = fn(val, t)
			}
			return val
		}
	}
}

// FoldWithIndex returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func FoldWithIndex[T any](fn func(acc T, t T, i int) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(val, t, i)
			}
			return val
		}
	}
}

// FoldWithIndexAndSlice returns the provided initializer if the slice
// is empty, otherwise folds the slice from left to right onto the
// initializer based on the accumulator function.
func FoldWithIndexAndSlice[T any](fn func(acc T, t T, i int, ts []T) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(val, t, i, ts)
			}
			return val
		}
	}
}

// FoldRight returns the provided initializer if the slice is empty,
// otherwise folds the slice from right to left onto the initializer
// based on the accumulator function.
func FoldRight[T any](fn func(acc T, t T) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, t := range Reverse(ts) {
				val = fn(val, t)
			}
			return val
		}
	}
}

// FoldRightWithIndex returns the provided initializer if the slice
// is empty, otherwise folds the slice from right to left onto the
// initializer based on the accumulator function.
func FoldRightWithIndex[T any](fn func(acc T, t T, i int) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(Zip(ts)) {
				val = fn(val, ZipValue(pair), ZipIndex(pair))
			}
			return val
		}
	}
}

// FoldRightWithIndexAndSlice returns the provided initializer if
// the slice is empty, otherwise folds the slice from right to left
// onto the initializer based on the accumulator function.
func FoldRightWithIndexAndSlice[T any](fn func(acc T, t T, i int, ts []T) T) func(T) func([]T) T {
	return func(init T) func([]T) T {
		return func(ts []T) T {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(Zip(ts)) {
				val = fn(val, ZipValue(pair), ZipIndex(pair), ts)
			}
			return val
		}
	}
}

// Head returns None if the slice is empty, otherwise returns
// The first element of the list.
func Head[T any](ts []T) option.Option[T] {
	if IsEmpty(ts) {
		return option.None[T]()
	}
	return option.Some(ts[0])
}

// Index returns the index of the first element to satisfy
// the predicate, or None if no such element exists.
func Index[T any](predicate func(T) bool) func([]T) option.Option[int] {
	return func(ts []T) option.Option[int] {
		for i, t := range ts {
			if predicate(t) {
				return option.Some(i)
			}
		}
		return option.None[int]()
	}
}

// Indexes returns the indexes of every element which
// satisfies the predicate.
func Indexes[T any](predicate func(T) bool) func([]T) []int {
	return func(ts []T) []int {
		out := []int{}
		for i, t := range ts {
			if predicate(t) {
				out = append(out, i)
			}
		}
		return out
	}
}

// Init returns None if the slice is empty, otherwise returns
// all elements of the slice except the final element.
func Init[T any](ts []T) option.Option[[]T] {
	if IsEmpty(ts) {
		return option.None[[]T]()
	}
	return option.Some(ts[0 : Length(ts)-1])
}

// Last returns None if the slice is empty, otherwise
// returns the last element of the slice.
func Last[T any](ts []T) option.Option[T] {
	if IsEmpty(ts) {
		return option.None[T]()
	}
	return option.Some(ts[Length(ts)-1])
}

// Length returns the length of the slice
func Length[T any](ts []T) int {
	return len(ts)
}

// Map applied the given function to every element of the slice and
// returns a new slice with the mapped elements.
func Map[T, U any](fn func(T) U) func([]T) []U {
	return func(ts []T) []U {
		out := []U{}
		for _, t := range ts {
			out = append(out, fn(t))
		}
		return out
	}
}

// Range returns a slice of integers from lower (inclusive) to upper (exclusive).
func Range(lower int) func(int) []int {
	return func(upper int) []int {
		out := []int{}
		for i := lower; i < upper; i++ {
			out = append(out, i)
		}
		return out
	}
}

// Reduce returns None if the slice is empty, otherwise folds the slice
// from left to right based on the accumulator function with the first
// element of the slice as the initial value.
func Reduce[T any](fn func(acc T, t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val, ts[i])
		}
		return option.Some(val)
	}
}

// ReduceWithIndex returns None if the slice is empty, otherwise folds
// the slice from left to right based on the accumulator function with
// the first element of the slice as the initial value.
func ReduceWithIndex[T any](fn func(acc T, t T, i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val, ts[i], i)
		}
		return option.Some(val)
	}
}

// ReduceWithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from left to right based on the accumulator function
// with the first element of the slice as the initial value.
func ReduceWithIndexAndSlice[T any](fn func(acc T, t T, i int, slice []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val, ts[i], i, ts)
		}
		return option.Some(val)
	}
}

// ReduceRight returns None if the slice is empty, otherwise folds the slice
// from right to left based on the accumulator function with the last
// element of the slice as the initial value.
func ReduceRight[T any](fn func(acc T, t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val, ts[i])
		}
		return option.Some(val)
	}
}

// ReduceRightWithIndex returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func ReduceRightWithIndex[T any](fn func(acc T, t T, i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val, ts[i], i)
		}
		return option.Some(val)
	}
}

// ReduceRightWithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func ReduceRightWithIndexAndSlice[T any](fn func(acc T, t T, i int, ts []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val, ts[i], i, ts)
		}
		return option.Some(val)
	}
}

// Repeat returns a slice with the value repeated n times.
func Repeat[T any](n int) func(T) []T {
	return func(value T) []T {
		out := []T{}
		for i := 0; i < n; i++ {
			out = append(out, value)
		}
		return out
	}
}

// Reverse returns the slice with the order of the elements reversed.
func Reverse[T any](ts []T) []T {
	out := []T{}
	tsLen := Length(ts)
	for i := range ts {
		out = append(out, ts[tsLen-1-i])
	}
	return out
}

// Some returns true if any element of the slice satisfies the predicate.
func Some[T any](predicate func(T) bool) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if predicate(t) {
				return true
			}
		}
		return false
	}
}

// Sort returns the sorted form of the slice using the provded less
// than function while keeping the original order of equal elements.
func Sort[T any](lt func(T) func(T) bool) func([]T) []T {
	return func(ts []T) []T {
		out := Copy(ts)
		goComparator := func(i, j int) bool {
			return lt(out[i])(out[j])
		}
		sort.SliceStable(out, goComparator)
		return out
	}
}

// Tail returns None if the slice is empty, otherwise returns
// a slice with all elements except for the first element.
func Tail[T any](ts []T) option.Option[[]T] {
	if IsEmpty(ts) {
		return option.None[[]T]()
	}
	return option.Some(ts[1:])
}

// Take returns the first n elements of the slice, or the
// full slice if n > length(slice).
func Take[T any](n int) func([]T) []T {
	return func(ts []T) []T {
		if n < 0 {
			return []T{}
		}
		out := []T{}
		for i := 0; i < n && i < Length(ts); i++ {
			out = append(out, ts[i])
		}
		return out
	}
}

// Zip returns an (index, value) zipped version of the slice.
func Zip[T any](ts []T) []ZippedPair[T] {
	out := []ZippedPair[T]{}
	for i, t := range ts {
		out = append(out, ZippedPair[T]{i, t})
	}
	return out
}

// ZippedPair represents an (index, value) zipped pair of a slice.
type ZippedPair[T any] struct {
	index int
	value T
}

// ZipIndex returns the index of the zipped pair.
func ZipIndex[T any](pair ZippedPair[T]) int {
	return pair.index
}

// ZipValue returns the value of the zipped pair.
func ZipValue[T any](pair ZippedPair[T]) T {
	return pair.value
}
