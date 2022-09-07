package function

// Flip2 returns an equivalent 2-arity function, but with the argument order reversed.
func Flip2[A, B, C any](fn func(A, B) C) func(B, A) C {
	return func(b B, a A) C {
		return fn(a, b)
	}
}

// Flip3 returns an equivalent 3-arity function, but with the argument order reversed.
func Flip3[A, B, C, D any](fn func(A, B, C) D) func(C, B, A) D {
	return func(c C, b B, a A) D {
		return fn(a, b, c)
	}
}

// Flip4 returns an equivalent 4-arity function, but with the argument order reversed.
func Flip4[A, B, C, D, E any](fn func(A, B, C, D) E) func(D, C, B, A) E {
	return func(d D, c C, b B, a A) E {
		return fn(a, b, c, d)
	}
}

// Flip5 returns an equivalent 5-arity function, but with the argument order reversed.
func Flip5[A, B, C, D, E, F any](fn func(A, B, C, D, E) F) func(E, D, C, B, A) F {
	return func(e E, d D, c C, b B, a A) F {
		return fn(a, b, c, d, e)
	}
}

// Flip6 returns an equivalent 6-arity function, but with the argument order reversed.
func Flip6[A, B, C, D, E, F, G any](fn func(A, B, C, D, E, F) G) func(F, E, D, C, B, A) G {
	return func(f F, e E, d D, c C, b B, a A) G {
		return fn(a, b, c, d, e, f)
	}
}

// Flip7 returns an equivalent 7-arity function, but with the argument order reversed.
func Flip7[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F, G) H) func(G, F, E, D, C, B, A) H {
	return func(g G, f F, e E, d D, c C, b B, a A) H {
		return fn(a, b, c, d, e, f, g)
	}
}

// Flip8 returns an equivalent 8-arity function, but with the argument order reversed.
func Flip8[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G, H) I) func(H, G, F, E, D, C, B, A) I {
	return func(h H, g G, f F, e E, d D, c C, b B, a A) I {
		return fn(a, b, c, d, e, f, g, h)
	}
}

// Flip9 returns an equivalent 9-arity function, but with the argument order reversed.
func Flip9[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H, I) J) func(I, H, G, F, E, D, C, B, A) J {
	return func(i I, h H, g G, f F, e E, d D, c C, b B, a A) J {
		return fn(a, b, c, d, e, f, g, h, i)
	}
}
