package params

import (
	"fmt"
	"strings"
)

const (
	MODE_RANDOM  = "random"
	MODE_EXPLORE = "explore"
	MODE_TAXI    = "taxi"
)

func canonicalizeMode(pMode string) (string, error) {
	mode := strings.ToLower(pMode)
	if mode != MODE_RANDOM && mode != MODE_EXPLORE && mode != MODE_TAXI {
		return "", fmt.Errorf("unknown mode \"%s\"", mode)
	}
	return mode, nil
}
