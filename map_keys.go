package gogcoll

import "golang.org/x/exp/maps"

type Keys[K comparable, V any] map[K]V

func (s Keys[K, V]) Each(f Proc1[K]) {
	for k := range s {
		f(k)
	}
}

func (s Keys[K, V]) Iter() Iterator[K] {
	return s
}

func (s Keys[K, V]) Seq() Seq[K] {
	return sliceSeq(maps.Keys(s))
}

type fixedValueKeysAdder[K comparable, V any] struct {
	fixed V
	m     map[K]V
}

func (f *fixedValueKeysAdder[K, V]) Add(v K) {
	f.m[v] = f.fixed
}

func (s Keys[K, V]) Adder(fixed V) Adder[K] {
	return &fixedValueKeysAdder[K, V]{
		fixed: fixed,
		m:     s,
	}
}
