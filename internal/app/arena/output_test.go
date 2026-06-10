package arena

import (
	"bytes"
	"strings"
	"testing"
)

func TestAppPrintlnWritesToInjectedOutput(t *testing.T) {
	output := &bytes.Buffer{}

	app := NewApp(
		WithOutput(output),
	)

	app.println("Hello", "DevArena")

	result := output.String()

	if !strings.Contains(result, "Hello DevArena") {
		t.Fatalf("expected output to contain %q, got %q", "Hello DevArena", result)
	}
}
