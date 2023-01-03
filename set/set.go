package set

import (
	"fmt"
	"strings"
)

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

// Add returns a union of the given set and element.
func Add[T comparable](elem T) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		out.elements[elem] = struct{}{}
		return out
	}
}

// Copy returns a value copy of the given set.
func Copy[T comparable](set Set[T]) Set[T] {
	elems := make(map[T]struct{})
	for elem := range set.elements {
		elems[elem] = struct{}{}
	}
	return Set[T]{elems}
}

// Contains returns true if the set contains the target
// otherwise returns false.
func Contains[T comparable](target T) func(Set[T]) bool {
	return func(s Set[T]) bool {
		_, found := s.elements[target]
		return found
	}
}

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

// Disjoint returns true if the two sets have an empty intersection.
func Disjoint[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		return IsEmpty(Intersection(other)(s))
	}
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

// Filter returns a Set of all elements that satisfy the predicate.
func Filter[T comparable](predicate func(T) bool) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := New[T]()
		for elem := range s.elements {
			if predicate(elem) {
				out.elements[elem] = struct{}{}
			}
		}
		return out
	}
}

// FromSlice creates a set from the given slice.
func FromSlice[T comparable](init []T) Set[T] {
	elems := make(map[T]struct{})
	for _, t := range init {
		elems[t] = struct{}{}
	}
	return Set[T]{elems}
}

// Intersection returns the intersection of the two sets.
func Intersection[T comparable](other Set[T]) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := New[T]()
		for elem := range s.elements {
			if _, exists := other.elements[elem]; exists {
				out.elements[elem] = struct{}{}
			}
		}
		return out
	}
}

// IsEmpty returns true if the Set has no elements.
func IsEmpty[T comparable](s Set[T]) bool {
	return Size(s) == 0
}

// IsStrictSubset returns true if every value of the other set exists
// in the given set but the sets are not equal.
func IsStrictSubset[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				return false
			}
		}
		return Size(s) != Size(other)
	}
}

// IsStrictSuperset returns true if every value of the given set exists
// in the other set but the sets are not equal.
func IsStrictSuperset[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				return false
			}
		}
		return Size(s) != Size(other)
	}
}

// IsSubset returns true if every value of the other set exists
// in the given set.
func IsSubset[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range other.elements {
			if _, exists := s.elements[elem]; !exists {
				return false
			}
		}
		return true
	}
}

// IsSuperset returns true if every value of the given set exists
// in the other set.
func IsSuperset[T comparable](other Set[T]) func(Set[T]) bool {
	return func(s Set[T]) bool {
		for elem := range s.elements {
			if _, exists := other.elements[elem]; !exists {
				return false
			}
		}
		return true
	}
}

// Map returns a Set with the function applied to all elements of
// the given set. The resulting set may have fewer elements if multiple
// existing elements map to the same new element.
func Map[T, U comparable](fn func(T) U) func(Set[T]) Set[U] {
	return func(s Set[T]) Set[U] {
		out := New[U]()
		for elem := range s.elements {
			out.elements[fn(elem)] = struct{}{}
		}
		return out
	}
}

// New creates a Set with the given initial values if provided.
func New[T comparable](init ...T) Set[T] {
	elems := make(map[T]struct{})
	for _, t := range init {
		elems[t] = struct{}{}
	}
	return Set[T]{elems}
}

// Remove returns the Set with the given element removed.
func Remove[T comparable](elem T) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := Copy(s)
		delete(out.elements, elem)
		return out
	}
}

// Size returns the number of elements in the Set.
func Size[T comparable](s Set[T]) int {
	return len(s.elements)
}

// ToSlice returns a slice containing all elements of the Set.
func ToSlice[T comparable](s Set[T]) []T {
	out := []T{}
	for elem := range s.elements {
		out = append(out, elem)
	}
	return out
}

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

// Xor returns the symmetric difference of the two sets.
func Xor[T comparable](other Set[T]) func(Set[T]) Set[T] {
	return func(s Set[T]) Set[T] {
		out := New[T]()
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
