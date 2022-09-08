package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func ExampleFlip2() {
	fn := func(a, b int) int { return a - b }

	equal := fp.Flip2(fn)(2, 1) == fn(1, 2)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip3() {
	fn := func(a, b, c int) int { return a - b - c }

	equal := fp.Flip3(fn)(3, 2, 1) == fn(1, 2, 3)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip4() {
	fn := func(a, b, c, d int) int { return a - b - c - d }

	equal := fp.Flip4(fn)(4, 3, 2, 1) == fn(1, 2, 3, 4)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip5() {
	fn := func(a, b, c, d, e int) int { return a - b - c - d - e }

	equal := fp.Flip5(fn)(5, 4, 3, 2, 1) == fn(1, 2, 3, 4, 5)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip6() {
	fn := func(a, b, c, d, e, f int) int { return a - b - c - d - e - f }

	equal := fp.Flip6(fn)(6, 5, 4, 3, 2, 1) == fn(1, 2, 3, 4, 5, 6)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip7() {
	fn := func(a, b, c, d, e, f, g int) int { return a - b - c - d - e - f - g }

	equal := fp.Flip7(fn)(7, 6, 5, 4, 3, 2, 1) == fn(1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip8() {
	fn := func(a, b, c, d, e, f, g, h int) int { return a - b - c - d - e - f - g - h }

	equal := fp.Flip8(fn)(8, 7, 6, 5, 4, 3, 2, 1) == fn(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}

func ExampleFlip9() {
	fn := func(a, b, c, d, e, f, g, h, i int) int { return a - b - c - d - e - f - g - h - i }

	equal := fp.Flip9(fn)(9, 8, 7, 6, 5, 4, 3, 2, 1) == fn(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Flipped functions are equal: %v\n", equal)

	// Output:
	// Flipped functions are equal: true
}
