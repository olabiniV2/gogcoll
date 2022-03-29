package gogcoll

type Slice[T any] []T

func (s *Slice[T]) Iter() Iterator[T] {
	return s
}

func (s *Slice[T]) Each(f Proc1[T]) {
	if s == nil {
		return
	}

	for _, v := range *s {
		f(v)
	}
}

func (s *Slice[T]) Seq() Seq[T] {
	return sliceSeq(*s)
}

func sliceSeq[T any](sl []T) Seq[T] {
	current := -1
	next := func() T {
		current += 1
		if current < len(sl) {
			return sl[current]
		}
		return zero[T]()
	}

	hasNext := func() bool {
		return current+1 < len(sl)
	}

	return createFunctionSequence(next, hasNext)
}

func Append[T any](vs []T, v T) []T {
	return append(vs, v)
}
