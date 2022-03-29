package gogcoll

type Seq[T any] interface {
	Iterator[T]
	BasicSeq[T]

	Into(Adder[T])
	IntoSlice() []T
	Filter(Predicate[T]) Seq[T]
}

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
