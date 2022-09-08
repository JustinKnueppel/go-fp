package option_test

import (
	"fmt"
	"testing"

	"github.com/JustinKnueppel/go-fp/option"
)

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
	o1 := option.Some(1)
	if option.IsSome(o1) {
		fmt.Printf("Option contains %d\n", option.Unwrap(o1))
	}

	o2 := option.None[int]()
	if option.IsSome(o2) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option contains 1
}

func ExampleIsSomeAnd() {
	greaterThanOne := func(x int) bool { return x > 1 }

	o1 := option.Some(2)
	if option.IsSomeAnd(greaterThanOne, o1) {
		fmt.Printf("%d is greater than 1\n", option.Unwrap(o1))
	}

	o2 := option.Some(0)
	if option.IsSomeAnd(greaterThanOne, o2) {
		fmt.Println("This won't print!")
	}

	o3 := option.None[int]()
	if option.IsSomeAnd(greaterThanOne, o3) {
		fmt.Println("This won't print either!")
	}

	// Output:
	// 2 is greater than 1
}

func ExampleIsNone() {
	o1 := option.Some(1)
	if option.IsNone(o1) {
		fmt.Println("This won't print!", option.Unwrap(o1))
	}

	o2 := option.None[int]()
	if option.IsNone(o2) {
		fmt.Println("Option is None")
	}

	// Output:
	// Option is None
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
			val := option.Expect(tc.msg, tc.value)
			if val != tc.inner {
				t.Fail()
			}
		})
	}
}

func ExampleExpect() {
	o1 := option.Some(1)
	fmt.Printf("Option contains %d\n", option.Expect("no value", o1))

	// This will panic with error message "no value"
	// o2 := option.None[int]()
	// fmt.Printf("Option contains %d\n", option.Expect(o2, "no value"))

	// Output:
	// Option contains 1
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
	o1 := option.Some(1)
	fmt.Printf("Option contains %d\n", option.Unwrap(o1))

	// This will panic
	// o2 := option.None[int]()
	// fmt.Printf("Option contains %d\n", option.Unwrap(o2))

	// Output:
	// Option contains 1
}

func ExampleUnwrapOr() {
	o1 := option.Some(1)
	fmt.Printf("Option contains %d\n", option.UnwrapOr(10, o1))

	o2 := option.None[int]()
	fmt.Printf("Option contains %d\n", option.UnwrapOr(10, o2))

	// Output:
	// Option contains 1
	// Option contains 10
}

func ExampleUnwrapOrElse() {
	constant := 7
	computeDefault := func() int { return constant }

	o1 := option.Some(1)
	fmt.Printf("Returned %d\n", option.UnwrapOrElse(computeDefault, o1))

	o2 := option.None[int]()
	fmt.Printf("Returned %d\n", option.UnwrapOrElse(computeDefault, o2))

	// Output:
	// Returned 1
	// Returned 7
}

func ExampleUnwrapOrDefault() {
	o1 := option.Some(1)
	fmt.Printf("Returned %d\n", option.UnwrapOrDefault(o1))

	o2 := option.None[int]()
	fmt.Printf("Returned %d\n", option.UnwrapOrDefault(o2))

	// Output:
	// Returned 1
	// Returned 0
}

func ExampleMap() {
	double := func(x int) int { return x * 2 }

	o1 := option.Some(1)
	mapped1 := option.Map(double, o1)
	if option.IsSome(mapped1) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(mapped1))
	}

	o2 := option.None[int]()
	mapped2 := option.Map(double, o2)
	if option.IsSome(mapped2) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 2
}

func ExampleBind() {
	onlyGreaterThanOne := func(x int) option.Option[int] {
		if x > 1 {
			return option.Some(x)
		}
		return option.None[int]()
	}

	o1 := option.Some(2)
	mapped1 := option.Bind(onlyGreaterThanOne, o1)
	if option.IsSome(mapped1) {
		fmt.Printf("Option has value: %d", option.Unwrap(mapped1))
	}

	o2 := option.Some(0)
	mapped2 := option.Bind(onlyGreaterThanOne, o2)
	if option.IsSome(mapped2) {
		fmt.Println("This won't print!")
	}

	o3 := option.None[int]()
	mapped3 := option.Bind(onlyGreaterThanOne, o3)
	if option.IsSome(mapped3) {
		fmt.Println("This won't print either!")
	}

	// Output:
	// Option has value: 2
}

func ExampleInspect() {
	o1 := option.Some(1)
	option.Inspect(func(x int) {
		fmt.Printf("Option has value: %d\n", x)
	}, o1)

	o2 := option.None[int]()
	option.Inspect(func(_ int) {
		fmt.Println("This won't print!")
	}, o2)

	// Output:
	// Option has value: 1
}

func ExampleMapOr() {
	double := func(x int) int { return x * 2 }

	o1 := option.Some(1)
	val1 := option.MapOr(3, double, o1)
	fmt.Printf("Returned: %d\n", val1)

	o2 := option.None[int]()
	val2 := option.MapOr(3, double, o2)
	fmt.Printf("Returned: %d\n", val2)

	// Output:
	// Returned: 2
	// Returned: 3
}

func ExampleMapOrElse() {
	constant := 10
	calculateDefault := func() int { return constant }
	double := func(x int) int { return x * 2 }

	o1 := option.Some(1)
	val1 := option.MapOrElse(calculateDefault, double, o1)
	fmt.Printf("Returned: %d\n", val1)

	o2 := option.None[int]()
	val2 := option.MapOrElse(calculateDefault, double, o2)
	fmt.Printf("Returned: %d\n", val2)

	// Output:
	// Returned: 2
	// Returned: 10
}

func ExampleAnd() {
	o1 := option.Some(1)
	optb1 := option.Some(2)
	val1 := option.And(optb1, o1)
	if option.IsSome(val1) {
		fmt.Printf("Option has value: %d", option.Unwrap(val1))
	}

	o2 := option.Some(1)
	optb2 := option.None[int]()
	val2 := option.And(optb2, o2)
	if option.IsSome(val2) {
		fmt.Println("This won't print!")
	}

	o3 := option.None[int]()
	optb3 := option.Some(2)
	val3 := option.And(optb3, o3)
	if option.IsSome(val3) {
		fmt.Println("This won't print!")
	}

	o4 := option.None[int]()
	optb4 := option.None[int]()
	val4 := option.And(optb4, o4)
	if option.IsSome(val4) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 2
}

func ExampleFilter() {
	greaterThanOne := func(x int) bool { return x > 1 }

	o1 := option.Some(2)
	filtered1 := option.Filter(greaterThanOne, o1)
	if option.IsSome(filtered1) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(filtered1))
	}

	o2 := option.Some(0)
	filtered2 := option.Filter(greaterThanOne, o2)
	if option.IsSome(filtered2) {
		fmt.Println("This won't print!")
	}

	o3 := option.None[int]()
	filtered3 := option.Filter(greaterThanOne, o3)
	if option.IsSome(filtered3) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 2
}

func ExampleOr() {
	o1 := option.Some(1)
	optb1 := option.Some(2)
	val1 := option.Or(optb1, o1)
	if option.IsSome(val1) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val1))
	}

	o2 := option.Some(1)
	optb2 := option.None[int]()
	val2 := option.Or(optb2, o2)
	if option.IsSome(val2) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val2))
	}

	o3 := option.None[int]()
	optb3 := option.Some(2)
	val3 := option.Or(optb3, o3)
	if option.IsSome(val3) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val3))
	}

	o4 := option.None[int]()
	optb4 := option.None[int]()
	val4 := option.Or(optb4, o4)
	if option.IsSome(val4) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 1
	// Option has value: 1
	// Option has value: 2
}

func ExampleOrElse() {
	constant := 10

	o1 := option.Some(1)
	optb1 := func() option.Option[int] { return option.Some(2) }
	val1 := option.OrElse(optb1, o1)
	if option.IsSome(val1) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val1))
	}

	o2 := option.Some(1)
	optb2 := func() option.Option[int] { return option.None[int]() }
	val2 := option.OrElse(optb2, o2)
	if option.IsSome(val2) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val2))
	}

	o3 := option.None[int]()
	optb3 := func() option.Option[int] { return option.Some(2) }
	val3 := option.OrElse(optb3, o3)
	if option.IsSome(val3) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val3))
	}

	o4 := option.None[int]()
	optb4 := func() option.Option[int] { return option.Some(constant) }
	val4 := option.OrElse(optb4, o4)
	if option.IsSome(val4) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val4))
	}

	o5 := option.None[int]()
	optb5 := func() option.Option[int] { return option.None[int]() }
	val5 := option.OrElse(optb5, o5)
	if option.IsSome(val5) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 1
	// Option has value: 1
	// Option has value: 2
	// Option has value: 10
}

func ExampleXor() {
	o1 := option.Some(1)
	optb1 := option.Some(2)
	val1 := option.Xor(optb1, o1)
	if option.IsSome(val1) {
		fmt.Println("This won't print!")
	}

	o2 := option.Some(1)
	optb2 := option.None[int]()
	val2 := option.Xor(optb2, o2)
	if option.IsSome(val2) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val2))
	}

	o3 := option.None[int]()
	optb3 := option.Some(2)
	val3 := option.Xor(optb3, o3)
	if option.IsSome(val3) {
		fmt.Printf("Option has value: %d\n", option.Unwrap(val3))
	}

	o4 := option.None[int]()
	optb4 := option.None[int]()
	val4 := option.Xor(optb4, o4)
	if option.IsSome(val4) {
		fmt.Println("This won't print!")
	}

	// Output:
	// Option has value: 1
	// Option has value: 2
}

func ExampleContains() {
	o1 := option.Some(1)
	contains1 := option.Contains(1, o1)
	fmt.Printf("Option 1 contains: %v\n", contains1)

	o2 := option.Some(2)
	contains2 := option.Contains(1, o2)
	fmt.Printf("Option 2 contains: %v\n", contains2)

	o3 := option.None[int]()
	contains3 := option.Contains(1, o3)
	fmt.Printf("Option 3 contains: %v\n", contains3)

	// Output:
	// Option 1 contains: true
	// Option 2 contains: false
	// Option 3 contains: false
}

func ExampleCopy() {
	o1 := option.Some(1)
	copy1 := option.Copy(o1)
	fmt.Printf("Option 1 value equal: %v\n", option.Equal(copy1, o1))
	fmt.Printf("Option 1 reference equal: %v\n", &copy1 == &o1)

	o2 := option.None[int]()
	copy2 := option.Copy(o2)
	fmt.Printf("Option 2 value equal: %v\n", option.Equal(copy2, o2))
	fmt.Printf("Option 2 reference equal: %v\n", &copy2 == &o2)

	// Output:
	// Option 1 value equal: true
	// Option 1 reference equal: false
	// Option 2 value equal: true
	// Option 2 reference equal: false
}

func ExampleEqual() {
	o1 := option.Some(1)
	optb1 := option.Some(1)
	fmt.Printf("Option 1s equal: %v\n", option.Equal(optb1, o1))

	o2 := option.Some(1)
	optb2 := option.Some(2)
	fmt.Printf("Option 2s equal: %v\n", option.Equal(optb2, o2))

	o3 := option.None[int]()
	optb3 := option.Some(1)
	fmt.Printf("Option 3s equal: %v\n", option.Equal(optb3, o3))

	o4 := option.None[int]()
	optb4 := option.None[int]()
	fmt.Printf("Option 4s equal: %v\n", option.Equal(optb4, o4))

	// Output:
	// Option 1s equal: true
	// Option 2s equal: false
	// Option 3s equal: false
	// Option 4s equal: true
}

func ExampleFlatten() {
	o1 := option.Some(option.Some(1))
	expected1 := option.Some(1)
	flattened1 := option.Flatten(o1)
	fmt.Printf("Option 1s equal: %v\n", option.Equal(expected1, flattened1))

	o2 := option.Some(option.None[int]())
	expected2 := option.None[int]()
	flattened2 := option.Flatten(o2)
	fmt.Printf("Option 2s equal: %v\n", option.Equal(expected2, flattened2))

	o3 := option.None[option.Option[int]]()
	expected3 := option.None[int]()
	flattened3 := option.Flatten(o3)
	fmt.Printf("Option 3s equal: %v\n", option.Equal(expected3, flattened3))

	// Output:
	// Option 1s equal: true
	// Option 2s equal: true
	// Option 3s equal: true
}
