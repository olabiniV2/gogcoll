package gogcoll

func Reduce[T any, R any](values Iterator[T], initial R, f func(R, T) R) R {
	acc := initial

	values.Each(func(value T) {
		acc = f(acc, value)
	})

	return acc
}
