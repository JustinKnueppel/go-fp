# Set [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/set.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/set)

A `Set` represents a group of unique items of the same type in no particular order. There can be no duplicated items within a `Set`. This module provides functions that can help to work with this data structure.

```go
import (
	"github.com/JustinKnueppel/go-fp/set"
)
```

## Usage

Below are a few examples of how to use `Set`s and how they could be useful. Small examples for every function in the package can be found in the [package documentation](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/set).

### Authorization

`Set`s work perfectly for checking and modifying authorization scopes because scopes should never be duplicated and have no particular order.

```go
const (
  adminScope = "admin"
  userScope = "user"
)

type scope = string
type person struct {
  name string
  scopes set.Set[scope]
}

func isAdmin(p person) bool {
  return set.Member(adminScope)(p.scopes)
}

func isUser(p person) bool {
  return set.Member(userScope)(p.scopes)
}

func giveAdmin(p *person) {
  p.scopes = set.Insert(adminScope)(p.scopes)
}
```

### Easily removing duplicates

While creating a new data structure to remove duplicates may not be the fastest, it can be a declarative way to remove duplicates from a list. Note that this does not guarantee the same order.

```go
func uniqueNumbers(nums []int) []int {
  return fp.Pipe2(
    set.FromSlice[int],
    set.ToSlice[int],
  )(nums)
}
```
