package maps_test

import (
	"fmt"
	"strings"

	"github.com/JustinKnueppel/go-fp/either"
	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/maps"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/set"
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
		maps.Alter[int](safeDivFrom12)(4),
		maps.ToAscSlice[int, int](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, int]]),
	)(maps.FromSlice([]tuple.Pair[int, int]{
		tuple.NewPair[int, int](0)(1),
		tuple.NewPair[int, int](2)(0),
		tuple.NewPair[int, int](4)(6),
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
	// [(0 1) (2 0) (4 2)]
	// [(0 1) (2 0)]
}

func ExampleAssocs() {
	keysLt := fp.Curry2(func(p1, p2 tuple.Pair[int, string]) bool { return tuple.Fst(p1) < tuple.Fst(p2) })

	fp.Pipe3(
		maps.Assocs[int, string],
		slice.SortBy(keysLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe3(
		maps.Assocs[int, string],
		slice.SortBy(keysLt),
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

func ExampleDifferenceWithKey() {
	eqLen := fp.Curry3(func(k string, x int, s string) option.Option[int] {
		if k == "special" {
			return option.Some(-1)
		}

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
		maps.DifferenceWithKey(eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1"),
		tuple.NewPair[string, string]("bar")("1"),
	}))

	fp.Pipe3(
		maps.DifferenceWithKey(eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("12345"),
		tuple.NewPair[string, string]("bar")("12345"),
	}))

	fp.Pipe3(
		maps.DifferenceWithKey(eqLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.DifferenceWithKey(eqLen)(maps.Empty[string, int]()),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1234"),
		tuple.NewPair[string, string]("new")("1234"),
	}))

	fp.Pipe3(
		maps.DifferenceWithKey(eqLen)(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("baz")(2),
			tuple.NewPair[string, int]("special")(2),
		})),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1234"),
		tuple.NewPair[string, string]("new")("1234"),
		tuple.NewPair[string, string]("special")("1234"),
	}))

	// Output:
	// [(baz 3) (foo 1)]
	// [(baz 3)]
	// [(bar 2) (baz 3) (foo 1)]
	// []
	// [(baz 2) (special -1)]
}

func ExampleElems() {
	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort[int],
		fp.Inspect(printAny[[]int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("bar")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort[int],
		fp.Inspect(printAny[[]int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("baz")(2),
	}))

	fp.Pipe3(
		maps.Elems[string, int],
		slice.Sort[int],
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

func ExampleFromSlice() {
	fp.Pipe3(
		maps.FromSlice[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{})

	fp.Pipe3(
		maps.FromSlice[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	})

	fp.Pipe3(
		maps.FromSlice[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(3),
	})

	// Output:
	// []
	// [(bar 2) (foo 1)]
	// [(bar 2) (foo 3)]
}

func ExampleFromSliceWith() {
	fp.Pipe3(
		maps.FromSliceWith[string](operator.Add[int]),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{})

	fp.Pipe3(
		maps.FromSliceWith[string](operator.Add[int]),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	})

	fp.Pipe3(
		maps.FromSliceWith[string](operator.Add[int]),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(3),
	})

	// Output:
	// []
	// [(bar 2) (foo 1)]
	// [(bar 2) (foo 4)]
}

func ExampleFromSliceWithKey() {
	weightedAdd := fp.Curry3(func(k string, cur, next int) int { return cur + len(k)*next })

	fp.Pipe3(
		maps.FromSliceWithKey(weightedAdd),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{})

	fp.Pipe3(
		maps.FromSliceWithKey(weightedAdd),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	})

	fp.Pipe3(
		maps.FromSliceWithKey(weightedAdd),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(3),
	})

	// Output:
	// []
	// [(bar 2) (foo 1)]
	// [(bar 2) (foo 10)]
}

func ExampleInsert() {
	fp.Pipe3(
		maps.Insert[string, int]("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Insert[string, int]("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	// Output:
	// [(foo 4)]
	// [(foo 4)]
}

func ExampleInsertWith() {
	fp.Pipe3(
		maps.InsertWith[string](operator.Add[int])("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.InsertWith[string](operator.Add[int])("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	// Output:
	// [(foo 4)]
	// [(foo 5)]
}

func ExampleInsertWithKey() {
	weightedAdd := fp.Curry3(func(k string, new, old int) int { return old + len(k)*new })

	fp.Pipe3(
		maps.InsertWithKey(weightedAdd)("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.InsertWithKey(weightedAdd)("foo")(4),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	// Output:
	// [(foo 4)]
	// [(foo 13)]
}

func ExampleInsertLookupWithKey() {
	weightedAdd := fp.Curry3(func(k string, new, old int) int { return old + len(k)*new })

	fp.Pipe3(
		maps.InsertLookupWithKey(weightedAdd)("foo")(4),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.InsertLookupWithKey(weightedAdd)("foo")(4),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	// Output:
	// (None [(foo 4)])
	// (Some 1 [(foo 13)])
}

func ExampleIntersection() {
	baseMap := maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	})

	fp.Pipe3(
		maps.Intersection[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("bar")("world"),
	}))

	fp.Pipe3(
		maps.Intersection[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("new")("world"),
	}))

	fp.Pipe3(
		maps.Intersection[string, int, string](baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	// Output:
	// [(bar 2) (foo 1)]
	// [(foo 1)]
	// []
}

func ExampleIntersectionWith() {
	addLen := fp.Curry2(func(x int, s string) int { return x + len(s) })

	baseMap := maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	})

	fp.Pipe3(
		maps.IntersectionWith[string](addLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("bar")("world"),
	}))

	fp.Pipe3(
		maps.IntersectionWith[string](addLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("new")("world"),
	}))

	fp.Pipe3(
		maps.IntersectionWith[string](addLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	// Output:
	// [(bar 7) (foo 6)]
	// [(foo 6)]
	// []
}

func ExampleIntersectionWithKey() {
	weightedAddLen := fp.Curry3(func(k string, x int, s string) int { return len(k) * (x + len(s)) })

	baseMap := maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	})

	fp.Pipe3(
		maps.IntersectionWithKey(weightedAddLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("bar")("world"),
	}))

	fp.Pipe3(
		maps.IntersectionWithKey(weightedAddLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("hello"),
		tuple.NewPair[string, string]("new")("world"),
	}))

	fp.Pipe3(
		maps.IntersectionWithKey(weightedAddLen)(baseMap),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	// Output:
	// [(bar 21) (foo 18)]
	// [(foo 18)]
	// []
}

func ExampleNull() {
	m := maps.Empty[string, int]()
	fmt.Println(maps.Null(m))

	m = maps.Singleton[string, int]("foo")(1)
	fmt.Println(maps.Null(m))

	// Output:
	// true
	// false
}

func ExampleIsSubmapOf() {
	fp.Pipe2(
		maps.IsSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	fp.Pipe2(
		maps.IsSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	// Output:
	// true
	// true
	// false
	// true
	// false
}

func ExampleIsSubmapOfBy() {
	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe2(
		maps.IsSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(-1),
	}))

	// Output:
	// true
	// true
	// false
	// true
	// true
	// false
}

func ExampleIsProperSubmapOf() {
	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOf(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// false
	// true
	// false
	// false
	// false
	// true
	// false
}

func ExampleIsProperSubmapOfBy() {
	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe2(
		maps.IsProperSubmapOfBy[string](operator.Leq[int])(maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("bar")(2),
		})),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(-1),
	}))

	// Output:
	// false
	// true
	// false
	// false
	// true
	// false
}

func ExampleKeys() {
	fp.Pipe3(
		maps.Keys[string, int],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Keys[string, int],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe3(
		maps.Keys[string, int],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// []
	// [bar foo]
	// [bar foo]
}

func ExampleKeysOrdered() {
	fp.Pipe2(
		maps.KeysOrdered[string, int](strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.KeysOrdered[string, int](strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.KeysOrdered[string, int](strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output
	// []
	// [(bar 2) (foo 1)]
	// [(bar 2) (foo 1)]
}

func ExampleKeysSet() {
	fp.Pipe4(
		maps.KeysSet[string, int],
		set.ToSlice[string],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe4(
		maps.KeysSet[string, int],
		set.ToSlice[string],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe4(
		maps.KeysSet[string, int],
		set.ToSlice[string],
		slice.SortBy(strLt),
		fp.Inspect(printAny[[]string]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// []
	// [bar foo]
	// [bar foo]
}

func ExampleLookup() {
	fp.Pipe2(
		maps.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
	}))

	fp.Pipe2(
		maps.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// None
	// Some 1
	// None
}

func ExampleMap() {
	fp.Pipe3(
		maps.Map[string](operator.Inc[int]),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Map[string](operator.Inc[int]),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// []
	// [(bar 3) (foo 2)]
}

func ExampleMapWithKey() {
	addKeyLen := fp.Curry2(func(k string, x int) int { return len(k) + x })

	fp.Pipe3(
		maps.MapWithKey(addKeyLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.MapWithKey(addKeyLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// []
	// [(bar 5) (foo 4)]
}

func ExampleMapAccum() {
	trackSumAndAdd5 := fp.Curry2(func(acc, x int) tuple.Pair[int, int] { return tuple.NewPair[int, int](acc + x)(x + 5) })

	fp.Pipe3(
		maps.MapAccum[string](trackSumAndAdd5)(0),
		tuple.MapRight[int](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[int, []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.MapAccum[string](trackSumAndAdd5)(0),
		tuple.MapRight[int](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[int, []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// (0 [])
	// (3 [(bar 7) (foo 6)])
}

func ExampleMapAccumOrdered() {
	trackAppendX := fp.Curry2(func(acc, x string) tuple.Pair[string, string] { return tuple.NewPair[string, string](acc + x)(x + "X") })

	fp.Pipe3(
		maps.MapAccumOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.MapAccumOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1"),
		tuple.NewPair[string, string]("bar")("2"),
	}))

	// Output:
	// (_ [])
	// (_21 [(bar 2X) (foo 1X)])
}

func ExampleMapAccumWithKey() {
	trackSumAndAddKeyLen := fp.Curry3(func(acc int, k string, x int) tuple.Pair[int, int] {
		return tuple.NewPair[int, int](acc + x)(x + len(k))
	})

	fp.Pipe3(
		maps.MapAccumWithKey(trackSumAndAddKeyLen)(0),
		tuple.MapRight[int](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[int, []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.MapAccumWithKey(trackSumAndAddKeyLen)(0),
		tuple.MapRight[int](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[int, []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// (0 [])
	// (3 [(bar 5) (foo 4)])
}

func ExampleMapAccumWithKeyOrdered() {
	trackAppendX := fp.Curry3(func(acc, k, x string) tuple.Pair[string, string] {
		return tuple.NewPair[string, string](acc + "," + k + ":" + x)(x + "X")
	})

	fp.Pipe3(
		maps.MapAccumWithKeyOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.MapAccumWithKeyOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1"),
		tuple.NewPair[string, string]("bar")("2"),
	}))

	// Output:
	// (_ [])
	// (_,bar:2,foo:1 [(bar 2X) (foo 1X)])
}

func ExampleMapAccumRWithKeyOrdered() {
	trackAppendX := fp.Curry3(func(acc, k, x string) tuple.Pair[string, string] {
		return tuple.NewPair[string, string](acc + "," + k + ":" + x)(x + "X")
	})

	fp.Pipe3(
		maps.MapAccumRWithKeyOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{}))

	fp.Pipe3(
		maps.MapAccumRWithKeyOrdered[string, string, string, string](strLt)(trackAppendX)("_"),
		tuple.MapRight[string](maps.ToAscSlice[string, string](strLt)),
		fp.Inspect(printAny[tuple.Pair[string, []tuple.Pair[string, string]]]),
	)(maps.FromSlice([]tuple.Pair[string, string]{
		tuple.NewPair[string, string]("foo")("1"),
		tuple.NewPair[string, string]("bar")("2"),
	}))

	// Output:
	// (_ [])
	// (_,foo:1,bar:2 [(bar 2X) (foo 1X)])
}

func ExampleMapEither() {
	leftEven := func(x int) either.Either[int, int] {
		if x%2 == 0 {
			return either.Left[int, int](x)
		}
		return either.Right[int](x)
	}

	fp.Pipe4(
		maps.MapEither[string](leftEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe4(
		maps.MapEither[string](leftEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe4(
		maps.MapEither[string](leftEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(2),
	}))

	fp.Pipe4(
		maps.MapEither[string](leftEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// ([] [])
	// ([] [(bar 3) (foo 1)])
	// ([(bar 2) (foo 0)] [])
	// ([(bar 2)] [(foo 1)])
}

func ExampleMapEitherWithKey() {
	leftEvenKey := fp.Curry2(func(k string, x int) either.Either[int, int] {
		if len(k)%2 == 0 {
			return either.Left[int, int](x)
		}
		return either.Right[int](x)
	})

	fp.Pipe4(
		maps.MapEitherWithKey(leftEvenKey),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe4(
		maps.MapEitherWithKey(leftEvenKey),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("a")(1),
		tuple.NewPair[string, int]("abc")(2),
	}))

	fp.Pipe4(
		maps.MapEitherWithKey(leftEvenKey),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("ab")(1),
		tuple.NewPair[string, int]("abcd")(2),
	}))

	fp.Pipe4(
		maps.MapEitherWithKey(leftEvenKey),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("a")(1),
		tuple.NewPair[string, int]("ab")(2),
	}))

	// Output:
	// ([] [])
	// ([] [(a 1) (abc 2)])
	// ([(ab 1) (abcd 2)] [])
	// ([(ab 2)] [(a 1)])
}

func ExampleMapKeys() {
	mod5 := func(x int) int { return x % 5 }

	fp.Pipe3(
		maps.MapKeys[int, int, string](intLt)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe3(
		maps.MapKeys[int, int, string](intLt)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](5)("foo"),
		tuple.NewPair[int, string](6)("bar"),
	}))

	fp.Pipe3(
		maps.MapKeys[int, int, string](intLt)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](0)("foo"),
		tuple.NewPair[int, string](5)("bar"),
	}))

	fp.Pipe3(
		maps.MapKeys[int, int, string](intLt)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](5)("bar"),
		tuple.NewPair[int, string](0)("foo"),
	}))

	// Output:
	// []
	// [(0 foo) (1 bar)]
	// [(0 bar)]
	// [(0 bar)]
}

func ExampleMapKeysWith() {
	strAppend := fp.Curry2(func(s1, s2 string) string { return s1 + s2 })
	mod5 := func(x int) int { return x % 5 }

	fp.Pipe3(
		maps.MapKeysWith[int, int, string](intLt)(strAppend)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe3(
		maps.MapKeysWith[int, int, string](intLt)(strAppend)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](5)("foo"),
		tuple.NewPair[int, string](6)("bar"),
	}))

	fp.Pipe3(
		maps.MapKeysWith[int, int, string](intLt)(strAppend)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](0)("foo"),
		tuple.NewPair[int, string](5)("bar"),
	}))

	fp.Pipe3(
		maps.MapKeysWith[int, int, string](intLt)(strAppend)(mod5),
		maps.ToAscSlice[int, string](intLt),
		fp.Inspect(printAny[[]tuple.Pair[int, string]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](5)("bar"),
		tuple.NewPair[int, string](0)("foo"),
	}))

	// Output:
	// []
	// [(0 foo) (1 bar)]
	// [(0 barfoo)]
	// [(0 barfoo)]
}

func ExampleMapOption() {
	safeDiv12 := func(x int) option.Option[int] {
		if x == 0 {
			return option.None[int]()
		}
		return option.Some(12 / x)
	}

	fp.Pipe3(
		maps.MapOption[string](safeDiv12),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.MapOption[string](safeDiv12),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(2),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe3(
		maps.MapOption[string](safeDiv12),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(3),
	}))

	// Output:
	// []
	// [(bar 4) (foo 6)]
	// [(bar 4)]
}

func ExampleMapOptionWithKey() {
	safeDivKeyLen := fp.Curry2(func(k string, x int) option.Option[int] {
		if x == 0 {
			return option.None[int]()
		}
		return option.Some(len(k) / x)
	})

	fp.Pipe3(
		maps.MapOptionWithKey(safeDivKeyLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.MapOptionWithKey(safeDivKeyLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("abcd")(2),
		tuple.NewPair[string, int]("abc")(3),
	}))

	fp.Pipe3(
		maps.MapOptionWithKey(safeDivKeyLen),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("abcd")(0),
		tuple.NewPair[string, int]("abc")(3),
	}))

	// Output:
	// []
	// [(abc 1) (abcd 2)]
	// [(abc 1)]
}

func ExampleMember() {
	fp.Pipe2(
		maps.Member[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.Member[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(3),
	}))

	fp.Pipe2(
		maps.Member[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(0),
	}))

	// Output:
	// false
	// true
	// false
}

func ExampleNotMember() {
	fp.Pipe2(
		maps.NotMember[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.NotMember[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(3),
	}))

	fp.Pipe2(
		maps.NotMember[string, int]("foo"),
		fp.Inspect(printAny[bool]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(0),
	}))

	// Output:
	// true
	// false
	// true
}

func ExamplePartition() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe4(
		maps.Partition[string](isEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe4(
		maps.Partition[string](isEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(2),
	}))

	fp.Pipe4(
		maps.Partition[string](isEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe4(
		maps.Partition[string](isEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(1),
	}))

	// Output:
	// ([] [])
	// ([(bar 2) (foo 0)] [])
	// ([] [(bar 3) (foo 1)])
	// ([(foo 0)] [(bar 1)])
}

func ExamplePartitionWithKey() {
	sumKvIsEven := fp.Curry2(func(k string, x int) bool { return (len(k)+x)%2 == 0 })

	fp.Pipe4(
		maps.PartitionWithKey(sumKvIsEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe4(
		maps.PartitionWithKey(sumKvIsEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(2),
	}))

	fp.Pipe4(
		maps.PartitionWithKey(sumKvIsEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(3),
	}))

	fp.Pipe4(
		maps.PartitionWithKey(sumKvIsEven),
		tuple.MapLeft[map[string]int, map[string]int](maps.ToAscSlice[string, int](strLt)),
		tuple.MapRight[[]tuple.Pair[string, int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[string, int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(0),
		tuple.NewPair[string, int]("bar")(1),
	}))

	// Output:
	// ([] [])
	// ([] [(bar 2) (foo 0)])
	// ([(bar 3) (foo 1)] [])
	// ([(bar 1)] [(foo 0)])
}

func ExampleSingleton() {
	m := maps.Singleton[string, int]("foo")(1)
	fmt.Println(maps.ToSlice(m))

	// Output:
	// [(foo 1)]
}

func ExampleSize() {
	fp.Pipe2(
		maps.Size[string, int],
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.Size[string, int],
		fp.Inspect(printAny[int]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	}))

	// Output:
	// 0
	// 2
}

func ExampleSplit() {
	fp.Pipe4(
		maps.Split[int, string](intLt)(5),
		tuple.MapLeft[map[int]string, map[int]string](maps.ToAscSlice[int, string](intLt)),
		tuple.MapRight[[]tuple.Pair[int, string]](maps.ToAscSlice[int, string](intLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[int, string], []tuple.Pair[int, string]]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{}))

	fp.Pipe4(
		maps.Split[int, string](intLt)(5),
		tuple.MapLeft[map[int]string, map[int]string](maps.ToAscSlice[int, string](intLt)),
		tuple.MapRight[[]tuple.Pair[int, string]](maps.ToAscSlice[int, string](intLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[int, string], []tuple.Pair[int, string]]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](1)("foo"),
		tuple.NewPair[int, string](8)("bar"),
	}))

	fp.Pipe4(
		maps.Split[int, string](intLt)(5),
		tuple.MapLeft[map[int]string, map[int]string](maps.ToAscSlice[int, string](intLt)),
		tuple.MapRight[[]tuple.Pair[int, string]](maps.ToAscSlice[int, string](intLt)),
		fp.Inspect(printAny[tuple.Pair[[]tuple.Pair[int, string], []tuple.Pair[int, string]]]),
	)(maps.FromSlice([]tuple.Pair[int, string]{
		tuple.NewPair[int, string](1)("foo"),
		tuple.NewPair[int, string](5)("bar"),
		tuple.NewPair[int, string](8)("baz"),
	}))

	// Output:
	// ([] [])
	// ([(1 foo)] [(8 bar)])
	// ([(1 foo)] [(8 baz)])
}

func ExampleToAscSlice() {
	fp.Pipe2(
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// []
	// [(bar 2) (baz 3) (foo 1)]
}

func ExampleToDescSlice() {
	fp.Pipe2(
		maps.ToDescSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe2(
		maps.ToDescSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// []
	// [(foo 1) (baz 3) (bar 2)]
}

func ExampleToSlice() {
	keySort := fp.On[tuple.Pair[string, int]](strLt)(tuple.Fst[string, int])

	fp.Pipe3(
		maps.ToSlice[string, int],
		slice.SortBy(keySort),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.ToSlice[string, int],
		slice.SortBy(keySort),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// []
	// [(bar 2) (baz 3) (foo 1)]
}

func ExampleUnion() {
	fp.Pipe3(
		maps.Union((maps.FromSlice([]tuple.Pair[string, int]{}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.Union((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
			tuple.NewPair[string, int]("baz")(3),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Union((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 2) (baz 3) (foo 1)]
}

func ExampleUnionWith() {
	fp.Pipe3(
		maps.UnionWith[string](operator.Add[int])((maps.FromSlice([]tuple.Pair[string, int]{}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UnionWith[string](operator.Add[int])((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
			tuple.NewPair[string, int]("baz")(3),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.UnionWith[string](operator.Add[int])((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 12) (baz 3) (foo 1)]
}

func ExampleUnionWithKey() {
	weightedAdd := fp.Curry3(func(k string, v1, v2 int) int { return len(k) * (v1 + v2) })

	fp.Pipe3(
		maps.UnionWithKey(weightedAdd)((maps.FromSlice([]tuple.Pair[string, int]{}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UnionWithKey(weightedAdd)((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
			tuple.NewPair[string, int]("baz")(3),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.UnionWithKey(weightedAdd)((maps.FromSlice([]tuple.Pair[string, int]{
			tuple.NewPair[string, int]("foo")(1),
			tuple.NewPair[string, int]("bar")(2),
		}))),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 2) (baz 3) (foo 1)]
	// [(bar 36) (baz 3) (foo 1)]
}

func ExampleUnions() {
	fp.Pipe3(
		maps.Unions[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]map[string]int{})

	fp.Pipe3(
		maps.Unions[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]map[string]int{
		{"foo": 1, "bar": 2},
		{"foo": 1, "bar": 2},
	})

	fp.Pipe3(
		maps.Unions[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]map[string]int{
		{"foo": 1, "bar": 2},
		{"foo": 2, "bar": 2},
	})

	fp.Pipe3(
		maps.Unions[string, int],
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)([]map[string]int{
		{"foo": 1, "bar": 2},
		{"a": 1, "b": 2},
	})

	// Output:
	// []
	// [(bar 2) (foo 1)]
	// [(bar 2) (foo 1)]
	// [(a 1) (b 2) (bar 2) (foo 1)]
}

func ExampleUnionsWith() {
	fp.Pipe3(
		maps.UnionsWith[string](operator.StrAppend),
		maps.ToAscSlice[string, string](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, string]]),
	)([]map[string]string{})

	fp.Pipe3(
		maps.UnionsWith[string](operator.StrAppend),
		maps.ToAscSlice[string, string](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, string]]),
	)([]map[string]string{
		{"foo": "1", "bar": "2"},
		{"foo": "1", "bar": "2"},
	})

	fp.Pipe3(
		maps.UnionsWith[string](operator.StrAppend),
		maps.ToAscSlice[string, string](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, string]]),
	)([]map[string]string{
		{"foo": "1", "bar": "2"},
		{"foo": "2", "bar": "2"},
	})

	fp.Pipe3(
		maps.UnionsWith[string](operator.StrAppend),
		maps.ToAscSlice[string, string](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, string]]),
	)([]map[string]string{
		{"foo": "1", "bar": "2"},
		{"a": "1", "b": "2"},
	})

	// Output:
	// []
	// [(bar 22) (foo 11)]
	// [(bar 22) (foo 12)]
	// [(a 1) (b 2) (bar 2) (foo 1)]
}

func ExampleUpdate() {
	doubleIfOdd := func(x int) option.Option[int] {
		if x%2 != 0 {
			return option.Some(x * 2)
		}
		return option.None[int]()
	}

	fp.Pipe3(
		maps.Update[string](doubleIfOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.Update[string](doubleIfOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.Update[string](doubleIfOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.Update[string](doubleIfOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// []
	// [(bar 10) (baz 3)]
	// [(baz 3) (foo 2)]
	// [(baz 3)]
}

func ExampleUpdateWithKey() {
	doubleIfOddIfKeyOdd := fp.Curry2(func(k string, x int) option.Option[int] {
		if len(k)%2 == 0 {
			return option.Some(x)
		}
		if x%2 != 0 {
			return option.Some(x * 2)
		}
		return option.None[int]()
	})

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foos"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foos")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foos"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foos")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateWithKey(doubleIfOddIfKeyOdd)("foo"),
		maps.ToAscSlice[string, int](strLt),
		fp.Inspect(printAny[[]tuple.Pair[string, int]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// []
	// [(bar 10) (baz 3)]
	// [(baz 3) (foos 1)]
	// [(baz 3) (foos 10)]
	// [(baz 3) (foo 2)]
	// [(baz 3)]
}

func ExampleUpdateLookupWithKey() {
	doubleIfOddIfKeyOdd := fp.Curry2(func(k string, x int) option.Option[int] {
		if len(k)%2 == 0 {
			return option.Some(x)
		}
		if x%2 != 0 {
			return option.Some(x * 2)
		}
		return option.None[int]()
	})

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foo"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{}))

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foo"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foos"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foos")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foos"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foos")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foo"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("baz")(3),
	}))

	fp.Pipe3(
		maps.UpdateLookupWithKey(doubleIfOddIfKeyOdd)("foo"),
		tuple.MapRight[option.Option[int]](maps.ToAscSlice[string, int](strLt)),
		fp.Inspect(printAny[tuple.Pair[option.Option[int], []tuple.Pair[string, int]]]),
	)(maps.FromSlice([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(10),
		tuple.NewPair[string, int]("baz")(3),
	}))

	// Output:
	// (None [])
	// (None [(bar 10) (baz 3)])
	// (Some 1 [(baz 3) (foos 1)])
	// (Some 10 [(baz 3) (foos 10)])
	// (Some 2 [(baz 3) (foo 2)])
	// (Some 10 [(baz 3)])
}
