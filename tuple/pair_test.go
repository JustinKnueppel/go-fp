package tuple_test

import (
	"fmt"
	"strconv"
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

func TestString(t *testing.T) {
	pair := tuple.NewPair[string, int]("foo")(5)
	if pair.String() != "(foo 5)" {
		t.Fatal()
	}
}

func ExampleNewPair() {
	fp.Inspect(func(pair tuple.Pair[int, string]) {
		fmt.Printf("First element: %d\n", tuple.Fst(pair))
		fmt.Printf("Second element: %s\n", tuple.Snd(pair))
	})(tuple.NewPair[int, string](1)("hello"))

	// Output:
	// First element: 1
	// Second element: hello
}

func ExampleFst() {
	fp.Pipe2(
		tuple.Fst[int, string],
		fp.Inspect(func(x int) {
			fmt.Printf("First element: %d\n", x)
		}),
	)(tuple.NewPair[int, string](1)("hello"))

	// Output:
	// First element: 1
}

func ExampleSnd() {
	fp.Pipe2(
		tuple.Snd[int, string],
		fp.Inspect(func(s string) {
			fmt.Printf("Second element: %s\n", s)
		}),
	)(tuple.NewPair[int, string](1)("hello"))

	// Output:
	// Second element: hello
}

func ExampleMapLeft() {
	fp.Pipe2(
		tuple.MapLeft[int, string](fp.Compose2(strconv.Itoa, operator.Add(1))),
		fp.Inspect(func(pair tuple.Pair[string, string]) {
			fmt.Println(pair)
		}),
	)(tuple.NewPair[int, string](1)("foo"))

	fp.Pipe2(
		tuple.MapLeft[[]int, string](slice.Foldl(operator.Add[int])(0)),
		fp.Inspect(func(pair tuple.Pair[int, string]) {
			fmt.Println(pair)
		}),
	)(tuple.NewPair[[]int, string]([]int{1, 2, 3, 4})("bar"))

	// Output:
	// (2 foo)
	// (10 bar)
}

func ExampleMapRight() {
	fp.Pipe2(
		tuple.MapRight[string](fp.Compose2(strconv.Itoa, operator.Add(1))),
		fp.Inspect(func(pair tuple.Pair[string, string]) {
			fmt.Println(pair)
		}),
	)(tuple.NewPair[string, int]("foo")(1))

	fp.Pipe2(
		tuple.MapRight[string](slice.Foldl(operator.Add[int])(0)),
		fp.Inspect(func(pair tuple.Pair[string, int]) {
			fmt.Println(pair)
		}),
	)(tuple.NewPair[string, []int]("bar")([]int{1, 2, 3, 4}))

	// Output:
	// (foo 2)
	// (bar 10)
}

func ExamplePattern() {
	pair := tuple.NewPair[int, string](1)("foo")
	idx, val := tuple.Pattern(pair)
	fmt.Printf("%d: %s", idx, val)

	// Output:
	// 1: foo
}
