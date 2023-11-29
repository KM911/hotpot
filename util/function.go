package util

func Map[T any](array []T, callback func(T) T) []T {
	for i := 0; i < len(array); i++ {
		array[i] = callback(array[i])
	}
	return array
}
