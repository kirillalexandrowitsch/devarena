package main

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/battle"
	"github.com/rudyakovk/devarena/internal/enemy"
	"github.com/rudyakovk/devarena/internal/hero"
)

const (
	defaultHeroName           = "Ragnar"
	defaultHeroClass          = hero.HeroClassWarrior
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

func main() {
	fmt.Println("Welcome to DevArena")

	if !hero.IsValidHeroName(defaultHeroName) {
		fmt.Println("Invalid default hero name")
	}

	gameHero := hero.Hero{
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

	starterSword := hero.Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	gameHero.EquipWeapon(starterSword)

	strength, strengthExists := gameHero.Stat("strength")
	if strengthExists {
		gameHero.BonusDamage += strength / 5
	}

	temporaryBonus, temporaryBonusExists := gameHero.Stat("temporary_bonus")
	if temporaryBonusExists {
		gameHero.BonusDamage += temporaryBonus
		gameHero.RemoveStat("temporary_bonus")
	}

	fmt.Println("Hero class description:", hero.DescribeHeroClass(gameHero.Class))

	gameHero.PrintInfo()
	gameHero.PrintStats()
	gameHero.PrintAttacks()
	gameHero.PrintInventory("Hero inventory before battle:")

	gameEnemy := enemy.Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	arenaBattle := battle.Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
		Round: 0,
	}

	arenaBattle.Start()

	winnerName, defeatedName := arenaBattle.ResultNames()

	battleStatistics := battle.Statistics{
		Wins:   1,
		Losses: 0,
	}

	reward := battle.CalculateReward(gameHero.Level, defeatedName)
	gameHero.AddItem(reward.Item)

	rewardEvent := battle.NewRewardEvent(gameHero.Name, reward)

	fmt.Println(gameHero.Name, "received experience:", reward.Experience)
	fmt.Println(gameHero.Name, "received item:", reward.Item)

	fmt.Println("Reward event type:", rewardEvent.Type)
	fmt.Println("Reward event payload size:", rewardEvent.Payload.Size())

	fmt.Println("Battle winner:", winnerName)
	fmt.Println("Defeated opponent:", defeatedName)

	fmt.Println("Battle win rate:", battleStatistics.WinRatePercent())

	gameHero.PrintInventory("Hero inventory after battle:")

	fmt.Println("Final hero HP:", gameHero.HP)
	fmt.Println("Final enemy HP:", gameEnemy.HP)
}
