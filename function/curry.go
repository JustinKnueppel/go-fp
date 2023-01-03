package function

// Curry2 returns the curried form of the 2-arity funcion.
func Curry2[A, B, C any](fn func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return fn(a, b)
		}
	}
}

// Curry3 returns the curried form of the 3-arity funcion.
func Curry3[A, B, C, D any](fn func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return Curry2(func(b B, c C) D {
			return fn(a, b, c)
		})
	}
}

// Curry4 returns the curried form of the 4-arity funcion.
func Curry4[A, B, C, D, E any](fn func(A, B, C, D) E) func(A) func(B) func(C) func(D) E {
	return func(a A) func(B) func(C) func(D) E {
		return Curry3(func(b B, c C, d D) E {
			return fn(a, b, c, d)
		})
	}
}

// Curry5 returns the curried form of the 5-arity funcion.
func Curry5[A, B, C, D, E, F any](fn func(A, B, C, D, E) F) func(A) func(B) func(C) func(D) func(E) F {
	return func(a A) func(B) func(C) func(D) func(E) F {
		return Curry4(func(b B, c C, d D, e E) F {
			return fn(a, b, c, d, e)
		})
	}
}

// Curry6 returns the curried form of the 6-arity funcion.
func Curry6[A, B, C, D, E, F, G any](fn func(A, B, C, D, E, F) G) func(A) func(B) func(C) func(D) func(E) func(F) G {
	return func(a A) func(B) func(C) func(D) func(E) func(F) G {
		return Curry5(func(b B, c C, d D, e E, f F) G {
			return fn(a, b, c, d, e, f)
		})
	}
}

// Curry7 returns the curried form of the 7-arity funcion.
func Curry7[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F, G) H) func(A) func(B) func(C) func(D) func(E) func(F) func(G) H {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) H {
		return Curry6(func(b B, c C, d D, e E, f F, g G) H {
			return fn(a, b, c, d, e, f, g)
		})
	}
}

// Curry8 returns the curried form of the 8-arity funcion.
func Curry8[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G, H) I) func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I {
		return Curry7(func(b B, c C, d D, e E, f F, g G, h H) I {
			return fn(a, b, c, d, e, f, g, h)
		})
	}
}

// Curry9 returns the curried form of the 9-arity funcion.
func Curry9[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H, I) J) func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J {
		return Curry8(func(b B, c C, d D, e E, f F, g G, h H, i I) J {
			return fn(a, b, c, d, e, f, g, h, i)
		})
	}
}
