package slice_test

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

func printAny[T any](t T) {
	fmt.Println(t)
}

func ExampleAppend() {
	fp.Pipe2(
		slice.Append(3),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending to empty slice yields: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Append(3),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending to slice yields: %v\n", nums)
		}),
	)([]int{1, 2})

	// Output:
	// Appending to empty slice yields: [3]
	// Appending to slice yields: [1 2 3]
}

func ExampleAppendSlice() {
	fp.Pipe2(
		slice.AppendSlice([]int{}),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending empty slice to empty slice yields: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.AppendSlice([]int{3, 4}),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending to empty slice yields: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.AppendSlice([]int{}),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending empty slice to slice yields: %v\n", nums)
		}),
	)([]int{1, 2})

	fp.Pipe2(
		slice.AppendSlice([]int{3, 4}),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Appending to slice yields: %v\n", nums)
		}),
	)([]int{1, 2})

	// Output:
	// Appending empty slice to empty slice yields: []
	// Appending to empty slice yields: [3 4]
	// Appending empty slice to slice yields: [1 2]
	// Appending to slice yields: [1 2 3 4]
}

func ExampleAt() {
	fp.Pipe2(
		slice.At[int](-1),
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Returns None if index < 0: %v\n", option.IsNone(o))
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.At[int](3),
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Returns None if index >= length: %v\n", option.IsNone(o))
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.At[int](1),
		option.Inspect(func(x int) {
			fmt.Printf("Returns value at given index: %d\n", x)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Returns None if index < 0: true
	// Returns None if index >= length: true
	// Returns value at given index: 2
}

func ExampleBreak() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.Break(isEven),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(split), tuple.Snd(split))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Break(isEven),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(split), tuple.Snd(split))
		}),
	)([]int{2, 4})

	fp.Pipe2(
		slice.Break(isEven),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(split), tuple.Snd(split))
		}),
	)([]int{3, 5})

	fp.Pipe2(
		slice.Break(isEven),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(split), tuple.Snd(split))
		}),
	)([]int{3, 2, 3, 4, 5})

	// Output:
	// [[] []]
	// [[] [2 4]]
	// [[3 5] []]
	// [[3] [2 3 4 5]]
}

func ExampleStripPrefix() {
	fp.Pipe2(
		slice.StripPrefix([]int{1, 2}),
		fp.Inspect(printAny[option.Option[[]int]]),
	)([]int{})

	fp.Pipe2(
		slice.StripPrefix([]int{1, 2}),
		fp.Inspect(printAny[option.Option[[]int]]),
	)([]int{3, 2, 3, 4})

	fp.Pipe2(
		slice.StripPrefix([]int{1, 2}),
		fp.Inspect(printAny[option.Option[[]int]]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.StripPrefix([]int{1, 2}),
		fp.Inspect(printAny[option.Option[[]int]]),
	)([]int{1, 2, 3, 4})

	// Output:
	// None
	// None
	// Some []
	// Some [3 4]
}

func ExampleConcat() {
	fp.Pipe2(
		slice.Concat[int],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([][]int{})

	fp.Pipe2(
		slice.Concat[int],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([][]int{{1, 2}})

	fp.Pipe2(
		slice.Concat[int],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([][]int{{1, 2}, {3, 4}})

	// Output:
	// []
	// [1 2]
	// [1 2 3 4]
}

func ExampleConcatMap() {
	fp.Pipe2(
		slice.ConcatMap(slice.Range(1)),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.ConcatMap(slice.Range(1)),
		fp.Inspect(printAny[[]int]),
	)([]int{3, 4, 5})

	// Output:
	// []
	// [1 2 1 2 3 1 2 3 4]
}

func ExampleAnd() {
	fp.Pipe2(
		slice.And,
		fp.Inspect(printAny[bool]),
	)([]bool{})

	fp.Pipe2(
		slice.And,
		fp.Inspect(printAny[bool]),
	)([]bool{true, true})

	fp.Pipe2(
		slice.And,
		fp.Inspect(printAny[bool]),
	)([]bool{false})

	fp.Pipe2(
		slice.And,
		fp.Inspect(printAny[bool]),
	)([]bool{true, false})

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleOr() {
	fp.Pipe2(
		slice.Or,
		fp.Inspect(printAny[bool]),
	)([]bool{})

	fp.Pipe2(
		slice.Or,
		fp.Inspect(printAny[bool]),
	)([]bool{true, true})

	fp.Pipe2(
		slice.Or,
		fp.Inspect(printAny[bool]),
	)([]bool{false})

	fp.Pipe2(
		slice.Or,
		fp.Inspect(printAny[bool]),
	)([]bool{true, false})

	// Output:
	// false
	// true
	// false
	// true
}

func ExampleElem() {
	fp.Pipe2(
		slice.Elem(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Empty slice contained target: %v\n", contained)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Elem(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Slice 1 contained target: %v\n", contained)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Elem(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Slice 2 contained target: %v\n", contained)
		}),
	)([]int{1, 3})

	// Output:
	// Empty slice contained target: false
	// Slice 1 contained target: true
	// Slice 2 contained target: false
}

func ExampleNotElem() {
	fp.Pipe2(
		slice.NotElem(2),
		fp.Inspect(printAny[bool]),
	)([]int{})

	fp.Pipe2(
		slice.NotElem(2),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.NotElem(2),
		fp.Inspect(printAny[bool]),
	)([]int{1, 3})

	// Output:
	// true
	// false
	// true
}

func ExampleLookup() {
	fp.Pipe2(
		slice.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)([]tuple.Pair[string, int]{})

	fp.Pipe2(
		slice.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("bar")(2),
	})

	fp.Pipe2(
		slice.Lookup[string, int]("foo"),
		fp.Inspect(printAny[option.Option[int]]),
	)([]tuple.Pair[string, int]{
		tuple.NewPair[string, int]("foo")(1),
		tuple.NewPair[string, int]("bar")(2),
	})

	// Output:
	// None
	// None
	// Some 1
}

func ExampleCopy() {
	xs := []int{1, 2, 3}
	fp.Pipe2(
		slice.Copy[int],
		fp.Inspect(func(copied []int) {
			fmt.Printf("Slice values are equal: %v\n", slice.Equal(copied)(xs))
			fmt.Printf("Slice references are equal: %v\n", &copied == &xs)
		}),
	)(xs)

	ys := []int{1, 2, 3}
	fp.Pipe3(
		slice.Copy[int],
		func(list []int) []int {
			list[1] = 8
			return list
		},
		fp.Inspect(func(_ []int) {
			fmt.Printf("Original values are unchanged: %v\n", ys)
		}),
	)(ys)

	// Output:
	// Slice values are equal: true
	// Slice references are equal: false
	// Original values are unchanged: [1 2 3]
}

func ExampleDelete() {
	fp.Pipe2(
		slice.Delete(2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Delete on empty slice returns empty slice: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Delete(2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Delete returns input when element not found: %v\n", nums)
		}),
	)([]int{1, 3})

	fp.Pipe2(
		slice.Delete(2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Delete only removes first instance: %v\n", nums)
		}),
	)([]int{1, 2, 3, 2})

	// Output:
	// Delete on empty slice returns empty slice: []
	// Delete returns input when element not found: [1 3]
	// Delete only removes first instance: [1 3 2]
}

func ExampleDifference() {
	fp.Pipe2(
		slice.Difference([]int{}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Difference([]int{1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Difference([]int{4, 1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2})

	// Output:
	// []
	// [1 2 3 1]
	// [4 3 1]
}

func ExampleUnion() {
	fp.Pipe2(
		slice.Union([]int{}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Union([]int{1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Union([]int{4, 1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 5})

	// Output:
	// [1 2]
	// [1 2 3 1]
	// [4 1 2 3 1 5]
}

func ExampleIntersect() {
	fp.Pipe2(
		slice.Intersect([]int{}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Intersect([]int{1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Intersect([]int{4, 1, 2, 3, 1}),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 5})

	// Output:
	// []
	// []
	// [1 2 1]
}

func ExampleInsert() {
	fp.Pipe2(
		slice.Insert(4),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Insert(4),
		fp.Inspect(printAny[[]int]),
	)([]int{4})

	fp.Pipe2(
		slice.Insert(4),
		fp.Inspect(printAny[[]int]),
	)([]int{3, 1, 9, 1})

	// Output:
	// [4]
	// [4 4]
	// [3 1 4 9 1]
}

func ExampleDeleteAt() {
	fp.Pipe2(
		slice.DeleteAt[int](1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.DeleteAt[int](1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.DeleteAt[int](1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// []
	// [1]
	// [1 3]
}

func ExampleElemIndex() {
	fp.Pipe2(
		slice.ElemIndex(1),
		fp.Inspect(printAny[option.Option[int]]),
	)([]int{})

	fp.Pipe2(
		slice.ElemIndex(1),
		fp.Inspect(printAny[option.Option[int]]),
	)([]int{0, 2, 3})

	fp.Pipe2(
		slice.ElemIndex(1),
		fp.Inspect(printAny[option.Option[int]]),
	)([]int{1, 2, 3})

	// Output:
	// None
	// None
	// Some 0
}

func ExampleElemIndices() {
	fp.Pipe2(
		slice.ElemIndices(1),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.ElemIndices(1),
		fp.Inspect(printAny[[]int]),
	)([]int{2, 3})

	fp.Pipe2(
		slice.ElemIndices(1),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 1, 3, 1})

	// Output:
	// []
	// []
	// [0 1 3]
}

func ExampleUniqueBy() {
	type person struct {
		Name string
		Age  int
	}
	nameEq := fp.On[person](operator.Eq[string])(func(p person) string { return p.Name })

	fp.Pipe2(
		slice.UniqueBy(nameEq),
		fp.Inspect(printAny[[]person]),
	)([]person{})

	fp.Pipe2(
		slice.UniqueBy(nameEq),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 21},
		{"Barry", 20},
	})

	fp.Pipe2(
		slice.UniqueBy(nameEq),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 21},
		{"Barry", 20},
		{"Barry", 20},
		{"Jerry", 19},
	})

	// Output:
	// []
	// [{Jerry 21} {Barry 20}]
	// [{Jerry 21} {Barry 20}]
}

func ExampleDeleteBy() {
	type person struct {
		Name string
		Age  int
	}
	nameEqual := fp.Curry2(func(p1, p2 person) bool { return p1.Name == p2.Name })

	jerry := person{"Jerry", 18}

	fp.Pipe2(
		slice.DeleteBy(nameEqual)(jerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy on empty slice returns empty slice: %v\n", people)
		}),
	)([]person{})

	fp.Pipe2(
		slice.DeleteBy(nameEqual)(jerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy returns input when element not found: %v\n", people)
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.DeleteBy(nameEqual)(jerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy only removes first instance: %v\n", people)
		}),
	)([]person{{"Jerry", 21}, {"Tim", 30}, {"Jerry", 40}})

	// Output:
	// DeleteBy on empty slice returns empty slice: []
	// DeleteBy returns input when element not found: [{Tim 21}]
	// DeleteBy only removes first instance: [{Tim 30} {Jerry 40}]
}

func ExampleDeleteFirstsBy() {
	type person struct {
		Name string
		Age  int
	}
	personEq := fp.Curry2(func(p1, p2 person) bool { return p1.Name == p2.Name && p1.Age == p2.Age })

	fp.Pipe2(
		slice.DeleteFirstsBy(personEq)([]person{}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 18},
		{"Sam", 21},
	})

	fp.Pipe2(
		slice.DeleteFirstsBy(personEq)([]person{
			{"Jerry", 18},
			{"Sam", 21},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{})

	fp.Pipe2(
		slice.DeleteFirstsBy(personEq)([]person{
			{"Jerry", 18},
			{"Sam", 20},
			{"Sam", 21},
			{"Jerry", 18},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 18},
		{"Sam", 21},
	})

	// Output:
	// []
	// [{Jerry 18} {Sam 21}]
	// [{Sam 20} {Jerry 18}]
}

func ExampleDrop() {
	fp.Pipe2(
		slice.Drop[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Drop called on empty slice yields empty slice: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Drop[int](0),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Drop 0 returns original slice: %v\n", nums)
		}),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Drop[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Drop removes first 2 elements: %v\n", nums)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Drop[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Drop yields empty slice when n > length: %v\n", nums)
		}),
	)([]int{1})

	// Output:
	// Drop called on empty slice yields empty slice: []
	// Drop 0 returns original slice: [1 2]
	// Drop removes first 2 elements: [3]
	// Drop yields empty slice when n > length: []
}

func ExampleDropWhile() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.DropWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.DropWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{2, 4})

	fp.Pipe2(
		slice.DropWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{2, 4, 5, 6, 7})

	fp.Pipe2(
		slice.DropWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 4, 5, 6, 7})

	// Output:
	// []
	// []
	// [5 6 7]
	// [1 2 4 5 6 7]
}

func ExampleDropWhileEnd() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.DropWhileEnd(isEven),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.DropWhileEnd(isEven),
		fp.Inspect(printAny[[]int]),
	)([]int{6, 8})

	fp.Pipe2(
		slice.DropWhileEnd(isEven),
		fp.Inspect(printAny[[]int]),
	)([]int{2, 3})

	fp.Pipe2(
		slice.DropWhileEnd(isEven),
		fp.Inspect(printAny[[]int]),
	)([]int{2, 3, 6, 8})

	// Output:
	// []
	// []
	// [2 3]
	// [2 3]
}

func ExampleNull() {
	fp.Pipe2(
		slice.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Empty slice is empty: %v\n", empty)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Filled slice is empty: %v\n", empty)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice is empty: true
	// Filled slice is empty: false
}

func ExampleEqual() {
	fp.Pipe2(
		slice.Equal([]int{1, 2}),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Slices are equal: %v\n", equal)
		}),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Equal([]int{1, 2}),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Subsets are not equal: %v\n", equal)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Equal([]int{1, 2}),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Supersets are not equal: %v\n", equal)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Equal([]int{4, 5, 6}),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Slices with different elements are not equal: %v\n", equal)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Slices are equal: true
	// Subsets are not equal: false
	// Supersets are not equal: false
	// Slices with different elements are not equal: false
}

func ExampleAll() {
	is2 := func(x int) bool { return x == 2 }

	fp.Pipe2(
		slice.All(is2),
		fp.Inspect(func(every bool) {
			fmt.Printf("Every element passed in empty list: %v\n", every)
		}),
	)([]int{})

	fp.Pipe2(
		slice.All(is2),
		fp.Inspect(func(every bool) {
			fmt.Printf("Every element passed in slice 1: %v\n", every)
		}),
	)([]int{2, 2})

	fp.Pipe2(
		slice.All(is2),
		fp.Inspect(func(every bool) {
			fmt.Printf("Every element passed in slice 2: %v\n", every)
		}),
	)([]int{1, 2})

	// Output:
	// Every element passed in empty list: true
	// Every element passed in slice 1: true
	// Every element passed in slice 2: false
}

func ExampleFilter() {
	fp.Pipe2(
		slice.Filter(func(x int) bool { return x > 1 }),
		fp.Inspect(func(filtered []int) {
			fmt.Printf("Filtered slice: %v\n", filtered)
		}),
	)([]int{0, 1, 2, 3})

	fp.Pipe2(
		slice.Filter(func(x int) bool { return x > 1 }),
		fp.Inspect(func(filtered []int) {
			fmt.Printf("All filtered: %v\n", filtered)
		}),
	)([]int{0, 1})

	fp.Pipe2(
		slice.Filter(func(x int) bool { return x > 1 }),
		fp.Inspect(func(filtered []int) {
			fmt.Printf("None filtered: %v\n", filtered)
		}),
	)([]int{2, 3})

	// Output:
	// Filtered slice: [2 3]
	// All filtered: []
	// None filtered: [2 3]
}

func ExamplePartition() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.Partition(isEven),
		fp.Inspect(printAny[tuple.Pair[[]int, []int]]),
	)([]int{})

	fp.Pipe2(
		slice.Partition(isEven),
		fp.Inspect(printAny[tuple.Pair[[]int, []int]]),
	)([]int{1, 2, 3, 4, 5})

	// Output:
	// ([] [])
	// ([2 4] [1 3 5])
}

func ExampleFilterMap() {
	safeParseInt := func(s string) option.Option[int] {
		i, err := strconv.Atoi(s)
		if err != nil {
			return option.None[int]()
		}
		return option.Some(i)
	}

	fp.Pipe2(
		slice.FilterMap(safeParseInt),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FilterMap(safeParseInt),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]string{"hello"})

	fp.Pipe2(
		slice.FilterMap(safeParseInt),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]string{"hello", "5"})

	fp.Pipe2(
		slice.FilterMap(safeParseInt),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]string{"-1", "foo", "3", "bar", "100"})

	// Output:
	// []
	// []
	// [5]
	// [-1 3 100]
}

func ExampleFind() {
	type person struct {
		Name string
		Age  int
	}
	isJerry := func(p person) bool { return p.Name == "Jerry" }

	fp.Pipe2(
		slice.Find(isJerry),
		fp.Inspect(func(o option.Option[person]) {
			fmt.Printf("Returns None when element doesn't exist: %v\n", option.IsNone(o))
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.Find(isJerry),
		option.Inspect(func(p person) {
			fmt.Printf("Returns the first matching element: %v\n", p)
		}),
	)([]person{{"Tim", 21}, {"Jerry", 30}, {"Jerry", 40}})

	// Output:
	// Returns None when element doesn't exist: true
	// Returns the first matching element: {Jerry 30}
}

func ExampleBind() {
	repeatN := func(n int) []int { return slice.Replicate[int](n)(n) }

	fp.Pipe2(
		slice.Bind(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped empty slice: %v\n", xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Bind(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped slice where one map is empty: %v\n", xs)
		}),
	)([]int{1, 0, 3})

	fp.Pipe2(
		slice.Bind(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped slice: %v\n", xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Mapped empty slice: []
	// Mapped slice where one map is empty: [1 3 3 3]
	// Mapped slice: [1 2 2 3 3 3]
}

func ExampleFoldl() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Foldl(appendStr)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldl(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldl(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo
	// Multiple elements appended: start_foo_bar_baz
}

func ExampleFoldlWithIndex() {
	appendStrAndIndex := fp.Curry3(func(i int, x, y string) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.FoldlWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldlWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldlWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo0
	// Multiple elements appended: start_foo0_bar1_baz2
}

func ExampleFoldlWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(xs []string, i int, x, y string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.FoldlWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldlWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldlWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo01
	// Multiple elements appended: start_foo03_bar13_baz23
}

func ExampleFoldr() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Foldr(appendStr)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldr(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldr(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: foo_start
	// Multiple elements appended: foo_bar_baz_start
}

func ExampleFoldrWithIndex() {
	appendStrAndIndex := fp.Curry3(func(i int, x, acc string) string { return x + fmt.Sprint(i) + "_" + acc })

	fp.Pipe2(
		slice.FoldrWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldrWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldrWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: foo0_start
	// Multiple elements appended: foo0_bar1_baz2_start
}

func ExampleFoldrWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(xs []string, i int, x, acc string) string { return x + fmt.Sprintf("%d%d", i, len(xs)) + "_" + acc })

	fp.Pipe2(
		slice.FoldrWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldrWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldrWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: foo01_start
	// Multiple elements appended: foo03_bar13_baz23_start
}

func ExampleGroup() {
	fp.Pipe2(
		slice.Group[int],
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Group[int],
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 1})

	fp.Pipe2(
		slice.Group[int],
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 1, 4, 5, 5})

	fp.Pipe2(
		slice.Group[int],
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 3, 5, 7, 9})

	// Output:
	// []
	// [[1 1]]
	// [[1 1] [4] [5 5]]
	// [[1] [3] [5] [7] [9]]
}

func ExampleGroupBy() {
	close := fp.Curry2(func(x, y int) bool { return math.Abs(float64(x-y)) <= 1 })

	fp.Pipe2(
		slice.GroupBy(close),
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{})

	fp.Pipe2(
		slice.GroupBy(close),
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 2})

	fp.Pipe2(
		slice.GroupBy(close),
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 2, 4, 5, 9})

	fp.Pipe2(
		slice.GroupBy(close),
		fp.Inspect(func(groups [][]int) {
			fmt.Println(groups)
		}),
	)([]int{1, 3, 5, 7, 9})

	// Output:
	// []
	// [[1 2]]
	// [[1 2] [4 5] [9]]
	// [[1] [3] [5] [7] [9]]
}

func ExampleUnionBy() {
	type person struct {
		Name string
		Age  int
	}
	nameEq := fp.Curry2(func(p1, p2 person) bool { return p1.Name == p2.Name })

	fp.Pipe2(
		slice.UnionBy(nameEq)([]person{}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 17},
		{"Larry", 20},
	})

	fp.Pipe2(
		slice.UnionBy(nameEq)([]person{
			{"Jerry", 21},
			{"Sam", 18},
			{"Jerry", 21},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{})

	fp.Pipe2(
		slice.UnionBy(nameEq)([]person{
			{"Jerry", 21},
			{"Sam", 18},
			{"Jerry", 21},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 17},
		{"Larry", 20},
	})

	// Output:
	// [{Jerry 17} {Larry 20}]
	// [{Jerry 21} {Sam 18} {Jerry 21}]
	// [{Jerry 21} {Sam 18} {Jerry 21} {Larry 20}]
}

func ExampleIntersectBy() {
	type person struct {
		Name string
		Age  int
	}
	nameEq := fp.Curry2(func(p1, p2 person) bool { return p1.Name == p2.Name })

	fp.Pipe2(
		slice.IntersectBy(nameEq)([]person{}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 17},
		{"Larry", 20},
	})

	fp.Pipe2(
		slice.IntersectBy(nameEq)([]person{
			{"Jerry", 21},
			{"Sam", 18},
			{"Jerry", 21},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{})

	fp.Pipe2(
		slice.IntersectBy(nameEq)([]person{
			{"Jerry", 21},
			{"Sam", 18},
			{"Jerry", 21},
		}),
		fp.Inspect(printAny[[]person]),
	)([]person{
		{"Jerry", 17},
		{"Larry", 20},
	})

	// Output:
	// []
	// []
	// [{Jerry 21} {Jerry 21}]
}

func ExampleInsertBy() {
	fp.Pipe2(
		slice.InsertBy(operator.Leq[int])(5),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.InsertBy(operator.Leq[int])(5),
		fp.Inspect(printAny[[]int]),
	)([]int{9, 10})

	fp.Pipe2(
		slice.InsertBy(operator.Leq[int])(5),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.InsertBy(operator.Leq[int])(5),
		fp.Inspect(printAny[[]int]),
	)([]int{3, 1, 5, 1, 9})

	// Output:
	// [5]
	// [5 9 10]
	// [1 2 5]
	// [3 1 5 5 1 9]
}

func ExampleMaximumBy() {
	type person struct {
		Name string
		Age  int
	}
	ageLt := fp.Curry2(func(p1, p2 person) bool { return p1.Age < p2.Age })

	fp.Pipe2(
		slice.MaximumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{})

	fp.Pipe2(
		slice.MaximumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{
		{"Jerry", 21},
		{"Barry", 25},
		{"Terry", 25},
	})

	fp.Pipe2(
		slice.MaximumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{
		{"Jerry", 21},
		{"Barry", 20},
		{"Terry", 25},
	})

	// Output:
	// None
	// Some {Barry 25}
	// Some {Terry 25}
}

func ExampleMinimumBy() {
	type person struct {
		Name string
		Age  int
	}
	ageLt := fp.Curry2(func(p1, p2 person) bool { return p1.Age < p2.Age })

	fp.Pipe2(
		slice.MinimumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{})

	fp.Pipe2(
		slice.MinimumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{
		{"Jerry", 24},
		{"Barry", 21},
		{"Terry", 25},
	})

	fp.Pipe2(
		slice.MinimumBy(ageLt),
		fp.Inspect(printAny[option.Option[person]]),
	)([]person{
		{"Jerry", 21},
		{"Barry", 20},
		{"Terry", 20},
	})

	// Output:
	// None
	// Some {Barry 21}
	// Some {Barry 20}
}

func ExampleHead() {
	fp.Pipe2(
		slice.Head[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Empty list returns None: %v\n", option.IsNone(o))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Head[int],
		option.Inspect(func(first int) {
			fmt.Printf("Only element of list: %d\n", first)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Head[int],
		option.Inspect(func(first int) {
			fmt.Printf("First element of list: %d\n", first)
		}),
	)([]int{2, 3, 4})

	// Output:
	// Empty list returns None: true
	// Only element of list: 1
	// First element of list: 2
}

func ExampleIndex() {
	type person struct {
		Name string
		Age  int
	}
	isJerry := func(p person) bool { return p.Name == "Jerry" }

	fp.Pipe2(
		slice.FindIndex(isJerry),
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Returns None when element doesn't exist: %v\n", option.IsNone(o))
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.FindIndex(isJerry),
		option.Inspect(func(i int) {
			fmt.Printf("Returns the index of the first matching element: %v\n", i)
		}),
	)([]person{{"Tim", 21}, {"Jerry", 30}, {"Jerry", 40}})

	// Output:
	// Returns None when element doesn't exist: true
	// Returns the index of the first matching element: 1
}

func ExampleFindIndices() {
	type person struct {
		Name string
		Age  int
	}
	isJerry := func(p person) bool { return p.Name == "Jerry" }

	fp.Pipe2(
		slice.FindIndices(isJerry),
		fp.Inspect(func(indexes []int) {
			fmt.Printf("Empty slice yields empty slice: %v\n", indexes)
		}),
	)([]person{})

	fp.Pipe2(
		slice.FindIndices(isJerry),
		fp.Inspect(func(indexes []int) {
			fmt.Printf("No matches yields empty slice: %v\n", indexes)
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.FindIndices(isJerry),
		fp.Inspect(func(indexes []int) {
			fmt.Printf("Yields indexes of all matches in slice: %v\n", indexes)
		}),
	)([]person{{"Tim", 21}, {"Jerry", 30}, {"Jerry", 40}})

	// Output:
	// Empty slice yields empty slice: []
	// No matches yields empty slice: []
	// Yields indexes of all matches in slice: [1 2]
}

func ExampleInit() {
	fp.Pipe2(
		slice.Init[int],
		fp.Inspect(func(o option.Option[[]int]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Init[int],
		option.Inspect(func(nums []int) {
			fmt.Printf("1 element slice returns: %v\n", nums)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Init[int],
		option.Inspect(func(nums []int) {
			fmt.Printf("Init of slice returns: %v\n", nums)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice returns None: true
	// 1 element slice returns: []
	// Init of slice returns: [1 2]
}

func ExampleInits() {
	fp.Pipe2(
		slice.Inits[int],
		fp.Inspect(func(inits [][]int) {
			fmt.Println(inits)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Inits[int],
		fp.Inspect(func(inits [][]int) {
			fmt.Println(inits)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [[]]
	// [[] [1] [1 2] [1 2 3]]
}

func ExampleIntersperse() {
	fp.Pipe2(
		slice.Intersperse(0),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Intersperse(0),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Intersperse(0),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// []
	// [1]
	// [1 0 2 0 3]
}

func ExampleIntercalate() {
	fp.Pipe2(
		slice.Intercalate([]rune{',', ' '}),
		fp.Inspect(func(runes []rune) {
			fmt.Println(string(runes))
		}),
	)([][]rune{
		{'f', 'o', 'o'},
	})

	fp.Pipe2(
		slice.Intercalate([]rune{',', ' '}),
		fp.Inspect(func(runes []rune) {
			fmt.Println(string(runes))
		}),
	)([][]rune{})

	fp.Pipe2(
		slice.Intercalate([]rune{',', ' '}),
		fp.Inspect(func(runes []rune) {
			fmt.Println(string(runes))
		}),
	)([][]rune{
		{'f', 'o', 'o'},
		{'b', 'a', 'r'},
	})

	// Output:
	// foo
	//
	// foo, bar
}

func ExampleIterate() {
	fp.Pipe2(
		slice.Iterate(operator.Multiply(2))(1),
		fp.Inspect(func(iters []int) {
			fmt.Println(iters)
		}),
	)(0)

	fp.Pipe2(
		slice.Iterate(operator.Multiply(2))(1),
		fp.Inspect(func(iters []int) {
			fmt.Println(iters)
		}),
	)(5)

	// Output:
	// []
	// [2 4 8 16 32]
}

func ExampleLast() {
	fp.Pipe2(
		slice.Last[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Empty list returns None: %v\n", option.IsNone(o))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Last[int],
		option.Inspect(func(last int) {
			fmt.Printf("Only element of list: %d\n", last)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Last[int],
		option.Inspect(func(last int) {
			fmt.Printf("Last element of list: %d\n", last)
		}),
	)([]int{2, 3, 4})

	// Output:
	// Empty list returns None: true
	// Only element of list: 1
	// Last element of list: 4

}

func ExampleLength() {
	fp.Pipe2(
		slice.Length[int],
		fp.Inspect(func(length int) {
			fmt.Printf("Empty slice has length: %d\n", length)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Length[int],
		fp.Inspect(func(length int) {
			fmt.Printf("Slice has length: %d\n", length)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice has length: 0
	// Slice has length: 3
}

func ExampleMap() {
	fp.Pipe2(
		slice.Map(func(x int) int { return x * 2 }),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Empty slice mapped: %v\n", xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Map(func(x int) int { return x * 2 }),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped x's: %v\n", xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice mapped: []
	// Mapped x's: [2 4 6]
}

func ExampleMaximum() {
	fp.Pipe2(
		slice.Maximum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Maximum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Maximum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{3, 2, 1})

	// Output:
	// None
	// Some 1
	// Some 3
}

func ExampleMinimum() {
	fp.Pipe2(
		slice.Minimum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Minimum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Minimum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{3, 2, 1})

	fp.Pipe2(
		slice.Minimum[int],
		fp.Inspect(func(o option.Option[int]) {
			fmt.Println(o)
		}),
	)([]int{1, 2, 3})

	// Output:
	// None
	// Some 1
	// Some 1
	// Some 1
}

func ExamplePermutations() {
	intSliceLt := fp.Curry2(func(s1, s2 []int) bool {
		for i, x := range s1 {
			if x < s2[i] {
				return true
			} else if s2[i] < x {
				return false
			}
		}
		return false
	})

	fp.Pipe3(
		slice.Permutations[int],
		slice.Sort(intSliceLt),
		fp.Inspect(func(perms [][]int) {
			fmt.Println(perms)
		}),
	)([]int{})

	fp.Pipe3(
		slice.Permutations[int],
		slice.Sort(intSliceLt),
		fp.Inspect(func(perms [][]int) {
			fmt.Println(perms)
		}),
	)([]int{1})

	fp.Pipe3(
		slice.Permutations[int],
		slice.Sort(intSliceLt),
		fp.Inspect(func(perms [][]int) {
			fmt.Println(perms)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [[]]
	// [[1]]
	// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}

func ExamplePrepend() {
	fp.Pipe2(
		slice.Prepend(5),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Prepend(5),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [5]
	// [5 1 2 3]
}

func ExamplePrependSlice() {
	fp.Pipe2(
		slice.PrependSlice([]int{}),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.PrependSlice([]int{}),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.PrependSlice([]int{5, 6}),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.PrependSlice([]int{5, 6}),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// []
	// [1 2 3]
	// [5 6]
	// [5 6 1 2 3]
}

func ExampleProduct() {
	fp.Pipe2(
		slice.Product[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Product[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{2})

	fp.Pipe2(
		slice.Product[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{2, 3, 4})

	// Output:
	// 1
	// 2
	// 24
}

func ExampleRange() {
	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 1 has no elements: %v\n", xs)
	})(slice.Range(1)(1))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 2 has no elements: %v\n", xs)
	})(slice.Range(3)(1))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 3 has positive elements: %v\n", xs)
	})(slice.Range(1)(3))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 4 has negative elements: %v\n", xs)
	})(slice.Range(-3)(-1))

	// Output:
	// Slice 1 has no elements: []
	// Slice 2 has no elements: []
	// Slice 3 has positive elements: [1 2]
	// Slice 4 has negative elements: [-3 -2]
}

func ExampleFoldl1() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Foldl1(appendStr),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldl1(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldl1(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar_baz
}

func ExampleFoldl1WithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.Foldl1WithIndex(appendStrAndIndex),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldl1WithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldl1WithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar1_baz2
}

func ExampleFoldl1WithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.Foldl1WithIndexAndSlice(appendStrAndIndexAndLength),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldl1WithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldl1WithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar13_baz23
}

func ExampleFoldr1() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Foldr1(appendStr),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldr1(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldr1(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar_foo
}

func ExampleFoldr1WithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.Foldr1WithIndex(appendStrAndIndex),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldr1WithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldr1WithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar1_foo0
}

func ExampleFoldr1WithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.Foldr1WithIndexAndSlice(appendStrAndIndexAndLength),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Foldr1WithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Foldr1WithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar13_foo03
}

func ExampleReplicate() {
	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 1 has no elements: %v\n", xs)
	})(slice.Replicate[int](0)(3))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 2 has no elements: %v\n", xs)
	})(slice.Replicate[int](-2)(3))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 3 has multiple elements: %v\n", xs)
	})(slice.Replicate[int](4)(3))

	// Output:
	// Slice 1 has no elements: []
	// Slice 2 has no elements: []
	// Slice 3 has multiple elements: [3 3 3 3]
}

func ExampleCycle() {
	fp.Pipe2(
		slice.Cycle[int](3),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Cycle[int](3),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Cycle[int](0),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Cycle[int](-1),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 3})

	// Output:
	// []
	// [1 2 3 1 2 3 1 2 3]
	// []
	// []
}

func ExampleMapAccumL() {
	trackAdd := fp.Curry2(func(acc int, x int) tuple.Pair[int, int] {
		return tuple.NewPair[int, int](acc + x)(acc)
	})

	fp.Pipe2(
		slice.MapAccumL(trackAdd)(0),
		fp.Inspect(printAny[tuple.Pair[int, []int]]),
	)([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	trackAppend := fp.Curry2(func(s string, x int) tuple.Pair[string, string] {
		return tuple.NewPair[string, string](fmt.Sprintf("%s%d", s, x))(s)
	})

	fp.Pipe2(
		slice.MapAccumL(trackAppend)("0"),
		fp.Inspect(printAny[tuple.Pair[string, []string]]),
	)([]int{1, 2, 3, 4, 5})

	// Output:
	// (55 [0 1 3 6 10 15 21 28 36 45])
	// (012345 [0 01 012 0123 01234])
}

func ExampleMapAccumR() {
	trackAdd := fp.Curry2(func(acc int, x int) tuple.Pair[int, int] {
		return tuple.NewPair[int, int](acc + x)(acc)
	})

	fp.Pipe2(
		slice.MapAccumR(trackAdd)(0),
		fp.Inspect(printAny[tuple.Pair[int, []int]]),
	)([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	trackAppend := fp.Curry2(func(s string, x int) tuple.Pair[string, string] {
		return tuple.NewPair[string, string](fmt.Sprintf("%s%d", s, x))(s)
	})

	fp.Pipe2(
		slice.MapAccumR(trackAppend)("0"),
		fp.Inspect(printAny[tuple.Pair[string, []string]]),
	)([]int{1, 2, 3, 4, 5})

	// Output:
	// (55 [54 52 49 45 40 34 27 19 10 0])
	// (054321 [05432 0543 054 05 0])
}

func ExampleUnfoldr() {
	buildUntil0 := func(x int) option.Option[tuple.Pair[int, int]] {
		if x == 0 {
			return option.None[tuple.Pair[int, int]]()
		}
		return option.Some(tuple.NewPair[int, int](x)(x - 1))
	}

	fp.Pipe2(
		slice.Unfoldr(buildUntil0),
		fp.Inspect(printAny[[]int]),
	)(5)

	// Output:
	// [5 4 3 2 1]
}

func ExampleReverse() {
	fp.Pipe2(
		slice.Reverse[int],
		fp.Inspect(func(reversed []int) {
			fmt.Printf("Empty slice reversed: %v\n", reversed)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Reverse[int],
		fp.Inspect(func(reversed []int) {
			fmt.Printf("Slice with one element reversed: %v\n", reversed)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Reverse[int],
		fp.Inspect(func(reversed []int) {
			fmt.Printf("Slice with multiple elements reversed: %v\n", reversed)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice reversed: []
	// Slice with one element reversed: [1]
	// Slice with multiple elements reversed: [3 2 1]
}

func ExampleScanl() {
	fp.Pipe2(
		slice.Scanl(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Scanl(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{3})

	fp.Pipe2(
		slice.Scanl(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [1]
	// [1 4]
	// [1 2 4 7]
}

func ExampleScanl1() {
	fp.Pipe2(
		slice.Scanl1(operator.Add[int]),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Scanl1(operator.Add[int]),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 3, 4})

	fp.Pipe2(
		slice.Scanl1(operator.And),
		fp.Inspect(printAny[[]bool]),
	)([]bool{true, false, true, true})

	// Output:
	// []
	// [1 3 6 10]
	// [true false false false]
}

func ExampleScanr() {
	fp.Pipe2(
		slice.Scanr(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Scanr(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{3})

	fp.Pipe2(
		slice.Scanr(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [1]
	// [4 1]
	// [7 6 4 1]
}

func ExampleScanr1() {
	fp.Pipe2(
		slice.Scanr1(operator.Add[int]),
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Scanr1(operator.Add[int]),
		fp.Inspect(printAny[[]int]),
	)([]int{1, 2, 3, 4})

	fp.Pipe2(
		slice.Scanr1(operator.And),
		fp.Inspect(printAny[[]bool]),
	)([]bool{true, false, true, true})

	// Output:
	// []
	// [10 9 7 4]
	// [false false true true]
}

func ExampleSingleton() {
	fp.Pipe2(
		slice.Singleton[int],
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)(1)

	// Output:
	// [1]
}

func ExampleAny() {
	is2 := func(x int) bool { return x == 2 }

	fp.Pipe2(
		slice.Any(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in empty slice: %v\n", some)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Any(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in slice 1: %v\n", some)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Any(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in slice 2: %v\n", some)
		}),
	)([]int{1, 3})

	fp.Pipe2(
		slice.Any(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in slice 3: %v\n", some)
		}),
	)([]int{2, 2})

	// Output:
	// Some element existed in empty slice: false
	// Some element existed in slice 1: true
	// Some element existed in slice 2: false
	// Some element existed in slice 3: true
}

func ExampleSort() {
	firstCharLT := fp.Curry2(func(s1, s2 string) bool { return strings.Compare(s1[0:1], s2[0:1]) < 0 })

	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(sorted []string) {
			fmt.Printf("Sorted empty slice: %v\n", sorted)
		}),
	)([]string{})

	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(sorted []string) {
			fmt.Printf("Sorted when already sorted: %v\n", sorted)
		}),
	)([]string{"a", "b", "c"})

	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(sorted []string) {
			fmt.Printf("Sorted when not sorted: %v\n", sorted)
		}),
	)([]string{"b", "a", "c"})

	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(sorted []string) {
			fmt.Printf("Stable when sorted 1: %v\n", sorted)
		}),
	)([]string{"b", "alpha", "apple", "c"})

	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(sorted []string) {
			fmt.Printf("Stable when sorted 2: %v\n", sorted)
		}),
	)([]string{"b", "apple", "alpha", "c"})

	strs := []string{"b", "apple", "alpha", "c"}
	fp.Pipe2(
		slice.Sort(firstCharLT),
		fp.Inspect(func(_ []string) {
			fmt.Printf("Does not modify input: %v\n", strs)
		}),
	)(strs)

	// Output:
	// Sorted empty slice: []
	// Sorted when already sorted: [a b c]
	// Sorted when not sorted: [a b c]
	// Stable when sorted 1: [alpha apple b c]
	// Stable when sorted 2: [apple alpha b c]
	// Does not modify input: [b apple alpha c]
}

func ExampleSpan() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.Span(isEven),
		fp.Inspect(func(span tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(span), tuple.Snd(span))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Span(isEven),
		fp.Inspect(func(span tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(span), tuple.Snd(span))
		}),
	)([]int{2, 4})

	fp.Pipe2(
		slice.Span(isEven),
		fp.Inspect(func(span tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(span), tuple.Snd(span))
		}),
	)([]int{3, 5})

	fp.Pipe2(
		slice.Span(isEven),
		fp.Inspect(func(span tuple.Pair[[]int, []int]) {
			fmt.Printf("[%v %v]\n", tuple.Fst(span), tuple.Snd(span))
		}),
	)([]int{2, 2, 3, 4, 5})

	// Output:
	// [[] []]
	// [[2 4] []]
	// [[] [3 5]]
	// [[2 2] [3 4 5]]
}

func ExampleSplitAt() {
	fp.Pipe2(
		slice.SplitAt[int](1),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Println(split)
		}),
	)([]int{})

	fp.Pipe2(
		slice.SplitAt[int](1),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Println(split)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.SplitAt[int](3),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Println(split)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.SplitAt[int](2),
		fp.Inspect(func(split tuple.Pair[[]int, []int]) {
			fmt.Println(split)
		}),
	)([]int{1, 2, 3, 4})

	// Output:
	// ([] [])
	// ([1] [])
	// ([1 2 3] [])
	// ([1 2] [3 4])
}

func ExampleSubsequences() {
	fp.Pipe2(
		slice.Subsequences[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Subsequences[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(xs)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Subsequences[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [[]]
	// [[] [1]]
	// [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}

func ExampleSum() {
	fp.Pipe2(
		slice.Sum[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Sum[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{4})

	fp.Pipe2(
		slice.Sum[int],
		fp.Inspect(func(x int) {
			fmt.Println(x)
		}),
	)([]int{1, 2, 3, 4})

	// Output:
	// 0
	// 4
	// 10
}

func ExampleTail() {
	fp.Pipe2(
		slice.Tail[int],
		fp.Inspect(func(o option.Option[[]int]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]int{})

	fp.Pipe2(
		slice.Tail[int],
		option.Inspect(func(nums []int) {
			fmt.Printf("1 element slice returns: %v\n", nums)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Tail[int],
		option.Inspect(func(nums []int) {
			fmt.Printf("Tail of slice returns: %v\n", nums)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Empty slice returns None: true
	// 1 element slice returns: []
	// Tail of slice returns: [2 3]
}

func ExampleTails() {
	fp.Pipe2(
		slice.Tails[int],
		fp.Inspect(func(tails [][]int) {
			fmt.Println(tails)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Tails[int],
		fp.Inspect(func(tails [][]int) {
			fmt.Println(tails)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [[]]
	// [[1 2 3] [2 3] [3] []]
}

func ExampleIsPrefixOf() {
	fp.Pipe2(
		slice.IsPrefixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{})

	fp.Pipe2(
		slice.IsPrefixOf([]int{}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsPrefixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsPrefixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3, 4})

	fp.Pipe2(
		slice.IsPrefixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{5, 6, 3, 4})

	// Output:
	// false
	// true
	// true
	// true
	// false
}

func ExampleIsSuffixOf() {
	fp.Pipe2(
		slice.IsSuffixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{})

	fp.Pipe2(
		slice.IsSuffixOf([]int{}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsSuffixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsSuffixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3, 4})

	fp.Pipe2(
		slice.IsSuffixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{4, 1, 2, 3})

	// Output:
	// false
	// true
	// true
	// false
	// true
}

func ExampleIsInfixOf() {
	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{})

	fp.Pipe2(
		slice.IsInfixOf([]int{}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2, 3, 4})

	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{4, 1, 2, 3})

	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{4, 1, 2, 3, 6, 7})

	fp.Pipe2(
		slice.IsInfixOf([]int{1, 2, 3}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 4, 2, 3})

	// Output:
	// false
	// true
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsSubsequenceOf() {
	fp.Pipe2(
		slice.IsSubsequenceOf([]int{1, 2}),
		fp.Inspect(printAny[bool]),
	)([]int{})

	fp.Pipe2(
		slice.IsSubsequenceOf([]int{}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2})

	fp.Pipe2(
		slice.IsSubsequenceOf([]int{1, 2}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 3, 2})

	fp.Pipe2(
		slice.IsSubsequenceOf([]int{1, 2}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 3, 4})

	fp.Pipe2(
		slice.IsSubsequenceOf([]int{1, 2}),
		fp.Inspect(printAny[bool]),
	)([]int{1, 2})

	// Output:
	// false
	// true
	// true
	// false
	// true
}

func ExampleTake() {
	fp.Pipe2(
		slice.Take[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Empty slice returns empty slice: %v\n", nums)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Take[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Returns first 2 elements slice: %v\n", nums)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Take[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Returns all elements of slice: %v\n", nums)
		}),
	)([]int{1, 2})

	fp.Pipe2(
		slice.Take[int](2),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Returns all elements of slice: %v\n", nums)
		}),
	)([]int{1})

	fp.Pipe2(
		slice.Take[int](-5),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Negative numbers return empty slice: %v\n", nums)
		}),
	)([]int{1, 2})

	// Output:
	// Empty slice returns empty slice: []
	// Returns first 2 elements slice: [1 2]
	// Returns all elements of slice: [1 2]
	// Returns all elements of slice: [1]
	// Negative numbers return empty slice: []
}

func ExampleTakeWhile() {
	isEven := func(x int) bool { return x%2 == 0 }

	fp.Pipe2(
		slice.TakeWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.TakeWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 3})

	fp.Pipe2(
		slice.TakeWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{2, 4, 5, 6, 7})

	fp.Pipe2(
		slice.TakeWhile(isEven),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 4, 5, 6, 7})

	// Output:
	// []
	// []
	// [2 4]
	// []
}

func ExampleTranspose() {
	pp := func(xs []int) string { return fmt.Sprintf("%v", xs) }
	pp2d := func(xs [][]int) string {
		return fmt.Sprintf("[%s]", strings.Join(slice.Map(pp)(xs), "\n "))
	}

	fp.Pipe2(
		slice.Transpose[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(pp2d(xs))
		}),
	)([][]int{})

	fp.Pipe2(
		slice.Transpose[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(pp2d(xs))
		}),
	)([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}})

	fp.Pipe2(
		slice.Transpose[int],
		fp.Inspect(func(xs [][]int) {
			fmt.Println(pp2d(xs))
		}),
	)([][]int{
		{1, 2, 3},
		{4},
		{7, 8, 9}})

	// Output:
	// []
	// [[1 4 7]
	//  [2 5 8]
	//  [3 6 9]]
	// [[1 4 7]
	//  [2 8]
	//  [3 9]]
}

func ExampleUncons() {
	fp.Pipe3(
		slice.Uncons[int],
		option.Unwrap[tuple.Pair[int, []int]],
		fp.Inspect(func(pair tuple.Pair[int, []int]) {
			fmt.Println(pair)
		}),
	)([]int{1})

	fp.Pipe3(
		slice.Uncons[int],
		option.Unwrap[tuple.Pair[int, []int]],
		fp.Inspect(func(pair tuple.Pair[int, []int]) {
			fmt.Println(pair)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Uncons[int],
		fp.Inspect(func(o option.Option[tuple.Pair[int, []int]]) {
			fmt.Println(o)
		}),
	)([]int{})

	// Output:
	// (1 [])
	// (1 [2 3])
	// None
}

func ExampleZip() {
	fp.Pipe2(
		slice.Zip[int, int]([]int{1, 2, 3}),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Println(pairs)
		}),
	)([]int{4, 5, 6})

	fp.Pipe2(
		slice.Zip[int, int]([]int{1, 2, 3}),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Printf("Extra elements in longer list are dropped: %v\n", pairs)
		}),
	)([]int{11, 12, 13, 14})

	fp.Pipe2(
		slice.Zip[int, int]([]int{1, 2, 3, 4}),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Printf("Extra elements in longer list are dropped: %v\n", pairs)
		}),
	)([]int{11, 12, 13})

	// Output:
	// [(1 4) (2 5) (3 6)]
	// Extra elements in longer list are dropped: [(1 11) (2 12) (3 13)]
	// Extra elements in longer list are dropped: [(1 11) (2 12) (3 13)]
}

func ExampleUnzip() {
	newPair := func(x int, s string) tuple.Pair[int, string] { return tuple.NewPair[int, string](x)(s) }

	fp.Pipe2(
		slice.Unzip[int, string],
		fp.Inspect(func(pair tuple.Pair[[]int, []string]) {
			fmt.Println(pair)
		}),
	)([]tuple.Pair[int, string]{})

	fp.Pipe2(
		slice.Unzip[int, string],
		fp.Inspect(func(pair tuple.Pair[[]int, []string]) {
			fmt.Println(pair)
		}),
	)([]tuple.Pair[int, string]{newPair(1, "foo"), newPair(2, "bar"), newPair(3, "baz")})

	// Output:
	// ([] [])
	// ([1 2 3] [foo bar baz])
}

func ExampleZipWith() {
	fp.Pipe2(
		slice.ZipWith(operator.Add[int])([]int{1, 2, 3}),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{4, 5, 6})

	fp.Pipe2(
		slice.ZipWith(operator.Add[int])([]int{1, 2, 3}),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Extra elements in longer list are dropped: %v\n", xs)
		}),
	)([]int{11, 12, 13, 14})

	fp.Pipe2(
		slice.ZipWith(operator.Add[int])([]int{1, 2, 3, 4}),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Extra elements in longer list are dropped: %v\n", xs)
		}),
	)([]int{11, 12, 13})

	// Output:
	// [5 7 9]
	// Extra elements in longer list are dropped: [12 14 16]
	// Extra elements in longer list are dropped: [12 14 16]
}

func ExampleZipIndices() {
	fp.Pipe3(
		slice.ZipIndices[int],
		slice.Length[tuple.Pair[int, int]],
		fp.Inspect(func(length int) {
			fmt.Printf("Zipping empty slice returns empty slice: %d\n", length)
		}),
	)([]int{})

	fp.Pipe2(
		slice.ZipIndices[int],
		fp.Inspect(func(zipped []tuple.Pair[int, int]) {
			fmt.Printf("Zipping slice maintains length: %d\n", slice.Length(zipped))
			option.Inspect(func(pair tuple.Pair[int, int]) {
				fmt.Printf("Index of first element: %d\n", tuple.Fst(pair))
				fmt.Printf("Value of first element: %d\n", tuple.Snd(pair))
			})(slice.Head(zipped))
		}),
	)([]int{1, 2, 3})

	fp.Pipe3(
		slice.ZipIndices[int],
		slice.Map(tuple.Fst[int, int]),
		fp.Inspect(func(indices []int) {
			fmt.Printf("Zipping allows for getting a slice of indices: %v\n", indices)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Zipping empty slice returns empty slice: 0
	// Zipping slice maintains length: 3
	// Index of first element: 0
	// Value of first element: 1
	// Zipping allows for getting a slice of indices: [0 1 2]
}

func ExampleLines() {
	printLenAndElems := func(ss []string) {
		fmt.Printf("%d, %v\n", len(ss), ss)
	}

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("\n")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("one")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("one\n")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("one\n\n")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("one\ntwo")

	fp.Pipe2(
		slice.Lines,
		fp.Inspect(printLenAndElems),
	)("one\ntwo\n")

	// Output:
	// 0, []
	// 1, []
	// 1, [one]
	// 1, [one]
	// 2, [one ]
	// 2, [one two]
	// 2, [one two]
}

func ExampleWords() {
	fp.Pipe2(
		slice.Words,
		fp.Inspect(printAny[[]string]),
	)("")

	fp.Pipe2(
		slice.Words,
		fp.Inspect(printAny[[]string]),
	)("  \n")

	fp.Pipe2(
		slice.Words,
		fp.Inspect(printAny[[]string]),
	)(" hello\nworld foo \t bar ")

	// Output:
	// []
	// []
	// [hello world foo bar]
}

func ExampleUnlines() {
	fp.Pipe3(
		slice.Unlines,
		operator.Eq(""),
		fp.Inspect(printAny[bool]),
	)([]string{})

	fp.Pipe3(
		slice.Unlines,
		operator.Eq("foo\n"),
		fp.Inspect(printAny[bool]),
	)([]string{
		"foo",
	})

	fp.Pipe3(
		slice.Unlines,
		operator.Eq("foo\nbar\n"),
		fp.Inspect(printAny[bool]),
	)([]string{
		"foo",
		"bar",
	})

	// Output:
	// true
	// true
	// true
}

func ExampleUnwords() {
	fp.Pipe2(
		slice.Unwords,
		fp.Inspect(printAny[string]),
	)([]string{})

	fp.Pipe2(
		slice.Unwords,
		fp.Inspect(printAny[string]),
	)([]string{
		"foo",
	})

	fp.Pipe2(
		slice.Unwords,
		fp.Inspect(printAny[string]),
	)([]string{
		"foo",
		"bar",
	})

	// Output:
	//
	// foo
	// foo bar
}

func ExampleUnique() {
	fp.Pipe2(
		slice.Unique[int],
		fp.Inspect(printAny[[]int]),
	)([]int{})

	fp.Pipe2(
		slice.Unique[int],
		fp.Inspect(printAny[[]int]),
	)([]int{1})

	fp.Pipe2(
		slice.Unique[int],
		fp.Inspect(printAny[[]int]),
	)([]int{1, 1, 4, 3, 4, 4})

	// Output:
	// []
	// [1]
	// [1 4 3]
}
