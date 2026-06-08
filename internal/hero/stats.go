package hero

func (h Hero) Stat(name string) (int, bool) {
	value, exists := h.Stats[name]

	return value, exists
}
