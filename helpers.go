package gogcoll

func zero[T any]() T {
	return *new(T)
}
