package main

import "fmt"

const (
	defaultHeroClass          = "Warrior"
	defaultHeroLevel          = 1
	defaultHeroHP             = 100
	defaultHeroBaseDamage     = 15
	defaultHeroBonusDamage    = 5
	defaultHeroCriticalChance = 0.15

	defaultEnemyName = "Goblin"
	defaultEnemyHP   = 60
)

func calculateDamage(baseDamage int, bonusDamage int) int {
	return baseDamage + bonusDamage
}

func main() {
	fmt.Println("Welcome to DevArena")

	heroName := "Ragnar"
	heroClass := defaultHeroClass
	heroLevel := defaultHeroLevel
	heroHP := defaultHeroHP
	heroBaseDamage := defaultHeroBaseDamage
	heroBonusDamage := defaultHeroBonusDamage
	heroAlive := true
	heroCriticalChance := defaultHeroCriticalChance
	heroAttacks := [3]string{"Slash", "Pierce", "Heavy Strike"}
	heroInventory := []string{"Small Potion", "Wooden Shield"}
	heroStats := map[string]int{
		"strength": 10,
		"agility":  7,
		"stamina":  12,
	}
	heroTotalDamage := calculateDamage(heroBaseDamage, heroBonusDamage)

	enemyName := defaultEnemyName
	enemyHP := defaultEnemyHP

	inventoryBeforeBattle := make([]string, len(heroInventory))
	copy(inventoryBeforeBattle, heroInventory)

	fmt.Println("Hero name:", heroName)
	fmt.Println("Hero class:", heroClass)
	fmt.Println("Hero level:", heroLevel)
	fmt.Println("Hero HP:", heroHP)
	fmt.Println("Hero base damage:", heroBaseDamage)
	fmt.Println("Hero bonus damage:", heroBonusDamage)
	fmt.Println("Hero total damage:", heroTotalDamage)
	fmt.Println("Hero alive:", heroAlive)
	fmt.Println("Hero critical chance:", heroCriticalChance)

	fmt.Println("Hero stats:")
	for statName, statValue := range heroStats {
		fmt.Println(statName+":", statValue)
	}

	intellect, exists := heroStats["intellect"]
	if exists {
		fmt.Println("Intellect:", intellect)
	} else {
		fmt.Println("Intellect stat is not defined")
	}

	heroStats["strength"] = heroStats["strength"] + 2
	heroStats["intellect"] = 3

	fmt.Println("Hero stats after training:")
	for statName, statValue := range heroStats {
		fmt.Println(statName+":", statValue)
	}

	delete(heroStats, "intellect")

	fmt.Println("Hero stats after removing temporary intellect bonus:")
	for statName, statValue := range heroStats {
		fmt.Println(statName+":", statValue)
	}

	fmt.Println("Hero attacks:")
	for index, attack := range heroAttacks {
		fmt.Println("Attack", index+1, ":", attack)
	}

	fmt.Println("Hero inventory before battle:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length before reward:", len(heroInventory))
	fmt.Println("Inventory capacity before reward:", cap(heroInventory))

	heroInventory = append(heroInventory, "Rusty Sword")

	fmt.Println("Hero inventory after reward:")
	for index, item := range heroInventory {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory copy before battle still:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length after reward:", len(heroInventory))
	fmt.Println("Inventory capacity after reward:", cap(heroInventory))

	if heroHP > 0 {
		fmt.Println(heroName, "is ready to fight")
	} else {
		fmt.Println(heroName, "is defeated and cannot fight")
	}

	fmt.Println("Enemy name:", enemyName)
	fmt.Println("Enemy HP:", enemyHP)

	fmt.Println("Battle started")

	for round := 1; enemyHP > 0; round++ {
		attackIndex := (round - 1) % len(heroAttacks)
		selectedAttack := heroAttacks[attackIndex]

		enemyHP -= heroTotalDamage

		if enemyHP < 0 {
			enemyHP = 0
		}

		fmt.Println("Round:", round)
		fmt.Println(heroName, "uses", selectedAttack)
		fmt.Println(heroName, "hits", enemyName, "for", heroTotalDamage, "damage")
		fmt.Println(enemyName, "HP:", enemyHP)
	}

	fmt.Println("Battle finished")
	fmt.Println(enemyName, "is defeated")
	fmt.Println(heroName, "received item:", "Rusty Sword")
}
