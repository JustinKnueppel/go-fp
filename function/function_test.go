package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func ExampleDecrement() {
	fp.Pipe2(
		fp.Decrement,
		fp.Inspect(func(x int) {
			fmt.Printf("Decremented value: %d\n", x)
		}),
	)(1)

	// Output:
	// Decremented value: 0
}

func ExampleIncrement() {
	fp.Pipe2(
		fp.Increment,
		fp.Inspect(func(x int) {
			fmt.Printf("Incremented value: %d\n", x)
		}),
	)(1)

	// Output:
	// Incremented value: 2
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
