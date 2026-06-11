package hero

import "testing"

func TestCompactHeroSummarySizeIsNotLargerThanPadded(t *testing.T) {
	compactSize := CompactHeroSummarySize()
	paddedSize := PaddedHeroSummarySize()

	if compactSize > paddedSize {
		t.Fatalf("expected compact size %d not to be larger than padded size %d", compactSize, paddedSize)
	}
}

func TestCompactHeroSummaryFromHero(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Level: 3,
		Alive: true,
	}

	summary := CompactHeroSummaryFromHero(gameHero)

	if summary.ID != 1 {
		t.Fatalf("expected summary id %d, got %d", 1, summary.ID)
	}

	if summary.Level != 3 {
		t.Fatalf("expected summary level %d, got %d", 3, summary.Level)
	}

	if !summary.Alive {
		t.Fatal("expected summary to be alive")
	}
}

func TestPaddedHeroSummaryFromHero(t *testing.T) {
	gameHero := Hero{
		ID:    1,
		Level: 3,
		Alive: true,
	}

	summary := PaddedHeroSummaryFromHero(gameHero)

	if summary.ID != 1 {
		t.Fatalf("expected summary id %d, got %d", 1, summary.ID)
	}

	if summary.Level != 3 {
		t.Fatalf("expected summary level %d, got %d", 3, summary.Level)
	}

	if !summary.Alive {
		t.Fatal("expected summary to be alive")
	}
}
