package hero

import "unsafe"

type CompactHeroSummary struct {
	ID    HeroID
	Level int
	Alive bool
}

type PaddedHeroSummary struct {
	Alive bool
	ID    HeroID
	Level int
}

func CompactHeroSummarySize() uintptr {
	return unsafe.Sizeof(CompactHeroSummary{})
}

func PaddedHeroSummarySize() uintptr {
	return unsafe.Sizeof(PaddedHeroSummary{})
}

func CompactHeroSummaryFromHero(hero Hero) CompactHeroSummary {
	return CompactHeroSummary{
		ID:    hero.ID,
		Level: hero.Level,
		Alive: hero.Alive,
	}
}

func PaddedHeroSummaryFromHero(hero Hero) PaddedHeroSummary {
	return PaddedHeroSummary{
		Alive: hero.Alive,
		ID:    hero.ID,
		Level: hero.Level,
	}
}
