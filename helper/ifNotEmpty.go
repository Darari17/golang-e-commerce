package helper

func IfNotEmpty[T comparable](oldValue T, newValue T) T {
	var zero T
	if newValue != zero {
		return newValue
	}

	return oldValue
}
