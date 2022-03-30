package gogcoll

import "golang.org/x/exp/maps"

// Keys allow you to wrap a map and use the keys in the map as an Iterator or Seq.
type Keys[K comparable, V any] map[K]V

// Each will visit each key in the underlying map and yield the key to the procedure.
// The order of iteration is not defined.
func (s Keys[K, V]) Each(f Proc1[K]) {
	for k := range s {
		f(k)
	}
}

// Iter implements the Iterable interface
func (s Keys[K, V]) Iter() Iterator[K] {
	return s
}

// Seq returns a lazy sequence over the keys in the underlying map
func (s Keys[K, V]) Seq() Seq[K] {
	return sliceSeq(maps.Keys(s))
}

// fixedValueKeysAdder is the implementation that allows you to create an Adder
// from a Keys value. It will add the given value as a key with the same
// value evety time.
type fixedValueKeysAdder[K comparable, V any] struct {
	fixed V
	m     map[K]V
}

// Add implements the Adder interface. It will add the given argument as a key
// to the underlying map, and set the value to the fixed value.
func (f *fixedValueKeysAdder[K, V]) Add(v K) {
	f.m[v] = f.fixed
}

// Adder takes a fixed value and returns an implementation of Adder that will
// add new keys to the underlying map with the same fixed value.
func (s Keys[K, V]) Adder(fixed V) Adder[K] {
	return &fixedValueKeysAdder[K, V]{
		fixed: fixed,
		m:     s,
	}
}
