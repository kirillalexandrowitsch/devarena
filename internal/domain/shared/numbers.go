package shared

type OrderedNumber interface {
	~int | ~int64 | ~float64
}

func Max[T OrderedNumber](values []T) Selection[T] {
	if len(values) == 0 {
		var zero T

		return Selection[T]{
			Value: zero,
			Found: false,
		}
	}

	maxValue := values[0]

	for _, value := range values[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return Selection[T]{
		Value: maxValue,
		Found: true,
	}
}
