package function

// Compose2 returns a right-to-left composition of 2 functions.
func Compose2[A, B, C any](bc func(B) C, ab func(A) B) func(A) C {
	return func(a A) C {
		return bc(ab(a))
	}
}

// Compose3 returns a right-to-left composition of 3 functions.
func Compose3[A, B, C, D any](cd func(C) D, bc func(B) C, ab func(A) B) func(A) D {
	return func(a A) D {
		return cd(bc(ab(a)))
	}
}

// Compose4 returns a right-to-left composition of 4 functions.
func Compose4[A, B, C, D, E any](de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) E {
	return func(a A) E {
		return de(cd(bc(ab(a))))
	}
}

// Compose5 returns a right-to-left composition of 5 functions.
func Compose5[A, B, C, D, E, F any](ef func(E) F, de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) F {
	return func(a A) F {
		return ef(de(cd(bc(ab(a)))))
	}
}

// Compose6 returns a right-to-left composition of 6 functions.
func Compose6[A, B, C, D, E, F, G any](fg func(F) G, ef func(E) F, de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) G {
	return func(a A) G {
		return fg(ef(de(cd(bc(ab(a))))))
	}
}

// Compose7 returns a right-to-left composition of 7 functions.
func Compose7[A, B, C, D, E, F, G, H any](gh func(G) H, fg func(F) G, ef func(E) F, de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) H {
	return func(a A) H {
		return gh(fg(ef(de(cd(bc(ab(a)))))))
	}
}

// Compose8 returns a right-to-left composition of 8 functions.
func Compose8[A, B, C, D, E, F, G, H, I any](hi func(H) I, gh func(G) H, fg func(F) G, ef func(E) F, de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) I {
	return func(a A) I {
		return hi(gh(fg(ef(de(cd(bc(ab(a))))))))
	}
}

// Compose9 returns a right-to-left composition of 9 functions.
func Compose9[A, B, C, D, E, F, G, H, I, J any](ij func(I) J, hi func(H) I, gh func(G) H, fg func(F) G, ef func(E) F, de func(D) E, cd func(C) D, bc func(B) C, ab func(A) B) func(A) J {
	return func(a A) J {
		return ij(hi(gh(fg(ef(de(cd(bc(ab(a)))))))))
	}
}
