package gogcoll

// Seqable represents a type that can be lazily iterated over by generating
// a Seq from it
type Seqable[T any] interface {
	// Seq returns a sequence for this type
	Seq() Seq[T]
}
