package gogcoll

func Map[T any, U any, R Adder[U]](values Iterator[T], transform func(T) U, result R) R {
	return Reduce(values, result, func(acc R, value T) R {
		acc.Add(transform(value))
		return acc
	})
}

func MapIntoSlice[T any, U any](values Iterator[T], transform func(T) U) []U {
	return Map(values, transform, &sliceAdder[U]{[]U{}}).Unwrap()
}
