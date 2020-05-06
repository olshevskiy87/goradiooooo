package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetSystemPlayerCmd(customCmd []string) ([]string, error) {
	var cmd []string
	if len(customCmd) == 0 {
		if runtime.GOOS == "linux" {
			cmd = []string{"mpv", "--no-audio-display", "--no-video", "--really-quiet"}
		} else if runtime.GOOS == "darwin" {
			cmd = []string{"play"}
		} else {
			return nil, fmt.Errorf("unsupported operating system")
		}
	} else {
		cmd = make([]string, len(customCmd))
		copy(cmd, customCmd)
	}
	_, err := exec.LookPath(cmd[0])
	if err != nil {
		return nil, fmt.Errorf("unknown command \"%s\": %v", cmd[0], err)
	}
	return cmd, nil
}
