package logic

import (
	"fmt"
	"strings"
)

func SwitchStatementStandard(thing string) {
	switch thing {
	case "Hot Dog":
		fmt.Println("This is indeed a hotdog")
	default:
		fmt.Println("This is not a hotdog")
	}
}

func SwitchWithInit(thing string) bool {
	switch formatted := strings.ToLower(strings.Replace(thing, " ", "", 1)); formatted {
	case "hotdog":
		return true
	default:
		return false
	}
}

func SwitchWithCaseLogic(thing string) bool {
	switch {
	case strings.ToLower(strings.Replace(thing, " ", "", 5)) == "hotdog":
		return true
	default:
		return false
	}
}
