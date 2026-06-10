package shared

type Selection[T any] struct {
	Value T
	Found bool
}

func First[T any](items []T) Selection[T] {
	if len(items) == 0 {
		var zero T

		return Selection[T]{
			Value: zero,
			Found: false,
		}
	}

	return Selection[T]{
		Value: items[0],
		Found: true,
	}
}
