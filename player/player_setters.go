package player

import (
	"fmt"
	"os/exec"

	"github.com/olshevskiy87/goradiooooo/params"
)

func (r *RadioooooPlayer) SetSystemPlayerCmd(pCmd []string) error {
	if len(pCmd) == 0 {
		return fmt.Errorf("command is not specified")
	}
	_, err := exec.LookPath(pCmd[0])
	if err != nil {
		return fmt.Errorf("command \"%s\" is not available in your system", pCmd[0])
	}
	cmd := make([]string, len(pCmd))
	copy(cmd, pCmd)
	r.playerCmd = cmd

	return nil
}

func (r *RadioooooPlayer) SetParams(p *params.Params) {
	switch p.Mode {
	case params.MODE_RANDOM:
		r.url = URL_PLAY_RANDOM
	case params.MODE_EXPLORE:
		fallthrough
	case params.MODE_TAXI:
		r.url = URL_PLAY
	}
	r.params = p
}
