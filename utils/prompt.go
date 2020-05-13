package utils

import (
	"fmt"
)

func AskYN(msg string) bool {
	fmt.Printf("%s ", msg)

	var in string
	_, err := fmt.Scanln(&in)
	if err != nil {
		return false
	}
	return in == "y" || in == "Y"
}
