# Go functional programming

The purpose of this repository is to provide a set of utility functions to allow for more functional code to be written in Go. Every export will be either a function, or an immutable type that can be operated on with functions. If there is any missing functionality that you would like to see, feel free to leave a Github issue detailing the functionality and its use!

## `ComposeN`

The `ComposeN` function set performs right-to-left function composition on `N` functions. As a base set, there are supplied `Compose` functions for 2-9 functions to be composed. However, composition can be nested as `Compose` itself returns a function. For example, the two approaches are equivalent as function composition is associative:

```go
add1 := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
triple := func(x int) int { return x * 3 }

result4 := fp.Compose4(add1, double, add1, triple)(5) // performs triple -> add1 -> double -> add1
result2 := fp.Compose2(fp.Compose2(add1, double), fp.Compose2(add1, triple))(5) // performs (triple -> add1) -> (double -> add1)
```

These `Compose` functions allow for smaller functions to be composed together without needing to be called immediately.


## `PipeN`

The `PipeN` function set performs left-to-right function composition on `N` functions. As a base set, there are supplied `Pipe` functions for 2-9 functions to be composed. However, composition can be nested as `Pipe` itself returns a function. For example, the two approaches are equivalent as function composition is associative:

```go
add1 := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
triple := func(x int) int { return x * 3 }

result4 := fp.Pipe4(add1, double, add1, triple)(5) // performs add1 -> double -> add1 -> triple
result2 := fp.Pipe2(fp.Pipe2(add1, double), fp.Pipe2(add1, triple))(5) // performs (add1 -> double) -> (add1 -> triple)
```

These `Pipe` functions allow for smaller functions to be composed together without needing to be called immediately.

## `CurryN`

The `CurryN` function set curries an n-arity function into n nested 1-arity functions. As a base set, there are supplied `Curry` functions for 2-9 arguments. Here is an example of the usage.

```go
func add(a, b int) int { return a + b }

curried := fp.Curry2(add)
add1 := curried(1)

if add1(2) != 3 {
  t.Fail()
}
```

## `UncurryN`

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
