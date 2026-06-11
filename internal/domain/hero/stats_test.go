package hero

import "testing"

func TestStatReturnsValueWhenStatExists(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
		},
	}

	value, exists := gameHero.Stat("strength")

	if !exists {
		t.Fatal("expected stat to exist")
	}

	if value != 10 {
		t.Fatalf("expected stat value %d, got %d", 10, value)
	}
}

func TestStatReturnsFalseWhenStatDoesNotExist(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
		},
	}

	_, exists := gameHero.Stat("agility")

	if exists {
		t.Fatal("expected stat not to exist")
	}
}

func TestRemoveStatReturnsTrueWhenStatExists(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"temporary_bonus": 2,
		},
	}

	removed := gameHero.RemoveStat("temporary_bonus")

	if !removed {
		t.Fatal("expected stat to be removed")
	}

	if _, exists := gameHero.Stats["temporary_bonus"]; exists {
		t.Fatal("expected stat to be deleted from map")
	}
}

func TestRemoveStatReturnsFalseWhenStatDoesNotExist(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{},
	}

	removed := gameHero.RemoveStat("temporary_bonus")

	if removed {
		t.Fatal("expected stat not to be removed")
	}
}

func TestHighestStatValueReturnsHighestValue(t *testing.T) {
	stats := map[string]int{
		"strength": 10,
		"agility":  7,
		"stamina":  12,
	}

	selection := HighestStatValue(stats)

	if !selection.Found {
		t.Fatal("expected highest stat value to be found")
	}

	if selection.Value != 12 {
		t.Fatalf("expected highest stat value %d, got %d", 12, selection.Value)
	}
}

func TestHighestStatValueReturnsNotFoundForEmptyStats(t *testing.T) {
	selection := HighestStatValue(map[string]int{})

	if selection.Found {
		t.Fatal("expected highest stat value not to be found")
	}

	if selection.Value != 0 {
		t.Fatalf("expected zero stat value, got %d", selection.Value)
	}
}

func TestHeroHighestStatValueReturnsHighestValue(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	selection := gameHero.HighestStatValue()

	if !selection.Found {
		t.Fatal("expected highest hero stat value to be found")
	}

	if selection.Value != 12 {
		t.Fatalf("expected highest hero stat value %d, got %d", 12, selection.Value)
	}
}

func TestCloneStatsCreatesIndependentMap(t *testing.T) {
	stats := map[string]int{
		"strength": 10,
		"agility":  7,
	}

	cloned := CloneStats(stats)

	stats["strength"] = 99

	if cloned["strength"] != 10 {
		t.Fatalf("expected cloned strength %d, got %d", 10, cloned["strength"])
	}
}

func TestNewStatsSnapshotStoresCountAndClonesValues(t *testing.T) {
	stats := map[string]int{
		"strength": 10,
		"agility":  7,
	}

	snapshot := NewStatsSnapshot(stats)

	if snapshot.Count != 2 {
		t.Fatalf("expected snapshot count %d, got %d", 2, snapshot.Count)
	}

	stats["strength"] = 99

	if snapshot.Values["strength"] != 10 {
		t.Fatalf("expected snapshot strength %d, got %d", 10, snapshot.Values["strength"])
	}
}

func TestHeroStatsSnapshotReturnsStatsMetadata(t *testing.T) {
	gameHero := Hero{
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	snapshot := gameHero.StatsSnapshot()

	if snapshot.Count != 3 {
		t.Fatalf("expected snapshot count %d, got %d", 3, snapshot.Count)
	}

	if snapshot.Values["stamina"] != 12 {
		t.Fatalf("expected stamina value %d, got %d", 12, snapshot.Values["stamina"])
	}
}
