package gogcoll

type BasicSeq[T any] interface {
	Next() T
	HasNext() bool
}
