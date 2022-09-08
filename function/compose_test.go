package function_test

import (
	"fmt"
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func add1(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func strlen(s string) int {
	return len(s)
}

func TestCompose2(t *testing.T) {
	t.Run("composes right to left", func(t *testing.T) {
		initial := 2
		expected := 6

		if fp.Compose2(double, add1)(initial) != expected {
			t.Fail()
		}
	})

	t.Run("composes across types", func(t *testing.T) {
		initial := "hello"
		expected := 10

		if fp.Compose2(double, strlen)(initial) != expected {
			t.Fail()
		}
	})
}

func ExampleCompose2() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose2(double, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 8
}

func ExampleCompose3() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose3(double, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 10
}

func ExampleCompose4() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose4(double, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 12
}

func TestComposeNested(t *testing.T) {
	add1 := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	triple := func(x int) int { return x * 3 }

	result4 := fp.Compose4(add1, double, add1, triple)(5)
	result2 := fp.Compose2(fp.Compose2(add1, double), fp.Compose2(add1, triple))(5)

	if result2 != result4 {
		t.Fatal("nesting compose should be equal to larger compose")
	}
}

func ExampleCompose5() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose5(double, add1, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 14
}

func ExampleCompose6() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose6(double, add1, add1, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 16
}

func ExampleCompose7() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose7(double, add1, add1, add1, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 18
}

func ExampleCompose8() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose8(double, add1, add1, add1, add1, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 20
}

func ExampleCompose9() {
	double := func(x int) int { return x * 2 }
	add1 := func(x int) int { return x + 1 }

	result := fp.Compose9(double, add1, add1, add1, add1, add1, add1, add1, add1)(3)
	fmt.Printf("Final result: %d\n", result)

	// Output:
	// Final result: 22
}
