package logic

import (
	"errors"
)

func IfStatement(name string) error {
	if name == "wrong" {
		return errors.New("Wrong name!!")
	}
	return nil
}

func IfStatementWithInit(ageGuess int) (error, bool) {
	if invalidAge := (ageGuess > 160 || ageGuess < 0); invalidAge {
		return errors.New("Age cant be this high"), false
	}
	correctGuess := ageGuess == 33
	return nil, correctGuess
}
