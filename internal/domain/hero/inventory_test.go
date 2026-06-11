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

func TestNewInventorySnapshotStoresLengthAndCapacity(t *testing.T) {
	items := make([]string, 0, 5)
	items = append(items, "Small Potion", "Wooden Shield")

	snapshot := NewInventorySnapshot(items)

	if snapshot.Length != 2 {
		t.Fatalf("expected snapshot length %d, got %d", 2, snapshot.Length)
	}

	if snapshot.Capacity != 5 {
		t.Fatalf("expected snapshot capacity %d, got %d", 5, snapshot.Capacity)
	}
}

func TestNewInventorySnapshotClonesItems(t *testing.T) {
	items := []string{"Small Potion", "Wooden Shield"}

	snapshot := NewInventorySnapshot(items)

	items[0] = "Broken Potion"

	if snapshot.Items[0] != "Small Potion" {
		t.Fatalf("expected snapshot item %q, got %q", "Small Potion", snapshot.Items[0])
	}
}

func TestHeroInventorySnapshotReturnsInventoryMetadata(t *testing.T) {
	gameHero := Hero{
		Inventory: make([]string, 0, 4),
	}

	gameHero.Inventory = append(gameHero.Inventory, "Small Potion")

	snapshot := gameHero.InventorySnapshot()

	if snapshot.Length != 1 {
		t.Fatalf("expected snapshot length %d, got %d", 1, snapshot.Length)
	}

	if snapshot.Capacity != 4 {
		t.Fatalf("expected snapshot capacity %d, got %d", 4, snapshot.Capacity)
	}
}
