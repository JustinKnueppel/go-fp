package fp

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
		return func(b B) func(C) D {
			return func(c C) D {
				return fn(a, b, c)
			}
		}
	}
}

// Curry4 returns the curried form of the 4-arity funcion.
func Curry4[A, B, C, D, E any](fn func(A, B, C, D) E) func(A) func(B) func(C) func(D) E {
	return func(a A) func(B) func(C) func(D) E {
		return func(b B) func(C) func(D) E {
			return func(c C) func(D) E {
				return func(d D) E {
					return fn(a, b, c, d)
				}
			}
		}
	}
}

// Curry5 returns the curried form of the 5-arity funcion.
func Curry5[A, B, C, D, E, F any](fn func(A, B, C, D, E) F) func(A) func(B) func(C) func(D) func(E) F {
	return func(a A) func(B) func(C) func(D) func(E) F {
		return func(b B) func(C) func(D) func(E) F {
			return func(c C) func(D) func(E) F {
				return func(d D) func(E) F {
					return func(e E) F {
						return fn(a, b, c, d, e)
					}
				}
			}
		}
	}
}

// Curry6 returns the curried form of the 6-arity funcion.
func Curry6[A, B, C, D, E, F, G any](fn func(A, B, C, D, E, F) G) func(A) func(B) func(C) func(D) func(E) func(F) G {
	return func(a A) func(B) func(C) func(D) func(E) func(F) G {
		return func(b B) func(C) func(D) func(E) func(F) G {
			return func(c C) func(D) func(E) func(F) G {
				return func(d D) func(E) func(F) G {
					return func(e E) func(F) G {
						return func(f F) G {
							return fn(a, b, c, d, e, f)
						}
					}
				}
			}
		}
	}
}

// Curry7 returns the curried form of the 7-arity funcion.
func Curry7[A, B, C, D, E, F, G, H any](fn func(A, B, C, D, E, F, G) H) func(A) func(B) func(C) func(D) func(E) func(F) func(G) H {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) H {
		return func(b B) func(C) func(D) func(E) func(F) func(G) H {
			return func(c C) func(D) func(E) func(F) func(G) H {
				return func(d D) func(E) func(F) func(G) H {
					return func(e E) func(F) func(G) H {
						return func(f F) func(G) H {
							return func(g G) H {
								return fn(a, b, c, d, e, f, g)
							}
						}
					}
				}
			}
		}
	}
}

// Curry8 returns the curried form of the 8-arity funcion.
func Curry8[A, B, C, D, E, F, G, H, I any](fn func(A, B, C, D, E, F, G, H) I) func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) I {
		return func(b B) func(C) func(D) func(E) func(F) func(G) func(H) I {
			return func(c C) func(D) func(E) func(F) func(G) func(H) I {
				return func(d D) func(E) func(F) func(G) func(H) I {
					return func(e E) func(F) func(G) func(H) I {
						return func(f F) func(G) func(H) I {
							return func(g G) func(H) I {
								return func(h H) I {
									return fn(a, b, c, d, e, f, g, h)
								}
							}
						}
					}
				}
			}
		}
	}
}

// Curry9 returns the curried form of the 9-arity funcion.
func Curry9[A, B, C, D, E, F, G, H, I, J any](fn func(A, B, C, D, E, F, G, H, I) J) func(A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J {
	return func(a A) func(B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J {
		return func(b B) func(C) func(D) func(E) func(F) func(G) func(H) func(I) J {
			return func(c C) func(D) func(E) func(F) func(G) func(H) func(I) J {
				return func(d D) func(E) func(F) func(G) func(H) func(I) J {
					return func(e E) func(F) func(G) func(H) func(I) J {
						return func(f F) func(G) func(H) func(I) J {
							return func(g G) func(H) func(I) J {
								return func(h H) func(I) J {
									return func(i I) J {
										return fn(a, b, c, d, e, f, g, h, i)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
