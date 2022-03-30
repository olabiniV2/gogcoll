package gogcoll

// iteratorSeq is a Seq implementation that can use an underlying
// iterator. It uses a goroutine and a channel to be able to manage
// turning an eager algorithm into a lazy one, so it's a bit more
// expensive than native implementations of Seqs. However, sometimes
// this is necessary to be able to keep track of where you are in an
// iteration, since Golang doesn't always give this information. One
// good example of this is when iterating over a hash, and you want to
// avoid creating a slice of the keys or values and keep it in memory
// first.
type iteratorSeq[T any] struct {
	baseSeq[T]
	c         chan T
	hasNext   bool
	nextValue T
}

func (s *iteratorSeq[T]) readNext() {
	// assumes s.hasNext is false
	if s.c == nil {
		return
	}
	v, ok := <-s.c
	if ok {
		s.hasNext = true
		s.nextValue = v
	}
}

// Next is part of the implementation for the Seq interface
func (s *iteratorSeq[T]) Next() T {
	if !s.hasNext {
		s.readNext()
	}
	if !s.hasNext {
		return zero[T]()
	}
	s.hasNext = false
	return s.nextValue
}

// HasNext is part of the implementation for the Seq interface
func (s *iteratorSeq[T]) HasNext() bool {
	if s.hasNext {
		return true
	}
	s.readNext()
	return s.hasNext
}

func (s *iteratorSeq[T]) start(i Iterator[T]) {
	i.Each(func(val T) {
		s.c <- val
	})
	close(s.c)
	s.c = nil
}
