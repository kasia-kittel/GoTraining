package patterns


// Higher-Order Functions - Map-Filter-Reduce
// run function on every element

func mapFn[T any, E any](a []T, fn func(T) E) []E {
	result := make([]E, len(a))

	for i, v := range a {
		result[i] = fn(v)
	}

	return result
}

func filterFn[T any](a []T, fn func(T) bool) []T {
	var result []T

	for _, v := range a {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func reduceFn[T any, E any](a []T, fn func(E, T) E) E {
	var result E

	for _, v := range a {
		result = fn(result, v)
	}
	return result
}

// TODO concurrent implementations?

//Closures and currying
// Iterator
// monads / monoids