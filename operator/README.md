# Operator [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/operator.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/operator)

This package gives access to function versions of all of the built in operators. This can be very useful for mapping, partially applying, and working with built in operators in a functional context.

```go
import (
  "github.com/JustinKnueppel/go-fp/operator"
)
```

## Usage

All of the functions in this package are curried to make composition and partial application as easy as possible. Below are a few examples of how these functions can be used together, but small examples for each function can be found in the [package documentation](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/operator).

### Left shift all numbers in slice

```go
fp.Pipe2(
  slice.Map(fp.Flip2(operator.LShift[int])(1)),
  fp.Inspect(func(xs []int) {
    fmt.Println(xs)
  }),
)([]int{1, 2, 3})

// Output:
// [2 4 6]
```

## Add 5 to a slice of numbers

```go
fp.Pipe2(
  slice.Map(operator.Add(5)),
  fp.Inspect(func(xs []int) {
    fmt.Println(xs)
  }),
)([]int{1, 2, 3})

// Output:
// [6 7 8]
```

## Sum numbers with slice.Reduce

```go
fp.Pipe2(
  slice.Reduce(operator.Add[int]),
		option.Inspect(func(x int) {
    fmt.Println(x)
  }),
)([]int{1, 2, 3})

// Output:
// 6
```
