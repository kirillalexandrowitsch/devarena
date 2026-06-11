package hero

import "github.com/rudyakovk/devarena/internal/domain/shared"

type StatsSnapshot struct {
	Values map[string]int
	Count  int
}

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

func CloneStats(stats map[string]int) map[string]int {
	cloned := make(map[string]int, len(stats))

	for name, value := range stats {
		cloned[name] = value
	}

	return cloned
}

func NewStatsSnapshot(stats map[string]int) StatsSnapshot {
	return StatsSnapshot{
		Values: CloneStats(stats),
		Count:  len(stats),
	}
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

func (h Hero) StatsSnapshot() StatsSnapshot {
	return NewStatsSnapshot(h.Stats)
}
