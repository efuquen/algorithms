package iter

type Iterator[T any] chan T

type Iterable[T any] interface {
	Iter() Iterator[*T]
}
