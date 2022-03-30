package gogcoll

// filterSeq is a Seq implementation that is returned from most Filter methods on
// other sequences.
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

// Next is part of the implemention of BasicSeq
func (s *filterSeq[T]) Next() T {
	s.readNext()
	if s.hasNext {
		s.hasNext = false
		return s.nextValue
	}
	return zero[T]()
}

// HasNext is part of the implemention of BasicSeq
func (s *filterSeq[T]) HasNext() bool {
	s.readNext()
	return s.hasNext
}
