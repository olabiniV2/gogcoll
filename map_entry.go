package gogcoll

// Entry is a data structure that can be used to understand pieces of
// a dictionary/map data structure. Each Entry will have its own
// key and value of given types.
type Entry[K comparable, V any] struct {
	// Key is the key for the Entry
	Key K
	// Value is the value that the specific key maps to
	Value V
}

// AddingEntry is a wrapper type that allows you to use a map
// in a setting where you want a sequence or collection operation
// to add entries to a map, for example using Filter or Reduce.
type AddingEntry[K comparable, V any] map[K]V

// Add will add the key with the specified value to the underlying map
func (m AddingEntry[K, V]) Add(e Entry[K, V]) {
	m[e.Key] = e.Value
}

// Entries allow you to wrap a map and iterate or sequence over the key-value pairs in it
type Entries[K comparable, V any] map[K]V

// Each will visit each entry in the map and yield it to the given function.
func (s Entries[K, V]) Each(f Proc1[Entry[K, V]]) {
	for k, v := range s {
		f(Entry[K, V]{Key: k, Value: v})
	}
}

// Iter implements the Iterable interface
func (s Entries[K, V]) Iter() Iterator[Entry[K, V]] {
	return s
}
