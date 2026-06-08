package hero

func (h Hero) Stat(name string) (int, bool) {
	value, exists := h.Stats[name]

	return value, exists
}

func (h *Hero) RemoveStat(name string) bool {
	if h.Stats == nil {
		return false
	}

	_, exists := h.Stats[name]
	if !exists {
		return false
	}

	delete(h.Stats, name)

	return true
}
