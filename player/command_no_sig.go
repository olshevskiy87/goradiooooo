// +build !linux

package player

import (
	"os/exec"
)

func setSysProcAttrs(cmd *exec.Cmd) {
}
