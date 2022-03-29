package gogcoll

type filterSeq[T any] struct {
	baseSeq[T]
	p         Predicate[T]
	realSeq   BasicSeq[T]
	hasNext   bool
	nextValue T
}

func (s *filterSeq[T]) readNext() {
	if s.hasNext {
		return
	}
	for s.realSeq.HasNext() {
		v := s.realSeq.Next()
		if s.p(v) {
			s.hasNext = true
			s.nextValue = v
			return
		}
	}
}

func (s *filterSeq[T]) Next() T {
	s.readNext()
	if s.hasNext {
		s.hasNext = false
		return s.nextValue
	}
	return zero[T]()
}

func (s *filterSeq[T]) HasNext() bool {
	s.readNext()
	return s.hasNext
}
