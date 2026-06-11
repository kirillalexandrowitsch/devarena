package hero

import "testing"

func TestBuildHeroSummariesReturnsSummaries(t *testing.T) {
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

	summaries := BuildHeroSummaries(heroes)

	if len(summaries) != 2 {
		t.Fatalf("expected summaries length %d, got %d", 2, len(summaries))
	}

	if summaries[0].Name != "Ragnar" {
		t.Fatalf("expected first summary name %q, got %q", "Ragnar", summaries[0].Name)
	}
}

func TestBuildHeroSummariesIntoReusesDestination(t *testing.T) {
	heroes := []Hero{
		{
			ID:    1,
			Name:  "Ragnar",
			Class: HeroClassWarrior,
			Level: 3,
			Alive: true,
		},
	}

	destination := make([]HeroSummary, 0, 4)

	summaries := BuildHeroSummariesInto(heroes, destination)

	if len(summaries) != 1 {
		t.Fatalf("expected summaries length %d, got %d", 1, len(summaries))
	}

	if cap(summaries) != 4 {
		t.Fatalf("expected reused capacity %d, got %d", 4, cap(summaries))
	}
}

func TestCountAliveSummariesReturnsAliveCount(t *testing.T) {
	summaries := []HeroSummary{
		{
			Name:  "Ragnar",
			Alive: true,
		},
		{
			Name:  "Merlin",
			Alive: false,
		},
		{
			Name:  "Luna",
			Alive: true,
		},
	}

	aliveCount := CountAliveSummaries(summaries)

	if aliveCount != 2 {
		t.Fatalf("expected alive count %d, got %d", 2, aliveCount)
	}
}
