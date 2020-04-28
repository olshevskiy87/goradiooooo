package player

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
}

func New() *RadioooooPlayer {
	return &RadioooooPlayer{
		gorequest.New().Timeout(10*time.Second).
			Set("User-Agent", userAgent).
			Set("Accept", "application/json").
			Set("Content-Type", "application/json"),
	}
}

func (r *RadioooooPlayer) getNextSongLink() (string, error) {
	_, responseBody, errs := r.requestAgent.
		Post(URL_PLAY_RANDOM).
		Send(`{"mode":"random","moods":["SLOW","FAST","WEIRD"]}`).
		End()
	if errs != nil {
		return "", fmt.Errorf("could not perform request with url \"%s\"", URL_PLAY_RANDOM)
	}

	var songInfo map[string]interface{}
	err := json.Unmarshal([]byte(responseBody), &songInfo)
	if err != nil {
		return "", fmt.Errorf("could not parse json: %v", err)
	}

	links, ok := songInfo["links"]
	if !ok {
		return "", fmt.Errorf("no key \"links\" in song info")
	}

	linksMap := links.(map[string]interface{})
	for _, link := range linksMap {
		return link.(string), nil
	}
	return "", fmt.Errorf("no song links in song info")
}

func (r *RadioooooPlayer) PlayOneSong() {
	link, err := r.getNextSongLink()
	if err != nil {
		return
	}
	cmd := exec.Command("mpv", "--no-audio-display", link)
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
