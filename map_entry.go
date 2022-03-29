package gogcoll

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type AddingEntry[K comparable, V any] map[K]V

func (m AddingEntry[K, V]) Add(e Entry[K, V]) {
	m[e.Key] = e.Value
}
