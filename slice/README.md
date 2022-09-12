# Slice [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/slice.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/slice)

Slices are one of the core data types in Go, but most ways in which you work with slices in Go must be hand rolled. This package provides a set of utility functions to work with slices in a more functional manner.

```go
import (
  "github.com/JustinKnueppel/go-fp/slice"
)
```

## Usage

All of the functions in this package are curried to make composition as easy as possible. Below are a few examples of how these functions can be used together, but small examples for each function can be found in the [package documentation](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/slice).

### Find first 5 multiples of 7 over 100

```go
divisibleBy7 := func(x int) bool { return x%7 == 0 }
over100 := func(x int) bool { return x > 100 }

fp.Pipe4(
  slice.Filter(over100),
  slice.Filter(divisibleBy7),
  slice.Take[int](5),
  fp.Inspect(func(nums []int) {
    fmt.Printf("First 5 multiples of 7 over 100: %v\n", nums)
  }),
)(slice.Range(0)(200))

// Output:
// First 5 multiples of 7 over 100: [105 112 119 126 133]
```

### Display names of all adults

```go
type person struct {
  name string
  age  int
}
isAdult := func(p person) bool { return p.age >= 18 }
getName := func(p person) string { return p.name }
makeListItem := func(s string) string { return "- " + s }
AggregateNames := func(names string, name string) string { return names + "\n" + name }

fp.Pipe5(
  slice.Filter(isAdult),
  slice.Map(getName),
  slice.Map(makeListItem),
  slice.Reduce(AggregateNames),
  option.Inspect(func(namesString string) {
    fmt.Printf("Adults:\n%s\n", namesString)
  }),
)([]person{{"Tim", 16}, {"James", 18}, {"John", 9}, {"Jim", 30}})

// Output:
// Adults:
// - James
// - Jim
```

## `Fold` vs `Reduce`

The `Fold` function set allows for an initial value to be defined while the `Reduce` function set uses the first value in the slice as its initial value. Because of the need for a non-empty slice, the `Reduce` function set must return an `Option` to keep a full function mapping.

## Future functionality

Once we have some sort of `Pair` or `Tuple` construct, the following functions will be added

- `SplitAt` - Returns two slices split at the given index from the original
- `ZipWith` - returns a slice of tuples where the ith element is the ith element of the two provided slices with the remaining elements of the longer slice getting truncated
