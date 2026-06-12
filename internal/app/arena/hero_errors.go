package arena

import "github.com/rudyakovk/devarena/internal/domain/hero"

func describeHeroCreationError(err error) string {
	if err == nil {
		return ""
	}

	if err == hero.ErrHeroNameEmpty {
		return "default hero name is empty"
	}

	return "default hero creation failed: " + err.Error()
}
