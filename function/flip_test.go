package function_test

import (
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func TestFlip2(t *testing.T) {
	fn := func(a, b int) int { return a - b }

	if fp.Flip2(fn)(2, 1) != fn(1, 2) {
		t.Fatal("should flip the 2-arity function arguments")
	}
}

func TestFlip3(t *testing.T) {
	fn := func(a, b, c int) int { return a - b - c }

	if fp.Flip3(fn)(3, 2, 1) != fn(1, 2, 3) {
		t.Fatal("should flip the 3-arity function arguments")
	}
}

func TestFlip4(t *testing.T) {
	fn := func(a, b, c, d int) int { return a - b - c - d }

	if fp.Flip4(fn)(4, 3, 2, 1) != fn(1, 2, 3, 4) {
		t.Fatal("should flip the 4-arity function arguments")
	}
}

func TestFlip5(t *testing.T) {
	fn := func(a, b, c, d, e int) int { return a - b - c - d - e }

	if fp.Flip5(fn)(5, 4, 3, 2, 1) != fn(1, 2, 3, 4, 5) {
		t.Fatal("should flip the 5-arity function arguments")
	}
}

func TestFlip6(t *testing.T) {
	fn := func(a, b, c, d, e, f int) int { return a - b - c - d - e - f }

	if fp.Flip6(fn)(6, 5, 4, 3, 2, 1) != fn(1, 2, 3, 4, 5, 6) {
		t.Fatal("should flip the 6-arity function arguments")
	}
}

func TestFlip7(t *testing.T) {
	fn := func(a, b, c, d, e, f, g int) int { return a - b - c - d - e - f - g }

	if fp.Flip7(fn)(7, 6, 5, 4, 3, 2, 1) != fn(1, 2, 3, 4, 5, 6, 7) {
		t.Fatal("should flip the 7-arity function arguments")
	}
}

func TestFlip8(t *testing.T) {
	fn := func(a, b, c, d, e, f, g, h int) int { return a - b - c - d - e - f - g - h }

	if fp.Flip8(fn)(8, 7, 6, 5, 4, 3, 2, 1) != fn(1, 2, 3, 4, 5, 6, 7, 8) {
		t.Fatal("should flip the 8-arity function arguments")
	}
}

func TestFlip9(t *testing.T) {
	fn := func(a, b, c, d, e, f, g, h, i int) int { return a - b - c - d - e - f - g - h - i }

	if fp.Flip9(fn)(9, 8, 7, 6, 5, 4, 3, 2, 1) != fn(1, 2, 3, 4, 5, 6, 7, 8, 9) {
		t.Fatal("should flip the 9-arity function arguments")
	}
}
