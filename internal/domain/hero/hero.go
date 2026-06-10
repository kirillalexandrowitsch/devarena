package hero

import "fmt"

type Hero struct {
	ID    HeroID
	Name  string
	Class HeroClass
	Level int
	Alive bool
	CombatStats
	Attacks   [3]string
	Inventory []string
	Stats     map[string]int
	Weapon    Weapon
}

func (h Hero) TotalDamage() int {
	totalDamage := h.BaseDamage + h.BonusDamage

	if h.Weapon != nil {
		totalDamage += h.Weapon.DamageBonus()
	}

	return totalDamage
}

func (h Hero) PrintInfo() {
	fmt.Println("Hero ID:", h.ID)
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

func (h *Hero) AddItems(items ...string) {
	h.Inventory = EnsureInventoryCapacity(h.Inventory, len(h.Inventory)+len(items))

	for _, item := range items {
		h.AddItem(item)
	}
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

func (hero Hero) HasItem(item string) bool {
	return HasInventoryItem(hero.Inventory, item)
}
