package gogcoll

// Filter takes an iterator of values and a predicate function and a result
// which can be added into. It will go through all values, and add the
// elements for which the predicate returns true to the result.
func Filter[T any, R Adder[T]](values Iterator[T], p Predicate[T], result R) R {
	return Reduce(values, result, func(acc R, value T) R {
		if p(value) {
			acc.Add(value)
		}
		return acc
	})
}

// FilterIntoSlice uses Filter to add every element from the iterator where
// the predicate returns true to a slice, and return that slice.
func FilterIntoSlice[T any](values Iterator[T], p Predicate[T]) []T {
	return Filter(values, p, &sliceAdder[T]{[]T{}}).Unwrap()
}
