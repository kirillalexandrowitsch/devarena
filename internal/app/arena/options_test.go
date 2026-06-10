package arena

import "testing"

func TestNewAppUsesDefaultOptions(t *testing.T) {
	app := NewApp()

	if !app.printSessionReport {
		t.Fatal("expected session report to be enabled by default")
	}

	if !app.showInventoryInfo {
		t.Fatal("expected inventory info to be enabled by default")
	}

	if !app.showStatSummary {
		t.Fatal("expected stat summary to be enabled by default")
	}
}

func TestNewAppAppliesFunctionalOptions(t *testing.T) {
	app := NewApp(
		WithSessionReport(false),
		WithInventoryInfo(false),
		WithStatSummary(false),
	)

	if app.printSessionReport {
		t.Fatal("expected session report to be disabled")
	}

	if app.showInventoryInfo {
		t.Fatal("expected inventory info to be disabled")
	}

	if app.showStatSummary {
		t.Fatal("expected stat summary to be disabled")
	}
}
