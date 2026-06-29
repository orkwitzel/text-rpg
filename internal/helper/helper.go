package helper

func RemoveElementFromArray[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
