package maps

import (
	"github.com/JustinKnueppel/go-fp/either"
	fp "github.com/JustinKnueppel/go-fp/function"
	"github.com/JustinKnueppel/go-fp/option"
	"github.com/JustinKnueppel/go-fp/set"
	"github.com/JustinKnueppel/go-fp/slice"
	"github.com/JustinKnueppel/go-fp/tuple"
)

/* =========== Query =========== */

// Null returns true if the map has no elements.
func Null[K comparable, V any](m map[K]V) bool {
	return Size(m) == 0
}

// Size returns the number of keys in the map.
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}

// Member returns true if the key is a member of the map.
func Member[K comparable, V any](key K) func(map[K]V) bool {
	return func(m map[K]V) bool {
		for k := range m {
			if k == key {
				return true
			}
		}
		return false
	}
}

// NotMember returns true if the key is not a member of the map.
func NotMember[K comparable, V any](key K) func(map[K]V) bool {
	return func(m map[K]V) bool {
		for k := range m {
			if k == key {
				return false
			}
		}
		return true
	}
}

// Lookup returns None if the key is not present, otherwise returns the corresponding value.
func Lookup[K comparable, V any](key K) func(map[K]V) option.Option[V] {
	return func(m map[K]V) option.Option[V] {
		v, ok := m[key]
		if !ok {
			return option.None[V]()
		}
		return option.Some(v)
	}
}

// FindWithDefault returns the corresponding value or the fallback if the key is not present.
func FindWithDefault[K comparable, V any](fallback V) func(K) func(map[K]V) V {
	return func(key K) func(map[K]V) V {
		return func(m map[K]V) V {
			v, ok := m[key]
			if !ok {
				return fallback
			}
			return v
		}
	}
}

/* =========== Construction =========== */

// Empty returns an empty map.
func Empty[K comparable, V any]() map[K]V {
	return make(map[K]V)
}

// Singleton returns a map with the key and value given.
func Singleton[K comparable, V any](k K) func(V) map[K]V {
	return func(v V) map[K]V {
		return map[K]V{k: v}
	}
}

/* =========== Insertion =========== */

// Insert inserts the given key value pair, replacing the previous value if present.
func Insert[K comparable, V any](k K) func(V) func(map[K]V) map[K]V {
	return func(v V) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			out[k] = v
			return out
		}
	}
}

// InsertWith inserts the key value pair into the map if the key is not present.
// If the key is present, the value (combineFn(newVal)(oldVal)) is inserted instead.
func InsertWith[K comparable, V any](combineFn func(V) func(V) V) func(K) func(V) func(map[K]V) map[K]V {
	return func(k K) func(V) func(map[K]V) map[K]V {
		return func(v V) func(map[K]V) map[K]V {
			return func(m map[K]V) map[K]V {
				out := Copy(m)
				if cur, ok := out[k]; ok {
					out[k] = combineFn(v)(cur)
				} else {
					out[k] = v
				}
				return out
			}
		}
	}
}

// InsertWithKey inserts the key value pair into the map if the key is not present.
// If the key is present, the value (combineFn(key)(newVal)(oldVal)) is inserted instead.
func InsertWithKey[K comparable, V any](combineFn func(K) func(V) func(V) V) func(K) func(V) func(map[K]V) map[K]V {
	return func(k K) func(V) func(map[K]V) map[K]V {
		return func(v V) func(map[K]V) map[K]V {
			return func(m map[K]V) map[K]V {
				out := Copy(m)
				if cur, ok := out[k]; ok {
					out[k] = combineFn(k)(v)(cur)
				} else {
					out[k] = v
				}
				return out
			}
		}
	}
}

// InsertLookupWithKey inserts the key value pair into the map if the key is not present
// and returns (None, newMap). If the key is present, the value (combineFn(key)(newVal)(oldVal))
// is inserted instead and (Some(oldVal), newMap) is returned.
func InsertLookupWithKey[K comparable, V any](combineFn func(K) func(V) func(V) V) func(K) func(V) func(map[K]V) tuple.Pair[option.Option[V], map[K]V] {
	return func(k K) func(V) func(map[K]V) tuple.Pair[option.Option[V], map[K]V] {
		return func(v V) func(map[K]V) tuple.Pair[option.Option[V], map[K]V] {
			return func(m map[K]V) tuple.Pair[option.Option[V], map[K]V] {
				out := Copy(m)
				cur, ok := out[k]
				if ok {
					out[k] = combineFn(k)(v)(cur)
					return tuple.NewPair[option.Option[V], map[K]V](option.Some(cur))(out)
				}
				out[k] = v
				return tuple.NewPair[option.Option[V], map[K]V](option.None[V]())(out)
			}
		}
	}
}

/* =========== Delete/Update =========== */

// Delete deletes the key value pair at the given key if present.
func Delete[K comparable, V any](k K) func(map[K]V) map[K]V {
	return func(m map[K]V) map[K]V {
		out := Copy(m)
		delete(out, k)
		return out
	}
}

// Adjust applies the given function to the value at the given key if it exists.
func Adjust[K comparable, V any](adjustFn func(V) V) func(K) func(map[K]V) map[K]V {
	return func(k K) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			if cur, ok := out[k]; ok {
				out[k] = adjustFn(cur)
			}
			return out
		}
	}
}

// AdjustWithKey applies the given function to the value at the given key if it exists.
func AdjustWithKey[K comparable, V any](adjustFn func(K) func(V) V) func(K) func(map[K]V) map[K]V {
	return func(k K) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			if cur, ok := out[k]; ok {
				out[k] = adjustFn(k)(cur)
			}
			return out
		}
	}
}

// Update applies the given function to the value at the given key if it exists.
// If the function returns Some(x) then x will be inserted. If it returns None then
// the element will be deleted.
func Update[K comparable, V any](adjustFn func(V) option.Option[V]) func(K) func(map[K]V) map[K]V {
	return func(k K) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			if cur, ok := out[k]; ok {
				newVal := adjustFn(cur)
				if option.IsSome(newVal) {
					out[k] = option.Unwrap(newVal)
				} else {
					delete(out, k)
				}
			}
			return out
		}
	}
}

// UpdateWithKey applies the given function to the value at the given key if it exists.
// If the function returns Some(x) then x will be inserted. If it returns None then
// the element will be deleted.
func UpdateWithKey[K comparable, V any](adjustFn func(K) func(V) option.Option[V]) func(K) func(map[K]V) map[K]V {
	return func(k K) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			if cur, ok := out[k]; ok {
				newVal := adjustFn(k)(cur)
				if option.IsSome(newVal) {
					out[k] = option.Unwrap(newVal)
				} else {
					delete(out, k)
				}
			}
			return out
		}
	}
}

// UpdateLookupWithKey applies the given function to the value at the given key if it exists.
// If the function returns Some(x) then x will be inserted and (newVal, newMap) will be returned.
// If it returns None then the element will be deleted and (oldVal, newMap) will be returned.
// If the key does not exist, then (None, sameMap) will be returned.
func UpdateLookupWithKey[K comparable, V any](adjustFn func(K) func(V) option.Option[V]) func(K) func(map[K]V) tuple.Pair[option.Option[V], map[K]V] {
	return func(k K) func(map[K]V) tuple.Pair[option.Option[V], map[K]V] {
		return func(m map[K]V) tuple.Pair[option.Option[V], map[K]V] {
			out := Copy(m)
			if cur, ok := out[k]; ok {
				newValOpt := adjustFn(k)(cur)
				if option.IsSome(newValOpt) {
					newVal := option.Unwrap(newValOpt)
					out[k] = newVal
					return tuple.NewPair[option.Option[V], map[K]V](option.Some(newVal))(out)
				} else {
					delete(out, k)
					return tuple.NewPair[option.Option[V], map[K]V](option.Some(cur))(out)
				}
			}
			return tuple.NewPair[option.Option[V], map[K]V](option.None[V]())(out)
		}
	}
}

// Alter inserts the result of the given function if it returns Some,
// otherwise removes the element if present.
func Alter[K comparable, V any](alterFn func(option.Option[V]) option.Option[V]) func(K) func(map[K]V) map[K]V {
	return func(k K) func(map[K]V) map[K]V {
		return func(m map[K]V) map[K]V {
			out := Copy(m)
			newValOpt := alterFn(Lookup[K, V](k)(out))
			if option.IsSome(newValOpt) {
				out[k] = option.Unwrap(newValOpt)
			} else {
				delete(out, k)
			}
			return out
		}
	}
}

/* =========== Combine =========== */

// Union contains all elements of both maps, preferring the first map's element.
func Union[K comparable, V any](m1 map[K]V) func(map[K]V) map[K]V {
	return UnionWith[K](fp.Const[V, V])(m1)
}

// UnionWith contains all elements of both maps, and uses combine(m1[k], m2[k]) if k is in both maps.
func UnionWith[K comparable, V any](combineFn func(V) func(V) V) func(map[K]V) func(map[K]V) map[K]V {
	return func(m1 map[K]V) func(map[K]V) map[K]V {
		return func(m2 map[K]V) map[K]V {
			out := Copy(m1)
			for k, v2 := range m2 {
				if v1, ok := m1[k]; ok {
					out[k] = combineFn(v1)(v2)
				} else {
					out[k] = v2
				}
			}
			return out
		}
	}
}

// UnionWithKey contains all elements of both maps, and uses combine(k, m1[k], m2[k]) if k is in both maps.
func UnionWithKey[K comparable, V any](combineFn func(K) func(V) func(V) V) func(map[K]V) func(map[K]V) map[K]V {
	return func(m1 map[K]V) func(map[K]V) map[K]V {
		return func(m2 map[K]V) map[K]V {
			out := Copy(m1)
			for k, v2 := range m2 {
				if v1, ok := m1[k]; ok {
					out[k] = combineFn(k)(v1)(v2)
				} else {
					out[k] = v2
				}
			}
			return out
		}
	}
}

// Unions returns the union of a list of maps
func Unions[K comparable, V any](maps []map[K]V) map[K]V {
	return slice.Foldl(Union[K, V])(Empty[K, V]())(maps)
}

// UnionsWith returns the union of a slice of maps using combine(oldVal)(newVal) to compute the value when a duplicate key is found.
func UnionsWith[K comparable, V any](combineFn func(V) func(V) V) func([]map[K]V) map[K]V {
	return slice.Foldl(UnionWith[K](combineFn))(Empty[K, V]())
}

/* =========== Difference =========== */

// Difference contains all elements in the first map which do not exist in the second map.
func Difference[K comparable, V1, V2 any](m1 map[K]V1) func(map[K]V2) map[K]V1 {
	alwaysNone := fp.Curry2(func(_ V1, _ V2) option.Option[V1] { return option.None[V1]() })
	return DifferenceWith[K](alwaysNone)(m1)
}

// DifferenceWith contains the elements of m1 not in m2 as well as any elements for which combine(m1[k], m2[k]) = Some(v) and uses v as the updated value.
func DifferenceWith[K comparable, V1, V2 any](combineFn func(V1) func(V2) option.Option[V1]) func(map[K]V1) func(map[K]V2) map[K]V1 {
	return func(m1 map[K]V1) func(map[K]V2) map[K]V1 {
		return func(m2 map[K]V2) map[K]V1 {
			out := Empty[K, V1]()
			for k, v1 := range m1 {
				if v2, ok := m2[k]; ok {
					newVOpt := combineFn(v1)(v2)
					if option.IsSome(newVOpt) {
						out[k] = option.Unwrap(newVOpt)
					}
				} else {
					out[k] = v1
				}
			}
			return out
		}
	}
}

// DifferenceWithKey contains the elements of m1 not in m2 as well as any elements for which combine(k, m1[k], m2[k]) = Some(v) and uses v as the updated value.
func DifferenceWithKey[K comparable, V1, V2 any](combineFn func(K) func(V1) func(V2) option.Option[V1]) func(map[K]V1) func(map[K]V2) map[K]V1 {
	return func(m1 map[K]V1) func(map[K]V2) map[K]V1 {
		return func(m2 map[K]V2) map[K]V1 {
			out := Empty[K, V1]()
			for k, v1 := range m1 {
				if v2, ok := m2[k]; ok {
					newVOpt := combineFn(k)(v1)(v2)
					if option.IsSome(newVOpt) {
						out[k] = option.Unwrap(newVOpt)
					}
				} else {
					out[k] = v1
				}
			}
			return out
		}
	}
}

/* =========== Intersection =========== */

// Intersection returns all elements that exist in both sets, preferring the value of the first set.
func Intersection[K comparable, V1, V2 any](m1 map[K]V1) func(map[K]V2) map[K]V1 {
	return func(m2 map[K]V2) map[K]V1 {
		out := Empty[K, V1]()
		for k, v1 := range m1 {
			if _, ok := m2[k]; ok {
				out[k] = v1
			}
		}
		return out
	}
}

// IntersectionWith returns all elements that exist in both sets and sets the value at k to combine(m1[k], m2[k]).
func IntersectionWith[K comparable, V1, V2, V3 any](combineFn func(V1) func(V2) V3) func(map[K]V1) func(map[K]V2) map[K]V3 {
	return func(m1 map[K]V1) func(map[K]V2) map[K]V3 {
		return func(m2 map[K]V2) map[K]V3 {
			out := Empty[K, V3]()
			for k, v1 := range m1 {
				if v2, ok := m2[k]; ok {
					out[k] = combineFn(v1)(v2)
				}
			}
			return out
		}
	}
}

// IntersectionWithKey returns all elements that exist in both sets and sets the value at k to combine(k, m1[k], m2[k]).
func IntersectionWithKey[K comparable, V1, V2, V3 any](combineFn func(K) func(V1) func(V2) V3) func(map[K]V1) func(map[K]V2) map[K]V3 {
	return func(m1 map[K]V1) func(map[K]V2) map[K]V3 {
		return func(m2 map[K]V2) map[K]V3 {
			out := Empty[K, V3]()
			for k, v1 := range m1 {
				if v2, ok := m2[k]; ok {
					out[k] = combineFn(k)(v1)(v2)
				}
			}
			return out
		}
	}
}

/* =========== Traversal =========== */
/* =========== Map =========== */

// Map applies the given function to all values in the map.
func Map[K comparable, V1, V2 any](fn func(V1) V2) func(map[K]V1) map[K]V2 {
	return func(m map[K]V1) map[K]V2 {
		out := Empty[K, V2]()
		for k, v := range m {
			out[k] = fn(v)
		}
		return out
	}
}

// MapWithKey applies the given function to all values in the map.
func MapWithKey[K comparable, V1, V2 any](fn func(K) func(V1) V2) func(map[K]V1) map[K]V2 {
	return func(m map[K]V1) map[K]V2 {
		out := Empty[K, V2]()
		for k, v := range m {
			out[k] = fn(k)(v)
		}
		return out
	}
}

// MapAccum threads an accumulating argument through the map without a guaranteed order.
func MapAccum[K comparable, A, V1, V2 any](accumFn func(A) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
	return func(acc A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
		return func(m map[K]V1) tuple.Pair[A, map[K]V2] {
			outAcc := acc
			outMap := Empty[K, V2]()
			for k, v := range m {
				pair := accumFn(outAcc)(v)
				outAcc = tuple.Fst(pair)
				outMap[k] = tuple.Snd(pair)
			}
			return tuple.NewPair[A, map[K]V2](outAcc)(outMap)
		}
	}
}

// MapAccumOrdered threads an accumulating argument through the map with the keys ordered by the given function.
func MapAccumOrdered[K comparable, A, V1, V2 any](lt func(K) func(K) bool) func(func(A) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
	return func(accumFn func(A) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
		return func(acc A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
			return func(m map[K]V1) tuple.Pair[A, map[K]V2] {
				keys := KeysOrdered[K, V1](lt)(m)
				outAcc := acc
				outMap := Empty[K, V2]()
				for _, k := range keys {
					pair := accumFn(outAcc)(m[k])
					outAcc = tuple.Fst(pair)
					outMap[k] = tuple.Snd(pair)
				}
				return tuple.NewPair[A, map[K]V2](outAcc)(outMap)
			}
		}
	}
}

// MapAccumWithKey threads an accumulating argument through the map.
func MapAccumWithKey[K comparable, A, V1, V2 any](accumFn func(A) func(K) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
	return func(acc A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
		return func(m map[K]V1) tuple.Pair[A, map[K]V2] {
			outAcc := acc
			outMap := Empty[K, V2]()
			for k, v := range m {
				pair := accumFn(outAcc)(k)(v)
				outAcc = tuple.Fst(pair)
				outMap[k] = tuple.Snd(pair)
			}
			return tuple.NewPair[A, map[K]V2](outAcc)(outMap)
		}
	}
}

// MapAccumWithKeyOrdered threads an accumulating argument through the map in ascending order of keys
// according to the given less than function.
func MapAccumWithKeyOrdered[K comparable, A, V1, V2 any](lt func(K) func(K) bool) func(func(A) func(K) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
	return func(accumFn func(A) func(K) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
		return func(acc A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
			return func(m map[K]V1) tuple.Pair[A, map[K]V2] {
				keys := KeysOrdered[K, V1](lt)(m)
				outAcc := acc
				outMap := Empty[K, V2]()
				for _, k := range keys {
					pair := accumFn(outAcc)(k)(m[k])
					outAcc = tuple.Fst(pair)
					outMap[k] = tuple.Snd(pair)
				}
				return tuple.NewPair[A, map[K]V2](outAcc)(outMap)
			}
		}
	}
}

// MapAccumRWithKeyOrdered threads an accumulating argument through the map in descending order of keys
// according to the given less than function.
func MapAccumRWithKeyOrdered[K comparable, A, V1, V2 any](lt func(K) func(K) bool) func(func(A) func(K) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
	return func(accumFn func(A) func(K) func(V1) tuple.Pair[A, V2]) func(A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
		return func(acc A) func(map[K]V1) tuple.Pair[A, map[K]V2] {
			return func(m map[K]V1) tuple.Pair[A, map[K]V2] {
				keys := slice.Reverse(KeysOrdered[K, V1](lt)(m))
				outAcc := acc
				outMap := Empty[K, V2]()
				for _, k := range keys {
					pair := accumFn(outAcc)(k)(m[k])
					outAcc = tuple.Fst(pair)
					outMap[k] = tuple.Snd(pair)
				}
				return tuple.NewPair[A, map[K]V2](outAcc)(outMap)
			}
		}
	}
}

// MapKeys applies the function to each key in ascending order. If two keys end up
// with the same value, the value of the new key will override the existing value.
func MapKeys[K1, K2 comparable, V any](lt func(K1) func(K1) bool) func(func(K1) K2) func(map[K1]V) map[K2]V {
	return func(fn func(K1) K2) func(map[K1]V) map[K2]V {
		return func(m map[K1]V) map[K2]V {
			out := Empty[K2, V]()
			for _, k := range KeysOrdered[K1, V](lt)(m) {
				out[fn(k)] = m[k]
			}
			return out
		}
	}
}

// MapKeysWith applies the function to each key in ascending order. If two keys end up
// with the same value, the values are combined according to the combination function combineFn(newVal)(curVal).
func MapKeysWith[K1, K2 comparable, V any](lt func(K1) func(K1) bool) func(func(V) func(V) V) func(func(K1) K2) func(map[K1]V) map[K2]V {
	return func(combineFn func(V) func(V) V) func(func(K1) K2) func(map[K1]V) map[K2]V {
		return func(fn func(K1) K2) func(map[K1]V) map[K2]V {
			return func(m map[K1]V) map[K2]V {
				out := Empty[K2, V]()
				for _, k := range KeysOrdered[K1, V](lt)(m) {
					newK := fn(k)
					if cur, ok := out[newK]; ok {
						out[newK] = combineFn(m[k])(cur)
					} else {
						out[newK] = m[k]
					}
				}
				return out
			}
		}
	}
}

/* =========== Fold =========== */

// Fold right folds the values in the map with the given function. Order not guaranteed.
// See FoldrWithKey for the ability to order keys before folding.
func Fold[K comparable, V, A any](fn func(V) func(A) A) func(A) func(map[K]V) A {
	return func(a A) func(map[K]V) A {
		return fp.Compose2(slice.Foldr(fn)(a), Elems[K, V])
	}
}

// FoldWithKey right folds the keys/values in the map with the given function. Order is not guaranteed.
// See FoldrWithKey for the ability to order keys before folding.
func FoldWithKey[K comparable, V, A any](fn func(K) func(V) func(A) A) func(A) func(map[K]V) A {
	return func(a A) func(map[K]V) A {
		return func(m map[K]V) A {
			out := a
			for k, v := range m {
				out = fn(k)(v)(out)
			}
			return out
		}
	}
}

// FoldrWithKey post-order folds the values in the map with the given function from the value at the lowest key to the highest.
func FoldrWithKey[K comparable, V, A any](lt func(K) func(K) bool) func(fn func(K) func(V) func(A) A) func(A) func(map[K]V) A {
	return func(fn func(K) func(V) func(A) A) func(A) func(map[K]V) A {
		return func(a A) func(map[K]V) A {
			return func(m map[K]V) A {
				out := a
				for _, k := range slice.Reverse(KeysOrdered[K, V](lt)(m)) {
					out = fn(k)(m[k])(out)
				}
				return out
			}
		}
	}
}

// FoldlWithKey pre-order folds the values in the map with the given function from the value at the highest key to the lowest.
func FoldlWithKey[K comparable, V, A any](lt func(K) func(K) bool) func(fn func(A) func(K) func(V) A) func(A) func(map[K]V) A {
	return func(fn func(A) func(K) func(V) A) func(A) func(map[K]V) A {
		return func(a A) func(map[K]V) A {
			return func(m map[K]V) A {
				out := a
				for _, k := range KeysOrdered[K, V](lt)(m) {
					out = fn(out)(k)(m[k])
				}
				return out
			}
		}
	}
}

/* =========== Conversion =========== */

// Elems returns a slice of all values in the map.
func Elems[K comparable, V any](m map[K]V) []V {
	out := []V{}
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

// Keys returns a slice of all keys in the map.
func Keys[K comparable, V any](m map[K]V) []K {
	out := []K{}
	for k := range m {
		out = append(out, k)
	}
	return out
}

// KeysOrdered returns the keys in the map ordered accoridng to the less than function.
func KeysOrdered[K comparable, V any](lt func(K) func(K) bool) func(map[K]V) []K {
	return func(m map[K]V) []K {
		keys := []K{}
		for k := range m {
			keys = append(keys, k)
		}
		return slice.SortBy(lt)(keys)
	}
}

// KeysSet returns a set of all keys in the map.
func KeysSet[K comparable, V any](m map[K]V) set.Set[K] {
	out := set.Empty[K]()
	for k := range m {
		out = set.Insert(k)(out)
	}
	return out
}

// Assocs returns all key/value pairs in the map.
func Assocs[K comparable, V any](m map[K]V) []tuple.Pair[K, V] {
	pairs := []tuple.Pair[K, V]{}
	for k, v := range m {
		pairs = append(pairs, tuple.NewPair[K, V](k)(v))
	}
	return pairs
}

// AssocsOrdered returns all key/value pairs in the map in ascending key order according
// to the less than function.
func AssocsOrdered[K comparable, V any](lt func(K) func(K) bool) func(map[K]V) []tuple.Pair[K, V] {
	return func(m map[K]V) []tuple.Pair[K, V] {
		keys := KeysOrdered[K, V](lt)(m)
		pairs := []tuple.Pair[K, V]{}
		for _, k := range keys {
			pairs = append(pairs, tuple.NewPair[K, V](k)(m[k]))
		}
		return pairs
	}
}

/* =========== Slices =========== */

// ToSlice converts the map to to a slice of (k, v) pairs.
func ToSlice[K comparable, V any](m map[K]V) []tuple.Pair[K, V] {
	out := []tuple.Pair[K, V]{}
	for k, v := range m {
		out = append(out, tuple.NewPair[K, V](k)(v))
	}
	return out
}

// FromSlice builds a map from the given slice of (k, v) pairs. If a key is duplicated, the final
// value for that key is contained.
func FromSlice[K comparable, V any](pairs []tuple.Pair[K, V]) map[K]V {
	out := Empty[K, V]()
	for _, pair := range pairs {
		out[tuple.Fst(pair)] = tuple.Snd(pair)
	}
	return out
}

// FromSliceWith builds a map from a slice, inserting combineFn(current)(newVal) when a key is encountered multiple times.
func FromSliceWith[K comparable, V any](combineFn func(V) func(V) V) func([]tuple.Pair[K, V]) map[K]V {
	return func(pairs []tuple.Pair[K, V]) map[K]V {
		out := Empty[K, V]()
		for _, pair := range pairs {
			k, v := tuple.Pattern(pair)
			if oldV, ok := out[k]; ok {
				out[k] = combineFn(oldV)(v)
			} else {
				out[k] = v
			}
		}
		return out
	}
}

// FromSliceWithKey builds a map from a slice, inserting combineFn(k)(current)(newVal) when a key is encountered multiple times.
func FromSliceWithKey[K comparable, V any](combineFn func(K) func(V) func(V) V) func([]tuple.Pair[K, V]) map[K]V {
	return func(pairs []tuple.Pair[K, V]) map[K]V {
		out := Empty[K, V]()
		for _, pair := range pairs {
			k, v := tuple.Pattern(pair)
			if oldV, ok := out[k]; ok {
				out[k] = combineFn(k)(oldV)(v)
			} else {
				out[k] = v
			}
		}
		return out
	}
}

/* =========== Ordered slices =========== */

// ToAscSlice converts the map to a slice ascending order by keys.
func ToAscSlice[K comparable, V any](lt func(K) func(K) bool) func(map[K]V) []tuple.Pair[K, V] {
	return func(m map[K]V) []tuple.Pair[K, V] {
		out := []tuple.Pair[K, V]{}
		for _, k := range KeysOrdered[K, V](lt)(m) {
			out = append(out, tuple.NewPair[K, V](k)(m[k]))
		}
		return out
	}
}

// ToDescSlice converts the map to a slice descending order by keys.
func ToDescSlice[K comparable, V any](lt func(K) func(K) bool) func(map[K]V) []tuple.Pair[K, V] {
	return func(m map[K]V) []tuple.Pair[K, V] {
		return slice.Reverse(ToAscSlice[K, V](lt)(m))
	}
}

/* =========== Filter =========== */

// Filter filters all values that satisfy the predicate.
func Filter[K comparable, V any](predicate func(V) bool) func(map[K]V) map[K]V {
	return func(m map[K]V) map[K]V {
		out := Empty[K, V]()
		for k, v := range m {
			if predicate(v) {
				out[k] = v
			}
		}
		return out
	}
}

// FilterWithKey filters all values that satisfy the predicate.
func FilterWithKey[K comparable, V any](predicate func(K) func(V) bool) func(map[K]V) map[K]V {
	return func(m map[K]V) map[K]V {
		out := Empty[K, V]()
		for k, v := range m {
			if predicate(k)(v) {
				out[k] = v
			}
		}
		return out
	}
}

// Partition partitions the map according to the predicate. The first map contains all elements that satisfy
// the predicate, the second contains all elements which fail the predicate.
func Partition[K comparable, V any](predicate func(V) bool) func(map[K]V) tuple.Pair[map[K]V, map[K]V] {
	return func(m map[K]V) tuple.Pair[map[K]V, map[K]V] {
		passes := Empty[K, V]()
		fails := Empty[K, V]()
		for k, v := range m {
			if predicate(v) {
				passes[k] = v
			} else {
				fails[k] = v
			}
		}
		return tuple.NewPair[map[K]V, map[K]V](passes)(fails)
	}
}

// PartitionWithKey partitions the map according to the predicate. The first map contains all elements that satisfy
// the predicate, the second contains all elements which fail the predicate.
func PartitionWithKey[K comparable, V any](predicate func(K) func(V) bool) func(map[K]V) tuple.Pair[map[K]V, map[K]V] {
	return func(m map[K]V) tuple.Pair[map[K]V, map[K]V] {
		passes := Empty[K, V]()
		fails := Empty[K, V]()
		for k, v := range m {
			if predicate(k)(v) {
				passes[k] = v
			} else {
				fails[k] = v
			}
		}
		return tuple.NewPair[map[K]V, map[K]V](passes)(fails)
	}
}

// MapOption maps the values and collects only the Some results.
func MapOption[K comparable, V1, V2 any](fn func(V1) option.Option[V2]) func(map[K]V1) map[K]V2 {
	return func(m map[K]V1) map[K]V2 {
		out := Empty[K, V2]()
		for k, v := range m {
			opt := fn(v)
			if option.IsSome(opt) {
				out[k] = option.Unwrap(opt)
			}
		}
		return out
	}
}

// MapOptionWithKey maps the keys/values and collects only the Some results.
func MapOptionWithKey[K comparable, V1, V2 any](fn func(K) func(V1) option.Option[V2]) func(map[K]V1) map[K]V2 {
	return func(m map[K]V1) map[K]V2 {
		out := Empty[K, V2]()
		for k, v := range m {
			opt := fn(k)(v)
			if option.IsSome(opt) {
				out[k] = option.Unwrap(opt)
			}
		}
		return out
	}
}

// MapEither maps values and separates the Left and Right results into two maps.
func MapEither[K comparable, V, L, R any](fn func(V) either.Either[L, R]) func(map[K]V) tuple.Pair[map[K]L, map[K]R] {
	return func(m map[K]V) tuple.Pair[map[K]L, map[K]R] {
		left := Empty[K, L]()
		right := Empty[K, R]()
		for k, v := range m {
			e := fn(v)
			if either.IsLeft(e) {
				left[k] = either.UnwrapLeft(e)
			} else {
				right[k] = either.Unwrap(e)
			}
		}
		return tuple.NewPair[map[K]L, map[K]R](left)(right)
	}
}

// MapEitherWithKey maps keys/values and separates the Left and Right results into two maps.
func MapEitherWithKey[K comparable, V, L, R any](fn func(K) func(V) either.Either[L, R]) func(map[K]V) tuple.Pair[map[K]L, map[K]R] {
	return func(m map[K]V) tuple.Pair[map[K]L, map[K]R] {
		left := Empty[K, L]()
		right := Empty[K, R]()
		for k, v := range m {
			e := fn(k)(v)
			if either.IsLeft(e) {
				left[k] = either.UnwrapLeft(e)
			} else {
				right[k] = either.Unwrap(e)
			}
		}
		return tuple.NewPair[map[K]L, map[K]R](left)(right)
	}
}

// Split splits the map into two maps (map1, map2), where all keys in map1 are smaller than the
// given key, and all keys in map2 are greater than the target key.
func Split[K comparable, V any](lt func(K) func(K) bool) func(K) func(map[K]V) tuple.Pair[map[K]V, map[K]V] {
	return func(target K) func(map[K]V) tuple.Pair[map[K]V, map[K]V] {
		return func(m map[K]V) tuple.Pair[map[K]V, map[K]V] {
			less := Empty[K, V]()
			greater := Empty[K, V]()
			for k, v := range m {
				if lt(k)(target) {
					less[k] = v
				} else if lt(target)(k) {
					greater[k] = v
				}
			}
			return tuple.NewPair[map[K]V, map[K]V](less)(greater)
		}
	}
}

/* =========== Submap =========== */

// IsSubmapOf returns true if all keys in m1 are in m2, and m1[k] == m2[k] for all keys in m1.
func IsSubmapOf[K, V comparable](m1 map[K]V) func(map[K]V) bool {
	return func(m2 map[K]V) bool {
		for k, v1 := range m1 {
			if v2, ok := m2[k]; !ok || v1 != v2 {
				return false
			}
		}
		return true
	}
}

// IsSubmapOfBy returns true if all keys in m1 are in m2, and eqFn(m1[k])(m2[k]) == true for all keys in m1.
func IsSubmapOfBy[K comparable, V1, V2 any](eqFn func(V1) func(V2) bool) func(map[K]V1) func(map[K]V2) bool {
	return func(m1 map[K]V1) func(map[K]V2) bool {
		return func(m2 map[K]V2) bool {
			for k, v1 := range m1 {
				if v2, ok := m2[k]; !ok || !eqFn(v1)(v2) {
					return false
				}
			}
			return true
		}
	}
}

// IsProperSubmapOf returns true if m1 != m2, all keys in m1 are in m2, and m1[k] == m2[k] for all keys in m1.
func IsProperSubmapOf[K, V comparable](m1 map[K]V) func(map[K]V) bool {
	return func(m2 map[K]V) bool {
		if Size(m2) <= Size(m1) {
			return false
		}
		for k, v1 := range m1 {
			if v2, ok := m2[k]; !ok || v1 != v2 {
				return false
			}
		}
		return true
	}
}

// IsProperSubmapOfBy returns true if m1 != m2, all keys in m1 are in m2, and eqFn(m1[k])(m2[k]) == true for all keys in m1.
func IsProperSubmapOfBy[K comparable, V1, V2 any](eqFn func(V1) func(V2) bool) func(map[K]V1) func(map[K]V2) bool {
	return func(m1 map[K]V1) func(map[K]V2) bool {
		return func(m2 map[K]V2) bool {
			if Size(m2) <= Size(m1) {
				return false
			}
			for k, v1 := range m1 {
				if v2, ok := m2[k]; !ok || !eqFn(v1)(v2) {
					return false
				}
			}
			return true
		}
	}
}

// Copy returns a shallow copy of the map.
func Copy[K comparable, V any](m map[K]V) map[K]V {
	out := make(map[K]V)
	for k, v := range m {
		out[k] = v
	}
	return out
}
