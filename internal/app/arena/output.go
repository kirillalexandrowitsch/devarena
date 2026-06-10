package arena

import "fmt"

func (app *App) println(values ...any) {
	fmt.Fprintln(app.output, values...)
}
