package functions

import (
	"errors"
)

func ErrorReturn(boom bool) error {
	if !boom {
		return nil
	}
	return errors.New("Pop goes the weasel!")
}
