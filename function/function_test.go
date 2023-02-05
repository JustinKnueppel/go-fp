package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

func ExampleConst() {
	fp.Pipe2(
		fp.Const[int, string](5),
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)("Hello world!")

	// Output:
	// 5
}

func ExampleId() {
	fp.Pipe2(
		fp.Id[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)(5)

	// Output:
	// 5
}

func ExampleInspect() {
	double := func(x int) int { return x * 2 }
	fp.Pipe2(
		double,
		fp.Inspect(func(x int) {
			fmt.Printf("X has value: %d\n", x)
		}),
	)(3)

	// Output:
	// X has value: 6
}

func ExampleOn() {
	length := func(s string) int { return len(s) }
	newPair := func(x int, s string) tuple.Pair[int, string] { return tuple.NewPair[int, string](x)(s) }

	fp.Pipe2(
		// Give the ability to sort these tuples based on the length of the string
		slice.SortBy(fp.On[tuple.Pair[int, string]](operator.Lt[int])(fp.Compose2(length, tuple.Snd[int, string]))),
		fp.Inspect(func(pairs []tuple.Pair[int, string]) {
			fmt.Println(pairs)
		}),
	)([]tuple.Pair[int, string]{newPair(3, "h"), newPair(2, "foo"), newPair(4, "ts")})

	// Output:
	// [(3 h) (4 ts) (2 foo)]
}

func ExampleTupled() {
	fp.Pipe2(
		fp.Tupled(operator.Add[int]),
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)(tuple.NewPair[int, int](4)(5))

	fp.Pipe3(
		slice.Zip[int, int]([]int{9, 8, 7}),
		slice.Map(fp.Tupled(operator.Add[int])),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// 9
	// [10 10 10]
}

func ExampleUntupled() {
	multiplePairs := func(p tuple.Pair[int, int]) int { return tuple.Fst(p) * tuple.Snd(p) }

	fp.Pipe2(
		fp.Untupled(multiplePairs)(10),
		fp.Inspect(func(product int) {
			fmt.Println(product)
		}),
	)(10)

	// Output:
	// 100
}
