# Function package [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/function.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/function)

This package contains utility functions for working with functions themselves. as well as a few basic utility functions. Some examples and descriptions can be found below, otherwise there are small examples for every function in the package documentation.

```go
import (
  fp "github.com/JustinKnueppel/fp-go/function"
)
```

## Contents

- [ComposeN](#compose)
- [CurryN](#curry)
- [FlipN](#flip)
- [Inspect](#inspect)
- [PipeN](#pipe)
- [UncurryN](#uncurry)

### <a name="compose">`ComposeN`</a>

The `ComposeN` function set performs right-to-left function composition on `N` functions. As a base set, there are supplied `Compose` functions for 2-9 functions to be composed. However, composition can be nested as `Compose` itself returns a function. For example, the two approaches are equivalent as function composition is associative:

```go
add1 := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
triple := func(x int) int { return x * 3 }

result4 := fp.Compose4(add1, double, add1, triple)(5) // performs triple -> add1 -> double -> add1
result2 := fp.Compose2(fp.Compose2(add1, double), fp.Compose2(add1, triple))(5) // performs (triple -> add1) -> (double -> add1)
```

These `Compose` functions allow for smaller functions to be composed together without needing to be called immediately.

### <a name="curry">`CurryN`</a>

The `CurryN` function set curries an n-arity function into n nested 1-arity functions. As a base set, there are supplied `Curry` functions for 2-9 arguments. Here is an example of the usage.

```go
func add(a, b int) int { return a + b }

curried := fp.Curry2(add)
add1 := curried(1)

if add1(2) != 3 {
  t.Fail()
}
```

### <a name="flip">`FlipN`</a>

The `FlipN` function set reverses the arguments of the n-arity curried function. As a base set, there are supplied `Flip` functions for 2-9 arity functions. Here is an example usage.

```go
func greet(greeting string) func(name string) string {
  return func(name string) {
    return fmt.Sprintf("%s, %s!\n", greeting, name)
  }
}

greet("Hello")("John") // "Hello, John!"

greetJohn := fp.Flip2(greet)("John")
greetJohn("Hello") // "Hello, John!"
greetJohn("Sup") // "Sup, John!"
```

### <a name="inspect">`Inspect`</a>

The `Inspect` gives a method for passing a piped value to a closure without breaking the pipe loop. In Go, a function with no return value defined does not have some sort of `Void` or `Unit` type, but is rather completely unusable. This causes functions without a return type to break `Pipe` or `Compose` pipelines. `Inspect` addresses this issue by allowing such functions to be called and forwarding along the argument.

```go
double := func(x int) int { return x * 2 }
fp.Pipe2(
  double,
  fp.Inspect(func(x int) {
    fmt.Printf("X has value: %d\n", x)
  }),
)(3)

// Outputs:
// X has value: 6
```

### <a name="pipe">`PipeN`</a>

The `PipeN` function set performs left-to-right function composition on `N` functions. As a base set, there are supplied `Pipe` functions for 2-9 functions to be composed. However, composition can be nested as `Pipe` itself returns a function. For example, the two approaches are equivalent as function composition is associative:

```go
add1 := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
triple := func(x int) int { return x * 3 }

result4 := fp.Pipe4(add1, double, add1, triple)(5) // performs add1 -> double -> add1 -> triple
result2 := fp.Pipe2(fp.Pipe2(add1, double), fp.Pipe2(add1, triple))(5) // performs (add1 -> double) -> (add1 -> triple)
```

These `Pipe` functions allow for smaller functions to be composed together without needing to be called immediately.

### <a name="uncurry">`UncurryN`</a>

The `UncurryN` function set uncurries n nested 1-arity functions into an n-arity function. As a base set, there are supplied `Uncurry` functions for 2-9 levels of nesting. Here is an example of the usage.

```go
func curriedAdd(a int) func(int) int {
  return func(b int) int {
    return a + b
  }
}

add := fp.Uncurry2(curriedAdd)

if add(1, 2) != 3 {
  t.Fail()
}
```
