package fp_test

import (
	"testing"

	"github.com/JustinKnueppel/go-fp"
)

func TestUncurry2(t *testing.T) {
	fn := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	if fp.Uncurry2(fn)(1, 2) != fn(1)(2) {
		t.Fatal("should uncurry the 2-arity function")
	}
}

func TestUncurry3(t *testing.T) {
	fn := func(a int) func(int) func(int) int {
		return func(b int) func(int) int {
			return func(c int) int {
				return a + b + c
			}
		}
	}

	if fp.Uncurry3(fn)(1, 2, 3) != fn(1)(2)(3) {
		t.Fatal("should uncurry the 3-arity function")
	}
}

func TestUncurry4(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) int {
			return func(c int) func(int) int {
				return func(d int) int {
					return a + b + c + d
				}
			}
		}
	}

	if fp.Uncurry4(fn)(1, 2, 3, 4) != fn(1)(2)(3)(4) {
		t.Fatal("should uncurry the 4-arity function")
	}
}

func TestUncurry5(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) func(int) int {
			return func(c int) func(int) func(int) int {
				return func(d int) func(int) int {
					return func(e int) int {
						return a + b + c + d + e
					}
				}
			}
		}
	}

	if fp.Uncurry5(fn)(1, 2, 3, 4, 5) != fn(1)(2)(3)(4)(5) {
		t.Fatal("should uncurry the 5-arity function")
	}
}

func TestUncurry6(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) func(int) func(int) int {
			return func(c int) func(int) func(int) func(int) int {
				return func(d int) func(int) func(int) int {
					return func(e int) func(int) int {
						return func(f int) int {
							return a + b + c + e + f
						}
					}
				}
			}
		}
	}

	if fp.Uncurry6(fn)(1, 2, 3, 4, 5, 6) != fn(1)(2)(3)(4)(5)(6) {
		t.Fatal("should uncurry the 6-arity function")
	}
}

func TestUncurry7(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) func(int) func(int) func(int) int {
			return func(c int) func(int) func(int) func(int) func(int) int {
				return func(d int) func(int) func(int) func(int) int {
					return func(e int) func(int) func(int) int {
						return func(f int) func(int) int {
							return func(g int) int {
								return a + b + c + d + e + f + g
							}
						}
					}
				}
			}
		}
	}

	if fp.Uncurry7(fn)(1, 2, 3, 4, 5, 6, 7) != fn(1)(2)(3)(4)(5)(6)(7) {
		t.Fatal("should uncurry the 7-arity function")
	}
}

func TestUncurry8(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) func(int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) func(int) func(int) func(int) func(int) int {
			return func(c int) func(int) func(int) func(int) func(int) func(int) int {
				return func(d int) func(int) func(int) func(int) func(int) int {
					return func(e int) func(int) func(int) func(int) int {
						return func(f int) func(int) func(int) int {
							return func(g int) func(int) int {
								return func(h int) int {
									return a + b + c + d + e + f + g + h
								}
							}
						}
					}
				}
			}
		}
	}

	if fp.Uncurry8(fn)(1, 2, 3, 4, 5, 6, 7, 8) != fn(1)(2)(3)(4)(5)(6)(7)(8) {
		t.Fatal("should uncurry the 8-arity function")
	}
}

func TestUncurry9(t *testing.T) {
	fn := func(a int) func(int) func(int) func(int) func(int) func(int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) func(int) func(int) func(int) func(int) func(int) int {
			return func(c int) func(int) func(int) func(int) func(int) func(int) func(int) int {
				return func(d int) func(int) func(int) func(int) func(int) func(int) int {
					return func(e int) func(int) func(int) func(int) func(int) int {
						return func(f int) func(int) func(int) func(int) int {
							return func(g int) func(int) func(int) int {
								return func(h int) func(int) int {
									return func(i int) int {
										return a + b + c + d + e + f + g + h + i
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if fp.Uncurry9(fn)(1, 2, 3, 4, 5, 6, 7, 8, 9) != fn(1)(2)(3)(4)(5)(6)(7)(8)(9) {
		t.Fatal("should uncurry the 9-arity function")
	}
}
