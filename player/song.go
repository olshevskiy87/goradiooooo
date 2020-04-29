package player

import "fmt"

type Song struct {
	Links   map[string]string
	Artist  string
	Album   string
	Title   string
	Year    string
	Country string
}

func (s *Song) String() string {
	artist := "-"
	if s.Artist != "" {
		if s.Country != "" {
			artist = fmt.Sprintf("%s (%s)", s.Artist, s.Country)
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
		title = fmt.Sprintf(", title: %s", s.Title)
	}
	return fmt.Sprintf("artist: %s%s%s", artist, album, title)
}
