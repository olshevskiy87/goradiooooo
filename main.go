package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olshevskiy87/goradiooooo/player"
)

const (
	DELAY_BETWEEN_SONGS = 2 // seconds
)

var (
	Version = "-"
)

func play(p *player.RadioooooPlayer) {
	song, err := p.GetNextSongLink()
	if err != nil {
		return
	}
	fmt.Printf("play -> %s\n", song)
	err = p.Play(song)
	if err != nil {
		return
	}
}

func main() {
	fmt.Printf("version: %s\n", Version)
	fmt.Printf("press Ctrl-C to exit\n")

	radioPlayer, err := player.New()
	if err != nil {
		fmt.Printf("could not initialize player: %v\n", err)
		os.Exit(1)
	}
	for {
		play(radioPlayer)
		time.Sleep(DELAY_BETWEEN_SONGS * time.Second)
	}
}
