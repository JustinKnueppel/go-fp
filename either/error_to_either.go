package either

// ErrorToEither0 returns a 0-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither0[T any](fn func() (T, error)) func() Either[error, T] {
	return func() Either[error, T] {
		t, err := fn()
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither1 returns a 1-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither1[A, T any](fn func(A) (T, error)) func(A) Either[error, T] {
	return func(a A) Either[error, T] {
		t, err := fn(a)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither2 returns a 2-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither2[A, B, T any](fn func(A, B) (T, error)) func(A, B) Either[error, T] {
	return func(a A, b B) Either[error, T] {
		t, err := fn(a, b)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither3 returns a 3-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither3[A, B, C, T any](fn func(A, B, C) (T, error)) func(A, B, C) Either[error, T] {
	return func(a A, b B, c C) Either[error, T] {
		t, err := fn(a, b, c)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither4 returns a 4-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither4[A, B, C, D, T any](fn func(A, B, C, D) (T, error)) func(A, B, C, D) Either[error, T] {
	return func(a A, b B, c C, d D) Either[error, T] {
		t, err := fn(a, b, c, d)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither5 returns a 5-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither5[A, B, C, D, E, T any](fn func(A, B, C, D, E) (T, error)) func(A, B, C, D, E) Either[error, T] {
	return func(a A, b B, c C, d D, e E) Either[error, T] {
		t, err := fn(a, b, c, d, e)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither6 returns a 6-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither6[A, B, C, D, E, F, T any](fn func(A, B, C, D, E, F) (T, error)) func(A, B, C, D, E, F) Either[error, T] {
	return func(a A, b B, c C, d D, e E, f F) Either[error, T] {
		t, err := fn(a, b, c, d, e, f)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither7 returns a 7-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither7[A, B, C, D, E, F, G, T any](fn func(A, B, C, D, E, F, G) (T, error)) func(A, B, C, D, E, F, G) Either[error, T] {
	return func(a A, b B, c C, d D, e E, f F, g G) Either[error, T] {
		t, err := fn(a, b, c, d, e, f, g)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither8 returns a 8-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither8[A, B, C, D, E, F, G, H, T any](fn func(A, B, C, D, E, F, G, H) (T, error)) func(A, B, C, D, E, F, G, H) Either[error, T] {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) Either[error, T] {
		t, err := fn(a, b, c, d, e, f, g, h)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}

// ErrorToEither9 returns a 9-arity function that returns a value or an error, and
// converts the return value to an Either[error, T].
func ErrorToEither9[A, B, C, D, E, F, G, H, I, T any](fn func(A, B, C, D, E, F, G, H, I) (T, error)) func(A, B, C, D, E, F, G, H, I) Either[error, T] {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) Either[error, T] {
		t, err := fn(a, b, c, d, e, f, g, h, i)
		if err != nil {
			return Left[error, T](err)
		}
		return Right[error](t)
	}
}
