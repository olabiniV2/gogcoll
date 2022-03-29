package gogcoll

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
