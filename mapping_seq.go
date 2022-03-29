package gogcoll

type mappingSeq[T, U any] struct {
	baseSeq[U]
	f       Func1[T, U]
	realSeq BasicSeq[T]
}

func (s *mappingSeq[T, U]) Next() U {
	return s.f(s.realSeq.Next())
}

func (s *mappingSeq[T, U]) HasNext() bool {
	return s.realSeq.HasNext()
}

func LazyMap[T, U any](values BasicSeq[T], f Func1[T, U]) Seq[U] {
	sq := &mappingSeq[T, U]{
		realSeq: values,
		f:       f,
	}
	sq.self = sq
	return sq
}
