package set

import (
	"fmt"
	"strings"

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
func Equal[T comparable](other Set[T]) func(Set[T]) bool {
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

//TODO: Singleton

// FromSlice creates a set from the given slice.
func FromSlice[T comparable](init []T) Set[T] {
	elems := make(map[T]struct{})
	for _, t := range init {
		elems[t] = struct{}{}
	}
	return Set[T]{elems}
}

//TODO: Powerset

/* ============ Insertion ============ */

// Insert returns a union of the given set and element.
func Insert[T comparable](elem T) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		out.elements[elem] = struct{}{}
		return out
	}
}

/* ============ Deletion ============ */

// Delete returns the Set with the given element removed.
func Delete[T comparable](elem T) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		delete(out.elements, elem)
		return out
	}
}

/* ============ Query ============ */

// Member returns true if the set contains the target
// otherwise returns false.
func Member[T comparable](target T) func(Set[T]) bool {
	return func(s Set[T]) bool {
		_, found := s.elements[target]
		return found
	}
}

//TODO: NotMember

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
func IsSubsetOf[T comparable](other Set[T]) func(Set[T]) bool {
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
func IsProperSubsetOf[T comparable](other Set[T]) func(Set[T]) bool {
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
func IsSupersetOf[T comparable](other Set[T]) func(Set[T]) bool {
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
func IsProperSupersetOf[T comparable](other Set[T]) func(Set[T]) bool {
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
func Disjoint[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		return Null(Intersection(other)(s))
	}
}

/* ============ Combine ============ */

// Union returns the union of the two Sets.
func Union[T comparable](other Set[T]) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		for elem := range other.elements {
			out.elements[elem] = struct{}{}
		}
		return out
	}
}

//TODO: Unions

// Difference returns the set difference of the two sets.
func Difference[T comparable](subtrahend Set[T]) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		for elem := range subtrahend.elements {
			delete(out.elements, elem)
		}
		return out
	}
}

// Intersection returns the intersection of the two sets.
func Intersection[T comparable](other Set[T]) func(Set[T]) Set[T] {
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
func Xor[T comparable](other Set[T]) func(Set[T]) Set[T] {
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
func CartesianProduct[T, U comparable](s1 Set[T]) func(Set[U]) Set[tuple.Pair[T, U]] {
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

//TODO: DisjointUnion

/* ============ Filter ============ */

// Filter returns a Set of all elements that satisfy the predicate.
func Filter[T comparable](predicate func(T) bool) func(Set[T]) Set[T] {
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

//TODO: Partition

//TODO: Split

/* ============ Map ============ */

// Map returns a Set with the function applied to all elements of
// the given set. The resulting set may have fewer elements if multiple
// existing elements map to the same new element.
func Map[T, U comparable](fn func(T) U) func(Set[T]) Set[U] {
	return func(s Set[T]) Set[U] {
		out := Empty[U]()
		for elem := range s.elements {
			out.elements[fn(elem)] = struct{}{}
		}
		return out
	}
}

/* ============ Folds ============ */

//TODO: Foldl

//TODO: Foldr

/* ============ Conversion ============ */

//TODO: Elems

// ToSlice returns a slice containing all elements of the Set.
func ToSlice[T comparable](s Set[T]) []T {
	out := []T{}
	for elem := range s.elements {
		out = append(out, elem)
	}
	return out
}

//TODO: ToAscSlice

//TODO: ToDescSlice
