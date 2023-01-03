package tuple_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/tuple"
)

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
