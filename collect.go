package iters

// Collects the results from the transformation performed on the input data.
// It takes an iterator as input and fills up a slice with values from the iterator.
//
//   mapped := Map([4, 5], func(x int) int {
// 	         return x * x
//   })
//
//   result := Collect(mapped)
func Collect[T any](iter Iterator[T]) []T {
	var items []T

	for iter.Next() {
		items = append(items, iter.Value())
	}

	return items
}
