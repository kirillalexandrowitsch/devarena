package main

func recoverPanic(report func(value any)) {
	value := recover()
	if value == nil {
		return
	}

	report(value)
}
