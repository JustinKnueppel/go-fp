package set

import (
	"fmt"
	"sort"
	"strings"

	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/operator"
	"github.com/JustinKnueppel/go-fp/tuple"
)

/* ============ Set type ============ */

// Set represents a collection of unique elements.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// String is used only for properly printing a Set.
func (s Set[T]) String() string {
	strSlice := []string{}
	for t := range s.elements {
		strSlice = append(strSlice, fmt.Sprintf("%v", t))
	}
	return fmt.Sprintf("{%s}", strings.Join(strSlice, " "))
}

/* ============ Basic functions ============ */

// Copy returns a value copy of the given set.
func Copy[T comparable](set Set[T]) Set[T] {
	elems := make(map[T]struct{})
	for elem := range set.elements {
		elems[elem] = struct{}{}
	}
	return Set[T]{elems}
}

// Equal returns true if both sets have the same elements.
func Equal[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				return false
			}
		}
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				return false
			}
		}
		return true
	}
}

/* ============ Construction ============ */

// Empty creates an empty Set.
func Empty[T comparable]() Set[T] {
	elems := make(map[T]struct{})
	return Set[T]{elems}
}

// Singleton returns a Set with the single element contained.
func Singleton[T comparable](t T) Set[T] {
	elems := make(map[T]struct{})
	return Insert(t)(Set[T]{elems})
}

// FromSlice creates a set from the given slice.
func FromSlice[T comparable](init []T) Set[T] {
	elems := make(map[T]struct{})
	for _, t := range init {
		elems[t] = struct{}{}
	}
	return Set[T]{elems}
}

// Powerset returns a slice of all subsets.
func PowerSet[T comparable](s Set[T]) []Set[T] {
	if Null(s) {
		return []Set[T]{{}}
	}

	x := ToSlice(s)[0]
	xs := Delete(x)(s)
	powerSetRest := PowerSet(xs)
	for _, y := range powerSetRest {
		powerSetRest = append(powerSetRest, Insert(x)(y))
	}
	return powerSetRest
}

/* ============ Insertion ============ */

// Insert returns a union of the given set and element.
func Insert[T comparable](elem T) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		out.elements[elem] = struct{}{}
		return out
	}
}

/* ============ Deletion ============ */

// Delete returns the Set with the given element removed.
func Delete[T comparable](elem T) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		delete(out.elements, elem)
		return out
	}
}

/* ============ Query ============ */

// Member returns true if the set contains the target
// otherwise returns false.
func Member[T comparable](target T) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		_, found := s.elements[target]
		return found
	}
}

// NotMember returns true if the set does not contain the target.
func NotMember[T comparable](target T) func(s Set[T]) bool {
	return fp.Compose2(operator.Not, Member(target))
}

// Null returns true if the Set has no elements.
func Null[T comparable](s Set[T]) bool {
	return Size(s) == 0
}

// Size returns the number of elements in the Set.
func Size[T comparable](s Set[T]) int {
	return len(s.elements)
}

// IsSubsetOf returns true if every value of the other set exists
// in the given set.
func IsSubsetOf[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				return false
			}
		}
		return true
	}
}

// IsProperSubsetOf returns true if every value of the other set exists
// in the given set but the sets are not equal.
func IsProperSubsetOf[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				return false
			}
		}
		return Size(s) != Size(other)
	}
}

// IsSupersetOf returns true if every value of the given set exists
// in the other set.
func IsSupersetOf[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				return false
			}
		}
		return true
	}
}

// IsProperSupersetOf returns true if every value of the given set exists
// in the other set but the sets are not equal.
func IsProperSupersetOf[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				return false
			}
		}
		return Size(s) != Size(other)
	}
}

// Disjoint returns true if the two sets have an empty intersection.
func Disjoint[T comparable](other Set[T]) func(s Set[T]) bool {
	return func(s Set[T]) bool {
		return Null(Intersection(other)(s))
	}
}

/* ============ Combine ============ */

// Union returns the union of the two Sets.
func Union[T comparable](other Set[T]) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		for elem := range other.elements {
			out.elements[elem] = struct{}{}
		}
		return out
	}
}

// Unions returns the union of all Sets in the slice.
func Unions[T comparable](sets []Set[T]) Set[T] {
	out := Empty[T]()
	for _, s := range sets {
		out = Union(s)(out)
	}
	return out
}

// Difference returns the set difference of the two sets.
func Difference[T comparable](minuend Set[T]) func(subtrahend Set[T]) Set[T] {
	return func(subtrahend Set[T]) Set[T] {
		out := Copy(minuend)
		for elem := range subtrahend.elements {
			delete(out.elements, elem)
		}
		return out
	}
}

// Intersection returns the intersection of the two sets.
func Intersection[T comparable](other Set[T]) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Empty[T]()
		for elem := range s.elements {
			if _, exists := other.elements[elem]; exists {
				out.elements[elem] = struct{}{}
			}
		}
		return out
	}
}

// Xor returns the symmetric difference of the two sets.
func Xor[T comparable](other Set[T]) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Empty[T]()
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				out.elements[elem] = struct{}{}
			}
		}
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				out.elements[elem] = struct{}{}
			}
		}
		return out
	}
}

// CartesianProduct combines the two Sets into pairs as a new Set.
func CartesianProduct[T, U comparable](s1 Set[T]) func(s2 Set[U]) Set[tuple.Pair[T, U]] {
	return func(s2 Set[U]) Set[tuple.Pair[T, U]] {
		out := Empty[tuple.Pair[T, U]]()
		for x := range s1.elements {
			for y := range s2.elements {
				out = Insert(tuple.NewPair[T, U](x)(y))(out)
			}
		}
		return out
	}
}

/* ============ Filter ============ */

// Filter returns a Set of all elements that satisfy the predicate.
func Filter[T comparable](predicate func(T) bool) func(s Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Empty[T]()
		for elem := range s.elements {
			if predicate(elem) {
				out.elements[elem] = struct{}{}
			}
		}
		return out
	}
}

// Partition splits the Set into two Sets with the left having all elements which satisfy the predicate,
// and the right having all which do not.
func Partition[T comparable](predicate func(T) bool) func(s Set[T]) tuple.Pair[Set[T], Set[T]] {
	return func(s Set[T]) tuple.Pair[Set[T], Set[T]] {
		return tuple.NewPair[Set[T], Set[T]](Filter(predicate)(s))(Filter(fp.Compose2(operator.Not, predicate))(s))
	}
}

// Split splits the set into a pair (s1, s2) where s1 has all elements less than x, and s2
// has all elements greater than x using the provided less than function.
func Split[T comparable](lt func(T) func(T) bool) func(x T) func(s Set[T]) tuple.Pair[Set[T], Set[T]] {
	return func(x T) func(Set[T]) tuple.Pair[Set[T], Set[T]] {
		return func(s Set[T]) tuple.Pair[Set[T], Set[T]] {
			left := Empty[T]()
			right := Empty[T]()

			for e := range s.elements {
				if lt(e)(x) {
					left = Insert(e)(left)
				}
				if lt(x)(e) {
					right = Insert(e)(right)
				}
			}

			return tuple.NewPair[Set[T], Set[T]](left)(right)
		}
	}
}

/* ============ Map ============ */

// Map returns a Set with the function applied to all elements of
// the given set. The resulting set may have fewer elements if multiple
// existing elements map to the same new element.
func Map[T, U comparable](fn func(T) U) func(s Set[T]) Set[U] {
	return func(s Set[T]) Set[U] {
		out := Empty[U]()
		for elem := range s.elements {
			out.elements[fn(elem)] = struct{}{}
		}
		return out
	}
}

/* ============ Conversion ============ */

// Elems returns a slice containing all elements of the Set. Alias to ToSlice.
func Elems[T comparable](s Set[T]) []T {
	return ToSlice(s)
}

// ToSlice returns a slice containing all elements of the Set.
func ToSlice[T comparable](s Set[T]) []T {
	out := []T{}
	for elem := range s.elements {
		out = append(out, elem)
	}
	return out
}

// ToAscSlice returns a slice containing all elements of the Set ordered least
// to greatest based on the given less than function.
func ToAscSlice[T comparable](lt func(T) func(T) bool) func(s Set[T]) []T {
	return func(s Set[T]) []T {
		elems := ToSlice(s)
		goComparator := func(i, j int) bool {
			return lt(elems[i])(elems[j])
		}
		sort.SliceStable(elems, goComparator)
		return elems
	}
}

// ToDescSlice returns a slice containing all elements of the Set ordered greatest
// to least based on the given less than function.
func ToDescSlice[T comparable](lt func(T) func(T) bool) func(s Set[T]) []T {
	return func(s Set[T]) []T {
		elems := ToSlice(s)
		goComparator := func(i, j int) bool {
			return lt(elems[j])(elems[i])
		}
		sort.SliceStable(elems, goComparator)
		return elems
	}
}
