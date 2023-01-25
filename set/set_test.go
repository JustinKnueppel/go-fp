package set_test

import (
	"fmt"
	"testing"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/set"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

func TestString(t *testing.T) {
	s := set.FromSlice([]int{1})
	if s.String() != "{1}" {
		t.Fatal()
	}
}

func ExampleInsert() {
	fp.Pipe2(
		set.Insert(2),
		fp.Inspect(func(s set.Set[int]) {
			fmt.Printf("Empty set now contains new value: %v\n", set.Equal(set.FromSlice([]int{2}))(s))
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Insert(2),
		fp.Inspect(func(s set.Set[int]) {
			fmt.Printf("Set now contains new value: %v\n", set.Equal(set.FromSlice([]int{1, 2}))(s))
		}),
	)(set.FromSlice([]int{1}))

	fp.Pipe2(
		set.Insert(1),
		fp.Inspect(func(s set.Set[int]) {
			fmt.Printf("Set does not change size when duplicate added: %v\n", set.Equal(set.FromSlice([]int{1}))(s))
		}),
	)(set.FromSlice([]int{1}))

	// Output:
	// Empty set now contains new value: true
	// Set now contains new value: true
	// Set does not change size when duplicate added: true
}

func ExampleCartesianProduct() {
	pairsLt := fp.Curry2(func(p1, p2 tuple.Pair[int, int]) bool {
		return tuple.Fst(p1) < tuple.Fst(p2) ||
			(tuple.Fst(p1) == tuple.Fst(p2) && tuple.Snd(p1) < tuple.Snd(p2))
	})

	fp.Pipe4(
		set.CartesianProduct[int, int](set.FromSlice([]int{1, 2, 3})),
		set.ToSlice[tuple.Pair[int, int]],
		slice.Sort(pairsLt),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Println(pairs)
		}),
	)(set.FromSlice([]int{}))

	fp.Pipe4(
		set.CartesianProduct[int, int](set.FromSlice([]int{1, 2})),
		set.ToSlice[tuple.Pair[int, int]],
		slice.Sort(pairsLt),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Println(pairs)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe4(
		set.CartesianProduct[int, int](set.FromSlice([]int{1, 2, 3})),
		set.ToSlice[tuple.Pair[int, int]],
		slice.Sort(pairsLt),
		fp.Inspect(func(pairs []tuple.Pair[int, int]) {
			fmt.Println(pairs)
		}),
	)(set.FromSlice([]int{4, 5, 6}))

	// Output:
	// []
	// [(1 1) (1 2) (2 1) (2 2)]
	// [(1 4) (1 5) (1 6) (2 4) (2 5) (2 6) (3 4) (3 5) (3 6)]
}

func ExampleCopy() {
	s := set.FromSlice([]int{1, 2})
	fp.Pipe2(
		set.Copy[int],
		fp.Inspect(func(newSet set.Set[int]) {
			fmt.Printf("Two sets are equal: %v\n", set.Equal(newSet)(s))
			fmt.Printf("Two sets have same reference: %v\n", &newSet == &s)
			set.Insert(3)(newSet)
			fmt.Printf("Old set was not affected by add: %v\n", set.Equal(set.FromSlice([]int{1, 2}))(s))
		}),
	)(s)

	// Output:
	// Two sets are equal: true
	// Two sets have same reference: false
	// Old set was not affected by add: true
}

func ExampleMember() {
	fp.Pipe2(
		set.Member(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Empty set contained element: %v\n", contained)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Member(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Set 1 contained element: %v\n", contained)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.Member(2),
		fp.Inspect(func(contained bool) {
			fmt.Printf("Set 2 contained element: %v\n", contained)
		}),
	)(set.FromSlice([]int{1, 3}))

	// Output:
	// Empty set contained element: false
	// Set 1 contained element: true
	// Set 2 contained element: false
}

func ExampleDifference() {
	fp.Pipe3(
		set.Difference(set.Empty[int]()),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Difference between empty sets is empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Difference(set.Empty[int]()),
		set.Equal(set.FromSlice([]int{1})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Difference between set and empty set is original set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1}))

	fp.Pipe3(
		set.Difference(set.FromSlice([]int{1})),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Difference between empty set and non-empty set is empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Difference(set.FromSlice([]int{1, 2})),
		set.Equal(set.FromSlice([]int{3})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Difference removes elements of other set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	// Output:
	// Difference between empty sets is empty set: true
	// Difference between set and empty set is original set: true
	// Difference between empty set and non-empty set is empty set: true
	// Difference removes elements of other set: true
}

func ExampleDisjoint() {
	fp.Pipe2(
		set.Disjoint(set.Empty[int]()),
		fp.Inspect(func(disjoint bool) {
			fmt.Printf("Two empty sets are disjoint: %v\n", disjoint)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Disjoint(set.Empty[int]()),
		fp.Inspect(func(disjoint bool) {
			fmt.Printf("Any set is disjoint with the empty set: %v\n", disjoint)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	fp.Pipe2(
		set.Disjoint(set.FromSlice([]int{2})),
		fp.Inspect(func(disjoint bool) {
			fmt.Printf("Two sets are disjoint if they share no elements: %v\n", disjoint)
		}),
	)(set.FromSlice([]int{1, 3}))

	fp.Pipe2(
		set.Disjoint(set.FromSlice([]int{2, 3})),
		fp.Inspect(func(disjoint bool) {
			fmt.Printf("Two sets are not disjoint if they share elements: %v\n", disjoint)
		}),
	)(set.FromSlice([]int{1, 3}))

	// Output:
	// Two empty sets are disjoint: true
	// Any set is disjoint with the empty set: true
	// Two sets are disjoint if they share no elements: true
	// Two sets are not disjoint if they share elements: false
}

func ExampleEqual() {
	fp.Pipe2(
		set.Equal(set.Empty[int]()),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Two empty sets are equal: %v\n", equal)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Two sets with the same elements are equal: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Order does not affect equality: %v\n", equal)
		}),
	)(set.FromSlice([]int{2, 1}))

	fp.Pipe2(
		set.Equal(set.FromSlice([]int{3, 4})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Two sets with different elements are not equal: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.Equal(set.FromSlice([]int{1})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Strict subsets are not equal: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Strict supersets are not equal: %v\n", equal)
		}),
	)(set.FromSlice([]int{1}))

	// Output:
	// Two empty sets are equal: true
	// Two sets with the same elements are equal: true
	// Order does not affect equality: true
	// Two sets with different elements are not equal: false
	// Strict subsets are not equal: false
	// Strict supersets are not equal: false
}

func ExampleFilter() {
	greaterThan2 := func(x int) bool { return x > 2 }

	fp.Pipe3(
		set.Filter(greaterThan2),
		set.Equal(set.Empty[int]()),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Empty set returns empty set: %v\n", equal)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Filter(greaterThan2),
		set.Equal(set.FromSlice([]int{3})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Elements which fail to satisfy the predicate are removed: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	// Output:
	// Empty set returns empty set: true
	// Elements which fail to satisfy the predicate are removed: true
}

func ExampleFromSlice() {
	fp.Pipe3(
		set.FromSlice[int],
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Empty slice creates empty set: %v\n", empty)
		}),
	)([]int{})

	fp.Pipe3(
		set.FromSlice[int],
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("All elements of slice are in set: %v\n", equal)
		}),
	)([]int{1, 2})

	fp.Pipe3(
		set.FromSlice[int],
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Order of elements does not matter: %v\n", equal)
		}),
	)([]int{2, 1})

	fp.Pipe3(
		set.FromSlice[int],
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Duplicate elements are present once: %v\n", equal)
		}),
	)([]int{1, 2, 2})

	// Output:
	// Empty slice creates empty set: true
	// All elements of slice are in set: true
	// Order of elements does not matter: true
	// Duplicate elements are present once: true
}

func ExampleIntersection() {
	fp.Pipe3(
		set.Intersection(set.Empty[int]()),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Intersection of two empty sets is empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Intersection(set.Empty[int]()),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Intersection with empty set is empty set: %v\n", empty)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	fp.Pipe3(
		set.Intersection(set.FromSlice([]int{1, 2})),
		set.Equal(set.FromSlice([]int{2})),
		fp.Inspect(func(empty bool) {
			fmt.Printf("Intersection contains overlapping entries: %v\n", empty)
		}),
	)(set.FromSlice([]int{2, 3}))

	// Output:
	// Intersection of two empty sets is empty set: true
	// Intersection with empty set is empty set: true
	// Intersection contains overlapping entries: true
}

func ExampleNull() {
	fp.Pipe2(
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Set with no elements is empty: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Set with elements is not empty: %v\n", empty)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Set with no elements is empty: true
	// Set with elements is not empty: false
}

func ExampleIsProperSubsetOf() {
	fp.Pipe2(
		set.IsProperSubsetOf(set.Empty[int]()),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Empty set is a strict subset of empty set: %v\n", subset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsProperSubsetOf(set.Empty[int]()),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Empty set is a strict subset of any non-empty set: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsProperSubsetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Equal set is a strict subset: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsProperSubsetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Set is a strict subset if all elements exist in base set: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	fp.Pipe2(
		set.IsProperSubsetOf(set.FromSlice([]int{1, 2, 3})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Superset is not a strict subset: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Empty set is a strict subset of empty set: false
	// Empty set is a strict subset of any non-empty set: true
	// Equal set is a strict subset: false
	// Set is a strict subset if all elements exist in base set: true
	// Superset is not a strict subset: false
}

func ExampleIsProperSupersetOf() {
	fp.Pipe2(
		set.IsProperSupersetOf(set.Empty[int]()),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Empty set is a strict superset of empty set: %v\n", superset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsProperSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Any non-empty set is a strict superset of empty set: %v\n", superset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsProperSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Equal set is a strict superset: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsProperSupersetOf(set.FromSlice([]int{1, 2, 3})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Set is a strict superset if all elements of base set exist in set: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsProperSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Subset is not a strict superset: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	// Output:
	// Empty set is a strict superset of empty set: false
	// Any non-empty set is a strict superset of empty set: true
	// Equal set is a strict superset: false
	// Set is a strict superset if all elements of base set exist in set: true
	// Subset is not a strict superset: false
}

func ExampleIsSubsetOf() {
	fp.Pipe2(
		set.IsSubsetOf(set.Empty[int]()),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Empty set is a subset of empty set: %v\n", subset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsSubsetOf(set.Empty[int]()),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Empty set is a subset of any set: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsSubsetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Equal set is a subset: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsSubsetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Set is a subset if all elements exist in base set: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	fp.Pipe2(
		set.IsSubsetOf(set.FromSlice([]int{1, 2, 3})),
		fp.Inspect(func(subset bool) {
			fmt.Printf("Superset is not a subset: %v\n", subset)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Empty set is a subset of empty set: true
	// Empty set is a subset of any set: true
	// Equal set is a subset: true
	// Set is a subset if all elements exist in base set: true
	// Superset is not a subset: false
}

func ExampleIsSupersetOf() {
	fp.Pipe2(
		set.IsSupersetOf(set.Empty[int]()),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Empty set is a superset of empty set: %v\n", superset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Empty set is a superset of any set: %v\n", superset)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.IsSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Equal set is a superset: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsSupersetOf(set.FromSlice([]int{1, 2, 3})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Set is a superset if all elements in base set exist in the set: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe2(
		set.IsSupersetOf(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(superset bool) {
			fmt.Printf("Subset is not a superset: %v\n", superset)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	// Output:
	// Empty set is a superset of empty set: true
	// Empty set is a superset of any set: true
	// Equal set is a superset: true
	// Set is a superset if all elements in base set exist in the set: true
	// Subset is not a superset: false
}

func ExampleMap() {
	toString := func(x int) string { return fmt.Sprint(x) }
	even := func(x int) bool { return x%2 == 0 }

	fp.Pipe3(
		set.Map(toString),
		set.Null[string],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Map of empty set is empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Map(toString),
		set.Equal(set.FromSlice([]string{"1", "2"})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Map affects all elements of set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe3(
		set.Map(even),
		set.Equal(set.FromSlice([]bool{true, false})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Map can reduce size of set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2, 3, 4}))

	// Output:
	// Map of empty set is empty set: true
	// Map affects all elements of set: true
	// Map can reduce size of set: true
}

func ExampleEmpty() {
	fp.Pipe2(
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("No arguments creates empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Inspect(func(s set.Set[int]) {
		fmt.Printf("New set contains 1: %v\n", set.Member(1)(s))
		fmt.Printf("New set contains 2: %v\n", set.Member(2)(s))
		fmt.Printf("New set has size 2: %d\n", set.Size(s))
	})(set.FromSlice([]int{1, 2}))

	fp.Inspect(func(s set.Set[int]) {
		fmt.Printf("Duplicates are removed: %d\n", set.Size(s))
	})(set.FromSlice([]int{1, 2, 2}))

	// Output:
	// No arguments creates empty set: true
	// New set contains 1: true
	// New set contains 2: true
	// New set has size 2: 2
	// Duplicates are removed: 2
}

func ExampleDelete() {
	fp.Pipe3(
		set.Delete(1),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Removing element from empty set returns empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Delete(1),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Removing last element from set returns empty set: %v\n", empty)
		}),
	)(set.FromSlice([]int{1}))

	fp.Pipe3(
		set.Delete(1),
		set.Equal(set.FromSlice([]int{2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Removing element from set returns remaining set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Removing element from empty set returns empty set: true
	// Removing last element from set returns empty set: true
	// Removing element from set returns remaining set: true
}

func ExampleSize() {
	fp.Pipe2(
		set.Size[int],
		fp.Inspect(func(size int) {
			fmt.Printf("Empty set has size: %d\n", size)
		}),
	)(set.Empty[int]())

	fp.Pipe2(
		set.Size[int],
		fp.Inspect(func(size int) {
			fmt.Printf("Non-empty set has size: %d\n", size)
		}),
	)(set.FromSlice([]int{1, 2, 3}))

	// Output:
	// Empty set has size: 0
	// Non-empty set has size: 3
}

func ExampleToSlice() {
	fp.Pipe2(
		set.ToSlice[int],
		fp.Inspect(func(nums []int) {
			fmt.Printf("Empty set returns empty slice: %v\n", nums)
		}),
	)(set.Empty[int]())

	ltComparator := fp.Curry2(func(x1, x2 int) bool { return x1 < x2 })
	fp.Pipe3(
		set.ToSlice[int],
		slice.Sort(ltComparator),
		fp.Inspect(func(nums []int) {
			fmt.Printf("Set returns slice in non-guaranteed order: %v\n", nums)
		}),
	)(set.FromSlice([]int{2, 3, 1}))

	// Output:
	// Empty set returns empty slice: []
	// Set returns slice in non-guaranteed order: [1 2 3]
}

func ExampleUnion() {
	fp.Pipe3(
		set.Union(set.Empty[int]()),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Union of empty sets is empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Union(set.Empty[int]()),
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Union with empty set is original set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe3(
		set.Union(set.FromSlice([]int{1, 2})),
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Union with self is original set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe3(
		set.Union(set.FromSlice([]int{3, 4})),
		set.Equal(set.FromSlice([]int{1, 2, 3, 4})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Union with other set keeps all elements from both: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Union of empty sets is empty set: true
	// Union with empty set is original set: true
	// Union with self is original set: true
	// Union with other set keeps all elements from both: true
}

func ExampleXor() {
	fp.Pipe3(
		set.Xor(set.Empty[int]()),
		set.Null[int],
		fp.Inspect(func(empty bool) {
			fmt.Printf("Xor with empty sets yields empty set: %v\n", empty)
		}),
	)(set.Empty[int]())

	fp.Pipe3(
		set.Xor(set.Empty[int]()),
		set.Equal(set.FromSlice([]int{1, 2})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Xor with empty set yields original set: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	fp.Pipe3(
		set.Xor(set.FromSlice([]int{2, 3})),
		set.Equal(set.FromSlice([]int{1, 3})),
		fp.Inspect(func(equal bool) {
			fmt.Printf("Xor excludes common elements: %v\n", equal)
		}),
	)(set.FromSlice([]int{1, 2}))

	// Output:
	// Xor with empty sets yields empty set: true
	// Xor with empty set yields original set: true
	// Xor excludes common elements: true
}
