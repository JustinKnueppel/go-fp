package fp

// Pipe2 returns a left-to-right composition of 2 functions.
func Pipe2[A, B, C any](ab func(A) B, bc func(B) C) func(A) C {
	return func(a A) C {
		return bc(ab(a))
	}
}

// Pipe3 returns a left-to-right composition of 3 functions.
func Pipe3[A, B, C, D any](ab func(A) B, bc func(B) C, cd func(C) D) func(A) D {
	return func(a A) D {
		return cd(bc(ab(a)))
	}
}

// Pipe4 returns a left-to-right composition of 4 functions.
func Pipe4[A, B, C, D, E any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E) func(A) E {
	return func(a A) E {
		return de(cd(bc(ab(a))))
	}
}

// Pipe5 returns a left-to-right composition of 5 functions.
func Pipe5[A, B, C, D, E, F any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F) func(A) F {
	return func(a A) F {
		return ef(de(cd(bc(ab(a)))))
	}
}

// Pipe6 returns a left-to-right composition of 6 functions.
func Pipe6[A, B, C, D, E, F, G any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F, fg func(F) G) func(A) G {
	return func(a A) G {
		return fg(ef(de(cd(bc(ab(a))))))
	}
}

// Pipe7 returns a left-to-right composition of 7 functions.
func Pipe7[A, B, C, D, E, F, G, H any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F, fg func(F) G, gh func(G) H) func(A) H {
	return func(a A) H {
		return gh(fg(ef(de(cd(bc(ab(a)))))))
	}
}

// Pipe8 returns a left-to-right composition of 8 functions.
func Pipe8[A, B, C, D, E, F, G, H, I any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F, fg func(F) G, gh func(G) H, hi func(H) I) func(A) I {
	return func(a A) I {
		return hi(gh(fg(ef(de(cd(bc(ab(a))))))))
	}
}

// Pipe9 returns a left-to-right composition of 9 functions.
func Pipe9[A, B, C, D, E, F, G, H, I, J any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F, fg func(F) G, gh func(G) H, hi func(H) I, ij func(I) J) func(A) J {
	return func(a A) J {
		return ij(hi(gh(fg(ef(de(cd(bc(ab(a)))))))))
	}
}
