package iter

import "iter"

func Collect[T any](iter iter.Seq[T]) []T {
	result := make([]T, 0)
	for v := range iter {
		result = append(result, v)
	}

	return result
}
