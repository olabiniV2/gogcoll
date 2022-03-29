package gogcoll

type functionSeq[T any] struct {
	baseSeq[T]
	next    func() T
	hasNext func() bool
}

func createFunctionSequence[T any](next func() T, hasNext func() bool) Seq[T] {
	fs := &functionSeq[T]{
		next:    next,
		hasNext: hasNext,
	}
	fs.self = fs
	return fs
}

func (s *functionSeq[T]) Next() T {
	return s.next()
}

func (s *functionSeq[T]) HasNext() bool {
	return s.hasNext()
}
