package hero

type HeroSummary struct {
	ID    HeroID
	Name  string
	Class HeroClass
	Level int
	Alive bool
}

func NewHeroSummary(hero Hero) HeroSummary {
	return HeroSummary{
		ID:    hero.ID,
		Name:  hero.Name,
		Class: hero.Class,
		Level: hero.Level,
		Alive: hero.Alive,
	}
}

func NewHeroSummaryPointer(hero Hero) *HeroSummary {
	summary := NewHeroSummary(hero)

	return &summary
}

func (hero Hero) Summary() HeroSummary {
	return NewHeroSummary(hero)
}

func (hero Hero) SummaryPointer() *HeroSummary {
	return NewHeroSummaryPointer(hero)
}
