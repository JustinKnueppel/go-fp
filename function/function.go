package function

// Decrement returns the given number decremented.
func Decrement(x int) int {
	return x - 1
}

// Increment returns the given number incremented.
func Increment(x int) int {
	return x + 1
}

// Inspect calls the given function with the provided value
// and returns the unchanged value.
func Inspect[T any](fn func(T)) func(T) T {
	return func(t T) T {
		fn(t)
		return t
	}
}
