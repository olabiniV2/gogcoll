package gogcoll

type sliceAdder[T any] struct {
	s []T
}

func (a *sliceAdder[T]) Unwrap() []T {
	return a.s
}

func (a *sliceAdder[T]) Add(v T) {
	a.s = append(a.s, v)
}
