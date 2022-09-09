# Option

This package offers an `Option` generic type largely inspired by the `Option` type from [Rust](https://doc.rust-lang.org/stable/std/option/).

```go
import (
  "github.com/JustinKnueppel/fp-go/option"
)
```

## Purpose

The only exported type of this package is `Option`. This type can never be instantiated directly, but rather through the two constructors: `Some(t T)` or `None()`. The former represents an `Option` with a value `t` of generic type `T`, and the latter represents the absense of a value. Using an `Option` allows for the idea of something being "nullable" without actually doing null checks or error checks while keeping full type safety, and without using pointers to allow base types to be null. The zero value of an `Option` is `None`.

To work with `Option`s, you will often inject functionality into the option type rather than immediately trying to pull the type out of the `Option`. For example, if you have a function that may or may not return an `int`. Normally this would be implemented by either returning a `(int, error)` tuple where the error represents the lack of an `int`, or a `*int` could be used with `nil` as the return value when no `int` is present. Using these two paradigms if we wanted to double the value we would have to do something like the following:

```go
// Error method
func FirstCalculation() (int, error) {...}

v1, err := FirstCalculation()
if err != nil {
  return 0, err
}

return v1 * 2, nil

// Pointer method
func SomeCalculation() *int {...}

v1 := SomeCalculation()
if v1 == nil {
  return nil
}

return v1 * 2
```

In both of these cases, we have to propogate the means with which we checked this error despite having to manully handle the error ourselves. With an `Option`, we apply the logic into the `Option` itself like so:

```go
func SomeCalculation() Option[int] {...}

val := SomeCalculation()
doubled := option.Map(val, func(x int) int {x * 2})
return doubled // Option[int]
```

If `SomeCalculation` returned a value, we will double that value, otherwise we will automatically propogate the `None` which was returned.

## Usage

`Option` types should be instantiated via the `Some(t T)` and `None` constructors only (though the default value for an `Option` is a `None`). All functions in this package are curried, and (other than the constructors) take the target option as their final argument. This is to ensure the greatest ease with function composition. A few examples of how to use the package follow, and small examples for each function can be found in the package documentation.

### Safe division

In standard Go, a division function is liable to panic if a divisor of 0 is passed. Even if this is handled in the division function, it would be necessary to work with a returned error every time the division function is used. With an `Option` type, we can encapsulate this uncertainty in the type system itself.

```go
func Divide(dividend, divisor int) option.Option[int] {
  if divisor == 0 {
    return option.None()
  }
  return option.Some(dividend/divisor)
}

Divide(6, 2) // Some(3)
Divide(6, 0) // None
```

### Chaining maps

The following example shows how multiple functions can be chained together to modify the `Option`. If in either the `couldReturnNone` or `unsafeMap` functions a `None` is returned, then `composeAll` will return `None`. Otherwise, all of the maps will be applied one after the other to create a `Some` with the final value.

```go
func couldReturnNone(int) option.Option[int] {}
func firstMap(int) int {}
func secondMap(int) string {}
func unsafeMap(string) option.Option[int] {}
func finalMap(int) int {}

func main() {
  function.Pipe6(
		couldReturnNone,
		option.Map(firstMap),
		option.Map(secondMap),
		option.Bind(unsafeMap),
		option.Map(finalMap),
		option.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)(6)
}
```

### Working with contained value

For the most part, functionality should be injected into the `Option` rather than trying to pull the inner value out of it. This can be achieved via `Map`s, `AndThen`s, `Inspect`s, and many more utility functions. However, if eventually you need to try to get the value out of the `Option`, it will be less convenient than in the Rust counterpart of this package as Go does not have pattern matching. Instead we must use the `IsSome`, `IsNone`, and `UnWrap*` methods. Note that the `Unwrap` method panics if called on a `None` type, and should always be guarded by an `IsSome` or `IsNone` check. To specify the error message with which `Unwrap` panics, the `Expect` method can be used instead.

```go
func getAnOption() option.Option[int] {}

func main() {
  opt := getAnOption()
  if option.IsSome(opt) {
    val := option.Unwrap(opt)
    fmt.Println(val + 2)
  } else {
    fmt.Println("No value exists")
  }
}
```
