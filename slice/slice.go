package slice

import (
	"sort"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/tuple"
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

// Break splits the slice on the first element that does not satisfy the predicate
func Break[T any](predicate func(T) bool) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if IsEmpty(ts) {
			return tuple.NewPair[[]T, []T]([]T{})([]T{})
		}
		return fp.Pipe2(
			Index(predicate),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))([]T{}))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](ts[:i])(ts[i:])
			}),
		)(ts)
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
func Fold[T, U any](fn func(acc U) func(t T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, t := range ts {
				val = fn(val)(t)
			}
			return val
		}
	}
}

// FoldWithIndex returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func FoldWithIndex[T, U any](fn func(acc U) func(t T) func(i int) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(val)(t)(i)
			}
			return val
		}
	}
}

// FoldWithIndexAndSlice returns the provided initializer if the slice
// is empty, otherwise folds the slice from left to right onto the
// initializer based on the accumulator function.
func FoldWithIndexAndSlice[T, U any](fn func(acc U) func(t T) func(i int) func(ts []T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(val)(t)(i)(ts)
			}
			return val
		}
	}
}

// FoldRight returns the provided initializer if the slice is empty,
// otherwise folds the slice from right to left onto the initializer
// based on the accumulator function.
func FoldRight[T, U any](fn func(acc U) func(t T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, t := range Reverse(ts) {
				val = fn(val)(t)
			}
			return val
		}
	}
}

// FoldRightWithIndex returns the provided initializer if the slice
// is empty, otherwise folds the slice from right to left onto the
// initializer based on the accumulator function.
func FoldRightWithIndex[T, U any](fn func(acc U) func(t T) func(i int) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(Zip(ts)) {
				val = fn(val)(tuple.Snd(pair))(tuple.Fst(pair))
			}
			return val
		}
	}
}

// FoldRightWithIndexAndSlice returns the provided initializer if
// the slice is empty, otherwise folds the slice from right to left
// onto the initializer based on the accumulator function.
func FoldRightWithIndexAndSlice[T, U any](fn func(acc U) func(t T) func(i int) func(ts []T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if IsEmpty(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(Zip(ts)) {
				val = fn(val)(tuple.Snd(pair))(tuple.Fst(pair))(ts)
			}
			return val
		}
	}
}

// Group is a special case of GroupBy where elements are grouped if
// they are equivalent and adjacent.
func Group[T comparable](ts []T) [][]T {
	return GroupBy(operator.Eq[T])(ts)
}

// GroupBy returns a list of lists whose concatenation is the original list.
// The elements are grouped based on their evaluation of the predicate function.
func GroupBy[T any](predicate func(T) func(T) bool) func([]T) [][]T {
	return func(ts []T) [][]T {
		if IsEmpty(ts) {
			return [][]T{}
		}
		x := fp.Pipe2(Head[T], option.Unwrap[T])(ts)
		xs := fp.Pipe2(Tail[T], option.Unwrap[[]T])(ts)
		s := Span(predicate(x))(xs)
		fitIn := tuple.Fst(s)
		rest := tuple.Snd(s)
		return AppendSlice(GroupBy(predicate)(rest))([][]T{AppendSlice(fitIn)([]T{x})})
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

// Inits returns the initial segments of a slice, shortest first.
func Inits[T any](ts []T) [][]T {
	return Prepend([]T{})(
		Map(func(i int) []T {
			return ts[0 : i+1]
		})(Range(0)(Length(ts))),
	)
}

func Intersperse[T any](t T) func([]T) []T {
	return fp.Pipe3(
		Fold(fp.Curry2(func(list []T, x T) []T {
			return append(list, x, t)
		}))([]T{}),
		Init[T],
		option.UnwrapOr([]T{}),
	)
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

// Prepend inserts the given element at the beginning of the slice.
func Prepend[T any](t T) func([]T) []T {
	return func(ts []T) []T {
		return append([]T{t}, ts...)
	}
}

// PrependSlice inserts each element of the given slice at the beginning of the target slice.
func PrependSlice[T any](xs []T) func([]T) []T {
	return func(ts []T) []T {
		return append(xs, ts...)
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
func Reduce[T any](fn func(acc T) func(t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])
		}
		return option.Some(val)
	}
}

// ReduceWithIndex returns None if the slice is empty, otherwise folds
// the slice from left to right based on the accumulator function with
// the first element of the slice as the initial value.
func ReduceWithIndex[T any](fn func(acc T) func(t T) func(i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])(i)
		}
		return option.Some(val)
	}
}

// ReduceWithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from left to right based on the accumulator function
// with the first element of the slice as the initial value.
func ReduceWithIndexAndSlice[T any](fn func(acc T) func(t T) func(i int) func(slice []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])(i)(ts)
		}
		return option.Some(val)
	}
}

// ReduceRight returns None if the slice is empty, otherwise folds the slice
// from right to left based on the accumulator function with the last
// element of the slice as the initial value.
func ReduceRight[T any](fn func(acc T) func(t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val)(ts[i])
		}
		return option.Some(val)
	}
}

// ReduceRightWithIndex returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func ReduceRightWithIndex[T any](fn func(acc T) func(t T) func(i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val)(ts[i])(i)
		}
		return option.Some(val)
	}
}

// ReduceRightWithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func ReduceRightWithIndexAndSlice[T any](fn func(acc T) func(t T) func(i int) func(ts []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if IsEmpty(ts) {
			return option.None[T]()
		}
		sliceLen := Length(ts)
		val := ts[sliceLen-1]
		for _, i := range Reverse(Range(0)(sliceLen - 1)) {
			val = fn(val)(ts[i])(i)(ts)
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

// Scan applies the transformation function to the initial argument
// and the first argument in the list, then feeds each additional list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func Scan[T, U any](fn func(U) func(T) U) func(U) func([]T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Inits[T],
			Map(Fold(fn)(init)),
		)
	}
}

// Scan applies the transformation function to the initial argument
// and the last argument in the list, then feeds each additional previous list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func ScanRight[T, U any](fn func(U) func(T) U) func(U) func([]T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Tails[T],
			Map(Fold(fn)(init)),
		)
	}
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

// Span splits the slice on the first element that satisfies the predicate
func Span[T any](predicate func(T) bool) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if IsEmpty(ts) {
			return tuple.NewPair[[]T, []T]([]T{})([]T{})
		}
		return fp.Pipe2(
			Index(fp.Compose2(operator.Not, predicate)),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))([]T{}))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](ts[:i])(ts[i:])
			}),
		)(ts)
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

// Tails returns the initial segments of the slice, with the shortest last
func Tails[T any](ts []T) [][]T {
	return Append([]T{})(Map(func(i int) []T {
		return ts[i:Length(ts)]
	})(Range(0)(Length(ts))))
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
func Zip[T any](ts []T) []tuple.Pair[int, T] {
	out := []tuple.Pair[int, T]{}
	for i, t := range ts {
		out = append(out, tuple.NewPair[int, T](i)(t))
	}
	return out
}
