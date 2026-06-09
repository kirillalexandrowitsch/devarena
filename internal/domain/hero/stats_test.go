package hero

import "testing"

func TestHeroStatReturnsExistingStat(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
		},
	}

	value, exists := gameHero.Stat("strength")

	if !exists {
		t.Fatal("expected strength stat to exist")
	}

	if value != 10 {
		t.Fatalf("expected strength stat %d, got %d", 10, value)
	}
}

func TestHeroStatReturnsFalseForMissingStat(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
		},
	}

	value, exists := gameHero.Stat("intellect")

	if exists {
		t.Fatal("expected intellect stat not to exist")
	}

	if value != 0 {
		t.Fatalf("expected missing stat value %d, got %d", 0, value)
	}
}

func TestHeroStatReturnsFalseForNilStats(t *testing.T) {
	gameHero := Hero{}

	value, exists := gameHero.Stat("strength")

	if exists {
		t.Fatal("expected stat not to exist for nil stats map")
	}

	if value != 0 {
		t.Fatalf("expected missing stat value %d, got %d", 0, value)
	}
}

func TestHeroRemoveStatDeletesExistingStat(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength":        10,
			"temporary_bonus": 3,
		},
	}

	removed := gameHero.RemoveStat("temporary_bonus")

	if !removed {
		t.Fatal("expected temporary_bonus stat to be removed")
	}

	_, exists := gameHero.Stats["temporary_bonus"]
	if exists {
		t.Fatal("expected temporary_bonus stat to be missing after removal")
	}

	if gameHero.Stats["strength"] != 10 {
		t.Fatalf("expected strength stat to remain %d, got %d", 10, gameHero.Stats["strength"])
	}
}

func TestHeroRemoveStatReturnsFalseForMissingStat(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
		},
	}

	removed := gameHero.RemoveStat("temporary_bonus")

	if removed {
		t.Fatal("expected missing stat removal to return false")
	}
}

func TestHeroRemoveStatReturnsFalseForNilStats(t *testing.T) {
	gameHero := Hero{}

	removed := gameHero.RemoveStat("temporary_bonus")

	if removed {
		t.Fatal("expected nil stats removal to return false")
	}
}
