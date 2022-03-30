package gogcoll

// zero generically returns the zero value for any type
func zero[T any]() T {
	// This idiom is a bit harder to understand than the
	// one where you create a new variable with a type and returns it,
	// but apparently this is the idiomatic way - especially since it's a one-liner.
	return *new(T)
}
