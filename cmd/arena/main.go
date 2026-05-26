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

type Weapon interface {
	Name() string
	DamageBonus() int
}

type Sword struct {
	Title string
	Bonus int
}

func (s Sword) Name() string {
	return s.Title
}

func (s Sword) DamageBonus() int {
	return s.Bonus
}

type Axe struct {
	Title string
	Bonus int
}

func (a Axe) Name() string {
	return a.Title
}

func (a Axe) DamageBonus() int {
	return a.Bonus
}

type Hero struct {
	Name           string
	Class          string
	Level          int
	HP             int
	BaseDamage     int
	BonusDamage    int
	Alive          bool
	CriticalChance float64
	Attacks        [3]string
	Inventory      []string
	Stats          map[string]int
	Weapon         Weapon
}

type Enemy struct {
	Name string
	HP   int
}

type Battle struct {
	Hero  *Hero
	Enemy *Enemy
	Round int
}

func (h Hero) TotalDamage() int {
	totalDamage := h.BaseDamage + h.BonusDamage

	if h.Weapon != nil {
		totalDamage += h.Weapon.DamageBonus()
	}

	return totalDamage
}

func (h Hero) PrintInfo() {
	fmt.Println("Hero name:", h.Name)
	fmt.Println("Hero class:", h.Class)
	fmt.Println("Hero level:", h.Level)
	fmt.Println("Hero HP:", h.HP)
	fmt.Println("Hero base damage:", h.BaseDamage)
	fmt.Println("Hero bonus damage:", h.BonusDamage)

	if h.Weapon != nil {
		fmt.Println("Hero weapon:", h.Weapon.Name())
		fmt.Println("Weapon damage bonus:", h.Weapon.DamageBonus())
	} else {
		fmt.Println("Hero weapon: none")
	}

	fmt.Println("Hero total damage:", h.TotalDamage())
	fmt.Println("Hero alive:", h.Alive)
	fmt.Println("Hero critical chance:", h.CriticalChance)
}

func (h Hero) PrintStats() {
	fmt.Println("Hero stats:")
	for statName, statValue := range h.Stats {
		fmt.Println(statName+":", statValue)
	}
}

func (h Hero) PrintAttacks() {
	fmt.Println("Hero attacks:")
	for index, attack := range h.Attacks {
		fmt.Println("Attack", index+1, ":", attack)
	}
}

func (h Hero) PrintInventory(title string) {
	fmt.Println(title)
	for index, item := range h.Inventory {
		fmt.Println("Item", index+1, ":", item)
	}
}

func (h *Hero) AddItem(item string) {
	h.Inventory = append(h.Inventory, item)
}

func (h *Hero) EquipWeapon(weapon Weapon) {
	h.Weapon = weapon
}

func (h *Hero) TakeDamage(damage int) {
	h.HP -= damage

	if h.HP <= 0 {
		h.HP = 0
		h.Alive = false
	}
}

func (e *Enemy) TakeDamage(damage int) {
	e.HP -= damage

	if e.HP < 0 {
		e.HP = 0
	}
}

func (b *Battle) Start() {
	fmt.Println("Battle started")

	for b.Enemy.HP > 0 {
		b.Round++

		attackIndex := (b.Round - 1) % len(b.Hero.Attacks)
		selectedAttack := b.Hero.Attacks[attackIndex]

		b.Enemy.TakeDamage(b.Hero.TotalDamage())

		fmt.Println("Round:", b.Round)
		fmt.Println(b.Hero.Name, "uses", selectedAttack)

		if b.Hero.Weapon != nil {
			fmt.Println(b.Hero.Name, "attacks with", b.Hero.Weapon.Name())
		}

		fmt.Println(b.Hero.Name, "hits", b.Enemy.Name, "for", b.Hero.TotalDamage(), "damage")
		fmt.Println(b.Enemy.Name, "HP:", b.Enemy.HP)
	}

	fmt.Println("Battle finished")
	fmt.Println(b.Enemy.Name, "is defeated")
}

func damageHeroCopy(hero Hero, damage int) {
	hero.HP -= damage
	fmt.Println("Inside damageHeroCopy HP:", hero.HP)
}

func damageHeroOriginal(hero *Hero, damage int) {
	hero.HP -= damage

	if hero.HP <= 0 {
		hero.HP = 0
		hero.Alive = false
	}

	fmt.Println("Inside damageHeroOriginal HP:", hero.HP)
}

func printWeaponInfo(weapon Weapon) {
	fmt.Println("Weapon name:", weapon.Name())
	fmt.Println("Weapon damage bonus:", weapon.DamageBonus())
}

func main() {
	fmt.Println("Welcome to DevArena")

	hero := Hero{
		Name:           "Ragnar",
		Class:          defaultHeroClass,
		Level:          defaultHeroLevel,
		HP:             defaultHeroHP,
		BaseDamage:     defaultHeroBaseDamage,
		BonusDamage:    defaultHeroBonusDamage,
		Alive:          true,
		CriticalChance: defaultHeroCriticalChance,
		Attacks:        [3]string{"Slash", "Pierce", "Heavy Strike"},
		Inventory:      []string{"Small Potion", "Wooden Shield"},
		Stats: map[string]int{
			"strength": 10,
			"agility":  7,
			"stamina":  12,
		},
	}

	starterSword := Sword{
		Title: "Starter Sword",
		Bonus: 4,
	}

	battleAxe := Axe{
		Title: "Battle Axe",
		Bonus: 7,
	}

	fmt.Println("Available weapons:")
	printWeaponInfo(starterSword)
	printWeaponInfo(battleAxe)

	hero.EquipWeapon(starterSword)

	enemy := Enemy{
		Name: defaultEnemyName,
		HP:   defaultEnemyHP,
	}

	battle := Battle{
		Hero:  &hero,
		Enemy: &enemy,
		Round: 0,
	}

	inventoryBeforeBattle := make([]string, len(hero.Inventory))
	copy(inventoryBeforeBattle, hero.Inventory)

	hero.PrintInfo()
	hero.PrintStats()

	intellect, exists := hero.Stats["intellect"]
	if exists {
		fmt.Println("Intellect:", intellect)
	} else {
		fmt.Println("Intellect stat is not defined")
	}

	hero.Stats["strength"] = hero.Stats["strength"] + 2
	hero.Stats["intellect"] = 3

	fmt.Println("Hero stats after training:")
	hero.PrintStats()

	delete(hero.Stats, "intellect")

	fmt.Println("Hero stats after removing temporary intellect bonus:")
	hero.PrintStats()

	hero.PrintAttacks()

	fmt.Println("Hero inventory before battle:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length before reward:", len(hero.Inventory))
	fmt.Println("Inventory capacity before reward:", cap(hero.Inventory))

	hero.AddItem("Rusty Sword")

	hero.PrintInventory("Hero inventory after reward:")

	fmt.Println("Inventory copy before battle still:")
	for index, item := range inventoryBeforeBattle {
		fmt.Println("Item", index+1, ":", item)
	}

	fmt.Println("Inventory length after reward:", len(hero.Inventory))
	fmt.Println("Inventory capacity after reward:", cap(hero.Inventory))

	if hero.HP > 0 {
		fmt.Println(hero.Name, "is ready to fight")
	} else {
		fmt.Println(hero.Name, "is defeated and cannot fight")
	}

	fmt.Println("Pointer demo:")
	fmt.Println("Hero HP before damageHeroCopy:", hero.HP)
	damageHeroCopy(hero, 10)
	fmt.Println("Hero HP after damageHeroCopy:", hero.HP)

	fmt.Println("Hero HP before damageHeroOriginal:", hero.HP)
	damageHeroOriginal(&hero, 10)
	fmt.Println("Hero HP after damageHeroOriginal:", hero.HP)

	fmt.Println("Enemy name:", enemy.Name)
	fmt.Println("Enemy HP:", enemy.HP)

	battle.Start()

	fmt.Println(hero.Name, "received item:", "Rusty Sword")
	fmt.Println("Final hero HP:", hero.HP)
	fmt.Println("Final enemy HP:", enemy.HP)
}
