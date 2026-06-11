package hero

import "testing"

func TestNewHeroSummaryReturnsValue(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := NewHeroSummary(gameHero)

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}

	if summary.Level != 3 {
		t.Fatalf("expected summary level %d, got %d", 3, summary.Level)
	}
}

func TestNewHeroSummaryPointerReturnsPointer(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := NewHeroSummaryPointer(gameHero)

	if summary == nil {
		t.Fatal("expected summary pointer not to be nil")
	}

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}

func TestHeroSummaryReturnsIndependentValue(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := gameHero.Summary()

	gameHero.Name = "Changed"

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}

func TestHeroSummaryPointerReturnsIndependentPointer(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := gameHero.SummaryPointer()

	gameHero.Name = "Changed"

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}
