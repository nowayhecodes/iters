package iters

type Iterator[T any] interface {
	Next() bool
	Value() T
}
