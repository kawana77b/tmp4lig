package util

func Map[T any, U any](arr []T, fn func(T) U) []U {
	var newArr []U
	for _, item := range arr {
		newArr = append(newArr, fn(item))
	}
	return newArr
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	var newArr []T
	for _, item := range arr {
		if fn(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}
