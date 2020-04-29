package player

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	URL_PLAY_RANDOM = "https://radiooooo.app/play/random"
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
}

func New() (*RadioooooPlayer, error) {
	playerCmd, err := getSystemPlayerCmd()
	if err != nil {
		return nil, fmt.Errorf("could not get system player: %v", err)
	}
	return &RadioooooPlayer{
		gorequest.New().Timeout(10*time.Second).
			Set("User-Agent", userAgent).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json"),
		playerCmd,
	}, nil
}

func (r *RadioooooPlayer) SetSystemPlayerCmd(cmd []string) {
	r.playerCmd = cmd
}

func (r *RadioooooPlayer) GetNextSongLink() (*Song, error) {
	_, responseBody, errs := r.requestAgent.
		Post(URL_PLAY_RANDOM).
		Send(`{"mode":"random","moods":["SLOW","FAST","WEIRD"]}`).
		End()
	if errs != nil {
		return nil, fmt.Errorf("could not perform request with url \"%s\"", URL_PLAY_RANDOM)
	}

	var songInfo map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &songInfo)
	if err != nil {
		return nil, fmt.Errorf("could not parse json: %v", err)
	}

	links, ok := songInfo["links"]
	if !ok {
		return nil, fmt.Errorf("no key \"links\" in song info")
	}

	linksStr := links.(map[string]interface{})
	linksLen := len(linksStr)
	if linksLen == 0 {
		return nil, fmt.Errorf("no song links in song info")
	}

	song := &Song{
		Links: make(map[string]string, linksLen),
	}
	for format, link := range linksStr {
		song.Links[format] = link.(string)
	}
	if artist, ok := songInfo["artist"]; ok {
		song.Artist = artist.(string)
	}
	if album, ok := songInfo["album"]; ok {
		song.Album = album.(string)
	}
	if title, ok := songInfo["title"]; ok {
		song.Title = title.(string)
	}
	if year, ok := songInfo["year"]; ok {
		song.Year = year.(string)
	}
	if country, ok := songInfo["country"]; ok {
		song.Country = country.(string)
	}

	return song, nil
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
