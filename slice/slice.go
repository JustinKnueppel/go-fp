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

// Concat flattens a multidimensional slice by one dimension.
func Concat[T any](ts [][]T) []T {
	out := []T{}

	for _, t := range ts {
		out = append(out, t...)
	}

	return out
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

// DeleteAt removes the element at the given index if the slice is sufficiently long.
func DeleteAt[T any](index int) func([]T) []T {
	return func(ts []T) []T {
		out := []T{}
		for i, t := range ts {
			if i != index {
				out = append(out, t)
			}
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

// FilterMap applies the transformation function, filters out any Nones,
// and returns the unwrapped Options remaining.
func FilterMap[T, U any](fn func(T) option.Option[U]) func([]T) []U {
	return fp.Pipe3(
		Map(fn),
		Filter(option.IsSome[U]),
		Map(option.Unwrap[U]),
	)
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
			for _, pair := range Reverse(ZipIndexes(ts)) {
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
			for _, pair := range Reverse(ZipIndexes(ts)) {
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

// Maximum returns the highest number in the slice if any elemnts exist, otherwise None.
func Maximum[T operator.Number](xs []T) option.Option[T] {
	gt := fp.Curry2(func(x, y T) T {
		if x > y {
			return x
		}
		return y
	})

	return Reduce(gt)(xs)
}

// Minimum returns the lowest number in the slice if any elemnts exist, otherwise None.
func Minimum[T operator.Number](xs []T) option.Option[T] {
	lt := fp.Curry2(func(x, y T) T {
		if x < y {
			return x
		}
		return y
	})

	return Reduce(lt)(xs)
}

// Permutations returns all n length permutations of the given slice.
func Permutations[T any](ts []T) [][]T {
	if IsEmpty(ts) {
		return [][]T{{}}
	}
	if Length(ts) == 1 {
		return [][]T{ts}
	}
	out := [][]T{}

	for i, t := range ts {
		for _, perm := range Permutations(DeleteAt[T](i)(ts)) {
			out = append(out, append([]T{t}, perm...))
		}
	}
	return out
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

// Product returns the product of the slice of numbers, or 1 if empty.
func Product[T operator.Number](xs []T) T {
	return Fold(operator.Multiply[T])(1)(xs)
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

// Singleton returns a slice with only the element given
func Singleton[T any](t T) []T {
	return []T{t}
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

// SplitAt returns the first n elements and the remaining elements as a pair.
// If the slice runs out of elements, empty slices will remain in the pairs.
func SplitAt[T any](i int) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		return tuple.NewPair[[]T, []T](Take[T](i)(ts))(Drop[T](i)(ts))
	}
}

// Subsequences returns the slice of all subsequences of the slice.
func Subsequences[T any](ts []T) [][]T {
	out := [][]T{{}}

	for _, t := range ts {
		for _, s := range out {
			out = append(out, append(s, t))
		}
	}

	return out
}

// Sum returns the sum of the elements in the slice, or 0 if empty.
func Sum[T operator.Number](ts []T) T {
	return Fold(operator.Add[T])(0)(ts)
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

// Transpose converts all existing rows to columns in the 2D slice,
// skipping elements if the given row is not long enough.
func Transpose[T any](ts [][]T) [][]T {
	return option.MapOr[int]([][]T{})(func(maxLen int) [][]T {
		return Map(func(i int) []T {
			return FilterMap(At[T](i))(ts)
		})(Range(0)(maxLen))
	})(Maximum(Map(Length[T])(ts)))
}

// Uncons returns a pair containing the head and tail of the slice.
func Uncons[T any](s []T) option.Option[tuple.Pair[T, []T]] {
	if IsEmpty(s) {
		return option.None[tuple.Pair[T, []T]]()
	}

	return option.Some(tuple.NewPair[T, []T](s[0])(s[1:]))
}

// Zip takes two slices and returns a slice of corresponding pairs.
func Zip[T, U any](ts []T) func([]U) []tuple.Pair[T, U] {
	return func(us []U) []tuple.Pair[T, U] {
		n := Length(ts)
		if n > Length(us) {
			n = Length(us)
		}
		out := []tuple.Pair[T, U]{}
		for i := 0; i < n; i++ {
			out = append(out, tuple.NewPair[T, U](ts[i])(us[i]))
		}
		return out
	}
}

// Unzip converts a zipped slice into two separate slices.
func Unzip[T, U any](pairs []tuple.Pair[T, U]) tuple.Pair[[]T, []U] {
	ts := []T{}
	us := []U{}
	for _, pair := range pairs {
		ts = append(ts, tuple.Fst(pair))
		us = append(us, tuple.Snd(pair))
	}
	return tuple.NewPair[[]T, []U](ts)(us)
}

// ZipWith combines the two slices using the given function.
func ZipWith[T, U, V any](fn func(T) func(U) V) func([]T) func([]U) []V {
	return func(ts []T) func([]U) []V {
		return func(us []U) []V {
			n := Length(ts)
			if n > Length(us) {
				n = Length(us)
			}
			out := []V{}
			for i := 0; i < n; i++ {
				out = append(out, fn(ts[i])(us[i]))
			}
			return out
		}
	}
}

// ZipIndexes returns an (index, value) zipped version of the slice.
func ZipIndexes[T any](ts []T) []tuple.Pair[int, T] {
	out := []tuple.Pair[int, T]{}
	for i, t := range ts {
		out = append(out, tuple.NewPair[int, T](i)(t))
	}
	return out
}
