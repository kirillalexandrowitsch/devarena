package arena

import (
	"errors"

	"github.com/rudyakovk/devarena/internal/domain/hero"
)

func describeHeroCreationError(err error) string {
	if err == nil {
		return ""
	}

	if errors.Is(err, hero.ErrHeroNameEmpty) {
		return "default hero name is empty"
	}

	return "default hero creation failed: " + err.Error()
}
