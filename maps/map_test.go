package maps_test

import (
	"fmt"
	"strings"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/maps"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

func intLt(x int) func(int) bool {
	return func(y int) bool {
		return x < y
	}
}

func strLt(s1 string) func(string) bool {
	return func(s2 string) bool {
		return strings.Compare(s1, s2) < 0
	}
}

func printPairSlice[K, V any](pairs []tuple.Pair[K, V]) {
	fmt.Println(pairs)
}

func ExampleAdjust() {
	add1 := func(x int) int { return x + 1 }

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(3)},
	))

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](1)(2),
	}))

	// Output:
	// []
	// [(0 1) (2 3)]
	// [(0 1) (1 3)]
}

func ExampleAdjustWithKey() {
	addKV := fp.Curry2(func(k, v int) int { return k + v })

	fp.Pipe3(
		maps.AdjustWithKey(addKV)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.AdjustWithKey(addKV)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(3)},
	))

	fp.Pipe3(
		maps.AdjustWithKey(addKV)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](4)(5)},
	))

	// Output:
	// []
	// [(0 1) (2 3)]
	// [(0 1) (4 9)]
}

func ExampleAlter() {
	safeDivFrom12 := option.Bind(func(v int) option.Option[int] {
		if v == 0 {
			return option.None[int]()
		}
		return option.Some(12 / v)
	})

	fp.Pipe3(
		maps.Alter[int](safeDivFrom12)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.Alter[int](safeDivFrom12)(3),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(0),
		tuple.NewPair[int, int](4)(5),
	}))

	fp.Pipe3(
		maps.Alter[int](safeDivFrom12)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printPairSlice[int, int]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(0),
		tuple.NewPair[int, int](4)(0),
	}))

	// Output:
	// []
	// [(0 1) (2 0) (4 5)]
	// [(0 1) (2 0)]
}

func ExampleAssocs() {
	keysLt := fp.Curry2(func(p1, p2 tuple.Pair[int, string]) bool { return tuple.Fst(p1) < tuple.Fst(p2) })

	fp.Pipe3(
		maps.Assocs[int, string],
		slice.Sort(keysLt),
		fp.Inspect(printPairSlice[int, string]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe3(
		maps.Assocs[int, string],
		slice.Sort(keysLt),
		fp.Inspect(printPairSlice[int, string]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](0)("foo"),
		tuple.NewPair[int, string](1)("bar"),
	}))

	// Output:
	// []
	// [(0 foo) (1 bar)]
}

func ExampleAssocsOrdered() {
	fp.Pipe2(
		maps.AssocsOrdered[int, string](intLt),
		fp.Inspect(printPairSlice[int, string]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe2(
		maps.AssocsOrdered[int, string](intLt),
		fp.Inspect(printPairSlice[int, string]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](0)("foo"),
		tuple.NewPair[int, string](1)("bar"),
	}))

	fp.Pipe2(
		maps.AssocsOrdered[int, string](intLt),
		fp.Inspect(printPairSlice[int, string]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](1)("bar"),
		tuple.NewPair[int, string](0)("foo"),
	}))

	// Output:
	// []
	// [(0 foo) (1 bar)]
	// [(0 foo) (1 bar)]
}

func ExampleCopy() {
	m := map[string]int{"foo": 1, "bar": 2}
	mCopy := maps.Copy(m)
	m["foo"] = 2

	printPairSlice(maps.ToAscSlice[string, int](strLt)(m))
	printPairSlice(maps.ToAscSlice[string, int](strLt)(mCopy))

	// Output:
	// [(bar 2) (foo 2)]
	// [(bar 2) (foo 1)]
}

func ExampleDelete() {
	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printPairSlice[string, int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printPairSlice[string, int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printPairSlice[string, int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	// Output:
	// []
	// [(bar 3)]
	// [(bar 3)]
}
