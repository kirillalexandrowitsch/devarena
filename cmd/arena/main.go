package main

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/battle"
	"github.com/rudyakovk/devarena/internal/enemy"
	"github.com/rudyakovk/devarena/internal/hero"
)

const (
	defaultHeroClass          = hero.HeroClassWarrior
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

func damageHeroCopy(hero hero.Hero, damage int) {
	hero.HP -= damage
	fmt.Println("Inside damageHeroCopy HP:", hero.HP)
}

func damageHeroOriginal(hero *hero.Hero, damage int) {
	hero.HP -= damage

	if hero.HP <= 0 {
		hero.HP = 0
		hero.Alive = false
	}

	fmt.Println("Inside damageHeroOriginal HP:", hero.HP)
}

func printWeaponInfo(weapon hero.Weapon) {
	fmt.Println("Weapon name:", weapon.Name())
	fmt.Println("Weapon damage bonus:", weapon.DamageBonus())
}

func printCalculatedDamage(calculator hero.DamageCalculator) {
	fmt.Println("Calculated damage from interface:", calculator.TotalDamage())
}

func addMethodSetReward(manager hero.InventoryManager) {
	manager.AddItem("Method Set Badge")
}

func printZeroValues() {
	var defaultName string
	var defaultLevel int
	var defaultCriticalChance float64
	var defaultAlive bool
	var defaultHeroPointer *hero.Hero
	var defaultInventory []string
	var defaultStats map[string]int
	var defaultHero hero.Hero

	fmt.Println("Zero values demo:")
	fmt.Println("Default string:", defaultName)
	fmt.Println("Default int:", defaultLevel)
	fmt.Println("Default float64:", defaultCriticalChance)
	fmt.Println("Default bool:", defaultAlive)
	fmt.Println("Default hero pointer:", defaultHeroPointer)
	fmt.Println("Default inventory slice:", defaultInventory)
	fmt.Println("Default stats map:", defaultStats)
	fmt.Println("Default hero:", defaultHero)
}

func printTypeConversions() {
	wins := 7
	losses := 3
	totalBattles := wins + losses

	winRate := float64(wins) / float64(totalBattles) * 100
	totalBattles64 := int64(totalBattles)
	winRateText := fmt.Sprintf("%.2f%%", winRate)

	fmt.Println("Type conversion demo:")
	fmt.Println("Wins:", wins)
	fmt.Println("Losses:", losses)
	fmt.Println("Total battles as int:", totalBattles)
	fmt.Println("Total battles as int64:", totalBattles64)
	fmt.Println("Win rate as float64:", winRate)
	fmt.Println("Win rate text:", winRateText)
}

func selectRewardItem(candidates []string) string {
	selectedReward := "Rusty Sword"

	for _, item := range candidates {
		if item == "" {
			continue
		}

		selectedReward = item
		break
	}

	return selectedReward
}

func calculateBattleReward(heroLevel int, defeatedEnemyName string) (rewardExperience int, rewardItem string) {
	baseExperience := 25
	levelBonus := heroLevel * 5

	rewardExperience = baseExperience + levelBonus

	rewardCandidates := []string{
		"",
		"Rusty Sword",
	}

	if defeatedEnemyName == "Goblin" {
		rewardCandidates = []string{
			"",
			"Goblin Dagger",
			"Rusty Sword",
		}
	}

	rewardItem = selectRewardItem(rewardCandidates)

	return
}

func main() {
	fmt.Println("Welcome to DevArena")

	printZeroValues()

	printTypeConversions()

	gameHero := hero.Hero{
		ID:    1,
		Name:  "Ragnar",
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
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	fmt.Println("Hero class description:", hero.DescribeHeroClass(gameHero.Class))

	starterSword := hero.Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	battleAxe := hero.Axe{
		Title: "Battle Axe",
		Bonus: 7,
	}

	fmt.Println("Available weapons:")
	printWeaponInfo(starterSword)
	printWeaponInfo(battleAxe)

	gameHero.EquipWeapon(starterSword)

	trainingSessions := 0

	calculateTrainingBonus := func(strength int) int {
		trainingSessions++

		return strength/2 + trainingSessions
	}

	firstTrainingBonus := calculateTrainingBonus(gameHero.Stats["strength"])
	secondTrainingBonus := calculateTrainingBonus(gameHero.Stats["strength"])
	trainingBonusDamage := firstTrainingBonus + secondTrainingBonus

	gameHero.BonusDamage += trainingBonusDamage

	fmt.Println("Training sessions:", trainingSessions)
	fmt.Println("First training bonus damage:", firstTrainingBonus)
	fmt.Println("Second training bonus damage:", secondTrainingBonus)
	fmt.Println("Total training bonus damage:", trainingBonusDamage)

	gameEnemy := enemy.Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	gameBattle := battle.Battle{
		Hero:  &gameHero,
		Enemy: &gameEnemy,
		Round: 0,
	}

	inventoryBeforeBattle := make([]string, len(gameHero.Inventory))
	copy(inventoryBeforeBattle, gameHero.Inventory)

	gameHero.PrintInfo()
	gameHero.PrintStats()

	fmt.Println("Method set demo:")
	printCalculatedDamage(gameHero)
	printCalculatedDamage(&gameHero)
	addMethodSetReward(&gameHero)

	intellect, exists := gameHero.Stats["intellect"]
	if exists {
		fmt.Println("Intellect:", intellect)
	} else {
		fmt.Println("Intellect stat is not defined")
	}

	gameHero.Stats["strength"] = gameHero.Stats["strength"] + 2
	gameHero.Stats["intellect"] = 3

	fmt.Println("Hero stats after training:")
	gameHero.PrintStats()

	delete(gameHero.Stats, "intellect")

	fmt.Println("Hero stats after removing temporary intellect bonus:")
	gameHero.PrintStats()

	gameHero.PrintAttacks()

	fmt.Println("Hero inventory before battle:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length before reward:", len(gameHero.Inventory))
	fmt.Println("Inventory capacity before reward:", cap(gameHero.Inventory))

	gameHero.AddItems("Rusty Sword", "Training Token")

	gameHero.PrintInventory("Hero inventory after reward:")

	fmt.Println("Inventory copy before battle still:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length after reward:", len(gameHero.Inventory))
	fmt.Println("Inventory capacity after reward:", cap(gameHero.Inventory))

	if gameHero.HP > 0 {
		fmt.Println(gameHero.Name, "is ready to fight")
	} else {
		fmt.Println(gameHero.Name, "is defeated and cannot fight")
	}

	fmt.Println("Pointer demo:")
	fmt.Println("Hero HP before damageHeroCopy:", gameHero.HP)
	damageHeroCopy(gameHero, 10)
	fmt.Println("Hero HP after damageHeroCopy:", gameHero.HP)

	fmt.Println("Hero HP before damageHeroOriginal:", gameHero.HP)
	damageHeroOriginal(&gameHero, 10)
	fmt.Println("Hero HP after damageHeroOriginal:", gameHero.HP)

	fmt.Println("Enemy name:", gameEnemy.Name)
	fmt.Println("Enemy HP:", gameEnemy.HP)

	gameBattle.Start()

	rewardExperience, rewardItem := calculateBattleReward(gameHero.Level, gameEnemy.Name)

	gameHero.AddItems(rewardItem)

	fmt.Println(gameHero.Name, "received experience:", rewardExperience)
	fmt.Println(gameHero.Name, "received item:", rewardItem)
	fmt.Println("Final hero HP:", gameHero.HP)
	fmt.Println("Final enemy HP:", gameEnemy.HP)
}
