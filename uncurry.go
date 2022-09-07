package fp

// Uncurry2 takes the curried form a 2-arity function and returns the uncurried form.
func Uncurry2[A, B, C any](fn func(A) func(B) C) func(A, B) C {
	return func(a A, b B) C {
		return fn(a)(b)
	}
}

// Uncurry3 takes the curried form a 3-arity function and returns the uncurried form.
func Uncurry3[A, B, C, D any](fn func(A) func(B) func(C) D) func(A, B, C) D {
	return func(a A, b B, c C) D {
		return fn(a)(b)(c)
	}
}

// Uncurry4 takes the curried form a 4-arity function and returns the uncurried form.
func Uncurry4[A, B, C, D, E any](fn func(A) func(B) func(C) func(D) E) func(A, B, C, D) E {
	return func(a A, b B, c C, d D) E {
		return fn(a)(b)(c)(d)
	}
}

// Uncurry5 takes the curried form a 5-arity function and returns the uncurried form.
func Uncurry5[A, B, C, D, E, F any](fn func(A) func(B) func(C) func(D) func(E) F) func(A, B, C, D, E) F {
	return func(a A, b B, c C, d D, e E) F {
		return fn(a)(b)(c)(d)(e)
	}
}

// Uncurry6 takes the curried form a 6-arity function and returns the uncurried form.
func Uncurry6[A, B, C, D, E, F, G any](fn func(A) func(B) func(C) func(D) func(E) func(F) G) func(A, B, C, D, E, F) G {
	return func(a A, b B, c C, d D, e E, f F) G {
		return fn(a)(b)(c)(d)(e)(f)
	}
}

// Uncurry7 takes the curried form a 7-arity function and returns the uncurried form.
func Uncurry7[A, B, C, D, E, F, G, H any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) H) func(A, B, C, D, E, F, G) H {
	return func(a A, b B, c C, d D, e E, f F, g G) H {
		return fn(a)(b)(c)(d)(e)(f)(g)
	}
}

// Uncurry8 takes the curried form a 8-arity function and returns the uncurried form.
func Uncurry8[A, B, C, D, E, F, G, H, I any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I) func(A, B, C, D, E, F, G, H) I {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) I {
		return fn(a)(b)(c)(d)(e)(f)(g)(h)
	}
}

// Uncurry9 takes the curried form a 9-arity function and returns the uncurried form.
func Uncurry9[A, B, C, D, E, F, G, H, I, J any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J) func(A, B, C, D, E, F, G, H, I) J {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) J {
		return fn(a)(b)(c)(d)(e)(f)(g)(h)(i)
	}
}
