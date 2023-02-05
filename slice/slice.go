package slice

import (
	"sort"
	"strings"
	"unicode"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/tuple"
)

/* =========== Basic functions =========== */

// Copy returns a value copy of the given slice.
func Copy[T any](slice []T) []T {
	out := []T{}
	out = append(out, slice...)
	return out
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

// Head returns None if the slice is empty, otherwise returns
// The first element of the list.
func Head[T any](ts []T) option.Option[T] {
	if Null(ts) {
		return option.None[T]()
	}
	return option.Some(ts[0])
}

// Last returns None if the slice is empty, otherwise
// returns the last element of the slice.
func Last[T any](ts []T) option.Option[T] {
	if Null(ts) {
		return option.None[T]()
	}
	return option.Some(ts[Length(ts)-1])
}

// Tail returns None if the slice is empty, otherwise returns
// a slice with all elements except for the first element.
func Tail[T any](ts []T) option.Option[[]T] {
	if Null(ts) {
		return option.None[[]T]()
	}
	return option.Some(ts[1:])
}

// Init returns None if the slice is empty, otherwise returns
// all elements of the slice except the final element.
func Init[T any](ts []T) option.Option[[]T] {
	if Null(ts) {
		return option.None[[]T]()
	}
	return option.Some(ts[0 : Length(ts)-1])
}

// Uncons returns a pair containing the head and tail of the slice.
func Uncons[T any](s []T) option.Option[tuple.Pair[T, []T]] {
	if Null(s) {
		return option.None[tuple.Pair[T, []T]]()
	}

	return option.Some(tuple.NewPair[T, []T](s[0])(s[1:]))
}

// Singleton returns a slice with only the element given
func Singleton[T any](t T) []T {
	return []T{t}
}

// Null returns true if the slice is empty.
func Null[T any](ts []T) bool {
	return Length(ts) == 0
}

// Length returns the length of the slice
func Length[T any](ts []T) int {
	return len(ts)
}

/* =========== Slice transformations =========== */

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

// FilterMap applies the transformation function, filters out any Nones,
// and returns the unwrapped Options remaining.
func FilterMap[T, U any](fn func(T) option.Option[U]) func([]T) []U {
	return fp.Pipe3(
		Map(fn),
		Filter(option.IsSome[U]),
		Map(option.Unwrap[U]),
	)
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

// Intersperse inserts the given element between every element of the slice.
func Intersperse[T any](t T) func([]T) []T {
	return fp.Pipe3(
		Foldl(fp.Curry2(func(list []T, x T) []T {
			return append(list, x, t)
		}))([]T{}),
		Init[T],
		option.UnwrapOr([]T{}),
	)
}

// Intercalate inserts the slice xs in between the slices in xs and concatenates the result.
func Intercalate[T any](xs []T) func([][]T) []T {
	return fp.Compose2(Concat[T], Intersperse(xs))
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

// Permutations returns all n length permutations of the given slice.
func Permutations[T any](ts []T) [][]T {
	if Null(ts) {
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

/* =========== Reducing slices (folds) =========== */

// Foldl returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func Foldl[T, U any](fn func(acc U) func(t T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
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

// FoldlWithIndex returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func FoldlWithIndex[T, U any](fn func(i int) func(acc U) func(t T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(i)(val)(t)
			}
			return val
		}
	}
}

// FoldlWithIndexAndSlice returns the provided initializer if the slice
// is empty, otherwise folds the slice from left to right onto the
// initializer based on the accumulator function.
func FoldlWithIndexAndSlice[T, U any](fn func(ts []T) func(i int) func(acc U) func(t T) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
				return init
			}
			val := init
			for i, t := range ts {
				val = fn(ts)(i)(val)(t)
			}
			return val
		}
	}
}

// Foldr returns the provided initializer if the slice is empty,
// otherwise folds the slice from right to left onto the initializer
// based on the accumulator function.
func Foldr[T, U any](fn func(t T) func(acc U) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
				return init
			}
			val := init
			for _, t := range Reverse(ts) {
				val = fn(t)(val)
			}
			return val
		}
	}
}

// FoldrWithIndex returns the provided initializer if the slice
// is empty, otherwise folds the slice from right to left onto the
// initializer based on the accumulator function.
func FoldrWithIndex[T, U any](fn func(i int) func(t T) func(acc U) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(ZipIndices(ts)) {
				val = fn(tuple.Fst(pair))(tuple.Snd(pair))(val)
			}
			return val
		}
	}
}

// FoldrWithIndexAndSlice returns the provided initializer if
// the slice is empty, otherwise folds the slice from right to left
// onto the initializer based on the accumulator function.
func FoldrWithIndexAndSlice[T, U any](fn func(ts []T) func(i int) func(t T) func(acc U) U) func(U) func([]T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
			if Null(ts) {
				return init
			}
			val := init
			for _, pair := range Reverse(ZipIndices(ts)) {
				val = fn(ts)(tuple.Fst(pair))(tuple.Snd(pair))(val)
			}
			return val
		}
	}
}

// Foldl1 returns None if the slice is empty, otherwise folds the slice
// from left to right based on the accumulator function with the first
// element of the slice as the initial value.
func Foldl1[T any](fn func(acc T) func(t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])
		}
		return option.Some(val)
	}
}

// Foldl1WithIndex returns None if the slice is empty, otherwise folds
// the slice from left to right based on the accumulator function with
// the first element of the slice as the initial value.
func Foldl1WithIndex[T any](fn func(acc T) func(t T) func(i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])(i)
		}
		return option.Some(val)
	}
}

// Foldl1WithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from left to right based on the accumulator function
// with the first element of the slice as the initial value.
func Foldl1WithIndexAndSlice[T any](fn func(acc T) func(t T) func(i int) func(slice []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(val)(ts[i])(i)(ts)
		}
		return option.Some(val)
	}
}

// Foldr1 returns None if the slice is empty, otherwise folds the slice
// from right to left based on the accumulator function with the last
// element of the slice as the initial value.
func Foldr1[T any](fn func(acc T) func(t T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
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

// Foldr1WithIndex returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func Foldr1WithIndex[T any](fn func(acc T) func(t T) func(i int) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
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

// Foldr1WithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func Foldr1WithIndexAndSlice[T any](fn func(acc T) func(t T) func(i int) func(ts []T) T) func([]T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
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

/* =========== Special folds =========== */

// Concat flattens a multidimensional slice by one dimension.
func Concat[T any](ts [][]T) []T {
	out := []T{}

	for _, t := range ts {
		out = append(out, t...)
	}

	return out
}

// ConcatMap maps a function over all elements in the slice and concatenates the resulting slices.
func ConcatMap[T, U any](fn func(T) []U) func([]T) []U {
	return fp.Compose2(Concat[U], Map(fn))
}

// And returns the conjunction of a slice of bools.
func And(xs []bool) bool {
	return Foldr(operator.And)(true)(xs)
}

// Or returns the disjunction of a slice of bools.
func Or(xs []bool) bool {
	return Foldr(operator.Or)(false)(xs)
}

// Any returns true if any element of the slice satisfies the predicate.
func Any[T any](predicate func(T) bool) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if predicate(t) {
				return true
			}
		}
		return false
	}
}

// All returns true if every element in the slice satisfies the predicate.
func All[T any](predicate func(T) bool) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if !predicate(t) {
				return false
			}
		}
		return true
	}
}

// Sum returns the sum of the elements in the slice, or 0 if empty.
func Sum[T operator.Number](ts []T) T {
	return Foldl(operator.Add[T])(0)(ts)
}

// Product returns the product of the slice of numbers, or 1 if empty.
func Product[T operator.Number](xs []T) T {
	return Foldl(operator.Multiply[T])(1)(xs)
}

// Maximum returns the highest number in the slice if any elemnts exist, otherwise None.
func Maximum[T operator.Number](xs []T) option.Option[T] {
	gt := fp.Curry2(func(x, y T) T {
		if x > y {
			return x
		}
		return y
	})

	return Foldl1(gt)(xs)
}

// Minimum returns the lowest number in the slice if any elemnts exist, otherwise None.
func Minimum[T operator.Number](xs []T) option.Option[T] {
	lt := fp.Curry2(func(x, y T) T {
		if x < y {
			return x
		}
		return y
	})

	return Foldl1(lt)(xs)
}

/* =========== Building slices =========== */

// Scanl applies the transformation function to the initial argument
// and the first argument in the list, then feeds each additional list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func Scanl[T, U any](fn func(U) func(T) U) func(U) func([]T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Inits[T],
			Map(Foldl(fn)(init)),
		)
	}
}

// Scanl1 returns a list of successive reduced values from the left without a starting element.
func Scanl1[T any](fn func(T) func(T) T) func([]T) []T {
	return func(ts []T) []T {
		if Null(ts) {
			return []T{}
		}
		x, xs := tuple.Pattern(option.Unwrap(Uncons(ts)))
		return Scanl(fn)(x)(xs)
	}
}

// Scan applies the transformation function to the initial argument
// and the last argument in the list, then feeds each additional previous list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func Scanr[T, U any](fn func(U) func(T) U) func(U) func([]T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Tails[T],
			Map(Foldl(fn)(init)),
		)
	}
}

// Scanr1 returns a list of successive reduced values from the right without a starting element.
func Scanr1[T any](fn func(T) func(T) T) func([]T) []T {
	return func(ts []T) []T {
		if Null(ts) {
			return []T{}
		}
		init := option.Unwrap(Init(ts))
		last := option.Unwrap(Last(ts))
		return Scanr(fn)(last)(init)
	}
}

/* =========== Building slices =========== */

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

// Iterate applies the given function to the given seed n times.
func Iterate[T any](fn func(T) T) func(T) func(int) []T {
	return func(seed T) func(int) []T {
		return func(n int) []T {
			iterFn := fp.Curry2(func(acc T, _ int) T { return fn(acc) })
			return Scanl(iterFn)(seed)(Range(0)(n))[1:]
		}
	}
}

// Replicate returns a slice with the value repeated n times.
func Replicate[T any](n int) func(T) []T {
	return func(value T) []T {
		out := []T{}
		for i := 0; i < n; i++ {
			out = append(out, value)
		}
		return out
	}
}

// Cycle returns a slice with the given slice repeated n times.
func Cycle[T any](n int) func([]T) []T {
	return func(ts []T) []T {
		return ConcatMap(fp.Const[[]T, int](ts))(Range(0)(n))
	}
}

// MapAccumL behaves like a combination of map and foldl. It applies a function each element of a slice, passing an accumulating parameter
// from left to right, and returning a final value of this accumulator together with the new slice.
func MapAccumL[A, T, U any](fn func(A) func(T) tuple.Pair[A, U]) func(A) func([]T) tuple.Pair[A, []U] {
	return func(init A) func([]T) tuple.Pair[A, []U] {
		return func(ts []T) tuple.Pair[A, []U] {
			acc := init
			s := []U{}
			for _, t := range ts {
				a, v := tuple.Pattern(fn(acc)(t))
				acc = a
				s = append(s, v)
			}
			return tuple.NewPair[A, []U](acc)(s)
		}
	}
}

// MapAccumR behaves like a combination of map and foldr. It applies a function each element of a slice, passing an accumulating parameter
// from right to left, and returning a final value of this accumulator together with the new slice.
func MapAccumR[A, T, U any](fn func(A) func(T) tuple.Pair[A, U]) func(A) func([]T) tuple.Pair[A, []U] {
	return func(init A) func([]T) tuple.Pair[A, []U] {
		return func(ts []T) tuple.Pair[A, []U] {
			acc := init
			s := []U{}
			for _, t := range Reverse(ts) {
				a, v := tuple.Pattern(fn(acc)(t))
				acc = a
				s = Prepend(v)(s)
			}
			return tuple.NewPair[A, []U](acc)(s)
		}
	}
}

/* =========== Unfolding =========== */

// Unfoldr is a dual to foldr. While foldr reduces a slice to a summary value, unfoldr builds a slice from a seed value. The function
// takes the elemnt and returns None if ti is done building the slice, or returns Some (t, u) where t is prepending to the slice, and b is
// used as the next element in a recursive call.
func Unfoldr[T, U any](fn func(U) option.Option[tuple.Pair[T, U]]) func(U) []T {
	return func(init U) []T {
		opt := fn(init)
		if option.IsNone(opt) {
			return []T{}
		}
		v, next := tuple.Pattern(option.Unwrap(opt))
		return Prepend(v)(Unfoldr(fn)(next))
	}
}

/* =========== Subslices =========== */

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

// SplitAt returns the first n elements and the remaining elements as a pair.
// If the slice runs out of elements, empty slices will remain in the pairs.
func SplitAt[T any](i int) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		return tuple.NewPair[[]T, []T](Take[T](i)(ts))(Drop[T](i)(ts))
	}
}

// TakeWhile takes elements from the beginning of the slice as long as the predicate returns true.
// all remaining elements will be dropped.
func TakeWhile[T any](predicate func(T) bool) func([]T) []T {
	return func(ts []T) []T {
		out := []T{}
		for i := 0; i < len(ts) && predicate(ts[i]); i++ {
			out = append(out, ts[i])
		}
		return out
	}
}

// DropWhile drops elements while the predicate remains true, and returns the remaining elements.
func DropWhile[T any](predicate func(T) bool) func([]T) []T {
	return func(ts []T) []T {
		i := 0
		for i < len(ts) && predicate(ts[i]) {
			i++
		}

		out := []T{}
		for i < len(ts) {
			out = append(out, ts[i])
			i++
		}
		return out
	}
}

// DropWhileEnd drops the largest suffix of a slice in which the given predicate holds for all elements.
func DropWhileEnd[T any](predicate func(T) bool) func([]T) []T {
	return fp.Compose3(Reverse[T], DropWhile(predicate), Reverse[T])
}

// Span splits the slice on the first element that satisfies the predicate
func Span[T any](predicate func(T) bool) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if Null(ts) {
			return tuple.NewPair[[]T, []T]([]T{})([]T{})
		}
		return fp.Pipe2(
			FindIndex(fp.Compose2(operator.Not, predicate)),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))([]T{}))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](ts[:i])(ts[i:])
			}),
		)(ts)
	}
}

// Break splits the slice on the first element that does not satisfy the predicate
func Break[T any](predicate func(T) bool) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if Null(ts) {
			return tuple.NewPair[[]T, []T]([]T{})([]T{})
		}
		return fp.Pipe2(
			FindIndex(predicate),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))([]T{}))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](ts[:i])(ts[i:])
			}),
		)(ts)
	}
}

// StripPrefix drops the given prefix from a slice. It returns None if the slice did not start with the given prefix,
// or Some with the remainder of the slice if it did.
func StripPrefix[T comparable](prefix []T) func([]T) option.Option[[]T] {
	return func(ts []T) option.Option[[]T] {
		if Length(ts) < Length(prefix) {
			return option.None[[]T]()
		}
		if !And(ZipWith(operator.Eq[T])(prefix)(ts)) {
			return option.None[[]T]()
		}
		return option.Some(ts[len(prefix):])
	}
}

// Group is a special case of GroupBy where elements are grouped if
// they are equivalent and adjacent.
func Group[T comparable](ts []T) [][]T {
	return GroupBy(operator.Eq[T])(ts)
}

// Inits returns the initial segments of a slice, shortest first.
func Inits[T any](ts []T) [][]T {
	return Prepend([]T{})(
		Map(func(i int) []T {
			return ts[0 : i+1]
		})(Range(0)(Length(ts))),
	)
}

// Tails returns the initial segments of the slice, with the shortest last
func Tails[T any](ts []T) [][]T {
	return Append([]T{})(Map(func(i int) []T {
		return ts[i:Length(ts)]
	})(Range(0)(Length(ts))))
}

/* =========== Predicates =========== */

// IsPrefixOf takes two slices and returns true if the first slice is a prefix of the second.
func IsPrefixOf[T comparable](prefix []T) func([]T) bool {
	return func(ts []T) bool {
		return Length(ts) >= Length(prefix) &&
			fp.Compose2(And, ZipWith(operator.Eq[T])(prefix))(ts)
	}
}

// IsSuffixOf takes two slices and returns true if the first slice is a suffix of the second.
func IsSuffixOf[T comparable](suffix []T) func([]T) bool {
	return func(ts []T) bool {
		tsLen := Length(ts)
		sLen := Length(suffix)
		return tsLen >= sLen &&
			fp.Compose2(And, ZipWith(operator.Eq[T])(suffix))(ts[tsLen-sLen:])
	}
}

// IsInfixOf takes two slices and returns true if the first list is contained, wholly and intact, anywhere within the second.
func IsInfixOf[T comparable](target []T) func([]T) bool {
	return fp.Compose2(Any(IsPrefixOf(target)), Tails[T])
}

// IsSubsequenceOf takes two slices and returns true if all the elements of the first slice occur, in order,
// in the second. The elements do not have to occur consecutively.
func IsSubsequenceOf[T comparable](target []T) func([]T) bool {
	return fp.Compose2(Any(Equal(target)), Subsequences[T])
}

/* =========== Searching slices =========== */

/* =========== By equality =========== */

// Elem returns true if the target exists in the slice.
func Elem[T comparable](target T) func([]T) bool {
	return func(ts []T) bool {
		for _, t := range ts {
			if t == target {
				return true
			}
		}
		return false
	}
}

// NotElem returns true if the target does not exist in the slice.
func NotElem[T comparable](target T) func([]T) bool {
	return fp.Compose2(operator.Not, Elem(target))
}

// Lookup looks up a key in an association slice.
func Lookup[K comparable, V any](key K) func([]tuple.Pair[K, V]) option.Option[V] {
	return func(pairs []tuple.Pair[K, V]) option.Option[V] {
		for _, pair := range pairs {
			k, v := tuple.Pattern(pair)
			if k == key {
				return option.Some(v)
			}
		}
		return option.None[V]()
	}
}

/* =========== By predicate =========== */

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

// Partition takes a predicate and a slice and returns the pair of slices
// of elemnts which do and do not satisfy the predicate, respectively.
func Partition[T any](predicate func(T) bool) func([]T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		passes := []T{}
		fails := []T{}

		for _, t := range ts {
			if predicate(t) {
				passes = append(passes, t)
			} else {
				fails = append(fails, t)
			}
		}

		return tuple.NewPair[[]T, []T](passes)(fails)
	}
}

/* =========== Indexing slices =========== */

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

// ElemIndex returns the index of the first element in the given slice which is equal to the query element,
// or None if there is no such element.
func ElemIndex[T comparable](target T) func([]T) option.Option[int] {
	return func(ts []T) option.Option[int] {
		for i, t := range ts {
			if t == target {
				return option.Some(i)
			}
		}
		return option.None[int]()
	}
}

// ElemIndices returns the indices of all elements in the given slice which are equal to the query element,
// or an empty slice if there is no such element.
func ElemIndices[T comparable](target T) func([]T) []int {
	return func(ts []T) []int {
		out := []int{}
		for i, t := range ts {
			if t == target {
				out = append(out, i)
			}
		}
		return out
	}
}

// FindIndex returns the index of the first element to satisfy
// the predicate, or None if no such element exists.
func FindIndex[T any](predicate func(T) bool) func([]T) option.Option[int] {
	return func(ts []T) option.Option[int] {
		for i, t := range ts {
			if predicate(t) {
				return option.Some(i)
			}
		}
		return option.None[int]()
	}
}

// FindIndices returns the indexes of every element which
// satisfies the predicate.
func FindIndices[T any](predicate func(T) bool) func([]T) []int {
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

/* =========== Zipping and unzipping slices =========== */

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

// ZipIndices returns an (index, value) zipped version of the slice.
func ZipIndices[T any](ts []T) []tuple.Pair[int, T] {
	out := []tuple.Pair[int, T]{}
	for i, t := range ts {
		out = append(out, tuple.NewPair[int, T](i)(t))
	}
	return out
}

/* =========== Special slices =========== */

/* =========== Functions on strings =========== */

// Lines splits a string based on newlines.
func Lines(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	runes := []rune(s)
	newLineIndex := ElemIndex('\n')(runes)
	if option.IsNone(newLineIndex) {
		return []string{s}
	}
	idx := option.Unwrap(newLineIndex)
	return Prepend(s[:idx])(Lines(s[idx+1:]))
}

// Words breaks a string up into a slice of words, which were delimited by white space.
func Words(s string) []string {
	toRunes := func(str string) []rune { return []rune(str) }
	toStr := func(rs []rune) string { return string(rs) }
	bothRuneEqual := fp.Curry2(func(r1, r2 rune) bool { return unicode.IsSpace(r1) == unicode.IsSpace(r2) })
	return fp.Compose4(Map(toStr), Filter(fp.Compose2(operator.Not, Any(unicode.IsSpace))), GroupBy(bothRuneEqual), toRunes)(s)
}

// Unlines concatenates a newline character to each string and then concatenates them.
func Unlines(strs []string) string {
	appendNewline := func(s string) string { return s + "\n" }
	return fp.Compose2(Foldl(operator.StrAppend)(""), Map(appendNewline))(strs)
}

// Unwords joins a slice of strings by a space.
func Unwords(strs []string) string {
	return strings.Join(strs, " ")
}

/* =========== "Set" operations =========== */

// Unique removes duplicates from the slice.
func Unique[T comparable](ts []T) []T {
	return UniqueBy(operator.Eq[T])(ts)
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

// Difference removes the first occurrence of each element of ys from xs.
func Difference[T comparable](xs []T) func(ys []T) []T {
	return Foldl(fp.Flip2(Delete[T]))(xs)
}

// Union appends the second slice to the first, with all elements y of the second slice
// which are equal to some element x in xs removed.
func Union[T comparable](xs []T) func(ys []T) []T {
	return UnionBy(operator.Eq[T])(xs)
}

// Intersect takes the slice intersection of the two slices. If the first slice contains duplicates, so will the result.
func Intersect[T comparable](xs []T) func(ys []T) []T {
	return IntersectBy(operator.Eq[T])(xs)
}

// InsertBy inserts x at the first postition where it is less than or equal to the next element.
func Insert[T operator.Number](x T) func([]T) []T {
	return InsertBy(operator.Leq[T])(x)
}

/* =========== Ordered slices =========== */

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

/* =========== Generalized functions =========== */

// UniqueBy takes a custom equality function and returns the slice with all duplicates removed.
func UniqueBy[T any](eq func(x1 T) func(x2 T) bool) func(xs []T) []T {
	return func(s []T) []T {
		uncons := Uncons(s)
		if option.IsNone(uncons) {
			return []T{}
		}
		x, xs := tuple.Pattern(option.Unwrap(uncons))
		return Prepend(x)(UniqueBy(eq)(Filter(fp.Compose2(operator.Not, eq(x)))(xs)))
	}
}

// DeleteBy removes the first element which satisfies the predicate eq(x)(y)
// if exists and returns the resulting slice.
func DeleteBy[T any](eq func(T) func(T) bool) func(x T) func(ys []T) []T {
	return func(x T) func([]T) []T {
		return func(ts []T) []T {
			found := false
			out := []T{}
			for _, y := range ts {
				if eq(x)(y) && !found {
					found = true
					continue
				}
				out = append(out, y)
			}
			return out
		}
	}
}

// DeleteFirstsBy takes a predicate and two slices and returns the first slice with
// the first occurrence of each element of the second slice removed.
func DeleteFirstsBy[T any](eq func(T) func(T) bool) func([]T) func([]T) []T {
	return Foldl(fp.Flip2(DeleteBy(eq)))
}

// GroupBy returns a list of lists whose concatenation is the original list.
// The elements are grouped based on their evaluation of the predicate function.
func GroupBy[T any](predicate func(T) func(T) bool) func([]T) [][]T {
	return func(ts []T) [][]T {
		if Null(ts) {
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

// UnionBy returns the second slice appended to the first with all elements y
// of the second slice which satisfy eq(x)(y) for some x in the first slice removed.
func UnionBy[T any](eq func(x T) func(y T) bool) func(xs []T) func(ys []T) []T {
	return func(xs []T) func([]T) []T {
		return func(ys []T) []T {
			yOut := Filter(func(y T) bool {
				return option.IsNone(Find(fp.Flip2(eq)(y))(xs))
			})(ys)
			return AppendSlice(yOut)(xs)
		}
	}
}

// IntersectBy returns the intersection of the two slices. If the first slice contains duplicate, the result will as well.
func IntersectBy[T any](eq func(x T) func(y T) bool) func(xs []T) func(ys []T) []T {
	return func(xs []T) func([]T) []T {
		return func(ys []T) []T {
			out := Filter(func(x T) bool {
				return option.IsSome(Find(eq(x))(ys))
			})(xs)
			return out
		}
	}
}

// InsertBy takes a custom <= function and inserts x at the first postition where it is less than
// or equal to the next element.
func InsertBy[T any](leq func(x T) func(y T) bool) func(x T) func(xs []T) []T {
	return func(x T) func([]T) []T {
		return func(xs []T) []T {
			yGeq := func(y T) bool { return leq(x)(y) }
			firstGeq := FindIndex(yGeq)(xs)
			if option.IsSome(firstGeq) {
				i := option.Unwrap(firstGeq)
				return Concat([][]T{xs[:i], {x}, xs[i:]})
			}
			return Append(x)(xs)
		}
	}
}

// MaximumBy returns None if the slice is empty, or Some x where x >= all
// other elements of the slice based on the provided less than function.
func MaximumBy[T any](lt func(x T) func(y T) bool) func(xs []T) option.Option[T] {
	pickMax := fp.Curry2(func(x, y T) T {
		if lt(x)(y) {
			return y
		}
		return x
	})
	return Foldl1(pickMax)
}

// MinimumBy returns None if the slice is empty, or Some x where x <= all
// other elements of the slice based on the provided less than function.
func MinimumBy[T any](lt func(x T) func(y T) bool) func(xs []T) option.Option[T] {
	pickMin := fp.Curry2(func(x, y T) T {
		if lt(y)(x) {
			return y
		}
		return x
	})
	return Foldl1(pickMin)
}

/* =========== Monadic functions =========== */

// Bind applies the function to each element of the slice,
// then flattens the slice.
func Bind[T, U any](fn func(T) []U) func([]T) []U {
	return func(ts []T) []U {
		out := []U{}
		for _, t := range ts {
			out = append(out, fn(t)...)
		}
		return out
	}
}
