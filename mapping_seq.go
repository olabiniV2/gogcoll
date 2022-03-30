package gogcoll

// mappingSeq is used to implement a lazy Map
type mappingSeq[T, U any] struct {
	baseSeq[U]
	f       Func1[T, U]
	realSeq BasicSeq[T]
}

// Next returns the next transformed value, implementing part of the Seq interface
func (s *mappingSeq[T, U]) Next() U {
	return s.f(s.realSeq.Next())
}

// HasNext implements part of the Seq interface
func (s *mappingSeq[T, U]) HasNext() bool {
	return s.realSeq.HasNext()
}

// LazyMap returns a new sequence that will yield values transformed by the function given
func LazyMap[T, U any](values BasicSeq[T], f Func1[T, U]) Seq[U] {
	sq := &mappingSeq[T, U]{
		realSeq: values,
		f:       f,
	}
	sq.self = sq
	return sq
}
