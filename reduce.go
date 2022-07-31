package iters

func Reduce[T, V any](iter Iterator[V], fn Reducer[T, V]) T {
	var acc T

	for iter.Next() {
		acc = fn(acc, iter.Value())
	}
	return acc
}
