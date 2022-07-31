package iters

func Collect[T any](iter Iterator[T]) []T {
	var items []T

	for iter.Next() {
		items = append(items, iter.Value())
	}

	return items
}
