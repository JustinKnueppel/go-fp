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
	out := Empty[T]()
	out = append(out, slice...)
	return out
}

// Equal returns true if the two slices have the same elements in the same order.
func Equal[T comparable](other []T) func(s []T) bool {
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
func Append[T any](t T) func(s []T) []T {
	return func(ts []T) []T {
		return append(ts, t)
	}
}

// AppendSlice returns the slice with all elements in other appended.
func AppendSlice[T any](other []T) func(s []T) []T {
	return func(ts []T) []T {
		return append(ts, other...)
	}
}

// Prepend inserts the given element at the beginning of the slice.
func Prepend[T any](t T) func(s []T) []T {
	return func(ts []T) []T {
		return append([]T{t}, ts...)
	}
}

// PrependSlice inserts each element of the given slice at the beginning of the target slice.
func PrependSlice[T any](xs []T) func(s []T) []T {
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
	return option.Map(func(tail []T) tuple.Pair[T, []T] {
		// If tail is Some, there is always an element in s
		// TODO: use applicative functions when implemented
		return tuple.NewPair[T, []T](s[0])(tail)
	})(Tail(s))
}

// Unappend returns a pair containing the init and last of the slice
func Unappend[T any](s []T) option.Option[tuple.Pair[[]T, T]] {
	return option.Map(func(init []T) tuple.Pair[[]T, T] {
		// If init is Some, there is always a final element in s
		// TODO: use applicative functions when implemented
		return tuple.NewPair[[]T, T](init)(option.Unwrap(Last(s)))
	})(Init(s))
}

// Empty returns an empty slice of the given type.
func Empty[T any]() []T {
	return []T{}
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
func Map[T, U any](fn func(T) U) func(s []T) []U {
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
func FilterMap[T, U any](fn func(T) option.Option[U]) func(s []T) []U {
	return fp.Pipe3(
		Map(fn),
		Filter(option.IsSome[U]),
		Map(option.Unwrap[U]),
	)
}

// Reverse returns the slice with the order of the elements reversed.
func Reverse[T any](ts []T) []T {
	return option.MapOr[tuple.Pair[T, []T]](Empty[T]())(func(uncons tuple.Pair[T, []T]) []T {
		x, xs := tuple.Pattern(uncons)
		return Append(x)(Reverse(xs))
	})(Uncons(ts))
}

// Intersperse inserts the given element between every element of the slice.
func Intersperse[T any](t T) func(s []T) []T {
	return fp.Pipe3(
		Foldl(fp.Curry2(func(list []T, x T) []T {
			return append(list, x, t)
		}))(Empty[T]()),
		Init[T],
		option.UnwrapOr(Empty[T]()),
	)
}

// Intercalate inserts the slice xs in between the slices in xs and concatenates the result.
func Intercalate[T any](xs []T) func(s [][]T) []T {
	return fp.Compose2(Concat[T], Intersperse(xs))
}

// Transpose converts all existing rows to columns in the 2D slice,
// skipping elements if the given row is not long enough.
func Transpose[T any](ts [][]T) [][]T {
	return option.MapOr[int](Empty[[]T]())(func(maxLen int) [][]T {
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
		return Singleton(Empty[T]())
	}

	return Concat(Map(func(pair tuple.Pair[int, T]) [][]T {
		i, t := tuple.Pattern(pair)
		return Map(Prepend(t))(Permutations(DeleteAt[T](i)(ts)))
	})(ZipIndices(ts)))
}

/* =========== Reducing slices (folds) =========== */

// Foldl returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func Foldl[T, U any](fn func(acc U) func(t T) U) func(init U) func(s []T) U {
	foldFn := fp.Curry4(func(_ []T, _ int, acc U, t T) U { return fn(acc)(t) })
	return FoldlWithIndexAndSlice(foldFn)
}

// FoldlWithIndex returns the provided initializer if the slice is empty,
// otherwise folds the slice from left to right onto the initializer
// based on the accumulator function.
func FoldlWithIndex[T, U any](fn func(i int) func(acc U) func(t T) U) func(init U) func(s []T) U {
	foldFn := fp.Curry4(func(_ []T, i int, acc U, t T) U { return fn(i)(acc)(t) })
	return FoldlWithIndexAndSlice(foldFn)
}

// FoldlWithIndexAndSlice returns the provided initializer if the slice
// is empty, otherwise folds the slice from left to right onto the
// initializer based on the accumulator function.
func FoldlWithIndexAndSlice[T, U any](fn func(ts []T) func(i int) func(acc U) func(t T) U) func(init U) func(s []T) U {
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
func Foldr[T, U any](fn func(t T) func(acc U) U) func(init U) func(s []T) U {
	foldFn := fp.Curry4(func(_ []T, _ int, t T, acc U) U { return fn(t)(acc) })
	return FoldrWithIndexAndSlice(foldFn)
}

// FoldrWithIndex returns the provided initializer if the slice
// is empty, otherwise folds the slice from right to left onto the
// initializer based on the accumulator function.
func FoldrWithIndex[T, U any](fn func(i int) func(t T) func(acc U) U) func(init U) func(s []T) U {
	foldFn := fp.Curry4(func(_ []T, i int, t T, acc U) U { return fn(i)(t)(acc) })
	return FoldrWithIndexAndSlice(foldFn)
}

// FoldrWithIndexAndSlice returns the provided initializer if
// the slice is empty, otherwise folds the slice from right to left
// onto the initializer based on the accumulator function.
func FoldrWithIndexAndSlice[T, U any](fn func(ts []T) func(i int) func(t T) func(acc U) U) func(init U) func(s []T) U {
	return func(init U) func([]T) U {
		return func(ts []T) U {
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
func Foldl1[T any](fn func(acc T) func(t T) T) func(s []T) option.Option[T] {
	foldFn := fp.Curry4(func(_ []T, _ int, acc T, t T) T { return fn(acc)(t) })
	return Foldl1WithIndexAndSlice(foldFn)
}

// Foldl1WithIndex returns None if the slice is empty, otherwise folds
// the slice from left to right based on the accumulator function with
// the first element of the slice as the initial value.
func Foldl1WithIndex[T any](fn func(i int) func(acc T) func(t T) T) func(s []T) option.Option[T] {
	foldFn := fp.Curry4(func(_ []T, i int, acc T, t T) T { return fn(i)(acc)(t) })
	return Foldl1WithIndexAndSlice(foldFn)
}

// Foldl1WithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from left to right based on the accumulator function
// with the first element of the slice as the initial value.
func Foldl1WithIndexAndSlice[T any](fn func(slice []T) func(i int) func(acc T) func(t T) T) func(s []T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if Null(ts) {
			return option.None[T]()
		}
		val := ts[0]
		for _, i := range Range(1)(Length(ts)) {
			val = fn(ts)(i)(val)(ts[i])
		}
		return option.Some(val)
	}
}

// Foldr1 returns None if the slice is empty, otherwise folds the slice
// from right to left based on the accumulator function with the last
// element of the slice as the initial value.
func Foldr1[T any](fn func(t T) func(acc T) T) func(s []T) option.Option[T] {
	foldFn := fp.Curry4(func(_ []T, _ int, t T, acc T) T { return fn(t)(acc) })
	return Foldr1WithIndexAndSlice(foldFn)
}

// Foldr1WithIndex returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func Foldr1WithIndex[T any](fn func(i int) func(t T) func(acc T) T) func(s []T) option.Option[T] {
	foldFn := fp.Curry4(func(_ []T, i int, t T, acc T) T { return fn(i)(t)(acc) })
	return Foldr1WithIndexAndSlice(foldFn)
}

// Foldr1WithIndexAndSlice returns None if the slice is empty, otherwise
// folds the slice from right to left based on the accumulator function
// with the last element of the slice as the initial value.
func Foldr1WithIndexAndSlice[T any](fn func(ts []T) func(i int) func(t T) func(acc T) T) func(s []T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		return option.Map(func(unappend tuple.Pair[[]T, T]) T {
			init, last := tuple.Pattern(unappend)
			return Foldr(fp.Curry2(func(pair tuple.Pair[int, T], acc T) T {
				return fn(ts)(tuple.Fst(pair))(tuple.Snd(pair))(acc)
			}))(last)(ZipIndices(init))
		})(Unappend(ts))
	}
}

/* =========== Special folds =========== */

// Concat flattens a multidimensional slice by one dimension.
func Concat[T any](ts [][]T) []T {
	out := Empty[T]()

	for _, t := range ts {
		out = append(out, t...)
	}

	return out
}

// ConcatMap maps a function over all elements in the slice and concatenates the resulting slices.
func ConcatMap[T, U any](fn func(T) []U) func(s []T) []U {
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
func Any[T any](predicate func(T) bool) func(s []T) bool {
	return fp.Compose2(Or, Map(predicate))
}

// All returns true if every element in the slice satisfies the predicate.
func All[T any](predicate func(T) bool) func(s []T) bool {
	return fp.Compose2(And, Map(predicate))
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
	return Foldl1(max[T])(xs)
}

// Minimum returns the lowest number in the slice if any elemnts exist, otherwise None.
func Minimum[T operator.Number](xs []T) option.Option[T] {
	return Foldl1(min[T])(xs)
}

/* =========== Building slices =========== */

// Scanl applies the transformation function to the initial argument
// and the first argument in the list, then feeds each additional list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func Scanl[T, U any](fn func(acc U) func(t T) U) func(init U) func(s []T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Inits[T],
			Map(Foldl(fn)(init)),
		)
	}
}

// Scanl1 returns a list of successive reduced values from the left without a starting element.
func Scanl1[T any](fn func(acc T) func(t T) T) func(s []T) []T {
	return fp.Compose2(option.MapOr[tuple.Pair[T, []T]](Empty[T]())(fp.Tupled(Scanl(fn))), Uncons[T])
}

// Scan applies the transformation function to the initial argument
// and the last argument in the list, then feeds each additional previous list item
// through the function with the previous result. Outputs all intermediate steps and final calculation.
func Scanr[T, U any](fn func(t T) func(acc U) U) func(init U) func(s []T) []U {
	return func(init U) func([]T) []U {
		return fp.Pipe2(
			Tails[T],
			Map(Foldr(fn)(init)),
		)
	}
}

// Scanr1 returns a list of successive reduced values from the right without a starting element.
func Scanr1[T any](fn func(t T) func(acc T) T) func(s []T) []T {
	return fp.Compose2(option.MapOr[tuple.Pair[[]T, T]](Empty[T]())(fp.Tupled(fp.Flip2(Scanr(fn)))), Unappend[T])
}

/* =========== Building slices =========== */

// Range returns a slice of integers from lower (inclusive) to upper (exclusive).
func Range(lower int) func(upper int) []int {
	return func(upper int) []int {
		out := []int{}
		for i := lower; i < upper; i++ {
			out = append(out, i)
		}
		return out
	}
}

// Iterate applies the given function to the given seed n times.
func Iterate[T any](n int) func(fn func(T) T) func(seed T) []T {
	return func(fn func(T) T) func(T) []T {
		return func(seed T) []T {
			iterFn := fp.Curry2(func(acc T, _ int) T { return fn(acc) })
			return Scanl(iterFn)(seed)(Range(0)(n))[1:]
		}
	}
}

// Replicate returns a slice with the value repeated n times.
func Replicate[T any](n int) func(value T) []T {
	return fp.Compose2(fp.Flip2(Map[int, T])(Range(0)(n)), fp.Const[T, int])
}

// Cycle returns a slice with the given slice repeated n times.
func Cycle[T any](n int) func(s []T) []T {
	return fp.Compose2(fp.Flip2(ConcatMap[int, T])(Range(0)(n)), fp.Const[[]T, int])
}

// MapAccumL behaves like a combination of map and foldl. It applies a function each element of a slice, passing an accumulating parameter
// from left to right, and returning a final value of this accumulator together with the new slice.
func MapAccumL[A, T, U any](fn func(acc A) func(t T) tuple.Pair[A, U]) func(init A) func(s []T) tuple.Pair[A, []U] {
	return fp.Compose2(
		Foldl(func(acc tuple.Pair[A, []U]) func(T) tuple.Pair[A, []U] {
			accumVal, mappedValues := tuple.Pattern(acc)
			return fp.Compose2(tuple.MapRight[A](fp.Flip2(Append[U])(mappedValues)), fn(accumVal))
		}),
		fp.Flip2(tuple.NewPair[A, []U])(Empty[U]()),
	)
}

// MapAccumR behaves like a combination of map and foldr. It applies a function each element of a slice, passing an accumulating parameter
// from right to left, and returning a final value of this accumulator together with the new slice.
func MapAccumR[A, T, U any](fn func(acc A) func(t T) tuple.Pair[A, U]) func(init A) func(s []T) tuple.Pair[A, []U] {
	return fp.Compose2(
		Foldr(fp.Flip2(func(acc tuple.Pair[A, []U]) func(T) tuple.Pair[A, []U] {
			accumVal, mappedValues := tuple.Pattern(acc)
			return fp.Compose2(tuple.MapRight[A](fp.Flip2(Prepend[U])(mappedValues)), fn(accumVal))
		})),
		fp.Flip2(tuple.NewPair[A, []U])(Empty[U]()),
	)
}

/* =========== Unfolding =========== */

// Unfoldr is a dual to foldr. While foldr reduces a slice to a summary value, unfoldr builds a slice from a seed value. The function
// takes the elemnt and returns None if ti is done building the slice, or returns Some (t, u) where t is prepending to the slice, and b is
// used as the next element in a recursive call.
func Unfoldr[T, U any](fn func(U) option.Option[tuple.Pair[T, U]]) func(seed U) []T {
	return fp.Compose2(
		option.MapOr[tuple.Pair[T, U]](Empty[T]())(func(pair tuple.Pair[T, U]) []T {
			v, next := tuple.Pattern(pair)
			return Prepend(v)(Unfoldr(fn)(next))
		}),
		fn,
	)
}

/* =========== Subslices =========== */

// Take returns the first n elements of the slice, or the
// full slice if n > length(slice).
func Take[T any](n int) func(s []T) []T {
	return func(ts []T) []T {
		return FilterMap(fp.Flip2(At[T])(ts))(Range(0)(min(n)(Length(ts))))
	}
}

// Drop returns the slice without its first n elements.
func Drop[T any](n int) func(s []T) []T {
	return func(ts []T) []T {
		return FilterMap(fp.Flip2(At[T])(ts))(Range(n)(Length(ts)))
	}
}

// SplitAt returns the first n elements and the remaining elements as a pair.
// If the slice runs out of elements, empty slices will remain in the pairs.
func SplitAt[T any](i int) func(s []T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		return tuple.NewPair[[]T, []T](Take[T](i)(ts))(Drop[T](i)(ts))
	}
}

// TakeWhile takes elements from the beginning of the slice as long as the predicate returns true.
// all remaining elements will be dropped.
func TakeWhile[T any](predicate func(T) bool) func(s []T) []T {
	return func(ts []T) []T {
		return option.MapOr[int](ts)(func(idx int) []T {
			return Take[T](idx)(ts)
		})(FindIndex(fp.Compose2(operator.Not, predicate))(ts))
	}
}

// DropWhile drops elements while the predicate remains true, and returns the remaining elements.
func DropWhile[T any](predicate func(T) bool) func(s []T) []T {
	return func(ts []T) []T {
		return option.MapOr[int](Empty[T]())(func(idx int) []T {
			return Drop[T](idx)(ts)
		})(FindIndex(fp.Compose2(operator.Not, predicate))(ts))
	}
}

// DropWhileEnd drops the largest suffix of a slice in which the given predicate holds for all elements.
func DropWhileEnd[T any](predicate func(T) bool) func(s []T) []T {
	return fp.Compose3(Reverse[T], DropWhile(predicate), Reverse[T])
}

// Span splits the slice on the first element that satisfies the predicate
func Span[T any](predicate func(T) bool) func(s []T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if Null(ts) {
			return tuple.NewPair[[]T, []T](Empty[T]())(Empty[T]())
		}
		return fp.Pipe2(
			FindIndex(fp.Compose2(operator.Not, predicate)),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))(Empty[T]()))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](Take[T](i)(ts))(Drop[T](i)(ts))
			}),
		)(ts)
	}
}

// Break splits the slice on the first element that does not satisfy the predicate
func Break[T any](predicate func(T) bool) func(s []T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		if Null(ts) {
			return tuple.NewPair[[]T, []T](Empty[T]())(Empty[T]())
		}
		return fp.Pipe2(
			FindIndex(predicate),
			option.MapOr[int](tuple.NewPair[[]T, []T](Copy(ts))(Empty[T]()))(func(i int) tuple.Pair[[]T, []T] {
				return tuple.NewPair[[]T, []T](Take[T](i)(ts))(Drop[T](i)(ts))
			}),
		)(ts)
	}
}

// StripPrefix drops the given prefix from a slice. It returns None if the slice did not start with the given prefix,
// or Some with the remainder of the slice if it did.
func StripPrefix[T comparable](prefix []T) func(s []T) option.Option[[]T] {
	return fp.Pipe4(
		option.Some[[]T],
		option.Filter(fp.On[[]T](operator.Leq[int])(Length[T])(prefix)),
		option.Filter(fp.Compose2(And, ZipWith(operator.Eq[T])(prefix))),
		option.Map(Drop[T](Length(prefix))),
	)
}

// Group is a special case of GroupBy where elements are grouped if
// they are equivalent and adjacent.
func Group[T comparable](ts []T) [][]T {
	return GroupBy(operator.Eq[T])(ts)
}

// Inits returns the initial segments of a slice, shortest first.
func Inits[T any](ts []T) [][]T {
	return Prepend(Empty[T]())(
		Map(func(i int) []T {
			return ts[0 : i+1]
		})(Range(0)(Length(ts))),
	)
}

// Tails returns the initial segments of the slice, with the shortest last
func Tails[T any](ts []T) [][]T {
	return Append(Empty[T]())(Map(func(i int) []T {
		return ts[i:Length(ts)]
	})(Range(0)(Length(ts))))
}

/* =========== Predicates =========== */

// IsPrefixOf takes two slices and returns true if the first slice is a prefix of the second.
func IsPrefixOf[T comparable](prefix []T) func(s []T) bool {
	return func(ts []T) bool {
		return Length(ts) >= Length(prefix) &&
			fp.Compose2(And, ZipWith(operator.Eq[T])(prefix))(ts)
	}
}

// IsSuffixOf takes two slices and returns true if the first slice is a suffix of the second.
func IsSuffixOf[T comparable](suffix []T) func(s []T) bool {
	return func(ts []T) bool {
		tsLen := Length(ts)
		sLen := Length(suffix)
		return tsLen >= sLen &&
			fp.Compose2(And, ZipWith(operator.Eq[T])(suffix))(ts[tsLen-sLen:])
	}
}

// IsInfixOf takes two slices and returns true if the first list is contained, wholly and intact, anywhere within the second.
func IsInfixOf[T comparable](target []T) func(s []T) bool {
	return fp.Compose2(Any(IsPrefixOf(target)), Tails[T])
}

// IsSubsequenceOf takes two slices and returns true if all the elements of the first slice occur, in order,
// in the second. The elements do not have to occur consecutively.
func IsSubsequenceOf[T comparable](target []T) func(s []T) bool {
	return fp.Compose2(Any(Equal(target)), Subsequences[T])
}

/* =========== Searching slices =========== */

/* =========== By equality =========== */

// Elem returns true if the target exists in the slice.
func Elem[T comparable](target T) func(s []T) bool {
	return ElemBy(operator.Eq[T])(target)
}

// NotElem returns true if the target does not exist in the slice.
func NotElem[T comparable](target T) func(s []T) bool {
	return fp.Compose2(operator.Not, Elem(target))
}

// Lookup looks up a key in an association slice.
func Lookup[K comparable, V any](key K) func(pairs []tuple.Pair[K, V]) option.Option[V] {
	findKey := func(pair tuple.Pair[K, V]) bool { return tuple.Fst(pair) == key }
	return fp.Compose2(option.Map(tuple.Snd[K, V]), Find(findKey))
}

/* =========== By predicate =========== */

// Find returns the first element to satisfy the predicate,
// or None if no such element exists.
func Find[T any](predicate func(T) bool) func(s []T) option.Option[T] {
	return fp.Compose2(Head[T], Filter(predicate))
}

// Filter returns a new slice with all elements of the slice
// which satisfy the predicate.
func Filter[T any](predicate func(T) bool) func(s []T) []T {
	return func(ts []T) []T {
		out := Empty[T]()
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
func Partition[T any](predicate func(T) bool) func(s []T) tuple.Pair[[]T, []T] {
	return func(ts []T) tuple.Pair[[]T, []T] {
		passes := Filter(predicate)(ts)
		fails := Filter(fp.Compose2(operator.Not, predicate))(ts)

		return tuple.NewPair[[]T, []T](passes)(fails)
	}
}

/* =========== Indexing slices =========== */

// At returns None if the index is out of bounds, otherwise
// returns the value at the given index.
func At[T any](index int) func(s []T) option.Option[T] {
	return func(ts []T) option.Option[T] {
		if index < 0 || index >= Length(ts) {
			return option.None[T]()
		}
		return option.Some(ts[index])
	}
}

// DeleteAt removes the element at the given index if the slice is sufficiently long.
func DeleteAt[T any](index int) func(s []T) []T {
	return fp.Compose3(Map(tuple.Snd[int, T]), Filter(fp.Compose2(operator.Neq(index), tuple.Fst[int, T])), ZipIndices[T])
}

// ElemIndex returns the index of the first element in the given slice which is equal to the query element,
// or None if there is no such element.
func ElemIndex[T comparable](target T) func(s []T) option.Option[int] {
	return fp.Compose2(Head[int], ElemIndices(target))
}

// ElemIndices returns the indices of all elements in the given slice which are equal to the query element,
// or an empty slice if there is no such element.
func ElemIndices[T comparable](target T) func(s []T) []int {
	return FindIndices(operator.Eq(target))
}

// FindIndex returns the index of the first element to satisfy
// the predicate, or None if no such element exists.
func FindIndex[T any](predicate func(T) bool) func(s []T) option.Option[int] {
	return fp.Compose2(Head[int], FindIndices(predicate))
}

// FindIndices returns the indexes of every element which
// satisfies the predicate.
func FindIndices[T any](predicate func(T) bool) func(s []T) []int {
	return fp.Compose3(Map(tuple.Fst[int, T]), Filter(fp.Compose2(predicate, tuple.Snd[int, T])), ZipIndices[T])
}

/* =========== Zipping and unzipping slices =========== */

// Zip takes two slices and returns a slice of corresponding pairs.
func Zip[T, U any](ts []T) func(us []U) []tuple.Pair[T, U] {
	return ZipWith(tuple.NewPair[T, U])(ts)
}

// Unzip converts a zipped slice into two separate slices.
func Unzip[T, U any](pairs []tuple.Pair[T, U]) tuple.Pair[[]T, []U] {
	ts := Map(tuple.Fst[T, U])(pairs)
	us := Map(tuple.Snd[T, U])(pairs)
	return tuple.NewPair[[]T, []U](ts)(us)
}

// ZipWith combines the two slices using the given function.
func ZipWith[T, U, V any](fn func(T) func(U) V) func(ts []T) func(us []U) []V {
	return func(ts []T) func([]U) []V {
		return func(us []U) []V {
			joinAtIndex := func(i int) V { return fn(ts[i])(us[i]) }
			n := min(Length(ts))(Length(us))
			return Map(joinAtIndex)(Range(0)(n))
		}
	}
}

// ZipIndices returns an (index, value) zipped version of the slice.
func ZipIndices[T any](ts []T) []tuple.Pair[int, T] {
	return Zip[int, T](Range(0)(Length(ts)))(ts)
}

/* =========== Special slices =========== */

/* =========== Functions on strings =========== */

// Lines splits a string based on newlines.
func Lines(s string) []string {
	if len(s) == 0 {
		return Empty[string]()
	}

	return option.MapOr[int](Singleton(s))(func(idx int) []string {
		return Prepend(s[:idx])(Lines(s[idx+1:]))
	})(ElemIndex('\n')([]rune(s)))
}

// Words breaks a string up into a slice of words, which were delimited by white space.
func Words(s string) []string {
	toRunes := func(str string) []rune { return []rune(str) }
	toStr := func(rs []rune) string { return string(rs) }
	bothRuneEqual := fp.On[rune](operator.Eq[bool])(unicode.IsSpace)
	return fp.Compose4(Map(toStr), Filter(fp.Compose2(operator.Not, Any(unicode.IsSpace))), GroupBy(bothRuneEqual), toRunes)(s)
}

// Unlines concatenates a newline character to each string and then concatenates them.
func Unlines(strs []string) string {
	appendNewline := fp.Flip2(operator.StrAppend)("\n")
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
func Delete[T comparable](target T) func(s []T) []T {
	return DeleteBy(operator.Eq[T])(target)
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
func Insert[T operator.Number](x T) func(s []T) []T {
	return InsertBy(operator.Leq[T])(x)
}

/* =========== Ordered slices =========== */

// Sort returns the slice sorted from least to greatest.
func Sort[T operator.Number](xs []T) []T {
	return SortBy(operator.Lt[T])(xs)
}

/* =========== Generalized functions =========== */

// UniqueBy takes a custom equality function and returns the slice with all duplicates removed.
func UniqueBy[T any](eq func(x1 T) func(x2 T) bool) func(xs []T) []T {
	return fp.Compose2(option.MapOr[tuple.Pair[T, []T]](Empty[T]())(func(uncons tuple.Pair[T, []T]) []T {
		x, xs := tuple.Pattern(uncons)
		return Prepend(x)(UniqueBy(eq)(Filter(fp.Compose2(operator.Not, eq(x)))(xs)))
	}), Uncons[T])
}

// DeleteBy removes the first element which satisfies the predicate eq(x)(y)
// if exists and returns the resulting slice.
func DeleteBy[T any](eq func(T) func(T) bool) func(x T) func(ys []T) []T {
	return func(x T) func([]T) []T {
		return func(ts []T) []T {
			return fp.Compose2(option.MapOr[int](ts)(fp.Flip2(DeleteAt[T])(ts)), FindIndex(eq(x)))(ts)
		}
	}
}

// DeleteFirstsBy takes a predicate and two slices and returns the first slice with
// the first occurrence of each element of the second slice removed.
func DeleteFirstsBy[T any](eq func(T) func(T) bool) func(s1 []T) func(s2 []T) []T {
	return Foldl(fp.Flip2(DeleteBy(eq)))
}

// GroupBy returns a list of lists whose concatenation is the original list.
// The elements are grouped based on their evaluation of the predicate function.
func GroupBy[T any](predicate func(T) func(T) bool) func(s []T) [][]T {
	return func(ts []T) [][]T {
		return option.MapOr[T](Empty[[]T]())(func(x T) [][]T {
			group := TakeWhile(predicate(x))(ts)
			rest := DropWhile(predicate(x))(ts)
			return Prepend(group)(GroupBy(predicate)(rest))
		})(Head(ts))
	}
}

// ElemBy determines if the target exists in the slice using the provided equality function.
func ElemBy[T any](eq func(x T) func(y T) bool) func(x T) func(xs []T) bool {
	return fp.Compose2(Any[T], eq)
}

// UnionBy returns the second slice appended to the first with all elements y
// of the second slice which satisfy eq(x)(y) for some x in the first slice removed.
func UnionBy[T any](eq func(x T) func(y T) bool) func(xs []T) func(ys []T) []T {
	return func(xs []T) func([]T) []T {
		return func(ys []T) []T {
			neq := fp.Curry2(func(x, y T) bool {
				return !eq(x)(y)
			})
			return AppendSlice(Foldr(fp.Compose2(Filter[T], neq))(ys)(xs))(xs)
		}
	}
}

// IntersectBy returns the intersection of the two slices. If the first slice contains duplicate, the result will as well.
func IntersectBy[T any](eq func(x T) func(y T) bool) func(xs []T) func(ys []T) []T {
	return func(xs []T) func([]T) []T {
		return func(ys []T) []T {
			return Filter(fp.Flip2(ElemBy(eq))(ys))(xs)
		}
	}
}

// InsertBy takes a custom <= function and inserts x at the first postition where it is less than
// or equal to the next element.
func InsertBy[T any](leq func(x T) func(y T) bool) func(x T) func(xs []T) []T {
	return func(x T) func([]T) []T {
		return func(xs []T) []T {
			return option.MapOr[int](Append(x)(xs))(func(i int) []T {
				return Concat([][]T{xs[:i], {x}, xs[i:]})
			})(FindIndex(leq(x))(xs))
		}
	}
}

// SortBy returns the sorted form of the slice using the provded less
// than function while keeping the original order of equal elements.
func SortBy[T any](lt func(T) func(T) bool) func(s []T) []T {
	return func(ts []T) []T {
		out := Copy(ts)
		goComparator := func(i, j int) bool {
			return lt(out[i])(out[j])
		}
		sort.SliceStable(out, goComparator)
		return out
	}
}

// MaximumBy returns None if the slice is empty, or Some x where x >= all
// other elements of the slice based on the provided less than function.
func MaximumBy[T any](lt func(x T) func(y T) bool) func(xs []T) option.Option[T] {
	return Foldl1(maxBy(lt))
}

// MinimumBy returns None if the slice is empty, or Some x where x <= all
// other elements of the slice based on the provided less than function.
func MinimumBy[T any](lt func(x T) func(y T) bool) func(xs []T) option.Option[T] {
	return Foldl1(minBy(lt))
}

/* =========== Monadic functions =========== */

// Bind applies the function to each element of the slice,
// then flattens the slice.
func Bind[T, U any](fn func(T) []U) func(s []T) []U {
	return Foldr(fp.Compose2(PrependSlice[U], fn))([]U{})
}

/* =========== Helper functions =========== */

// max returns the larger of the two arguments.
func max[T operator.Number](x T) func(y T) T {
	return maxBy(operator.Lt[T])(x)
}

// maxBy returns the larger of the two arguments given the provded less than function.
func maxBy[T any](lt func(x T) func(y T) bool) func(x T) func(y T) T {
	return func(x T) func(y T) T {
		return func(y T) T {
			if lt(x)(y) {
				return y
			}
			return x
		}
	}
}

// min returns the smaller of the two arguments.
func min[T operator.Number](x T) func(y T) T {
	return minBy(operator.Lt[T])(x)
}

// minBy returns the smaller of the two arguments given the provded less than function.
func minBy[T any](lt func(x T) func(y T) bool) func(x T) func(y T) T {
	return func(x T) func(y T) T {
		return func(y T) T {
			if lt(y)(x) {
				return y
			}
			return x
		}
	}
}

/* =========== Functor definitions ============= */

// Fmap maps a function over a slice. It is equivalent to Map.
func Fmap[T, U any](fn func(T) U) func(xs []T) []U {
	return Map(fn)
}

// ConstMap replaces all instances in the slice with the static value.
func ConstMap[T, U any](value U) func(xs []T) []U {
	return Map(fp.Const[U, T](value))
}
