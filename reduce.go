package iters

// Reduce values iterated over to a single value
//
//   var numbers []int = []int{1, 2, 3, 4, 5}
//   iter := NewSliceIterator(numbers)
//
//   result := Reduce(iter, func(accum, x int) int {
// 	         return accum + x
//   })
func Reduce[T, V any](iter Iterator[V], fn Reducer[T, V]) T {
	var acc T

	for iter.Next() {
		acc = fn(acc, iter.Value())
	}
	return acc
}
