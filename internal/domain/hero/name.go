package hero

const (
	MinHeroNameLength = 3
	MaxHeroNameLength = 20
)

func HeroNameLength(name string) int {
	return len([]rune(name))
}

func IsValidHeroName(name string) bool {
	nameLength := HeroNameLength(name)

	return nameLength >= MinHeroNameLength && nameLength <= MaxHeroNameLength
}
