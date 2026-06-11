package hero

import "testing"

func TestAllocateHeroSummariesReturnsSummaries(t *testing.T) {
	heroes := []Hero{
		{
			ID:    1,
			Name:  "Ragnar",
			Class: HeroClassWarrior,
			Level: 3,
			Alive: true,
		},
		{
			ID:    2,
			Name:  "Merlin",
			Class: HeroClassMage,
			Level: 5,
			Alive: true,
		},
	}

	summaries := AllocateHeroSummaries(heroes)

	if len(summaries) != 2 {
		t.Fatalf("expected summaries length %d, got %d", 2, len(summaries))
	}

	if cap(summaries) < 2 {
		t.Fatalf("expected summaries capacity at least %d, got %d", 2, cap(summaries))
	}
}

func TestReuseHeroSummariesReusesCapacity(t *testing.T) {
	heroes := []Hero{
		{
			ID:    1,
			Name:  "Ragnar",
			Class: HeroClassWarrior,
			Level: 3,
			Alive: true,
		},
	}

	summaries := make([]HeroSummary, 0, 8)

	reused := ReuseHeroSummaries(heroes, summaries)

	if len(reused) != 1 {
		t.Fatalf("expected reused summaries length %d, got %d", 1, len(reused))
	}

	if cap(reused) != 8 {
		t.Fatalf("expected reused capacity %d, got %d", 8, cap(reused))
	}
}

func TestIndexHeroesByIDReturnsIndex(t *testing.T) {
	heroes := []Hero{
		{
			ID:    1,
			Name:  "Ragnar",
			Class: HeroClassWarrior,
			Level: 3,
			Alive: true,
		},
		{
			ID:    2,
			Name:  "Merlin",
			Class: HeroClassMage,
			Level: 5,
			Alive: false,
		},
	}

	index := IndexHeroesByID(heroes)

	if len(index) != 2 {
		t.Fatalf("expected index length %d, got %d", 2, len(index))
	}

	if index[1].Name != "Ragnar" {
		t.Fatalf("expected hero name %q, got %q", "Ragnar", index[1].Name)
	}

	if index[2].Name != "Merlin" {
		t.Fatalf("expected hero name %q, got %q", "Merlin", index[2].Name)
	}
}
