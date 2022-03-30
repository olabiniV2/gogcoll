package gogcoll

// Set represents an unordered collection of objects.
// Sets can be compared for equality, and the elements manipulated
// in the expected ways
type Set[T comparable] interface {
	Iterable[T]
	Iterator[T]

	// Add will add the given element to the set. If the element already is in the set,
	// this method doesn't do anything.
	Add(T)
	// AddAll adds all the given arguments to the set, using the same logic as Add
	AddAll(...T)
	// Intersection returns a new Set that represents the intersection between this
	// set and the argument
	Intersection(Set[T]) Set[T]
	// Union returns a new Set that represents the union of this set and the argument
	Union(Set[T]) Set[T]
	// Difference returns a new Set containing only the elements from the receiver that
	// are not in the argument set
	Difference(Set[T]) Set[T]
	// SubSetOf returns true if the receiver Set is a sub set of the argument set
	SubSetOf(Set[T]) bool
	// Equal returns true if the receiver contains the same elements as the argument set
	Equal(Set[T]) bool
	// Contains returns true if the given argument is in the set
	Contains(T) bool
	// Size returns the number of elements in the set
	Size() int
	// Empty returns true if this is the empty set
	Empty() bool
	// ToSlice will return a slice with each element, in an undefined order
	ToSlice() []T
}

// simpleSet implements the Set interface using an underlying map from the Golang
// core library
type simpleSet[T comparable] struct {
	values  map[T]bool
	keyView Keys[T, bool]
}

// SetFrom will create a set containing the given values
func SetFrom[T comparable](values ...T) Set[T] {
	s := NewSet[T]()
	s.AddAll(values...)
	return s
}

// NewSet creates a new empty set
func NewSet[T comparable]() Set[T] {
	v := map[T]bool{}
	return &simpleSet[T]{
		values:  v,
		keyView: Keys[T, bool](v),
	}
}

// Add implements the Set interface for simpleSet
func (s *simpleSet[T]) Add(v T) {
	s.values[v] = true
}

// Size implements the Set interface for simpleSet
func (s *simpleSet[T]) Size() int {
	return len(s.values)
}

// Empty implements the Set interface for simpleSet
func (s *simpleSet[T]) Empty() bool {
	return s.Size() == 0
}

// AddAll implements the Set interface for simpleSet
func (s *simpleSet[T]) AddAll(values ...T) {
	for _, v := range values {
		s.Add(v)
	}
}

// Intersection implements the Set interface for simpleSet
func (s *simpleSet[T]) Intersection(other Set[T]) Set[T] {
	newSet := NewSet[T]()

	for k := range s.values {
		if other.Contains(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

// Difference implements the Set interface for simpleSet
func (s *simpleSet[T]) Difference(other Set[T]) Set[T] {
	newSet := NewSet[T]()

	for k := range s.values {
		if !other.Contains(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

// SubSetOf implements the Set interface for simpleSet
func (s *simpleSet[T]) SubSetOf(other Set[T]) bool {
	return s.Difference(other).Empty()
}

// Union implements the Set interface for simpleSet
func (s *simpleSet[T]) Union(other Set[T]) Set[T] {
	newSet := SetFrom(s.ToSlice()...)
	newSet.AddAll(other.ToSlice()...)
	return newSet
}

// Contains implements the Set interface for simpleSet
func (s *simpleSet[T]) Contains(v T) bool {
	return s.values[v]
}

// ToSlice implements the Set interface for simpleSet
func (s *simpleSet[T]) ToSlice() []T {
	return s.keyView.Seq().IntoSlice()
}

// Equal implements the Set interface for simpleSet
func (s *simpleSet[T]) Equal(other Set[T]) bool {
	return s.Difference(other).Empty() && other.Difference(s).Empty()
}

// Each implements the Iterator interface
func (s *simpleSet[T]) Each(f Proc1[T]) {
	if s == nil {
		return
	}

	s.keyView.Each(f)
}

// Iter implements the Iterable interface
func (s *simpleSet[T]) Iter() Iterator[T] {
	return s
}
