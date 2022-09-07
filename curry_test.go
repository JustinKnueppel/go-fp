package fp_test

import (
	"testing"

	"github.com/JustinKnueppel/go-fp"
)

func TestCurry2(t *testing.T) {
	fn := func(a, b int) int { return a + b }

	if fp.Curry2(fn)(1)(2) != fn(1, 2) {
		t.Fatal("should curry the 2-arity function")
	}
}

func TestCurry3(t *testing.T) {
	fn := func(a, b, c int) int { return a + b + c }

	if fp.Curry3(fn)(1)(2)(3) != fn(1, 2, 3) {
		t.Fatal("should curry the 3-arity function")
	}
}

func TestCurry4(t *testing.T) {
	fn := func(a, b, c, d int) int { return a + b + c + d }

	if fp.Curry4(fn)(1)(2)(3)(4) != fn(1, 2, 3, 4) {
		t.Fatal("should curry the 4-arity function")
	}
}

func TestCurry5(t *testing.T) {
	fn := func(a, b, c, d, e int) int { return a + b + c + d + e }

	if fp.Curry5(fn)(1)(2)(3)(4)(5) != fn(1, 2, 3, 4, 5) {
		t.Fatal("should curry the 5-arity function")
	}
}

func TestCurry6(t *testing.T) {
	fn := func(a, b, c, d, e, f int) int { return a + b + c + d + e + f }

	if fp.Curry6(fn)(1)(2)(3)(4)(5)(6) != fn(1, 2, 3, 4, 5, 6) {
		t.Fatal("should curry the 6-arity function")
	}
}

func TestCurry7(t *testing.T) {
	fn := func(a, b, c, d, e, f, g int) int { return a + b + c + d + e + f + g }

	if fp.Curry7(fn)(1)(2)(3)(4)(5)(6)(7) != fn(1, 2, 3, 4, 5, 6, 7) {
		t.Fatal("should curry the 7-arity function")
	}
}

func TestCurry8(t *testing.T) {
	fn := func(a, b, c, d, e, f, g, h int) int { return a + b + c + d + e + f + g + h }

	if fp.Curry8(fn)(1)(2)(3)(4)(5)(6)(7)(8) != fn(1, 2, 3, 4, 5, 6, 7, 8) {
		t.Fatal("should curry the 8-arity function")
	}
}

func TestCurry9(t *testing.T) {
	fn := func(a, b, c, d, e, f, g, h, i int) int { return a + b + c + d + e + f + g + h + i }

	if fp.Curry9(fn)(1)(2)(3)(4)(5)(6)(7)(8)(9) != fn(1, 2, 3, 4, 5, 6, 7, 8, 9) {
		t.Fatal("should curry the 9-arity function")
	}
}
