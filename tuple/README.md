# Tuple [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/tuple.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/tuple)

Tuples represent an immutable, fixed length, possible heterogeneous collection of values. Go does not have native support for tuples of any length despite de facto pairs being used as return types for functions with extreme frequency `(T, error)`. The provided type in this package `Pair[T, U]` mimics this 2-tuple that is returned by functions as best as possible. However, these tuples are not able to be used as supplying multiple arguments to an uncurried function such as in Haskell, or able to be pattern matched with Go syntax either (`x, err := f()`). They are however useful when it comes to zipping lists together, or returning multiple items from a function with a proper type.

```go
import (
  "github.com/JustinKnueppel/go-fp/tuple"
)
```

## Usage

Currently, the only type in this package is a `Pair[T, U]` with two helper functions to retrieve the first (`Fst`) or second (`Snd`) elements. Example tests can be found in the [package documentation](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/tuple).
