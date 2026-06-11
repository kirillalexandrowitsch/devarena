package hero

func AllocateHeroSummaries(heroes []Hero) []HeroSummary {
	summaries := make([]HeroSummary, 0, len(heroes))

	for _, gameHero := range heroes {
		summaries = append(summaries, gameHero.Summary())
	}

	return summaries
}

func ReuseHeroSummaries(heroes []Hero, summaries []HeroSummary) []HeroSummary {
	summaries = summaries[:0]

	for _, gameHero := range heroes {
		summaries = append(summaries, gameHero.Summary())
	}

	return summaries
}

func IndexHeroesByID(heroes []Hero) map[HeroID]Hero {
	index := make(map[HeroID]Hero, len(heroes))

	for _, gameHero := range heroes {
		index[gameHero.ID] = gameHero
	}

	return index
}
