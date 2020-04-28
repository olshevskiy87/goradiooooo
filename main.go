package main

import (
	"fmt"
	"time"

	"github.com/olshevskiy87/goradiooooo/player"
)

const (
	DELAY_BETWEEN_SONGS = 2 // seconds
)

var (
	Version = "-"
)

func main() {
	fmt.Printf("version: %s\n", Version)

	radioPlayer := player.New()
	for {
		radioPlayer.PlayOneSong()
		time.Sleep(DELAY_BETWEEN_SONGS * time.Second)
	}
}
