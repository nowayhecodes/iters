package iters

type mapIterator[T any] struct {
	source Iterator[T]
	mapper func(T) T
}

func (iter *mapIterator[T]) Next() bool {
	return iter.source.Next()
}

func (iter *mapIterator[T]) Value() T {
	value := iter.source.Value()
	return iter.mapper(value)
}

// It just creates a map iterator and returns it.
func Map[T any](iter Iterator[T], fn func(T) T) Iterator[T] {
	return &mapIterator[T]{
		iter, fn,
	}
}
