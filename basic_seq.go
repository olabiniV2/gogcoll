package gogcoll

// BasicSeq is the smallest possible lazy sequence interface.
// It allows you to pull the next element from the sequence, and
// check if there are more elements left in the sequence.
// BasicSeq is not defined to be concurrency safe and no implementation
// of it can be relied on to be concurrency safe, unless otherwise stated
// for that implementation.
type BasicSeq[T any] interface {
	// Next returns the next possible element from the sequence. If the sequence has
	// reached its end, it will return the zero value for the generic type T.
	// Every call to Next will move the sequence forward one step.
	Next() T
	// HasNext returns true if there are more elements in the sequence, and false otherwise.
	// It is safe to call HasNext multiple times without calling Next inbetween.
	HasNext() bool
}
