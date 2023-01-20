package option_test

import (
	"fmt"
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/option"
)

func TestString(t *testing.T) {
	some := option.Some(5)
	if some.String() != "Some 5" {
		t.Fatal()
	}

	none := option.None[int]()
	if none.String() != "None" {
		t.Fatal()
	}
}

func ExampleSome() {
	o := option.Some(1)
	fmt.Printf("Option is some: %v\n", option.IsSome(o))
	fmt.Printf("Option is none: %v\n", option.IsNone(o))
	fmt.Printf("Option value: %d\n", option.Unwrap(o))

	// Output:
	// Option is some: true
	// Option is none: false
	// Option value: 1
}

func ExampleNone() {
	o := option.None[int]()
	fmt.Printf("Option is some: %v\n", option.IsSome(o))
	fmt.Printf("Option is none: %v\n", option.IsNone(o))

	// Output:
	// Option is some: false
	// Option is none: true
}

func ExampleIsSome() {
	fp.Pipe2(
		option.IsSome[int],
		fp.Inspect(func(isSome bool) {
			fmt.Printf("Option 1 is some: %v\n", isSome)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.IsSome[int],
		fp.Inspect(func(isSome bool) {
			fmt.Printf("Option 2 is some: %v\n", isSome)
		}),
	)(option.None[int]())

	// Output:
	// Option 1 is some: true
	// Option 2 is some: false
}

func ExampleIsSomeAnd() {
	greaterThanOne := func(x int) bool { return x > 1 }

	fp.Pipe2(
		option.IsSomeAnd(greaterThanOne),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Option 1 has value greater than 1: %v\n", passed)
		}),
	)(option.Some(2))

	fp.Pipe2(
		option.IsSomeAnd(greaterThanOne),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Option 2 has value greater than 1: %v\n", passed)
		}),
	)(option.Some(0))

	fp.Pipe2(
		option.IsSomeAnd(greaterThanOne),
		fp.Inspect(func(passed bool) {
			fmt.Printf("Option 3 has value greater than 1: %v\n", passed)
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value greater than 1: true
	// Option 2 has value greater than 1: false
	// Option 3 has value greater than 1: false
}

func ExampleIsNone() {
	fp.Pipe2(
		option.IsNone[int],
		fp.Inspect(func(isNone bool) {
			fmt.Printf("Option 1 is None: %v\n", isNone)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.IsNone[int],
		fp.Inspect(func(isNone bool) {
			fmt.Printf("Option 2 is None: %v\n", isNone)
		}),
	)(option.None[int]())

	// Output:
	// Option 1 is None: false
	// Option 2 is None: true
}

func TestExpect(t *testing.T) {
	tests := map[string]struct {
		value         option.Option[int]
		inner         int
		msg           string
		expectedError bool
	}{
		"some": {
			value:         option.Some(1),
			inner:         1,
			msg:           "",
			expectedError: false,
		},
		"none": {
			value:         option.None[int](),
			inner:         0,
			msg:           "No value",
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
			val := option.Expect[int](tc.msg)(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleExpect() {
	fp.Pipe2(
		option.Expect[int]("no value"),
		fp.Inspect(func(x int) {
			fmt.Printf("Option 1 contains: %d\n", x)
		}),
	)(option.Some(1))

	// fp.Pipe2(
	// 	option.Expect[int]("no value"),
	// 	fp.Inspect(func(x int) {
	// 		fmt.Println("This will panic with message: no value")
	// 	}),
	// )(option.None[int]())

	// Output:
	// Option 1 contains: 1
}

func TestUnwrap(t *testing.T) {
	tests := map[string]struct {
		value         option.Option[int]
		inner         int
		expectedError bool
	}{
		"some": {
			value:         option.Some(1),
			inner:         1,
			expectedError: false,
		},
		"none": {
			value:         option.None[int](),
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
			val := option.Unwrap(tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleUnwrap() {
	fp.Pipe2(
		option.Unwrap[int],
		fp.Inspect(func(x int) {
			fmt.Printf("Option 1 contains: %d\n", x)
		}),
	)(option.Some(1))

	// fp.Pipe2(
	// 	option.Unwrap[int],
	// 	fp.Inspect(func(x int) {
	// 		fmt.Println("This will panic")
	// 	}),
	// )(option.None[int]())

	// Output:
	// Option 1 contains: 1
}

func ExampleUnwrapOr() {
	fp.Pipe2(
		option.UnwrapOr(10),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.UnwrapOr(10),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(option.None[int]())

	// Output:
	// Value 1: 1
	// Value 2: 10
}

func ExampleUnwrapOrElse() {
	constant := 7
	computeDefault := func() int { return constant }

	fp.Pipe2(
		option.UnwrapOrElse(computeDefault),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.UnwrapOrElse(computeDefault),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(option.None[int]())

	// Output:
	// Value 1: 1
	// Value 2: 7
}

func ExampleUnwrapOrDefault() {
	fp.Pipe2(
		option.UnwrapOrDefault[int],
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.UnwrapOrDefault[int],
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(option.None[int]())

	// Output:
	// Value 1: 1
	// Value 2: 0
}

func ExampleMap() {
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		option.Map(double),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Map(double),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 2
}

func ExampleBind() {
	onlyGreaterThanOne := func(x int) option.Option[int] {
		if x > 1 {
			return option.Some(x)
		}
		return option.None[int]()
	}

	fp.Pipe2(
		option.Bind(onlyGreaterThanOne),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(2))

	fp.Pipe2(
		option.Bind(onlyGreaterThanOne),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.Some(0))

	fp.Pipe2(
		option.Bind(onlyGreaterThanOne),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 2
}

func ExampleInspect() {
	o1 := option.Some(1)
	option.Inspect(func(x int) {
		fmt.Printf("Option has value: %d\n", x)
	})(o1)

	o2 := option.None[int]()
	option.Inspect(func(_ int) {
		fmt.Println("This won't print!")
	})(o2)

	// Output:
	// Option has value: 1
}

func ExampleMapOr() {
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		option.MapOr[int](3)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.MapOr[int](3)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(option.None[int]())

	// Output:
	// Value 1: 2
	// Value 2: 3
}

func ExampleMapOrElse() {
	constant := 10
	calculateDefault := func() int { return constant }
	double := func(x int) int { return x * 2 }

	fp.Pipe2(
		option.MapOrElse[int](calculateDefault)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 1: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.MapOrElse[int](calculateDefault)(double),
		fp.Inspect(func(x int) {
			fmt.Printf("Value 2: %d\n", x)
		}),
	)(option.None[int]())

	// Output:
	// Value 1: 2
	// Value 2: 10
}

func ExampleAnd() {
	fp.Pipe2(
		option.And[int](option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.And[int](option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.And[int](option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.And[int](option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 2
}

func ExampleFilter() {
	greaterThanOne := func(x int) bool { return x > 1 }

	fp.Pipe2(
		option.Filter(greaterThanOne),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(2))

	fp.Pipe2(
		option.Filter(greaterThanOne),
		option.Inspect(func(x int) {
			fmt.Printf("This won't print!")
		}),
	)(option.Some(0))

	fp.Pipe2(
		option.Filter(greaterThanOne),
		option.Inspect(func(x int) {
			fmt.Printf("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 2
}

func ExampleOr() {
	fp.Pipe2(
		option.Or(option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Or(option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Printf("Option 2 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Or(option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Printf("Option 3 has value: %d\n", x)
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.Or(option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 1
	// Option 2 has value: 1
	// Option 3 has value: 2
}

func ExampleOrElse() {
	constant := 10

	fp.Pipe2(
		option.OrElse(func() option.Option[int] { return option.Some(2) }),
		option.Inspect(func(x int) {
			fmt.Printf("Option 1 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.OrElse(func() option.Option[int] { return option.None[int]() }),
		option.Inspect(func(x int) {
			fmt.Printf("Option 2 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.OrElse(func() option.Option[int] { return option.Some(2) }),
		option.Inspect(func(x int) {
			fmt.Printf("Option 3 has value: %d\n", x)
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.OrElse(func() option.Option[int] { return option.Some(constant) }),
		option.Inspect(func(x int) {
			fmt.Printf("Option 4 has value: %d\n", x)
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.OrElse(func() option.Option[int] { return option.None[int]() }),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 1 has value: 1
	// Option 2 has value: 1
	// Option 3 has value: 2
	// Option 4 has value: 10
}

func ExampleXor() {
	fp.Pipe2(
		option.Xor(option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Xor(option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Printf("Option 2 has value: %d\n", x)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Xor(option.Some(2)),
		option.Inspect(func(x int) {
			fmt.Printf("Option 3 has value: %d\n", x)
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.Xor(option.None[int]()),
		option.Inspect(func(x int) {
			fmt.Println("This won't print!")
		}),
	)(option.None[int]())

	// Output:
	// Option 2 has value: 1
	// Option 3 has value: 2
}

func ExampleContains() {
	fp.Pipe2(
		option.Contains(1),
		fp.Inspect(func(contains bool) {
			fmt.Printf("Option 1 contains: %v\n", contains)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Contains(1),
		fp.Inspect(func(contains bool) {
			fmt.Printf("Option 2 contains: %v\n", contains)
		}),
	)(option.Some(2))

	fp.Pipe2(
		option.Contains(1),
		fp.Inspect(func(contains bool) {
			fmt.Printf("Option 3 contains: %v\n", contains)
		}),
	)(option.None[int]())

	// Output:
	// Option 1 contains: true
	// Option 2 contains: false
	// Option 3 contains: false
}

func ExampleCopy() {
	o1 := option.Some(1)
	fp.Pipe2(
		option.Copy[int],
		fp.Inspect(func(copy option.Option[int]) {
			fmt.Printf("Option 1 value equal: %v\n", option.Equal(copy)(o1))
			fmt.Printf("Option 1 reference equal: %v\n", &copy == &o1)
		}),
	)(o1)

	o2 := option.None[int]()
	fp.Pipe2(
		option.Copy[int],
		fp.Inspect(func(copy option.Option[int]) {
			fmt.Printf("Option 2 value equal: %v\n", option.Equal(copy)(o2))
			fmt.Printf("Option 2 reference equal: %v\n", &copy == &o2)
		}),
	)(o2)

	// Output:
	// Option 1 value equal: true
	// Option 1 reference equal: false
	// Option 2 value equal: true
	// Option 2 reference equal: false
}

func ExampleEqual() {
	fp.Pipe2(
		option.Equal(option.Some(1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 1s equal: %v\n", equal)
		}),
	)(option.Some(1))

	fp.Pipe2(
		option.Equal(option.Some(1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 2s equal: %v\n", equal)
		}),
	)(option.Some(2))

	fp.Pipe2(
		option.Equal(option.Some(1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 3s equal: %v\n", equal)
		}),
	)(option.None[int]())

	fp.Pipe2(
		option.Equal(option.None[int]()),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 4s equal: %v\n", equal)
		}),
	)(option.None[int]())

	// Output:
	// Option 1s equal: true
	// Option 2s equal: false
	// Option 3s equal: false
	// Option 4s equal: true
}

func ExampleFlatten() {
	fp.Pipe3(
		option.Flatten[int],
		option.Equal(option.Some(1)),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 1s equal: %v\n", equal)
		}),
	)(option.Some(option.Some(1)))

	fp.Pipe3(
		option.Flatten[int],
		option.Equal(option.None[int]()),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 2s equal: %v\n", equal)
		}),
	)(option.Some(option.None[int]()))

	fp.Pipe3(
		option.Flatten[int],
		option.Equal(option.None[int]()),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Option 3s equal: %v\n", equal)
		}),
	)(option.None[option.Option[int]]())

	// Output:
	// Option 1s equal: true
	// Option 2s equal: true
	// Option 3s equal: true
}
