package gogcoll

import "golang.org/x/exp/maps"

type Values[K comparable, V any] map[K]V

func (s Values[K, V]) Each(f Proc1[V]) {
	for _, v := range s {
		f(v)
	}
}

func (s Values[K, V]) Iter() Iterator[V] {
	return s
}

func (s Values[K, V]) Seq() Seq[V] {
	return sliceSeq(maps.Values(s))
}
