package player

import (
	"fmt"

	"github.com/olshevskiy87/goradiooooo/params"
)

type Song struct {
	Links   map[string]string
	Artist  string
	Album   string
	Title   string
	Year    string
	Country string
	Mood    string
}

func (s *Song) String() string {
	artist := "-"
	if s.Artist != "" {
		if s.Country != "" {
			countryName, err := params.GetCountryName(s.Country)
			if err != nil {
				countryName = s.Country
			}
			artist = fmt.Sprintf("%s (%s)", s.Artist, countryName)
		} else {
			artist = s.Artist
		}
	}
	album := ""
	if s.Album != "" {
		if s.Year != "" {
			album = fmt.Sprintf(", album: %s (%s)", s.Album, s.Year)
		} else {
			album = fmt.Sprintf(", album: %s", s.Album)
		}
	}
	title := ""
	if s.Title != "" {
		if s.Mood != "" {
			title = fmt.Sprintf(", title: %s (%s)", s.Title, s.Mood)
		} else {
			title = fmt.Sprintf(", title: %s", s.Title)
		}
	}
	return fmt.Sprintf("artist: %s%s%s", artist, album, title)
}
