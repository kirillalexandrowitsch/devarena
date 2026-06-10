package hero

import "github.com/rudyakovk/devarena/internal/domain/shared"

func (h Hero) Stat(name string) (int, bool) {
	value, exists := h.Stats[name]
	return value, exists
}

func (h *Hero) RemoveStat(name string) bool {
	if _, exists := h.Stats[name]; !exists {
		return false
	}

	delete(h.Stats, name)
	return true
}

func HighestStatValue(stats map[string]int) shared.Selection[int] {
	if len(stats) == 0 {
		return shared.Selection[int]{
			Found: false,
		}
	}

	values := make([]int, 0, len(stats))

	for _, value := range stats {
		values = append(values, value)
	}

	return shared.Max(values)
}

func (h Hero) HighestStatValue() shared.Selection[int] {
	return HighestStatValue(h.Stats)
}
