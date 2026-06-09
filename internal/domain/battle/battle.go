package battle

import (
	"fmt"

	"github.com/rudyakovk/devarena/internal/domain/enemy"
	"github.com/rudyakovk/devarena/internal/domain/hero"
)

type Battle struct {
	Hero  *hero.Hero
	Enemy *enemy.Enemy
	Round int
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
