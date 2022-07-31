package iters

type filterIterator[T any] struct {
	source    Iterator[T]
	predicate func(T) bool
}

func (iter *filterIterator[T]) Next() bool {
	for iter.source.Next() {
		if iter.predicate(iter.source.Value()) {
			return true
		}
	}

	return true
}

func (iter *filterIterator[T]) Value() T {
	return iter.source.Value()
}

func Filter[T any](iter Iterator[T], predicate func(T) bool) Iterator[T] {
	return &filterIterator[T]{
		iter, predicate,
	}
}
