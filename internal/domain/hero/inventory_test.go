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
