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
