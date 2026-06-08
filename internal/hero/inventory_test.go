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

func TestHeroCloneInventoryReturnsIndependentCopy(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion", "Wooden Shield"},
	}

	clonedInventory := gameHero.CloneInventory()
	clonedInventory[0] = "Large Potion"

	if gameHero.Inventory[0] != "Small Potion" {
		t.Fatalf("expected original inventory item to remain %q, got %q", "Small Potion", gameHero.Inventory[0])
	}

	if clonedInventory[0] != "Large Potion" {
		t.Fatalf("expected cloned inventory item to be %q, got %q", "Large Potion", clonedInventory[0])
	}
}

func TestHeroCloneInventoryPreservesLength(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion", "Wooden Shield", "Rusty Sword"},
	}

	clonedInventory := gameHero.CloneInventory()

	if len(clonedInventory) != len(gameHero.Inventory) {
		t.Fatalf("expected cloned inventory length %d, got %d", len(gameHero.Inventory), len(clonedInventory))
	}
}

func TestHeroCloneInventoryReturnsNilForNilInventory(t *testing.T) {
	gameHero := Hero{}

	clonedInventory := gameHero.CloneInventory()

	if clonedInventory != nil {
		t.Fatalf("expected nil cloned inventory, got %#v", clonedInventory)
	}
}
