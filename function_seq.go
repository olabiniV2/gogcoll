package gogcoll

// functionSeq is a generic Seq implementation that uses function fields
// for the Next and HasNext functionalities, to make it easier to create
// small Seq implementations.
type functionSeq[T any] struct {
	baseSeq[T]
	next    FixedFunction[T]
	hasNext FixedFunction[bool]
}

// createFunctionSequence takes functions for Next and HasNext and returns a Seq that will
// use the given functions to implement the full Seq functionality.
func createFunctionSequence[T any](next FixedFunction[T], hasNext FixedFunction[bool]) Seq[T] {
	fs := &functionSeq[T]{
		next:    next,
		hasNext: hasNext,
	}
	fs.self = fs
	return fs
}

// Next is part of the implementation of the Seq interface
func (s *functionSeq[T]) Next() T {
	return s.next()
}

// HasNext is part of the implementation of the Seq interface
func (s *functionSeq[T]) HasNext() bool {
	return s.hasNext()
}
