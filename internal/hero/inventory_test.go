package hero

import "testing"

func TestHeroEnsureInventoryCapacityExpandsCapacity(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion"},
	}

	gameHero.EnsureInventoryCapacity(3)

	wantLength := 1
	if len(gameHero.Inventory) != wantLength {
		t.Fatalf("expected inventory length %d, got %d", wantLength, len(gameHero.Inventory))
	}

	wantMinCapacity := 4
	if cap(gameHero.Inventory) < wantMinCapacity {
		t.Fatalf("expected inventory capacity at least %d, got %d", wantMinCapacity, cap(gameHero.Inventory))
	}

	if gameHero.Inventory[0] != "Small Potion" {
		t.Fatalf("expected existing inventory item to be preserved, got %q", gameHero.Inventory[0])
	}
}

func TestHeroEnsureInventoryCapacityIgnoresNonPositiveAdditionalItems(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion"},
	}

	initialCapacity := cap(gameHero.Inventory)

	gameHero.EnsureInventoryCapacity(0)

	if cap(gameHero.Inventory) != initialCapacity {
		t.Fatalf("expected inventory capacity %d, got %d", initialCapacity, cap(gameHero.Inventory))
	}
}

func TestHeroAddItemsUsesInventoryCapacity(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion"},
	}

	gameHero.AddItems("Wooden Shield", "Rusty Sword")

	wantLength := 3
	if len(gameHero.Inventory) != wantLength {
		t.Fatalf("expected inventory length %d, got %d", wantLength, len(gameHero.Inventory))
	}

	if cap(gameHero.Inventory) < wantLength {
		t.Fatalf("expected inventory capacity at least %d, got %d", wantLength, cap(gameHero.Inventory))
	}

	if gameHero.Inventory[1] != "Wooden Shield" {
		t.Fatalf("expected second inventory item %q, got %q", "Wooden Shield", gameHero.Inventory[1])
	}

	if gameHero.Inventory[2] != "Rusty Sword" {
		t.Fatalf("expected third inventory item %q, got %q", "Rusty Sword", gameHero.Inventory[2])
	}
}
