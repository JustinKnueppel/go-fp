package slice_test

import (
	"fmt"
	"math"
	"strings"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

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

func ExampleContains() {
	fp.Pipe2(
		slice.Contains(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Empty slice contained target: %v\n", contained)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Contains(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Slice 1 contained target: %v\n", contained)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Contains(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Slice 2 contained target: %v\n", contained)
		}),
	)([]int{1, 3})

	// Output:
	// Empty slice contained target: false
	// Slice 1 contained target: true
	// Slice 2 contained target: false
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

func ExampleDeleteBy() {
	type person struct {
		Name string
		Age  int
	}
	isJerry := func(p person) bool { return p.Name == "Jerry" }

	fp.Pipe2(
		slice.DeleteBy(isJerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy on empty slice returns empty slice: %v\n", people)
		}),
	)([]person{})

	fp.Pipe2(
		slice.DeleteBy(isJerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy returns input when element not found: %v\n", people)
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.DeleteBy(isJerry),
		fp.Inspect(func(people []person) {
			fmt.Printf("DeleteBy only removes first instance: %v\n", people)
		}),
	)([]person{{"Jerry", 21}, {"Tim", 30}, {"Jerry", 40}})

	// Output:
	// DeleteBy on empty slice returns empty slice: []
	// DeleteBy returns input when element not found: [{Tim 21}]
	// DeleteBy only removes first instance: [{Tim 30} {Jerry 40}]
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

func ExampleIsEmpty() {
	fp.Pipe2(
		slice.IsEmpty[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Empty slice is empty: %v\n", empty)
		}),
	)([]int{})

	fp.Pipe2(
		slice.IsEmpty[int],
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

func ExampleEvery() {
	is2 := func(x int) bool { return x == 2 }

	fp.Pipe2(
		slice.Every(is2),
		fp.Inspect(func(every bool) {
			fmt.Printf("Every element passed in empty list: %v\n", every)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Every(is2),
		fp.Inspect(func(every bool) {
			fmt.Printf("Every element passed in slice 1: %v\n", every)
		}),
	)([]int{2, 2})

	fp.Pipe2(
		slice.Every(is2),
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

func ExampleFlatMap() {
	repeatN := func(n int) []int { return slice.Repeat[int](n)(n) }

	fp.Pipe2(
		slice.FlatMap(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped empty slice: %v\n", xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.FlatMap(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped slice where one map is empty: %v\n", xs)
		}),
	)([]int{1, 0, 3})

	fp.Pipe2(
		slice.FlatMap(repeatN),
		fp.Inspect(func(xs []int) {
			fmt.Printf("Mapped slice: %v\n", xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// Mapped empty slice: []
	// Mapped slice where one map is empty: [1 3 3 3]
	// Mapped slice: [1 2 2 3 3 3]
}

func ExampleFold() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Fold(appendStr)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.Fold(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Fold(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo
	// Multiple elements appended: start_foo_bar_baz
}

func ExampleFoldWithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.FoldWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo0
	// Multiple elements appended: start_foo0_bar1_baz2
}

func ExampleFoldWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.FoldWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo01
	// Multiple elements appended: start_foo03_bar13_baz23
}

func ExampleFoldRight() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.FoldRight(appendStr)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldRight(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldRight(appendStr)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo
	// Multiple elements appended: start_baz_bar_foo
}

func ExampleFoldRightWithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.FoldRightWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldRightWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldRightWithIndex(appendStrAndIndex)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo0
	// Multiple elements appended: start_baz2_bar1_foo0
}

func ExampleFoldRightWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.FoldRightWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(init string) {
			fmt.Printf("Empty slice returns initial value: %s\n", init)
		}),
	)([]string{})

	fp.Pipe2(
		slice.FoldRightWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.FoldRightWithIndexAndSlice(appendStrAndIndexAndLength)("start"),
		fp.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns initial value: start
	// One element: start_foo01
	// Multiple elements appended: start_baz23_bar13_foo03
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
		slice.Index(isJerry),
		fp.Inspect(func(o option.Option[int]) {
			fmt.Printf("Returns None when element doesn't exist: %v\n", option.IsNone(o))
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.Index(isJerry),
		option.Inspect(func(i int) {
			fmt.Printf("Returns the index of the first matching element: %v\n", i)
		}),
	)([]person{{"Tim", 21}, {"Jerry", 30}, {"Jerry", 40}})

	// Output:
	// Returns None when element doesn't exist: true
	// Returns the index of the first matching element: 1
}

func ExampleIndexes() {
	type person struct {
		Name string
		Age  int
	}
	isJerry := func(p person) bool { return p.Name == "Jerry" }

	fp.Pipe2(
		slice.Indexes(isJerry),
		fp.Inspect(func(indexes []int) {
			fmt.Printf("Empty slice yields empty slice: %v\n", indexes)
		}),
	)([]person{})

	fp.Pipe2(
		slice.Indexes(isJerry),
		fp.Inspect(func(indexes []int) {
			fmt.Printf("No matches yields empty slice: %v\n", indexes)
		}),
	)([]person{{"Tim", 21}})

	fp.Pipe2(
		slice.Indexes(isJerry),
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

func ExampleReduce() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.Reduce(appendStr),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.Reduce(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.Reduce(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar_baz
}

func ExampleReduceWithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.ReduceWithIndex(appendStrAndIndex),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.ReduceWithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.ReduceWithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar1_baz2
}

func ExampleReduceWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.ReduceWithIndexAndSlice(appendStrAndIndexAndLength),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.ReduceWithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.ReduceWithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: foo_bar13_baz23
}

func ExampleReduceRight() {
	appendStr := fp.Curry2(func(x, y string) string { return x + "_" + y })

	fp.Pipe2(
		slice.ReduceRight(appendStr),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.ReduceRight(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.ReduceRight(appendStr),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar_foo
}

func ExampleReduceRightWithIndex() {
	appendStrAndIndex := fp.Curry3(func(x, y string, i int) string { return x + "_" + y + fmt.Sprint(i) })

	fp.Pipe2(
		slice.ReduceRightWithIndex(appendStrAndIndex),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.ReduceRightWithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.ReduceRightWithIndex(appendStrAndIndex),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar1_foo0
}

func ExampleReduceRightWithIndexAndSlice() {
	appendStrAndIndexAndLength := fp.Curry4(func(x, y string, i int, xs []string) string { return x + "_" + y + fmt.Sprintf("%d%d", i, len(xs)) })

	fp.Pipe2(
		slice.ReduceRightWithIndexAndSlice(appendStrAndIndexAndLength),
		fp.Inspect(func(o option.Option[string]) {
			fmt.Printf("Empty slice returns None: %v\n", option.IsNone(o))
		}),
	)([]string{})

	fp.Pipe2(
		slice.ReduceRightWithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("One element: %s\n", appended)
		}),
	)([]string{"foo"})

	fp.Pipe2(
		slice.ReduceRightWithIndexAndSlice(appendStrAndIndexAndLength),
		option.Inspect(func(appended string) {
			fmt.Printf("Multiple elements appended: %s\n", appended)
		}),
	)([]string{"foo", "bar", "baz"})

	// Output:
	// Empty slice returns None: true
	// One element: foo
	// Multiple elements appended: baz_bar13_foo03
}

func ExampleRepeat() {
	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 1 has no elements: %v\n", xs)
	})(slice.Repeat[int](0)(3))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 2 has no elements: %v\n", xs)
	})(slice.Repeat[int](-2)(3))

	fp.Inspect(func(xs []int) {
		fmt.Printf("Slice 3 has multiple elements: %v\n", xs)
	})(slice.Repeat[int](4)(3))

	// Output:
	// Slice 1 has no elements: []
	// Slice 2 has no elements: []
	// Slice 3 has multiple elements: [3 3 3 3]
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

func ExampleScan() {
	fp.Pipe2(
		slice.Scan(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Scan(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{3})

	fp.Pipe2(
		slice.Scan(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [1]
	// [1 4]
	// [1 2 4 7]
}

func ExampleScanRight() {
	fp.Pipe2(
		slice.ScanRight(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{})

	fp.Pipe2(
		slice.ScanRight(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{3})

	fp.Pipe2(
		slice.ScanRight(operator.Add[int])(1),
		fp.Inspect(func(xs []int) {
			fmt.Println(xs)
		}),
	)([]int{1, 2, 3})

	// Output:
	// [1]
	// [4 1]
	// [7 6 4 1]
}

func ExampleSome() {
	is2 := func(x int) bool { return x == 2 }

	fp.Pipe2(
		slice.Some(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in empty slice: %v\n", some)
		}),
	)([]int{})

	fp.Pipe2(
		slice.Some(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in slice 1: %v\n", some)
		}),
	)([]int{1, 2, 3})

	fp.Pipe2(
		slice.Some(is2),
		fp.Inspect(func(some bool) {
			fmt.Printf("Some element existed in slice 2: %v\n", some)
		}),
	)([]int{1, 3})

	fp.Pipe2(
		slice.Some(is2),
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

	// Output:
	// [(1 4) (2 5) (3 6)]
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

	// Output:
	// [5 7 9]
	// Extra elements in longer list are dropped: [12 14 16]
}

func ExampleZipIndexes() {
	fp.Pipe3(
		slice.ZipIndexes[int],
		slice.Length[tuple.Pair[int, int]],
		fp.Inspect(func(length int) {
			fmt.Printf("Zipping empty slice returns empty slice: %d\n", length)
		}),
	)([]int{})

	fp.Pipe2(
		slice.ZipIndexes[int],
		fp.Inspect(func(zipped []tuple.Pair[int, int]) {
			fmt.Printf("Zipping slice maintains length: %d\n", slice.Length(zipped))
			option.Inspect(func(pair tuple.Pair[int, int]) {
				fmt.Printf("Index of first element: %d\n", tuple.Fst(pair))
				fmt.Printf("Value of first element: %d\n", tuple.Snd(pair))
			})(slice.Head(zipped))
		}),
	)([]int{1, 2, 3})

	fp.Pipe3(
		slice.ZipIndexes[int],
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
