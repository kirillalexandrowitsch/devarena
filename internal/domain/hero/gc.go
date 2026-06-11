package hero

func BuildHeroSummaries(heroes []Hero) []HeroSummary {
	summaries := make([]HeroSummary, 0, len(heroes))

	for _, gameHero := range heroes {
		summaries = append(summaries, gameHero.Summary())
	}

	return summaries
}

func BuildHeroSummariesInto(heroes []Hero, destination []HeroSummary) []HeroSummary {
	destination = destination[:0]

	for _, gameHero := range heroes {
		destination = append(destination, gameHero.Summary())
	}

	return destination
}

func CountAliveSummaries(summaries []HeroSummary) int {
	aliveCount := 0

	for _, summary := range summaries {
		if summary.Alive {
			aliveCount++
		}
	}

	return aliveCount
}
