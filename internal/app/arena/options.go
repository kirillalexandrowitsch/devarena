package arena

import (
	"io"
	"os"
)

type Option func(*App)

type App struct {
	printSessionReport bool
	showInventoryInfo  bool
	showStatSummary    bool
	output             io.Writer
}

func NewApp(options ...Option) *App {
	app := &App{
		printSessionReport: true,
		showInventoryInfo:  true,
		showStatSummary:    true,
		output:             os.Stdout,
	}

	for _, option := range options {
		option(app)
	}

	return app
}

func WithSessionReport(enabled bool) Option {
	return func(app *App) {
		app.printSessionReport = enabled
	}
}

func WithInventoryInfo(enabled bool) Option {
	return func(app *App) {
		app.showInventoryInfo = enabled
	}
}

func WithStatSummary(enabled bool) Option {
	return func(app *App) {
		app.showStatSummary = enabled
	}
}

func WithOutput(output io.Writer) Option {
	return func(app *App) {
		if output == nil {
			return
		}

		app.output = output
	}
}
