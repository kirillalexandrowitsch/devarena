package hero

import "testing"

func TestBuildSummaryValueReturnsSummary(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := BuildSummaryValue(gameHero)

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}

func TestBuildSummaryPointerReturnsSummaryPointer(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	summary := BuildSummaryPointer(gameHero)

	if summary == nil {
		t.Fatal("expected summary pointer not to be nil")
	}

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}

func TestBuildSummaryAsAnyReturnsSummaryValue(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	value := BuildSummaryAsAny(gameHero)

	summary, ok := value.(HeroSummary)
	if !ok {
		t.Fatalf("expected HeroSummary, got %T", value)
	}

	if summary.Name != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", summary.Name)
	}
}

func TestBuildSummaryNameReaderCapturesSummary(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Name:  "Ragnar",
		Class: HeroClassWarrior,
		Level: 3,
		Alive: true,
	}

	readName := BuildSummaryNameReader(gameHero)

	if readName() != "Ragnar" {
		t.Fatalf("expected summary name %q, got %q", "Ragnar", readName())
	}
}
