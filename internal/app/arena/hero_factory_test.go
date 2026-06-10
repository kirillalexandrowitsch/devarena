package arena

import "testing"

func TestMustCreateDefaultHeroCreatesValidHero(t *testing.T) {
	gameHero := mustCreateDefaultHero()

	if gameHero.Name != defaultHeroName {
		t.Fatalf("expected hero name %q, got %q", defaultHeroName, gameHero.Name)
	}

	if gameHero.Class != defaultHeroClass {
		t.Fatalf("expected hero class %q, got %q", defaultHeroClass, gameHero.Class)
	}

	if gameHero.Level != defaultHeroLevel {
		t.Fatalf("expected hero level %d, got %d", defaultHeroLevel, gameHero.Level)
	}

	if !gameHero.Alive {
		t.Fatal("expected hero to be alive")
	}
}
