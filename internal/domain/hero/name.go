package hero

import "fmt"

const (
	MinHeroNameLength = 3
	MaxHeroNameLength = 20
)

func HeroNameLength(name string) int {
	return len([]rune(name))
}

func ValidateHeroName(name string) error {
	nameLength := HeroNameLength(name)

	if nameLength == 0 {
		return ErrHeroNameEmpty
	}

	if nameLength < MinHeroNameLength {
		return fmt.Errorf("hero name %q is too short: length %d is less than %d", name, nameLength, MinHeroNameLength)
	}

	if nameLength > MaxHeroNameLength {
		return fmt.Errorf("hero name %q is too long: length %d is greater than %d", name, nameLength, MaxHeroNameLength)
	}

	return nil
}

func IsValidHeroName(name string) bool {
	return ValidateHeroName(name) == nil
}
