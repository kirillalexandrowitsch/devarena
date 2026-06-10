package arena

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/domain/hero"
)

func mustCreateDefaultHero() hero.Hero {
	if !hero.IsValidHeroName(defaultHeroName) {
		panic(fmt.Sprintf("invalid default hero name: %q", defaultHeroName))
	}

	return hero.Hero{
		ID:    1,
		Name:  defaultHeroName,
		Class: defaultHeroClass,
		Level: defaultHeroLevel,
		Alive: true,
		CombatStats: hero.CombatStats{
			HP:             defaultHeroHP,
			BaseDamage:     defaultHeroBaseDamage,
			BonusDamage:    defaultHeroBonusDamage,
			CriticalChance: defaultHeroCriticalChance,
		},
		Attacks:   [3]string{"Slash", "Pierce", "Heavy Strike"},
		Inventory: []string{"Small Potion", "Wooden Shield"},
		Stats: map[string]int{
			"strength":        10,
			"agility":         7,
			"stamina":         12,
			"temporary_bonus": 2,
		},
	}
}
