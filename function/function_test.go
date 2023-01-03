package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

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
		slice.Sort(fp.On[tuple.Pair[int, string]](operator.Lt[int])(fp.Compose2(length, tuple.Snd[int, string]))),
		fp.Inspect(func(pairs []tuple.Pair[int, string]) {
			fmt.Println(slice.Map(pairToString[int, string])(pairs))
		}),
	)([]tuple.Pair[int, string]{newPair(3, "h"), newPair(2, "foo"), newPair(4, "ts")})

	// Output:
	// [(3 h) (4 ts) (2 foo)]
}

func pairToString[T, U any](pair tuple.Pair[T, U]) string {
	return fmt.Sprintf("(%v %v)", tuple.Fst(pair), tuple.Snd(pair))
}
