package gogcoll

type Seqable[T any] interface {
	Seq() Seq[T]
}
