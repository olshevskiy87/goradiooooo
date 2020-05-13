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

	if in == "y" || in == "Y" {
		return true
	}
	return false
}
