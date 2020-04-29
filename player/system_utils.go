package player

import (
	"fmt"
	"runtime"
)

func getSystemPlayerCmd() ([]string, error) {
	var cmd []string
	if runtime.GOOS == "linux" {
		cmd = []string{"mpv", "--no-audio-display"}
	} else if runtime.GOOS == "darwin" {
		cmd = []string{"play"}
	} else {
		return nil, fmt.Errorf("unsupported operating system")
	}
	return cmd, nil
}
