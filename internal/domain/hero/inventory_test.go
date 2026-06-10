package hero

import "testing"

func TestHasInventoryItemReturnsTrueWhenItemExists(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield"}

	if !HasInventoryItem(items, "Small Potion") {
		t.Fatal("expected inventory item to exist")
	}
}

func TestHasInventoryItemReturnsFalseWhenItemDoesNotExist(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield"}

	if HasInventoryItem(items, "Magic Ring") {
		t.Fatal("expected inventory item not to exist")
	}
}

func TestHeroHasItemReturnsTrueWhenItemExists(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion", "Wooden Shield"},
	}

	if !gameHero.HasItem("Wooden Shield") {
		t.Fatal("expected hero to have item")
	}
}

func TestHeroHasItemReturnsFalseWhenItemDoesNotExist(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion", "Wooden Shield"},
	}

	if gameHero.HasItem("Magic Ring") {
		t.Fatal("expected hero not to have item")
	}
}

func TestFirstInventoryItemReturnsFirstItemWhenInventoryIsNotEmpty(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield"}

	selection := FirstInventoryItem(items)

	if !selection.Found {
		t.Fatal("expected first inventory item to be found")
	}

	if selection.Value != "Small Potion" {
		t.Fatalf("expected first inventory item %q, got %q", "Small Potion", selection.Value)
	}
}

func TestFirstInventoryItemReturnsNotFoundWhenInventoryIsEmpty(t *testing.T) {
	selection := FirstInventoryItem([]string{})

	if selection.Found {
		t.Fatal("expected first inventory item not to be found")
	}
}

func TestHeroFirstInventoryItemReturnsFirstItem(t *testing.T) {
	gameHero := Hero{
		Inventory: []string{"Small Potion", "Wooden Shield"},
	}

	selection := gameHero.FirstInventoryItem()

	if !selection.Found {
		t.Fatal("expected first hero inventory item to be found")
	}

	if selection.Value != "Small Potion" {
		t.Fatalf("expected first hero inventory item %q, got %q", "Small Potion", selection.Value)
	}
}
