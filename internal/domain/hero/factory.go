package hero

import "fmt"

func NewHero(id HeroID, name string, class HeroClass, level int, stats CombatStats) (Hero, error) {
	if err := ValidateHeroName(name); err != nil {
		return Hero{}, fmt.Errorf("create hero: %w", err)
	}

	return Hero{
		ID:          id,
		Name:        name,
		Class:       class,
		Level:       level,
		Alive:       true,
		CombatStats: stats,
		Attacks: [3]string{
			"Slash",
			"Heavy Strike",
			"Shield Bash",
		},
		Inventory: []string{},
		Stats:     map[string]int{},
	}, nil
}
