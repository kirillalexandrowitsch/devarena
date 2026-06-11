package hero

import "fmt"

type Weapon interface {
	Name() string
	DamageBonus() int
}

type WeaponSnapshot struct {
	Name        string
	DamageBonus int
	DynamicType string
	Equipped    bool
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

func NewWeaponSnapshot(weapon Weapon) WeaponSnapshot {
	if weapon == nil {
		return WeaponSnapshot{
			Equipped: false,
		}
	}

	return WeaponSnapshot{
		Name:        weapon.Name(),
		DamageBonus: weapon.DamageBonus(),
		DynamicType: fmt.Sprintf("%T", weapon),
		Equipped:    true,
	}
}

func (hero Hero) WeaponSnapshot() WeaponSnapshot {
	return NewWeaponSnapshot(hero.Weapon)
}
