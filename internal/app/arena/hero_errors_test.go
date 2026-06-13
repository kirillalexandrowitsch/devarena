package arena

import (
	"fmt"
	"testing"

	"github.com/rudyakovk/devarena/internal/domain/hero"
)

func TestDescribeHeroCreationErrorReturnsEmptyStringForNil(t *testing.T) {
	message := describeHeroCreationError(nil)

	if message != "" {
		t.Fatalf("expected empty message, got %q", message)
	}
}

func TestDescribeHeroCreationErrorReturnsMessageForEmptyName(t *testing.T) {
	message := describeHeroCreationError(hero.ErrHeroNameEmpty)

	if message != "default hero name is empty" {
		t.Fatalf("expected empty name message, got %q", message)
	}
}

func TestDescribeHeroCreationErrorReturnsMessageForWrappedEmptyName(t *testing.T) {
	err := fmt.Errorf("create hero: %w", hero.ErrHeroNameEmpty)

	message := describeHeroCreationError(err)

	if message != "default hero name is empty" {
		t.Fatalf("expected empty name message, got %q", message)
	}
}

func TestDescribeHeroCreationErrorReturnsFallbackMessage(t *testing.T) {
	err := hero.ValidationError{Message: "hero name is too short"}

	message := describeHeroCreationError(err)

	expectedMessage := "default hero creation failed: hero name is too short"
	if message != expectedMessage {
		t.Fatalf("expected message %q, got %q", expectedMessage, message)
	}
}
