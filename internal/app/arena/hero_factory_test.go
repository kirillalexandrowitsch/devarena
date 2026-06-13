package arena

import (
	"errors"
	"testing"

	"github.com/rudyakovk/devarena/internal/domain/hero"
)

func TestCreateDefaultHeroReturnsHero(t *testing.T) {
	gameHero, err := createDefaultHero()

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if gameHero.Name != defaultHeroName {
		t.Fatalf("expected hero name %q, got %q", defaultHeroName, gameHero.Name)
	}

	if !gameHero.Alive {
		t.Fatal("expected default hero to be alive")
	}
}

func TestCreateDefaultHeroWithNameReturnsErrorChainForEmptyName(t *testing.T) {
	_, err := createDefaultHeroWithName("")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, hero.ErrHeroNameEmpty) {
		t.Fatalf("expected error chain to contain ErrHeroNameEmpty, got %v", err)
	}

	expectedMessage := "create default hero: create hero: hero name is empty"
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
	}

	domainErr := errors.Unwrap(err)
	if domainErr == nil {
		t.Fatal("expected wrapped domain error, got nil")
	}

	if domainErr.Error() != "create hero: hero name is empty" {
		t.Fatalf("expected domain error %q, got %q", "create hero: hero name is empty", domainErr.Error())
	}

	rootErr := errors.Unwrap(domainErr)
	if rootErr != hero.ErrHeroNameEmpty {
		t.Fatalf("expected root error ErrHeroNameEmpty, got %v", rootErr)
	}
}

func TestCreateDefaultHeroWithNameReturnsValidationErrorChainForShortName(t *testing.T) {
	_, err := createDefaultHeroWithName("Ra")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var validationErr hero.ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected error chain to contain ValidationError, got %v", err)
	}

	if validationErr.Field != "hero_name" {
		t.Fatalf("expected validation field %q, got %q", "hero_name", validationErr.Field)
	}

	expectedMessage := `create default hero: create hero: hero_name: value "Ra" is too short: length 2 is less than 3`
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
	}
}
