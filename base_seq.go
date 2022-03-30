package gogcoll

// baseSeq is an implementation of the fundamental operations of the `Seq` interface.
// By inheriting frmo baseSeq and setting `self` correctly, you can avoid the
// extra work of implementing all `Seq` methods yourself.
type baseSeq[T any] struct {
	self BasicSeq[T]
}

// Into is part of the implementation of the Seq interface
func (s *baseSeq[T]) Into(a Adder[T]) {
	s.Each(func(v T) {
		a.Add(v)
	})
}

// IntoSlice is part of the implementation of the Seq interface
func (s *baseSeq[T]) IntoSlice() []T {
	return Reduce[T, []T](s, []T{}, Append[T])
}

// Each is the implementation for the Iterator interface, which is also a part of the Seq interface
func (s *baseSeq[T]) Each(f Proc1[T]) {
	for s.self.HasNext() {
		f(s.self.Next())
	}
}

// Filter is part of the implementation of the Seq interface
func (s *baseSeq[T]) Filter(p Predicate[T]) Seq[T] {
	sq := &filterSeq[T]{
		realSeq: s.self,
		p:       p,
	}
	sq.self = sq
	return sq
}

// simpleFullSeq allows you to implement a full Seq even if you only have an implementation
// of the BasicSeq interface. This is the type that FullSeqFrom will return.
type simpleFullSeq[T any] struct {
	baseSeq[T]
}

// Next is part of the implementation of the BasicSeq interface
func (s *simpleFullSeq[T]) Next() T {
	return s.self.Next()
}

// HasNext is part of the implementation of the BasicSeq interface
func (s *simpleFullSeq[T]) HasNext() bool {
	return s.self.HasNext()
}

// FullSeqFrom takes a BasicSeq and returns a full Seq implementation using
// only the methods defined in BasicSeq
func FullSeqFrom[T any](s BasicSeq[T]) Seq[T] {
	// Might be a good idea to see if the sequence
	// already is a Seq and in that case return that.

	sq := &simpleFullSeq[T]{}
	sq.self = s
	return sq
}
