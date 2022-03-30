package gogcoll

// Iterable represents an object that can create an iterator
type Iterable[T any] interface {
	// Iter returns an Iterator for the type of the collection
	Iter() Iterator[T]
}

// Iterator represents an object that has the capacity to
// visit each elements. THis interface does not guarantee
// an ordered visit - only that each element will
// be visited once.
type Iterator[T any] interface {
	// Each will call the given procedure once for every element in the collection,
	// giving the element as an argument to the procedure.
	Each(Proc1[T])
}
