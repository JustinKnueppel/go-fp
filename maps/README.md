# Maps [![Go Reference](https://pkg.go.dev/badge/github.com/JustinKnueppel/go-fp/maps.svg)](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/maps)

Maps are built in structures in Go, but have very few built-in operations. Some of these operations also have their own unique syntax such as deleting from a map with `delete(map, key)`. This package is primarily a reimplementation of the [Haskell `Data.Map` package](https://hackage.haskell.org/package/containers-0.4.0.0/docs/Data-Map.html) to the best degree possible in Go.

```go
import (
  "github.com/JustinKnueppel/go-fp/maps"
)
```

Note the package name is `maps` not `map`, as `map` is a reserved keyword in Go. This follows the same convention as the standard library `strings` package.

## Usage

All of the functions in this package are curried to make composition as easy as possible. Below are a few examples of how these functions can be used together, but small examples for each function can be found in the [package documentation](https://pkg.go.dev/github.com/JustinKnueppel/go-fp/maps).

### Safe convert map of string prices to floats

```go
func strLt(s1 string) func(string) bool {
	return func(s2 string) bool {
		return strings.Compare(s1, s2) < 0
	}
}

func stringToFloat(s string) option.Option[float64] {
  if f, err := strconv.ParseFloat(s, 64); err == nil {
	  return option.Some(f)
  }
  return option.None[float64]()
}

func main() {
  fp.Pipe3(
		maps.MapOption[string](stringToFloat),
		maps.ToAscSlice[string, float64](strLt),
		fp.Inspect(func(pairs []tuple.Pair[string, float64]) {
			fmt.Println(pairs)
		}),
	)(map[string]string{
		"apple":  "2.99",
		"orange": "3.99",
	})

	// Output:
	// [(apple 2.99) (orange 3.99)]
}
```

### Collision safe hash table

```go
func main() {
  hashTable := maps.Empty[string, []int]()
	insert := fp.Curry2(func(k string, v int) func(map[string][]int) map[string][]int {
		return maps.InsertWith[string](slice.AppendSlice[int])(k)(slice.Singleton(v))
	})

	hashTable = insert("abcd")(5)(hashTable)
	hashTable = insert("xzy")(10)(hashTable)
	hashTable = insert("abcd")(15)(hashTable)

	fp.Pipe2(
		maps.ToAscSlice[string, []int](strLt),
		fp.Inspect(func(pairs []tuple.Pair[string, []int]) {
			fmt.Println(pairs)
		}),
	)(hashTable)

	// Output:
	// [(abcd [5 15]) (xzy [10])]
}
```
