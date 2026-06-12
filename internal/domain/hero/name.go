package hero

const (
	MinHeroNameLength = 3
	MaxHeroNameLength = 20
)

func HeroNameLength(name string) int {
	return len([]rune(name))
}

func ValidateHeroName(name string) error {
	nameLength := HeroNameLength(name)

	if nameLength < MinHeroNameLength {
		return ValidationError{Message: "hero name is too short"}
	}

	if nameLength > MaxHeroNameLength {
		return ValidationError{Message: "hero name is too long"}
	}

	return nil
}

func IsValidHeroName(name string) bool {
	return ValidateHeroName(name) == nil
}
