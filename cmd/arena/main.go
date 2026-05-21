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
	heroTotalDamage := calculateDamage(heroBaseDamage, heroBonusDamage)

	enemyName := defaultEnemyName
	enemyHP := defaultEnemyHP

	fmt.Println("Hero name:", heroName)
	fmt.Println("Hero class:", heroClass)
	fmt.Println("Hero level:", heroLevel)
	fmt.Println("Hero HP:", heroHP)
	fmt.Println("Hero base damage:", heroBaseDamage)
	fmt.Println("Hero bonus damage:", heroBonusDamage)
	fmt.Println("Hero total damage:", heroTotalDamage)
	fmt.Println("Hero alive:", heroAlive)
	fmt.Println("Hero critical chance:", heroCriticalChance)

	fmt.Println("Hero attacks:")
	fmt.Println("Attack 1:", heroAttacks[0])
	fmt.Println("Attack 2:", heroAttacks[1])
	fmt.Println("Attack 3:", heroAttacks[2])

	fmt.Println("Hero inventory before reward:", heroInventory)
	fmt.Println("Inventory length before reward:", len(heroInventory))
	fmt.Println("Inventory capacity before reward:", cap(heroInventory))

	heroInventory = append(heroInventory, "Rusty Sword")

	fmt.Println("Hero inventory after reward:", heroInventory)
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
