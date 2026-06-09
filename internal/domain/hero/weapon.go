package hero

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
