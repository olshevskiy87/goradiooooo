package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetSystemPlayerCmd() ([]string, error) {
	var cmd []string
	if runtime.GOOS == "linux" {
		cmd = []string{"mpv", "--no-audio-display"}
	} else if runtime.GOOS == "darwin" {
		cmd = []string{"play"}
	} else {
		return nil, fmt.Errorf("unsupported operating system")
	}
	_, err := exec.LookPath(cmd[0])
	if err != nil {
		return nil, fmt.Errorf("command \"%s\" is not available in your system", cmd[0])
	}
	return cmd, nil
}
