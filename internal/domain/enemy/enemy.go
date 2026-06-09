package enemy

type Enemy struct {
	Name string
	HP   int
}

func (e *Enemy) TakeDamage(damage int) {
	e.HP -= damage

	if e.HP < 0 {
		e.HP = 0
	}
}
