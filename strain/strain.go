package strain

// Keep returns an array of elements which satisfy predicate.
func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range collection {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Keep returns an array of elements which do not satisfy predicate.
func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range collection {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
