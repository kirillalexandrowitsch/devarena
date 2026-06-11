package hero

func BuildSummaryValue(hero Hero) HeroSummary {
	return NewHeroSummary(hero)
}

func BuildSummaryPointer(hero Hero) *HeroSummary {
	summary := NewHeroSummary(hero)

	return &summary
}

func BuildSummaryAsAny(hero Hero) any {
	summary := NewHeroSummary(hero)

	return summary
}

func BuildSummaryNameReader(hero Hero) func() string {
	summary := NewHeroSummary(hero)

	return func() string {
		return summary.Name
	}
}
