package gogcoll

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Each(Proc1[T])
}
