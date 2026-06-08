package hero

func (h *Hero) EnsureInventoryCapacity(additionalItems int) {
	if additionalItems <= 0 {
		return
	}

	requiredCapacity := len(h.Inventory) + additionalItems
	if cap(h.Inventory) >= requiredCapacity {
		return
	}

	expandedInventory := make([]string, 0, requiredCapacity)
	expandedInventory = append(expandedInventory, h.Inventory...)

	h.Inventory = expandedInventory
}

func (h Hero) CloneInventory() []string {
	if h.Inventory == nil {
		return nil
	}

	clonedInventory := make([]string, len(h.Inventory))
	copy(clonedInventory, h.Inventory)

	return clonedInventory
}
