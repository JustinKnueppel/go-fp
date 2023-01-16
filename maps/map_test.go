package maps_test

import (
	"fmt"
	"strings"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/maps"
	"github.com/JustinKnueppel/go-fp/operator"
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

func printAny[T any](t T) {
	fmt.Println(t)
}

func ExampleAdjust() {
	add1 := func(x int) int { return x + 1 }

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(3)},
	))

	fp.Pipe3(
		maps.Adjust[int](add1)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
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
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.AdjustWithKey(addKV)(1),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(3)},
	))

	fp.Pipe3(
		maps.AdjustWithKey(addKV)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
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
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{}))

	fp.Pipe3(
		maps.Alter[int](safeDivFrom12)(3),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(0),
		tuple.NewPair[int, int](4)(5),
	}))

	fp.Pipe3(
		maps.Alter[int](safeDivFrom12)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
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
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe3(
		maps.Assocs[int, string],
		slice.Sort(keysLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
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
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe2(
		maps.AssocsOrdered[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](0)("foo"),
		tuple.NewPair[int, string](1)("bar"),
	}))

	fp.Pipe2(
		maps.AssocsOrdered[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
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

	printAny(maps.ToAscSlice[string, int](strLt)(m))
	printAny(maps.ToAscSlice[string, int](strLt)(mCopy))

	// Output:
	// [(bar 2) (foo 2)]
	// [(bar 2) (foo 1)]
}

func ExampleDelete() {
	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe3(
		maps.Delete[string, int]("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	// Output:
	// []
	// [(bar 3)]
	// [(bar 3)]
}

func ExampleDifference() {
	baseMap := maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	})

	fp.Pipe3(
		maps.Difference[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("bar")("world"),
	}))

	fp.Pipe3(
		maps.Difference[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("new")("world"),
	}))

	fp.Pipe3(
		maps.Difference[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.Difference[string, int, string](maps.Empty[string, int]()),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("new")("world"),
	}))

	// Output:
	// [(baz 3)]
	// [(bar 2) (baz 3)]
	// [(bar 2) (baz 3) (foo 1)]
	// []
}

func ExampleDifferenceWith() {
	eqLen := fp.Curry2(func(x int, s string) option.Option[int] {
		if len(s) == x {
			return option.Some(x)
		}
		return option.None[int]()
	})

	baseMap := maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	})

	fp.Pipe3(
		maps.DifferenceWith[string](eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1"),
		tuple.NewPair[string, string]("bar")("1"),
	}))

	fp.Pipe3(
		maps.DifferenceWith[string](eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("12345"),
		tuple.NewPair[string, string]("bar")("12345"),
	}))

	fp.Pipe3(
		maps.DifferenceWith[string](eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.DifferenceWith[string](eqLen)(maps.Empty[string, int]()),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1234"),
		tuple.NewPair[string, string]("new")("1234"),
	}))

	// Output:
	// [(baz 3) (foo 1)]
	// [(baz 3)]
	// [(bar 2) (baz 3) (foo 1)]
	// []
}

func ExampleElems() {
	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort(intLt),
		fp.Inspect(printAny[[]int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("bar")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort(intLt),
		fp.Inspect(printAny[[]int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("baz")(2),
	}))

	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort(intLt),
		fp.Inspect(printAny[[]int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	// Output:
	// [1 2 3]
	// [2 2]
	// []
}

func ExampleEmpty() {
	m := maps.Empty[string, int]()
	fmt.Printf("Items: %v, size: %d", m, maps.Size(m))

	// Ouptut:
	// Items: [], size: 0
}

func ExampleFilter() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe3(
		maps.Filter[string](isEven),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Filter[string](isEven),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// []
	// [(bar 2)]
}

func ExampleFilterWithKey() {
	keyIsLen := fp.Curry2(func(s string, x int) bool { return len(s) == x })

	fp.Pipe3(
		maps.FilterWithKey(keyIsLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.FilterWithKey(keyIsLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	// Output:
	// []
	// [(bar 3)]
}

func ExampleFindWithDefault() {
	fp.Pipe2(
		maps.FindWithDefault[string](-1)("foo"),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.FindWithDefault[string](-1)("foo"),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe2(
		maps.FindWithDefault[string](-1)("foo"),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(3),
		tuple.NewPair[string, int]("baz")(1),
	}))

	// Output:
	// -1
	// 1
	// -1
}

func ExampleFold() {
	fp.Pipe2(
		maps.Fold[string](operator.Add[int])(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.Fold[string](operator.Add[int])(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
	}))

	fp.Pipe2(
		maps.Fold[string](operator.Add[int])(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("bar")(3),
		tuple.NewPair[string, int]("baz")(1),
	}))

	// Output:
	// 0
	// 2
	// 6
}

func ExampleFoldWithKey() {
	addTimesStrLen := fp.Curry3(func(k string, x, y int) int { return (len(k) * x) + y })

	fp.Pipe2(
		maps.FoldWithKey(addTimesStrLen)(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.FoldWithKey(addTimesStrLen)(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
	}))

	fp.Pipe2(
		maps.FoldWithKey(addTimesStrLen)(0),
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("123")(2),
		tuple.NewPair[string, int]("1234")(3),
		tuple.NewPair[string, int]("12345")(1),
	}))

	// Output:
	// 0
	// 6
	// 23
}

func ExampleFoldrWithKey() {
	addTimesStrLen := fp.Curry3(func(k, s, acc string) string { return fmt.Sprintf("%s,%s:%s", k, s, acc) })

	fp.Pipe2(
		maps.FoldrWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe2(
		maps.FoldrWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("x")("foo"),
	}))

	fp.Pipe2(
		maps.FoldrWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("a")("foo"),
		tuple.NewPair[string, string]("b")("bar"),
		tuple.NewPair[string, string]("c")("baz"),
	}))

	// Output:
	// _
	// x,foo:_
	// a,foo:b,bar:c,baz:_
}

func ExampleFoldlWithKey() {
	addTimesStrLen := fp.Curry3(func(acc, k, s string) string { return fmt.Sprintf("%s:%s,%s", acc, k, s) })

	fp.Pipe2(
		maps.FoldlWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe2(
		maps.FoldlWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("x")("foo"),
	}))

	fp.Pipe2(
		maps.FoldlWithKey[string, string, string](strLt)(addTimesStrLen)("_"),
		fp.Inspect(printAny[string]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("a")("foo"),
		tuple.NewPair[string, string]("b")("bar"),
		tuple.NewPair[string, string]("c")("baz"),
	}))

	// Output:
	// _
	// _:x,foo
	// _:a,foo:b,bar:c,baz
}
