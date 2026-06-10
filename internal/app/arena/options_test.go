package arena

import (
	"bytes"
	"testing"
)

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

	if app.output == nil {
		t.Fatal("expected output to be set by default")
	}
}

func TestNewAppAppliesFunctionalOptions(t *testing.T) {
	output := &bytes.Buffer{}

	app := NewApp(
		WithSessionReport(false),
		WithInventoryInfo(false),
		WithStatSummary(false),
		WithOutput(output),
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

	if app.output != output {
		t.Fatal("expected custom output to be applied")
	}
}

func TestWithOutputIgnoresNilOutput(t *testing.T) {
	app := NewApp(
		WithOutput(nil),
	)

	if app.output == nil {
		t.Fatal("expected nil output not to replace default output")
	}
}
