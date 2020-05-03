package player

import (
	"fmt"
)

type ErrorInJSON struct {
	msg string
}

func (e ErrorInJSON) Error() string {
	return fmt.Sprintf("error in JSON: %s", e.msg)
}

type ErrorPlayerNotSpecified struct {
	msg string
}

func (e ErrorPlayerNotSpecified) Error() string {
	return e.msg
}
