package hero

import "errors"

var ErrHeroNameEmpty = errors.New("hero name is empty")

type ValidationError struct {
	Field   string
	Message string
}

func (err ValidationError) Error() string {
	return err.Field + ": " + err.Message
}
