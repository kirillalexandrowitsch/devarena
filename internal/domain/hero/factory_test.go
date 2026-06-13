package hero

import (
	"errors"
	"testing"
)

func TestNewHeroReturnsHeroForValidInput(t *testing.T) {
	stats := CombatStats{
		HP:             100,
		BaseDamage:     15,
		BonusDamage:    5,
		CriticalChance: 0.15,
	}

	gameHero, err := NewHero(1, "Ragnar", HeroClassWarrior, 3, stats)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if gameHero.ID != 1 {
		t.Fatalf("expected hero id %d, got %d", 1, gameHero.ID)
	}

	if gameHero.Name != "Ragnar" {
		t.Fatalf("expected hero name %q, got %q", "Ragnar", gameHero.Name)
	}

	if !gameHero.Alive {
		t.Fatal("expected hero to be alive")
	}
}

func TestNewHeroReturnsWrappedErrorForEmptyName(t *testing.T) {
	stats := CombatStats{
		HP:             100,
		BaseDamage:     15,
		BonusDamage:    5,
		CriticalChance: 0.15,
	}

	gameHero, err := NewHero(1, "", HeroClassWarrior, 3, stats)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrHeroNameEmpty) {
		t.Fatalf("expected error chain to contain ErrHeroNameEmpty, got %v", err)
	}

	expectedMessage := "create hero: hero name is empty"
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
	}

	if gameHero.ID != 0 {
		t.Fatalf("expected zero hero id, got %d", gameHero.ID)
	}

	if gameHero.Name != "" {
		t.Fatalf("expected zero hero name, got %q", gameHero.Name)
	}

	if gameHero.Class != "" {
		t.Fatalf("expected zero hero class, got %q", gameHero.Class)
	}

	if gameHero.Level != 0 {
		t.Fatalf("expected zero hero level, got %d", gameHero.Level)
	}

	if gameHero.Alive {
		t.Fatal("expected zero hero to be not alive")
	}
}

func TestNewHeroReturnsWrappedValidationErrorForShortName(t *testing.T) {
	stats := CombatStats{
		HP:             100,
		BaseDamage:     15,
		BonusDamage:    5,
		CriticalChance: 0.15,
	}

	_, err := NewHero(1, "Ra", HeroClassWarrior, 3, stats)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var validationErr ValidationError
	if !errors.As(err, &validationErr) {
		t.Fatalf("expected ValidationError in chain, got %T", err)
	}

	if validationErr.Field != "hero_name" {
		t.Fatalf("expected field %q, got %q", "hero_name", validationErr.Field)
	}

	expectedMessage := `create hero: hero_name: value "Ra" is too short: length 2 is less than 3`
	if err.Error() != expectedMessage {
		t.Fatalf("expected error %q, got %q", expectedMessage, err.Error())
	}
}
