package gogcoll

// Map will turn the values in the given iterator into values of another type, using
// the transform function, and then adding these resulting values to the
// given Adder, and returning it.
func Map[T any, U any, R Adder[U]](values Iterator[T], transform func(T) U, result R) R {
	return Reduce(values, result, func(acc R, value T) R {
		acc.Add(transform(value))
		return acc
	})
}

// MapIntoSlice transforms the values given in the iterator using the transform function
// and returns a slice of these resulting values.
func MapIntoSlice[T any, U any](values Iterator[T], transform func(T) U) []U {
	return Map(values, transform, &sliceAdder[U]{[]U{}}).Unwrap()
}
