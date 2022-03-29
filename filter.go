package gogcoll

func Filter[T any, R Adder[T]](values Iterator[T], p func(T) bool, result R) R {
	return Reduce(values, result, func(acc R, value T) R {
		if p(value) {
			acc.Add(value)
		}
		return acc
	})
}

func FilterIntoSlice[T any](values Iterator[T], p func(T) bool) []T {
	return Filter(values, p, &sliceAdder[T]{[]T{}}).Unwrap()
}
