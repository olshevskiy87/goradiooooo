// +build linux

package player

import (
	"os/exec"
	"syscall"
)

func setSysProcAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGTERM,
	}
}
