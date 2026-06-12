package hero

import "errors"

var ErrHeroNameEmpty = errors.New("hero name is empty")

type ValidationError struct {
	Message string
}

func (err ValidationError) Error() string {
	return err.Message
}
