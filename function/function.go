package function

// Inspect calls the given function with the provided value
// and returns the unchanged value.
func Inspect[T any](fn func(T)) func(T) T {
	return func(t T) T {
		fn(t)
		return t
	}
}
