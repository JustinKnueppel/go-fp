package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
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
