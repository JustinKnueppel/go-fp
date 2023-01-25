package operator_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
)

func ExampleStrAppend() {
	fp.Pipe2(
		operator.StrAppend(""),
		fp.Inspect(func(s string) {
			fmt.Println(s)
		}),
	)("bar")

	fp.Pipe2(
		operator.StrAppend("foo"),
		fp.Inspect(func(s string) {
			fmt.Println(s)
		}),
	)("")

	fp.Pipe2(
		operator.StrAppend("foo"),
		fp.Inspect(func(s string) {
			fmt.Println(s)
		}),
	)("bar")

	// Output:
	// bar
	// foo
	// foobar
}

func ExampleAdd() {
	fp.Pipe2(
		operator.Add(5),
		fp.Inspect(func(x int) {
			fmt.Printf("5 + 3 = %d\n", x)
		}),
	)(3)

	// Output:
	// 5 + 3 = 8
}

func ExampleSubtract() {
	fp.Pipe2(
		operator.Subtract(5),
		fp.Inspect(func(x int) {
			fmt.Printf("5 - 3 = %d\n", x)
		}),
	)(3)

	// Output:
	// 5 - 3 = 2
}

func ExampleMultiply() {
	fp.Pipe2(
		operator.Multiply(5),
		fp.Inspect(func(x int) {
			fmt.Printf("5 * 3 = %d\n", x)
		}),
	)(3)

	// Output:
	// 5 * 3 = 15
}

func ExampleDivide() {
	fp.Pipe2(
		operator.Divide(10),
		fp.Inspect(func(x int) {
			fmt.Printf("10 / 2 = %d\n", x)
		}),
	)(2)

	// Output:
	// 10 / 2 = 5
}

func ExampleMod() {
	fp.Pipe2(
		operator.Mod(5),
		fp.Inspect(func(x int) {
			fmt.Printf("5 %% 3 = %d\n", x)
		}),
	)(3)

	// Output:
	// 5 % 3 = 2
}

func ExampleInc() {
	fp.Pipe2(
		operator.Inc[int],
		fp.Inspect(func(x int) {
			fmt.Printf("3++ = %d\n", x)
		}),
	)(3)

	// Output:
	// 3++ = 4
}

func ExampleDec() {
	fp.Pipe2(
		operator.Dec[int],
		fp.Inspect(func(x int) {
			fmt.Printf("3-- = %d\n", x)
		}),
	)(3)

	// Output:
	// 3-- = 2
}

func ExampleEq() {
	fp.Pipe2(
		operator.Eq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 == 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Eq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 == 2 = %v\n", x)
		}),
	)(2)

	// Output:
	// 2 == 3 = false
	// 2 == 2 = true
}

func ExampleNeq() {
	fp.Pipe2(
		operator.Neq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 != 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Neq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 != 2 = %v\n", x)
		}),
	)(2)

	// Output:
	// 2 != 3 = true
	// 2 != 2 = false
}

func ExampleGt() {
	fp.Pipe2(
		operator.Gt(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 > 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Gt(3),
		fp.Inspect(func(x bool) {
			fmt.Printf("3 > 2 = %v\n", x)
		}),
	)(2)

	fp.Pipe2(
		operator.Gt(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 > 2 = %v\n", x)
		}),
	)(2)

	// Output:
	// 2 > 3 = false
	// 3 > 2 = true
	// 2 > 2 = false
}

func ExampleGeq() {
	fp.Pipe2(
		operator.Geq(4),
		fp.Inspect(func(x bool) {
			fmt.Printf("4 >= 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Geq(3),
		fp.Inspect(func(x bool) {
			fmt.Printf("3 >= 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Geq(3),
		fp.Inspect(func(x bool) {
			fmt.Printf("3 >= 4 = %v\n", x)
		}),
	)(4)

	// Output:
	// 4 >= 3 = true
	// 3 >= 3 = true
	// 3 >= 4 = false
}

func ExampleLt() {
	fp.Pipe2(
		operator.Lt(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 < 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Lt(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 < 2 = %v\n", x)
		}),
	)(2)

	fp.Pipe2(
		operator.Lt(3),
		fp.Inspect(func(x bool) {
			fmt.Printf("3 < 2 = %v\n", x)
		}),
	)(2)

	// Output:
	// 2 < 3 = true
	// 2 < 2 = false
	// 3 < 2 = false
}

func ExampleLeq() {
	fp.Pipe2(
		operator.Leq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 <= 3 = %v\n", x)
		}),
	)(3)

	fp.Pipe2(
		operator.Leq(2),
		fp.Inspect(func(x bool) {
			fmt.Printf("2 <= 2 = %v\n", x)
		}),
	)(2)

	fp.Pipe2(
		operator.Leq(3),
		fp.Inspect(func(x bool) {
			fmt.Printf("3 <= 2 = %v\n", x)
		}),
	)(2)

	// Output:
	// 2 <= 3 = true
	// 2 <= 2 = true
	// 3 <= 2 = false
}

func ExampleAnd() {
	fp.Pipe2(
		operator.And(true),
		fp.Inspect(func(x bool) {
			fmt.Printf("true && true = %v\n", x)
		}),
	)(true)

	fp.Pipe2(
		operator.And(true),
		fp.Inspect(func(x bool) {
			fmt.Printf("true && false = %v\n", x)
		}),
	)(false)

	fp.Pipe2(
		operator.And(false),
		fp.Inspect(func(x bool) {
			fmt.Printf("false && false = %v\n", x)
		}),
	)(false)

	// Output:
	// true && true = true
	// true && false = false
	// false && false = false
}

func ExampleOr() {
	fp.Pipe2(
		operator.Or(true),
		fp.Inspect(func(x bool) {
			fmt.Printf("true || true = %v\n", x)
		}),
	)(true)

	fp.Pipe2(
		operator.Or(true),
		fp.Inspect(func(x bool) {
			fmt.Printf("true || false = %v\n", x)
		}),
	)(false)

	fp.Pipe2(
		operator.Or(false),
		fp.Inspect(func(x bool) {
			fmt.Printf("false || false = %v\n", x)
		}),
	)(false)

	// Output:
	// true || true = true
	// true || false = true
	// false || false = false
}

func ExampleNot() {
	fp.Pipe2(
		operator.Not,
		fp.Inspect(func(x bool) {
			fmt.Printf("!true = %v\n", x)
		}),
	)(true)

	fp.Pipe2(
		operator.Not,
		fp.Inspect(func(x bool) {
			fmt.Printf("!false = %v\n", x)
		}),
	)(false)

	// Output:
	// !true = false
	// !false = true
}

func ExampleBitAnd() {
	fp.Pipe2(
		operator.BitAnd(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 & 1 = %v\n", x)
		}),
	)(1)

	fp.Pipe2(
		operator.BitAnd(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 & 0 = %v\n", x)
		}),
	)(0)

	fp.Pipe2(
		operator.BitAnd(0),
		fp.Inspect(func(x int) {
			fmt.Printf("0 & 0 = %v\n", x)
		}),
	)(0)

	// Output:
	// 1 & 1 = 1
	// 1 & 0 = 0
	// 0 & 0 = 0
}

func ExampleBitOr() {
	fp.Pipe2(
		operator.BitOr(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 | 1 = %v\n", x)
		}),
	)(1)

	fp.Pipe2(
		operator.BitOr(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 | 0 = %v\n", x)
		}),
	)(0)

	fp.Pipe2(
		operator.BitOr(0),
		fp.Inspect(func(x int) {
			fmt.Printf("0 | 0 = %v\n", x)
		}),
	)(0)

	// Output:
	// 1 | 1 = 1
	// 1 | 0 = 1
	// 0 | 0 = 0
}

func ExampleBitXor() {
	fp.Pipe2(
		operator.BitXor(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 ^ 1 = %v\n", x)
		}),
	)(1)

	fp.Pipe2(
		operator.BitXor(1),
		fp.Inspect(func(x int) {
			fmt.Printf("1 ^ 0 = %v\n", x)
		}),
	)(0)

	fp.Pipe2(
		operator.BitXor(0),
		fp.Inspect(func(x int) {
			fmt.Printf("0 ^ 0 = %v\n", x)
		}),
	)(0)

	// Output:
	// 1 ^ 1 = 0
	// 1 ^ 0 = 1
	// 0 ^ 0 = 0
}

func ExampleRShift() {
	fp.Pipe2(
		operator.RShift(4),
		fp.Inspect(func(x int) {
			fmt.Printf("4 >> 1 = %d\n", x)
		}),
	)(1)

	// Output:
	// 4 >> 1 = 2
}

func ExampleLShift() {
	fp.Pipe2(
		operator.LShift(4),
		fp.Inspect(func(x int) {
			fmt.Printf("4 << 1 = %d\n", x)
		}),
	)(1)

	// Output:
	// 4 << 1 = 8
}
