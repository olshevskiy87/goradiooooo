package utils

import (
	"fmt"
)

type ErrorNoDefaultPlayer struct {
	command string
}

func (e ErrorNoDefaultPlayer) Error() string {
	return fmt.Sprintf("default command \"%s\" is not available in your system", e.command)
}
