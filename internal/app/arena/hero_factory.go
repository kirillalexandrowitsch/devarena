package arena

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/domain/hero"
)

func createDefaultHero() (hero.Hero, error) {
	return createDefaultHeroWithName(defaultHeroName)
}

func createDefaultHeroWithName(name string) (hero.Hero, error) {
	gameHero, err := hero.NewHero(
		1,
		name,
		defaultHeroClass,
		defaultHeroLevel,
		hero.CombatStats{
			HP:             defaultHeroHP,
			BaseDamage:     defaultHeroBaseDamage,
			BonusDamage:    defaultHeroBonusDamage,
			CriticalChance: defaultHeroCriticalChance,
		},
	)
	if err != nil {
		return hero.Hero{}, fmt.Errorf("create default hero: %w", err)
	}

	gameHero.Attacks = [3]string{
		"Slash",
		"Heavy Strike",
		"Shield Bash",
	}

	gameHero.Inventory = []string{
		"Small Potion",
		"Training Sword",
	}

	gameHero.Stats = map[string]int{
		"strength":        10,
		"agility":         7,
		"stamina":         12,
		"temporary_bonus": 2,
	}

	return gameHero, nil
}

func mustCreateDefaultHero() hero.Hero {
	gameHero, err := createDefaultHero()
	if err != nil {
		panic(describeHeroCreationError(err))
	}

	return gameHero
}
