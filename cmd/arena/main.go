package main

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/app/arena"
)

func main() {
	defer recoverPanic(func(value any) {
		fmt.Println("Arena stopped after panic:", value)
	})

	app := arena.NewApp(
		arena.WithSessionReport(true),
		arena.WithInventoryInfo(true),
		arena.WithStatSummary(true),
	)

	app.Run()
}
