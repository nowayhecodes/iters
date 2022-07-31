package iters

type Iterator[T any] interface {
	Next() bool
	Value() T
}

type Reducer[T, V any] func(acc T, value V) T
