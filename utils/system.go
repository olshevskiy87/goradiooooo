package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetSystemPlayerCmd() ([]string, error) {
	var cmd []string
	if runtime.GOOS == "linux" {
		cmd = []string{"mpv", "--no-audio-display", "--no-video", "--really-quiet"}
	} else if runtime.GOOS == "darwin" {
		cmd = []string{"play"}
	} else {
		return nil, fmt.Errorf("unsupported operating system")
	}
	_, err := exec.LookPath(cmd[0])
	if err != nil {
		return nil, &ErrorNoDefaultPlayer{cmd[0]}
	}
	return cmd, nil
}
