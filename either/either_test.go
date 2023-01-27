package either_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/JustinKnueppel/go-fp/either"
	fp "github.com/JustinKnueppel/go-fp/function"
)

func TestString(t *testing.T) {
	left := either.Left[string, int]("foo")
	if left.String() != "Left foo" {
		t.Fatal()
	}

	right := either.Right[string](5)
	if right.String() != "Right 5" {
		t.Fatal()
	}
}

func ExampleLeft() {
	e := either.Left[error, int](errors.New("Failed"))
	fmt.Printf("Either is left: %v\n", either.IsLeft(e))
	fmt.Printf("Either is right: %v\n", either.IsRight(e))

	// Output:
	// Either is left: true
	// Either is right: false
}

func ExampleRight() {
	e := either.Right[error](2)
	fmt.Printf("Either is left: %v\n", either.IsLeft(e))
	fmt.Printf("Either is right: %v\n", either.IsRight(e))
	fmt.Printf("Either value: %d\n", either.Unwrap(e))

	// Output:
	// Either is left: false
	// Either is right: true
	// Either value: 2
}

func ExampleIsLeft() {
	fp.Pipe2(
		either.IsLeft[error, int],
		fp.Inspect(func(isLeft bool) {
			fmt.Printf("Either 1 is left: %v\n", isLeft)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.IsLeft[error, int],
		fp.Inspect(func(isLeft bool) {
			fmt.Printf("Either 2 is left: %v\n", isLeft)
		}),
	)(either.Right[error](5))

	// Output:
	// Either 1 is left: true
	// Either 2 is left: false
}

func ExampleIsLeftAnd() {
	fp.Pipe2(
		either.IsLeftAnd[error, int](func(err error) bool {
			return err.Error() == "failed"
		}),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 1 is left and passed: %v\n", passed)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.IsLeftAnd[error, int](func(err error) bool {
			return err.Error() == "failed"
		}),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 2 is left and passed: %v\n", passed)
		}),
	)(either.Left[error, int](errors.New("crashed")))

	fp.Pipe2(
		either.IsLeftAnd[error, int](func(err error) bool {
			return err.Error() == "failed"
		}),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 3 is left and passed: %v\n", passed)
		}),
	)(either.Right[error](2))

	// Output:
	// Either 1 is left and passed: true
	// Either 2 is left and passed: false
	// Either 3 is left and passed: false
}

func ExampleIsRight() {
	fp.Pipe2(
		either.IsRight[error, int],
		fp.Inspect(func(isRight bool) {
			fmt.Printf("Either 1 is right: %v\n", isRight)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.IsRight[error, int],
		fp.Inspect(func(isRight bool) {
			fmt.Printf("Either 2 is right: %v\n", isRight)
		}),
	)(either.Right[error](5))

	// Output:
	// Either 1 is right: false
	// Either 2 is right: true
}

func ExampleIsRightAnd() {
	fp.Pipe2(
		either.IsRightAnd[error](func(x int) bool { return x == 2 }),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 1 is right and passed: %v\n", passed)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.IsRightAnd[error](func(x int) bool { return x == 2 }),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 2 is right and passed: %v\n", passed)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.IsRightAnd[error](func(x int) bool { return x == 2 }),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Either 3 is right and passed: %v\n", passed)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 is right and passed: true
	// Either 2 is right and passed: false
	// Either 3 is right and passed: false
}

func ExampleMap() {
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		either.Map[error](double),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.Map[error](double),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 2 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: 4
	// Either 2 has error: failed
}

func ExampleBind() {
	fp.Pipe2(
		either.Bind(func(x int) either.Either[error, int] {
			return either.Right[error](x * 2)
		}),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.Bind(func(x int) either.Either[error, int] {
			return either.Left[error, int](errors.New("crashed"))
		}),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 2 has error: %v\n", err)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.Bind(func(x int) either.Either[error, int] {
			return either.Right[error](x * 2)
		}),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 3 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.Bind(func(x int) either.Either[error, int] {
			return either.Left[error, int](errors.New("crashed"))
		}),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 4 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: 4
	// Either 2 has error: crashed
	// Either 3 has error: failed
	// Either 4 has error: failed
}

func ExampleMapOr() {
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		either.MapOr[error, int](10)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.MapOr[error, int](10)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Value 1: 4
	// Value 2: 10
}

func ExampleMapOrElse() {
	constant := 10
	fallbackFn := func() int { return constant }
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		either.MapOrElse[error, int](fallbackFn)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.MapOrElse[error, int](fallbackFn)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Value 1: 4
	// Value 2: 10
}

func ExampleMapLeft() {
	fp.Pipe2(
		either.MapLeft[error, int, string](func(err error) int {
			// Example converting to exit code
			return 1
		}),
		either.InspectLeft[int, string](func(x int) {
			fmt.Printf("Either 1 has exit code: %d\n", x)
		}),
	)(either.Left[error, string](errors.New("failed")))

	fp.Pipe2(
		either.MapLeft[error, int, string](func(err error) int {
			// Example converting to exit code
			return 1
		}),
		either.Inspect[int](func(s string) {
			fmt.Printf("Either 2 has value: %s\n", s)
		}),
	)(either.Right[error]("real output"))

	// Output:
	// Either 1 has exit code: 1
	// Either 2 has value: real output
}

func ExampleInspect() {
	either.Inspect[error](func(x int) {
		fmt.Printf("Either 1 has value: %d\n", x)
	})(either.Right[error](5))

	either.Inspect[error](func(x int) {
		fmt.Println("This won't print")
	})(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: 5
}

func ExampleInspectLeft() {
	either.InspectLeft[error, int](func(err error) {
		fmt.Printf("Either 1 has error: %v\n", err)
	})(either.Left[error, int](errors.New("failed")))

	either.InspectLeft[error, int](func(err error) {
		fmt.Println("This won't print!")
	})(either.Right[error](2))

	// Output:
	// Either 1 has error: failed
}

func TestExpect(t *testing.T) {
	tests := map[string]struct {
		value         either.Either[error, int]
		inner         int
		msg           string
		expectedError bool
	}{
		"right": {
			value:         either.Right[error](1),
			inner:         1,
			msg:           "won't use",
			expectedError: false,
		},
		"left": {
			value:         either.Left[error, int](errors.New("failed")),
			inner:         0,
			msg:           "error found",
			expectedError: true,
		},
	}

	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			defer func() {
				panicMsg := recover()
				if tc.expectedError && panicMsg != tc.msg {
					t.Fail()
				}
			}()
			val := either.Expect[error, int](tc.msg)(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleExpect() {
	fp.Pipe2(
		either.Expect[error, int]("expect failed"),
		fp.Inspect(func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](2))

	// fp.Pipe2(
	// 	either.Expect[error, int]("expect failed"),
	// 	fp.Inspect(func(x int) {
	// 		fmt.Println("This panics with message 'expect failed'")
	// 	}),
	// )(either.Left[error, int](errors.New("bad program")))

	// Output:
	// Either 1 has value: 2
}

func TestUnwrap(t *testing.T) {
	tests := map[string]struct {
		value         either.Either[error, int]
		inner         int
		expectedError bool
	}{
		"right": {
			value:         either.Right[error](1),
			inner:         1,
			expectedError: false,
		},
		"left": {
			value:         either.Left[error, int](errors.New("failed")),
			inner:         0,
			expectedError: true,
		},
	}

	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			defer func() {
				e := recover()
				if !tc.expectedError && e != nil {
					t.Fail()
				}
			}()
			val := either.Unwrap(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleUnwrap() {
	fp.Pipe2(
		either.Unwrap[error, int],
		fp.Inspect(func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](2))

	// fp.Pipe2(
	// 	either.Unwrap[error, int],
	// 	fp.Inspect(func(x int) {
	// 		fmt.Println("This panics")
	// 	}),
	// )(either.Left[error, int](errors.New("bad program")))

	// Output:
	// Either 1 has value: 2
}

func ExampleUnwrapOr() {
	fp.Pipe2(
		either.UnwrapOr[error](10),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.UnwrapOr[error](10),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Value 1: 2
	// Value 2: 10
}

func ExampleUnwrapOrElse() {
	constant := 10
	fallbackFn := func() int { return constant }

	fp.Pipe2(
		either.UnwrapOrElse[error](fallbackFn),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.UnwrapOrElse[error](fallbackFn),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Value 1: 2
	// Value 2: 10
}

func ExampleUnwrapOrDefault() {
	fp.Pipe2(
		either.UnwrapOrDefault[error, int],
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.UnwrapOrDefault[error, int],
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Value 1: 2
	// Value 2: 0
}

func TestExpectLeft(t *testing.T) {
	tests := map[string]struct {
		value         either.Either[int, string]
		inner         int
		msg           string
		expectedError bool
	}{
		"right": {
			value:         either.Right[int]("success"),
			inner:         1,
			msg:           "either was right",
			expectedError: true,
		},
		"left": {
			value:         either.Left[int, string](1),
			inner:         1,
			msg:           "won't be used",
			expectedError: false,
		},
	}

	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			defer func() {
				panicMsg := recover()
				if tc.expectedError && panicMsg != tc.msg {
					t.Fail()
				}
			}()
			val := either.ExpectLeft[int, string](tc.msg)(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleExpectLeft() {
	fp.Pipe2(
		either.ExpectLeft[error, int]("expect left failed"),
		fp.Inspect(func(err error) {
			fmt.Printf("Either 1 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// fp.Pipe2(
	// 	either.ExpectLeft[error, int]("expect left failed"),
	// 	fp.Inspect(func(err error) {
	// 		fmt.Println("This panics with message 'expect left failed'")
	// 	}),
	// )(either.Right[error](2))

	// Output:
	// Either 1 has error: failed
}

func TestUnwrapLeft(t *testing.T) {
	tests := map[string]struct {
		value         either.Either[int, string]
		inner         int
		expectedError bool
	}{
		"right": {
			value:         either.Right[int]("sucess"),
			inner:         1,
			expectedError: true,
		},
		"left": {
			value:         either.Left[int, string](1),
			inner:         1,
			expectedError: false,
		},
	}

	for tname, tc := range tests {
		t.Run(tname, func(t *testing.T) {
			defer func() {
				e := recover()
				if !tc.expectedError && e != nil {
					t.Fail()
				}
			}()
			val := either.UnwrapLeft(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleUnwrapLeftOr() {
	fp.Pipe2(
		either.UnwrapLeftOr[int, string](0),
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)(either.Left[int, string](1))

	fp.Pipe2(
		either.UnwrapLeftOr[int, string](0),
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)(either.Right[int]("foo"))

	// Output:
	// 1
	// 0
}

func ExampleUnwrapLeft() {
	fp.Pipe2(
		either.UnwrapLeft[error, int],
		fp.Inspect(func(err error) {
			fmt.Printf("Either 1 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// fp.Pipe2(
	// 	either.UnwrapLeft[error, int],
	// 	fp.Inspect(func(err error) {
	// 		fmt.Println("This will panic")
	// 	}),
	// )(either.Right[error](2))

	// Output:
	// Either 1 has error: failed
}

func ExampleAnd() {
	fp.Pipe2(
		either.And[error, int](either.Right[error]("success")),
		either.Inspect[error](func(s string) {
			fmt.Printf("Either 1 has value: %s\n", s)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.And[error, int](either.Left[error, string](errors.New("crashed"))),
		either.InspectLeft[error, string](func(err error) {
			fmt.Printf("Either 2 has error: %v\n", err)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.And[error, int](either.Right[error]("success")),
		either.InspectLeft[error, string](func(err error) {
			fmt.Printf("Either 3 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.And[error, int](either.Left[error, string](errors.New("crashed"))),
		either.InspectLeft[error, string](func(err error) {
			fmt.Printf("Either 4 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: success
	// Either 2 has error: crashed
	// Either 3 has error: failed
	// Either 4 has error: failed
}

func ExampleOr() {
	fp.Pipe2(
		either.Or(either.Right[error](2)),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.Or(either.Left[error, int](errors.New("crashed"))),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 2 has value: %d\n", x)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.Or(either.Right[error](2)),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 3 has value: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.Or(either.Left[error, int](errors.New("crashed"))),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 4 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: 1
	// Either 2 has value: 1
	// Either 3 has value: 2
	// Either 4 has error: crashed
}

func ExampleOrElse() {
	fp.Pipe2(
		either.OrElse(func(err error) either.Either[error, int] { return either.Right[error](2) }),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 1 has value: %d\n", x)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.OrElse(func(err error) either.Either[error, int] { return either.Left[error, int](errors.New("crashed")) }),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 2 has value: %d\n", x)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.OrElse(func(err error) either.Either[error, int] { return either.Right[error](2) }),
		either.Inspect[error](func(x int) {
			fmt.Printf("Either 3 has value: %d\n", x)
		}),
	)(either.Left[error, int](errors.New("failed")))

	fp.Pipe2(
		either.OrElse(func(err error) either.Either[error, int] { return either.Left[error, int](errors.New("crashed")) }),
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Either 4 has error: %v\n", err)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 has value: 1
	// Either 2 has value: 1
	// Either 3 has value: 2
	// Either 4 has error: crashed
}

func ExampleContains() {
	fp.Pipe2(
		either.Contains[error](2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 1 contained right value: %v\n", contained)
		}),
	)(either.Right[error](2))

	fp.Pipe2(
		either.Contains[error](2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 2 contained right value: %v\n", contained)
		}),
	)(either.Right[error](1))

	fp.Pipe2(
		either.Contains[error](2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 3 contained right value: %v\n", contained)
		}),
	)(either.Left[error, int](errors.New("failed")))

	// Output:
	// Either 1 contained right value: true
	// Either 2 contained right value: false
	// Either 3 contained right value: false
}

func ExampleContainsLeft() {
	fp.Pipe2(
		either.ContainsLeft[int, string](0),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 1 contained left value: %v\n", contained)
		}),
	)(either.Left[int, string](0))

	fp.Pipe2(
		either.ContainsLeft[int, string](0),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 2 contained left value: %v\n", contained)
		}),
	)(either.Left[int, string](1))

	fp.Pipe2(
		either.ContainsLeft[int, string](0),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Either 3 contained left value: %v\n", contained)
		}),
	)(either.Right[int]("other value"))

	// Output:
	// Either 1 contained left value: true
	// Either 2 contained left value: false
	// Either 3 contained left value: false
}

func ExampleCopy() {
	right := either.Right[int]("success")
	fp.Pipe2(
		either.Copy[int, string],
		fp.Inspect(func(e either.Either[int, string]) {
			fmt.Printf("Either 1 value equal: %v\n", either.Equal(right)(e))
			fmt.Printf("Either 1 reference equal: %v\n", &e == &right)
		}),
	)(right)

	left := either.Left[int, string](1)
	fp.Pipe2(
		either.Copy[int, string],
		fp.Inspect(func(e either.Either[int, string]) {
			fmt.Printf("Either 2 value equal: %v\n", either.Equal(left)(e))
			fmt.Printf("Either 2 reference equal: %v\n", &e == &left)
		}),
	)(left)

	// Output:
	// Either 1 value equal: true
	// Either 1 reference equal: false
	// Either 2 value equal: true
	// Either 2 reference equal: false
}

func ExampleFlatten() {
	fp.Pipe2(
		either.Flatten[error, int],
		either.Inspect[error](func(x int) {
			fmt.Printf("Flattened 1 has value: %d\n", x)
		}),
	)(either.Right[error](either.Right[error](2)))

	fp.Pipe2(
		either.Flatten[error, int],
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Flattened 2 has error: %v\n", err)
		}),
	)(either.Right[error](either.Left[error, int](errors.New("inner fail"))))

	fp.Pipe2(
		either.Flatten[error, int],
		either.InspectLeft[error, int](func(err error) {
			fmt.Printf("Flattened 3 has error: %v\n", err)
		}),
	)(either.Left[error, either.Either[error, int]](errors.New("outer fail")))

	// Output:
	// Flattened 1 has value: 2
	// Flattened 2 has error: inner fail
	// Flattened 3 has error: outer fail
}

func ExampleEqual() {
	fp.Pipe2(
		either.Equal(either.Right[int]("success")),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Either 1s equal: %v\n", equal)
		}),
	)(either.Right[int]("success"))

	fp.Pipe2(
		either.Equal(either.Right[int]("other message")),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Either 2s equal: %v\n", equal)
		}),
	)(either.Right[int]("success"))

	fp.Pipe2(
		either.Equal(either.Left[int, string](1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Either 3s equal: %v\n", equal)
		}),
	)(either.Right[int]("success"))

	fp.Pipe2(
		either.Equal(either.Left[int, string](1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Either 4s equal: %v\n", equal)
		}),
	)(either.Left[int, string](1))

	fp.Pipe2(
		either.Equal(either.Left[int, string](1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Either 5s equal: %v\n", equal)
		}),
	)(either.Left[int, string](0))

	// Output:
	// Either 1s equal: true
	// Either 2s equal: false
	// Either 3s equal: false
	// Either 4s equal: true
	// Either 5s equal: false
}

func ExampleLefts() {
	fp.Pipe2(
		either.Lefts[int, string],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]either.Either[int, string]{})

	fp.Pipe2(
		either.Lefts[int, string],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]either.Either[int, string]{
		either.Left[int, string](1),
		either.Right[int]("foo"),
		either.Left[int, string](3),
	})

	// Output:
	// []
	// [1 3]
}

func ExampleRights() {
	fp.Pipe2(
		either.Rights[int, string],
		fp.Inspect(func(xs []string) {
			fmt.Println(xs)
		}),
	)([]either.Either[int, string]{})

	fp.Pipe2(
		either.Rights[int, string],
		fp.Inspect(func(xs []string) {
			fmt.Println(xs)
		}),
	)([]either.Either[int, string]{
		either.Right[int]("foo"),
		either.Left[int, string](2),
		either.Right[int]("bar"),
	})

	// Output:
	// []
	// [foo bar]
}
