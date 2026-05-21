package main

import "fmt"

func calculateDamage(baseDamage int, bonusDamage int) int {
	return baseDamage + bonusDamage
}

func main() {
	fmt.Println("Welcome to DevArena")

	heroName := "Ragnar"
	heroClass := "Warrior"
	heroLevel := 1
	heroHP := 100
	heroBaseDamage := 15
	heroBonusDamage := 5
	heroAlive := true
	heroCriticalChance := 0.15

	heroTotalDamage := calculateDamage(heroBaseDamage, heroBonusDamage)

	fmt.Println("Hero name:", heroName)
	fmt.Println("Hero class:", heroClass)
	fmt.Println("Hero level:", heroLevel)
	fmt.Println("Hero HP:", heroHP)
	fmt.Println("Hero base damage:", heroBaseDamage)
	fmt.Println("Hero bonus damage:", heroBonusDamage)
	fmt.Println("Hero total damage:", heroTotalDamage)
	fmt.Println("Hero alive:", heroAlive)
	fmt.Println("Hero critical chance:", heroCriticalChance)
}
