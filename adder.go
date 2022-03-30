package gogcoll

// Adder is an interface that corresponds to a type that can mutably add an item
// to its underlying type. This type is used in places such as `Map` and `Reduce`
// in order to generically inject result data into arbitrary containers.
type Adder[T any] interface {
	// Add will take a value of the generic type T and add it to the underlying container.
	Add(T)
}
