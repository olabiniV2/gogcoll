package gogcoll

import "golang.org/x/exp/maps"

// Values allow you to wrap a map and iterate or sequence over the values in it
type Values[K comparable, V any] map[K]V

// Each will visit each value in the map and yield it to the given function
// If the map contains the same value more than once, the function will receive
// it once for every time it is available in the map
func (s Values[K, V]) Each(f Proc1[V]) {
	for _, v := range s {
		f(v)
	}
}

// Iter implements the Iterable interface
func (s Values[K, V]) Iter() Iterator[V] {
	return s
}

// Seq returns a lazy sequence over the values in the underlying map
func (s Values[K, V]) Seq() Seq[V] {
	return sliceSeq(maps.Values(s))
}
