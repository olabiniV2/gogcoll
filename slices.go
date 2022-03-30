package gogcoll

// Slice wraps a slice so that you can iterate or sequence over it
type Slice[T any] []T

// Iter implements the Iterable interface for a Slice
func (s Slice[T]) Iter() Iterator[T] {
	return s
}

// Each implements the Iterator interface for Slice
func (s Slice[T]) Each(f Proc1[T]) {
	for _, v := range s {
		f(v)
	}
}

// Seq implements the Seqable interface for Slice
func (s Slice[T]) Seq() Seq[T] {
	return sliceSeq(s)
}

// sliceSeq takes a slice and returns a Seq over that slice.
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

// Append wraps the built in append, so that you can pass it around as a
// function object
func Append[T any](vs []T, v T) []T {
	return append(vs, v)
}
