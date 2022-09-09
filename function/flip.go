package function

// Flip2 returns an equivalent 2-arity curried function, but with the argument order reversed.
func Flip2[A, B, C any](fn func(A) func(B) C) func(B) func(A) C {
	return Curry2(func(b B, a A) C {
		return fn(a)(b)
	})
}

// Flip3 returns an equivalent 3-arity curried function, but with the argument order reversed.
func Flip3[A, B, C, D any](fn func(A) func(B) func(C) D) func(C) func(B) func(A) D {
	return Curry3(func(c C, b B, a A) D {
		return fn(a)(b)(c)
	})
}

// Flip4 returns an equivalent 4-arity curried function, but with the argument order reversed.
func Flip4[A, B, C, D, E any](fn func(A) func(B) func(C) func(D) E) func(D) func(C) func(B) func(A) E {
	return Curry4(func(d D, c C, b B, a A) E {
		return fn(a)(b)(c)(d)
	})
}

// Flip5 returns an equivalent 5-arity curried function, but with the argument order reversed.
func Flip5[A, B, C, D, E, F any](fn func(A) func(B) func(C) func(D) func(E) F) func(E) func(D) func(C) func(B) func(A) F {
	return Curry5(func(e E, d D, c C, b B, a A) F {
		return fn(a)(b)(c)(d)(e)
	})
}

// Flip6 returns an equivalent 6-arity curried function, but with the argument order reversed.
func Flip6[A, B, C, D, E, F, G any](fn func(A) func(B) func(C) func(D) func(E) func(F) G) func(F) func(E) func(D) func(C) func(B) func(A) G {
	return Curry6(func(f F, e E, d D, c C, b B, a A) G {
		return fn(a)(b)(c)(d)(e)(f)
	})
}

// Flip7 returns an equivalent 7-arity curried function, but with the argument order reversed.
func Flip7[A, B, C, D, E, F, G, H any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) H) func(G) func(F) func(E) func(D) func(C) func(B) func(A) H {
	return Curry7(func(g G, f F, e E, d D, c C, b B, a A) H {
		return fn(a)(b)(c)(d)(e)(f)(g)
	})
}

// Flip8 returns an equivalent 8-arity curried function, but with the argument order reversed.
func Flip8[A, B, C, D, E, F, G, H, I any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I) func(H) func(G) func(F) func(E) func(D) func(C) func(B) func(A) I {
	return Curry8(func(h H, g G, f F, e E, d D, c C, b B, a A) I {
		return fn(a)(b)(c)(d)(e)(f)(g)(h)
	})
}

// Flip9 returns an equivalent 9-arity curried function, but with the argument order reversed.
func Flip9[A, B, C, D, E, F, G, H, I, J any](fn func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J) func(I) func(H) func(G) func(F) func(E) func(D) func(C) func(B) func(A) J {
	return Curry9(func(i I, h H, g G, f F, e E, d D, c C, b B, a A) J {
		return fn(a)(b)(c)(d)(e)(f)(g)(h)(i)
	})
}
