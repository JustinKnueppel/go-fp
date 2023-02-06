package operator

type Number interface {
	Float | Integer
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// StrAppend appends s2 to s1 (s1 + s2 in Go).
func StrAppend(s1 string) func(s2 string) string {
	return func(s2 string) string {
		return s1 + s2
	}
}

// Add takes two numbers and returns their sum.
func Add[T Number](x T) func(y T) T {
	return func(y T) T {
		return x + y
	}
}

// Subtract takes two numbers and returns their difference.
func Subtract[T Number](x T) func(y T) T {
	return func(y T) T {
		return x - y
	}
}

// Multiply takes two numbers and returns their product.
func Multiply[T Number](x T) func(y T) T {
	return func(y T) T {
		return x * y
	}
}

// Divide takes two numbers and returns their division.
func Divide[T Number](x T) func(y T) T {
	return func(y T) T {
		return x / y
	}
}

// Mod takes two numbers and returns their modulo.
func Mod(x int) func(mod int) int {
	return func(y int) int {
		return x % y
	}
}

// Inc increments the number by 1.
func Inc[T Number](x T) T {
	return x + 1
}

// Dec decrements the number by 1.
func Dec[T Number](x T) T {
	return x - 1
}

// Eq compares x and y for equality.
func Eq[T comparable](x T) func(y T) bool {
	return func(y T) bool {
		return x == y
	}
}

// Neq compares x and y for inequality.
func Neq[T comparable](x T) func(y T) bool {
	return func(y T) bool {
		return x != y
	}
}

// Gt determines if x > y.
func Gt[T Number](x T) func(y T) bool {
	return func(y T) bool {
		return x > y
	}
}

// Geq determines if x >= y.
func Geq[T Number](x T) func(y T) bool {
	return func(y T) bool {
		return x >= y
	}
}

// Lt determines if x < y.
func Lt[T Number](x T) func(y T) bool {
	return func(y T) bool {
		return x < y
	}
}

// Leq determines if x <= y.
func Leq[T Number](x T) func(y T) bool {
	return func(y T) bool {
		return x <= y
	}
}

// And returns true if both x and y are true.
func And(x bool) func(y bool) bool {
	return func(y bool) bool {
		return x && y
	}
}

// Or returns true if x, y, or both are true.
func Or(x bool) func(y bool) bool {
	return func(y bool) bool {
		return x || y
	}
}

// Not returns the negation of the boolean provided.
func Not(x bool) bool {
	return !x
}

// BitAnd returns the bitwise and operation on x and y.
func BitAnd[T Integer](x T) func(y T) T {
	return func(y T) T {
		return x & y
	}
}

// BitOr returns the bitwise or operation on x and y.
func BitOr[T Integer](x T) func(y T) T {
	return func(y T) T {
		return x | y
	}
}

// BitXor returns the bitwise xor operation on x and y.
func BitXor[T Integer](x T) func(y T) T {
	return func(y T) T {
		return x ^ y
	}
}

// RShift returns x binary right shifted by y.
func RShift[T Integer](x T) func(y T) T {
	return func(y T) T {
		return x >> y
	}
}

// LShift returns x binary left shifted by y.
func LShift[T Integer](x T) func(y T) T {
	return func(y T) T {
		return x << y
	}
}
