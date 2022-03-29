package gogcoll

type baseSeq[T any] struct {
	self BasicSeq[T]
}

func (s *baseSeq[T]) Into(a Adder[T]) {
	s.Each(func(v T) {
		a.Add(v)
	})
}

func (s *baseSeq[T]) IntoSlice() []T {
	return Reduce[T, []T](s, []T{}, Append[T])
}

func (s *baseSeq[T]) Each(f Proc1[T]) {
	for s.self.HasNext() {
		f(s.self.Next())
	}
}

func (s *baseSeq[T]) Filter(p Predicate[T]) Seq[T] {
	sq := &filterSeq[T]{
		realSeq: s.self,
		p:       p,
	}
	sq.self = sq
	return sq
}
