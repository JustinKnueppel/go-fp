package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func ExampleCurry2() {
	fn := func(a, b int) int { return a + b }

	equal := fp.Curry2(fn)(1)(2) == fn(1, 2)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry3() {
	fn := func(a, b, c int) int { return a + b + c }

	equal := fp.Curry3(fn)(1)(2)(3) == fn(1, 2, 3)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry4() {
	fn := func(a, b, c, d int) int { return a + b + c + d }

	equal := fp.Curry4(fn)(1)(2)(3)(4) == fn(1, 2, 3, 4)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry5() {
	fn := func(a, b, c, d, e int) int { return a + b + c + d + e }

	equal := fp.Curry5(fn)(1)(2)(3)(4)(5) == fn(1, 2, 3, 4, 5)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry6() {
	fn := func(a, b, c, d, e, f int) int { return a + b + c + d + e + f }

	equal := fp.Curry6(fn)(1)(2)(3)(4)(5)(6) == fn(1, 2, 3, 4, 5, 6)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry7() {
	fn := func(a, b, c, d, e, f, g int) int { return a + b + c + d + e + f + g }

	equal := fp.Curry7(fn)(1)(2)(3)(4)(5)(6)(7) == fn(1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry8() {
	fn := func(a, b, c, d, e, f, g, h int) int { return a + b + c + d + e + f + g + h }

	equal := fp.Curry8(fn)(1)(2)(3)(4)(5)(6)(7)(8) == fn(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}

func ExampleCurry9() {
	fn := func(a, b, c, d, e, f, g, h, i int) int { return a + b + c + d + e + f + g + h + i }

	equal := fp.Curry9(fn)(1)(2)(3)(4)(5)(6)(7)(8)(9) == fn(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Curried function equal: %v\n", equal)

	// Output:
	// Curried function equal: true
}
