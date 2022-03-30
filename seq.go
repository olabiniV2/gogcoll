package gogcoll

// Seq represents a lazy sequence that can also do some basic filtering
// and generating operations
type Seq[T any] interface {
	Iterator[T]
	BasicSeq[T]

	// Into will evaluate the full sequence and add it into the given adder
	Into(Adder[T])
	// IntoSlice will evaluate the full sequence and generate a slice of the result
	IntoSlice() []T
	// Filter will take the given predicate and return a new lazy sequence that
	// only yields values for which the predicate returns true.
	Filter(Predicate[T]) Seq[T]
}

// ToSeq will create a sequence from any iterator. This method uses goroutines
// and channels to turn the iteration into a sequence, and is thus not extremely
// efficient.
func ToSeq[T any](i Iterator[T]) Seq[T] {
	if ss, ok := i.(Seqable[T]); ok {
		return ss.Seq()
	}

	resultChan := make(chan T)
	seq := &iteratorSeq[T]{
		c: resultChan,
	}
	seq.self = seq
	go seq.start(i)
	return seq
}

// Take
// Drop
