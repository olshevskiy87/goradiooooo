package player

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/olshevskiy87/goradiooooo/params"
	"github.com/olshevskiy87/goradiooooo/utils"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_PLAY_RANDOM = "https://radiooooo.app/play/random"
	URL_PLAY        = "https://radiooooo.app/play"
)

var (
	userAgent = fmt.Sprintf(
		"%s_%s:goradiooooo (by /u/olshevskiy87)",
		runtime.GOOS, runtime.GOARCH,
	)
)

type RadioooooPlayer struct {
	requestAgent *gorequest.SuperAgent
	playerCmd    []string
	url          string
	params       *params.Params
}

func New(params *params.Params) (*RadioooooPlayer, error) {
	playerCmd, err := utils.GetSystemPlayerCmd()
	if err != nil {
		return nil, fmt.Errorf("could not get system player: %v", err)
	}
	requestAgent := gorequest.New().Timeout(10*time.Second).
		Set("User-Agent", userAgent).
		Set("Accept", "application/json").
		Set("Content-Type", "application/json")

	p := &RadioooooPlayer{
		requestAgent: requestAgent,
		playerCmd:    playerCmd,
		url:          "",
	}
	p.SetParams(params)
	return p, nil
}

func (r *RadioooooPlayer) Play(song *Song) error {
	var songLink string
	for _, link := range song.Links {
		songLink = link
	}
	if songLink == "" {
		return fmt.Errorf("song link is empty")
	}
	cmdName, cmdArgs := r.playerCmd[0], append(r.playerCmd[1:], songLink)
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGTERM,
	}
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not run command %s with args %v: %v", cmdName, cmdArgs, err)
	}
	return nil
}
