package hero

import "github.com/rudyakovk/devarena/internal/domain/shared"

type InventorySnapshot struct {
	Items    []string
	Length   int
	Capacity int
}

func EnsureInventoryCapacity(items []string, requiredCapacity int) []string {
	if cap(items) >= requiredCapacity {
		return items
	}

	expanded := make([]string, len(items), requiredCapacity)
	copy(expanded, items)

	return expanded
}

func CloneInventory(items []string) []string {
	cloned := make([]string, len(items))
	copy(cloned, items)

	return cloned
}

func HasInventoryItem(items []string, item string) bool {
	return shared.Contains(items, item)
}

func FirstInventoryItem(items []string) shared.Selection[string] {
	return shared.First(items)
}

func NewInventorySnapshot(items []string) InventorySnapshot {
	return InventorySnapshot{
		Items:    CloneInventory(items),
		Length:   len(items),
		Capacity: cap(items),
	}
}
