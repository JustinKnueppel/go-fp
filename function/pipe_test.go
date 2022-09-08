package function_test

import (
	"fmt"
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func TestPipe2(t *testing.T) {
	t.Run("composes left to right", func(t *testing.T) {
		initial := 2
		expected := 6

		if fp.Pipe2(add1, double)(initial) != expected {
			t.Fail()
		}
	})

	t.Run("composes across types", func(t *testing.T) {
		initial := "hello"
		expected := 10

		if fp.Pipe2(strlen, double)(initial) != expected {
			t.Fail()
		}
	})
}

func ExamplePipe2() {
	res := fp.Pipe2(double, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 5
}

func ExamplePipe3() {
	res := fp.Pipe3(double, double, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 9
}

func ExamplePipe4() {
	res := fp.Pipe4(double, double, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 10
}

func TestPipeNested(t *testing.T) {
	add1 := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	triple := func(x int) int { return x * 3 }

	result4 := fp.Pipe4(triple, add1, double, add1)(5)
	result2 := fp.Pipe2(fp.Pipe2(triple, add1), fp.Pipe2(double, add1))(5)

	if result2 != result4 {
		t.Fatal("nesting compose should be equal to larger compose")
	}
}

func ExamplePipe5() {
	res := fp.Pipe5(double, double, add1, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 11
}

func ExamplePipe6() {
	res := fp.Pipe6(double, double, add1, add1, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 12
}

func ExamplePipe7() {
	res := fp.Pipe7(double, double, add1, add1, add1, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 13
}

func ExamplePipe8() {
	res := fp.Pipe8(double, double, add1, add1, add1, add1, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 14
}

func ExamplePipe9() {
	res := fp.Pipe9(double, double, add1, add1, add1, add1, add1, add1, add1)(2)
	fmt.Printf("Result: %d\n", res)

	// Output:
	// Result: 15
}
