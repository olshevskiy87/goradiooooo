package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/olshevskiy87/goradiooooo/params"
	"github.com/olshevskiy87/goradiooooo/player"
)

type argsType struct {
	Mode      string          `arg:"--mode" help:"the song selection mode: random, explore, taxi"`
	Moods     []string        `arg:"--mood,-m,separate" help:"the song mood: SLOW, FAST, WEIRD. default: all moods"`
	Decades   []params.Decade `arg:"--decade,-d,separate" help:"the song decade from 1900 to 2020"`
	Countries []string        `arg:"--country,-c,separate" help:"3-letters country iso-code (e.g. GBR)"`
	SysPlayer string          `arg:"--player" help:"custom system player command"`
}

const (
	DELAY_BETWEEN_SONGS = 3 // seconds
)

var (
	Version = "-"
)

func main() {
	fmt.Printf("version: %s\n", Version)

	var args argsType
	args.Mode = params.MODE_RANDOM
	arg.MustParse(&args)

	playerParams, err := params.New(
		args.Mode,
		args.Moods,
		args.Decades,
		args.Countries,
	)
	if err != nil {
		fmt.Printf("could not initialize player params: %v\n", err)
		os.Exit(1)
	}
	if args.SysPlayer != "" {
		playerParams.Player = []string{args.SysPlayer}
	}

	radioPlayer, err := player.New(playerParams)
	if err != nil {
		fmt.Printf("could not initialize player: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("press Ctrl-C to exit\n")

	for {
		time.Sleep(DELAY_BETWEEN_SONGS * time.Second)

		song, err := radioPlayer.GetNextSongLink()
		if err != nil {
			fmt.Printf("could not get next song: %v\n", err)
			if _, ok := err.(*player.ErrorInJSON); ok {
				os.Exit(1)
			}
			continue
		}
		fmt.Printf("play -> %s\n", song)
		err = radioPlayer.Play(song)
		if err != nil {
			fmt.Printf("could not play song: %v\n", err)
			if _, ok := err.(*player.ErrorPlayerNotSpecified); ok {
				os.Exit(1)
			}
			continue
		}
	}
}
