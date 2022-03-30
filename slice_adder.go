package gogcoll

// sliceAdder is a helper object that can be used to wrap
// a slice and keep a pointer to it so you can later add to it.
// Adding to the slice will generate a new slice, so you have
// to unwrap the adder to get at the underlying slice again after
// all operations are done.
type sliceAdder[T any] struct {
	s []T
}

// Unwrap returns the current slice
func (a *sliceAdder[T]) Unwrap() []T {
	return a.s
}

// Add creates a new slice with the given element added to the end of it
func (a *sliceAdder[T]) Add(v T) {
	a.s = append(a.s, v)
}
