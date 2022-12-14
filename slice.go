package iters

type SliceIterator[T any] struct {
	Elements []T
	value    T
	index    int
}

// Creates a slice iterator over the elements []T
func NewSliceIterator[T any](elements []T) Iterator[T] {
	return &SliceIterator[T]{
		Elements: elements,
	}
}

func (iter *SliceIterator[T]) Next() bool {
	if iter.index < len(iter.Elements) {
		iter.value = iter.Elements[iter.index]
		iter.index++
		return true
	}
	return false
}

func (iter *SliceIterator[T]) Value() T {
	return iter.value
}
