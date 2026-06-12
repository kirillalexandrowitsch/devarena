package arena

import "github.com/rudyakovk/devarena/internal/domain/hero"

func mustCreateDefaultHero() hero.Hero {
	gameHero, err := hero.NewHero(
		1,
		defaultHeroName,
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
		panic(err)
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

	return gameHero
}
