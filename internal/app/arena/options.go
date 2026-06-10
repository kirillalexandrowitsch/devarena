package arena

type Option func(*App)

type App struct {
	printSessionReport bool
	showInventoryInfo  bool
	showStatSummary    bool
}

func NewApp(options ...Option) *App {
	app := &App{
		printSessionReport: true,
		showInventoryInfo:  true,
		showStatSummary:    true,
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
