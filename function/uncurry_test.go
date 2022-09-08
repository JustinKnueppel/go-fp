package function_test

import (
	"fmt"

	fp "github.com/JustinKnueppel/go-fp/function"
)

func ExampleUncurry2() {
	fn := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	equal := fp.Uncurry2(fn)(1, 2) == fn(1)(2)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry3() {
	fn := func(a int) func(int) func(int) int {
		return func(b int) func(int) int {
			return func(c int) int {
				return a + b + c
			}
		}
	}

	equal := fp.Uncurry3(fn)(1, 2, 3) == fn(1)(2)(3)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry4() {
	fn := func(a int) func(int) func(int) func(int) int {
		return func(b int) func(int) func(int) int {
			return func(c int) func(int) int {
				return func(d int) int {
					return a + b + c + d
				}
			}
		}
	}

	equal := fp.Uncurry4(fn)(1, 2, 3, 4) == fn(1)(2)(3)(4)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry5() {
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

	equal := fp.Uncurry5(fn)(1, 2, 3, 4, 5) == fn(1)(2)(3)(4)(5)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry6() {
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

	equal := fp.Uncurry6(fn)(1, 2, 3, 4, 5, 6) == fn(1)(2)(3)(4)(5)(6)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry7() {
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

	equal := fp.Uncurry7(fn)(1, 2, 3, 4, 5, 6, 7) == fn(1)(2)(3)(4)(5)(6)(7)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry8() {
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

	equal := fp.Uncurry8(fn)(1, 2, 3, 4, 5, 6, 7, 8) == fn(1)(2)(3)(4)(5)(6)(7)(8)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}

func ExampleUncurry9() {
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

	equal := fp.Uncurry9(fn)(1, 2, 3, 4, 5, 6, 7, 8, 9) == fn(1)(2)(3)(4)(5)(6)(7)(8)(9)
	fmt.Printf("Uncurried function is equal: %v\n", equal)

	// Output:
	// Uncurried function is equal: true
}
